package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	// "github.com/dgrijalva/jwt-go/request"
	"github.com/samber/lo"

	"strconv"

	"tk-error-handle/http"
	M "tk-error-handle/model"
)

var ctx = context.Background()

var sessionId = "bfa765bd3284cededa8eed1da8ad5ea3"

func main() {
	fmt.Println("========================================")
	fmt.Println("  TK商品异常处理工具")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("⚠️  注意：此脚本目前只处理图片异常，搜索出的异常数据只包含图片异常")
	fmt.Println()
	fmt.Println("📋 sessionId 获取步骤：")
	fmt.Println("   1. 登录 TK 商家中心")
	fmt.Println("   2. 按 F12 打开开发者工具")
	fmt.Println("   3. 点击「应用程序」(Application)")
	fmt.Println("   4. 双击「Cookie」")
	fmt.Println("   5. 选择 Cookie 下面的第一个选项")
	fmt.Println("   6. 在右侧表格的「名称」列中找到 sessionid")
	fmt.Println("   7. 复制其对应的值")
	fmt.Println()
	fmt.Print("请输入 sessionId: ")
	fmt.Scanln(&sessionId)
	if sessionId == "" {
		panic(errors.New("sessionId 不能为空"))
	}

	products := GetProductList()

	fmt.Println()
	for _, product := range products {
		appealStatus := "--"
		for _, item := range product.ReverseItem {
			if item.AppealOrderInfo != nil {
				appealStatus = "是"
				break
			}
		}
		fmt.Printf("SPU:%s   申诉状态：%s   货号：%s\n", product.SpuCode, appealStatus, product.ArticleNumber)
	}
	fmt.Printf("查询到%d个异常待处理，按回车开始处理\n", len(products))
	fmt.Scanln()

	for _, product := range products {
		fmt.Printf("当前处理spu %s\n", product.SpuCode)

		// productDesc, RelativeTaskIds = GetProductDesc(product.SpuCode)
		productDesc, _ := GetProductDesc(product.SpuCode)

		propList, propValueList := GetCategoryRelation(strconv.FormatInt(productDesc.CategoryID, 10))

		checkResult, uriToCheckResult := GetCheckProductResult(productDesc, propList, propValueList)

		// data, err := json.MarshalIndent(checkResult, "", "  ")
		// if err != nil {
		// 	panic(err)
		// }
		// os.WriteFile("qqqqq.json", data, 0644)

		// data, err = json.MarshalIndent(uriToCheckResult, "", "  ")
		// if err != nil {
		// 	panic(err)
		// }
		// os.WriteFile("wwwww.json", data, 0644)

		// time.Sleep(30 * time.Second)
	}
}

func GetProductList() []M.Product {
	var articleNumber = ""
	var pageNo = 1
	var pageSize = 10
	var excludeAppeal = 1

	fmt.Print("请输入查询关键字（默认为空）: ")
	fmt.Scanln(&articleNumber)

	fmt.Print("请输入页码（默认为1）: ")
	fmt.Scanln(&pageNo)

	fmt.Print("请输入每页数量（默认为10）: ")
	fmt.Scanln(&pageSize)

	fmt.Print("是否要排除处于申诉状态的?（默认排除，排除：1，不排除：0）: ")
	fmt.Scanln(&excludeAppeal)

	if excludeAppeal != 1 && excludeAppeal != 0 {
		panic(errors.New("输入内容不合法"))
	}

	request := M.ProductListRequest{
		Filter: M.Filter{
			ArticleNumber: articleNumber,
			ReverseStatus: 10,
			ReverseType:   []int{6}, // 默认只处理图片异常
		},
		PageInfo: M.PageInfo{
			PageNo:   pageNo,
			PageSize: pageSize,
		},
	}

	var response M.ProductListResponse

	err := http.Request("POST", "/reverse/list", sessionId, ctx, &request, &response)
	if err != nil {
		panic(err)
	}

	// 根据 excludeAppeal 决定是否排除申诉状态的产品
	if excludeAppeal == 1 {
		var filteredProducts []M.Product
		for _, product := range response.Products {
			hasAppeal := false
			// 检查是否有申诉状态
			for _, item := range product.ReverseItem {
				if item.AppealOrderInfo != nil {
					hasAppeal = true
					break
				}
			}
			// 如果没有申诉状态，则保留
			if !hasAppeal {
				filteredProducts = append(filteredProducts, product)
			}
		}
		return filteredProducts
	}

	return response.Products
}

func GetProductDesc(spuCode string) (M.ProductDesc, []int64) {
	request := M.ProductDescRequest{
		SpuCode:       spuCode,
		ReverseStatus: 10,
	}

	var response M.ProductDescResponse

	err := http.Request("POST", "/reverse/get_detail", sessionId, ctx, &request, &response)
	if err != nil {
		panic(err)
	}

	return response.Info.SpuDetail, response.RelativeTaskIds
}

func GetCategoryRelation(categoryId string) ([]M.Prop, []M.PropValue) {
	request := M.CategoryRelationRequest{
		CategoryIds: []string{categoryId},
		RegionList:  []string{"SA", "GB", "US", "FR", "DE", "IT", "ES", "MX", "JP"},
	}

	var response M.CategoryRelationResponse
	err := http.Request("POST", "/category/m_get_category_prop_relation", sessionId, ctx, &request, &response)
	if err != nil {
		panic(err)
	}

	return response.IdRelationMap[categoryId].PropList, response.IdRelationMap[categoryId].PropValueList
}

func GetCheckProductResult(productDesc M.ProductDesc, propList []M.Prop, propValueList []M.PropValue) ([]M.CheckResult, map[string]M.UriToCheckResult) {
	request := M.CheckProductRequest{
		CheckOption: M.CheckOption{
			CheckPrice:          false,
			CheckCertification:  false,
			CheckPackage:        false,
			CheckPic:            true,
			CheckProductDescPic: false,
		},
		ProductInfo: M.ProductInfo{},
	}

	propertiesV2 := lo.Map(productDesc.PropertiesV2, func(item M.PropertiesV2, _ int) M.CheckProductPropertiesV2 {
		propertyRelation, _ := lo.Find(propList, func(el M.Prop) bool {
			return el.PropertyId == item.Property.PropertyID
		})

		checkProductPropertyList := lo.Map(item.Property.PropertyValueList, func(el M.PropertyValue, _ int) M.CheckProductPropertyValue {
			propertyValueRelation, ok := lo.Find(propValueList, func(elItem M.PropValue) bool { return elItem.PropertyValueId == el.PropertyValueID })
			data := M.CheckProductPropertyValue{}
			data.PropertyValueCN = el.PropertyValueCN
			data.PropertyValueCode = el.PropertyValueCode
			data.PropertyValueEn = el.PropertyValueEn
			data.PropertyValueID = strconv.FormatInt(el.PropertyValueID, 10)
			if ok {
				TTSPropertyValueId := strconv.FormatInt(propertyValueRelation.TTSPropertyValueId, 10)
				if TTSPropertyValueId != "0" {
					data.TTSAttributeValueID = TTSPropertyValueId
				}
			}
			return data
		})

		checkProductProperty := M.CheckProductProperty{}
		checkProductProperty.PropertyID = strconv.FormatInt(item.Property.PropertyID, 10)
		checkProductProperty.PropertyCode = item.Property.PropertyCode
		checkProductProperty.PropertyValueList = checkProductPropertyList
		checkProductProperty.TTSAttributeID = strconv.FormatInt(propertyRelation.TTSPropertyId, 10)

		checkProductPropertiesV2 := M.CheckProductPropertiesV2{}
		checkProductPropertiesV2.Property = checkProductProperty
		checkProductPropertiesV2.Region = item.Region

		return checkProductPropertiesV2
	})

	salePropertyIDList := lo.Map(productDesc.SalePropertyList, func(item M.SalePropertyList, _ int) M.SalePropertyIDList {
		salePropertyIDItem := M.SalePropertyIDList{}
		salePropertyIDItem.PropertyID = strconv.FormatInt(item.Property.PropertyID, 10)
		salePropertyIDItem.TTSPropertyID = strconv.FormatInt(item.Property.TTSPropertyID, 10)
		return salePropertyIDItem
	})

	pictureList := lo.Map(productDesc.ProductMediaInfo.PicSet.SpuMaterial, func(item M.SpuMaterial, _ int) M.CheckProductPicture {
		extra := M.CheckProductMaterialExtra{}
		extra.Format = item.Material.Extra.Format
		extra.Height = strconv.FormatInt(item.Material.Extra.Height, 10)
		extra.Name = item.Material.Extra.Name
		extra.Resolution = item.Material.Extra.Resolution
		extra.Size = strconv.FormatInt(item.Material.Extra.Size, 10)
		extra.TargetHeight = strconv.FormatInt(item.Material.Extra.TargetHeight, 10)
		extra.TargetWidth = strconv.FormatInt(item.Material.Extra.TargetWidth, 10)
		extra.URIVa = item.Material.Extra.URIVa
		extra.VDuration = strconv.FormatInt(item.Material.Extra.VDuration, 10)
		extra.Width = strconv.FormatInt(item.Material.Extra.Width, 10)

		recognitionRes := lo.Map(item.Material.RecognitionRes, func(el M.RecognitionRe, _ int) M.CheckProductRecognition {
			checkProductRecognition := M.CheckProductRecognition{}
			checkProductRecognition.Actions = el.Actions
			checkProductRecognition.PicRecID = strconv.FormatInt(el.PicRecID, 10)
			checkProductRecognition.RecTimeMS = strconv.FormatInt(el.RecTimeMS, 10)
			checkProductRecognition.RecognitionAlgorithm = el.RecognitionAlgorithm
			checkProductRecognition.Score = el.Score
			checkProductRecognition.Status = el.Status
			checkProductRecognition.Type = el.Type
			return checkProductRecognition
		})

		material := M.CheckProductMaterial{
			Extra:          extra,
			ID:             strconv.FormatInt(item.Material.ID, 10),
			MaterialStatus: item.Material.MaterialStatus,
			MaterialType:   item.Material.MaterialType,
			Name:           item.Material.Name,
			ParentID:       strconv.FormatInt(item.Material.ParentID, 10),
			RecognitionRes: recognitionRes,
			SellerID:       strconv.FormatInt(item.Material.SellerID, 10),
			ShopID:         strconv.FormatInt(item.Material.ShopID, 10),
			Size:           strconv.FormatInt(item.Material.Size, 10),
			URI:            item.Material.URI,
			URLMap:         item.Material.URLMap,
			Vid:            item.Material.Vid,
		}

		checkProductPicture := M.CheckProductPicture{}
		checkProductPicture.ID = strconv.FormatInt(item.ID, 10)
		checkProductPicture.LinkType = item.LinkType
		checkProductPicture.MaterialShowType = item.MaterialShowType
		checkProductPicture.MaterialUseTypeCode = &item.MaterialUseTypeCode
		checkProductPicture.OrderNum = strconv.FormatInt(item.OrderNum, 10)
		checkProductPicture.Material = material

		return checkProductPicture
	})

	mediaInfo := M.CheckProductMediaInfo{
		PictureList: pictureList,
		PicType:     2, // TODO PicType 收集不到此值，默认为2
		PicSetType:  productDesc.ProductMediaInfo.PicSet.SetType,
	}

	manufacturerIDS := lo.Map(productDesc.ManufacturerInfos, func(item M.ManufacturerInfo, _ int) string {
		return item.ID
	})

	rpIDS := lo.Map(productDesc.RpModels, func(item M.RpModel, _ int) string {
		return item.ID
	})

	skcDetails := lo.Map(productDesc.SkcDetails, func(item M.SkcDetail, _ int) M.CheckProductSkcDetail {
		checkProductSkcDetail := M.CheckProductSkcDetail{}
		checkProductSkcDetail.Index = strconv.FormatInt(item.Index, 10)
		checkProductSkcDetail.SkcCode = item.SkcCode

		checkProductSkcDetailSaleProperty := M.CheckProductSaleProperty{}
		checkProductSkcDetailSaleProperty.PropertyValueID = strconv.FormatInt(item.SalePropertyValueInfo.PropertyValueId, 10)
		valueId := strconv.FormatInt(item.SalePropertyValueInfo.TTSPropertyValueId, 10)
		checkProductSkcDetailSaleProperty.TTSPropertyValueID = &valueId
		checkProductSkcDetail.SaleProperty = checkProductSkcDetailSaleProperty

		checkProductSkcDetailMediaInfo := M.CheckProductSkcDetailMediaInfo{}
		// checkProductSkcDetailMediaInfo.PictureList = item.PictureUrls
		checkProductSkcDetailMediaInfo.PictureList = []M.CheckProductPicture{} // TODO PictureList 这里的值应该是PictureUrls，但是由于类型定义过于冗余，现在赋值有问题，默认赋值为空数组
		checkProductSkcDetailMediaInfo.PicType = 2                             // TODO PicType 收集不到此值，默认为2
		checkProductSkcDetail.MediaInfo = checkProductSkcDetailMediaInfo

		checkProductSkcDetail.SkuDetails = lo.Map(item.SkuDetails, func(el M.SkuDetail, _ int) M.CheckProductSkuDetail {
			checkProductSkuDetail := M.CheckProductSkuDetail{}
			checkProductSkuDetail.SkuCode = el.SkuCode

			mediaInfo := M.CheckProductSkcDetailMediaInfo{}
			mediaInfo.PictureList = lo.Map(productDesc.ProductMediaInfo.PicSet.ObjectMaterial[el.SkuCode], func(elItem M.ObjectMaterialObj, _ int) M.CheckProductPicture {
				checkProductPicture := M.CheckProductPicture{}
				checkProductPicture.ID = strconv.FormatInt(elItem.ID, 10)
				checkProductPicture.MaterialShowType = elItem.MaterialShowType
				checkProductPicture.LinkType = elItem.LinkType
				checkProductPicture.OrderNum = strconv.FormatInt(elItem.OrderNum, 10)

				extra := M.CheckProductMaterialExtra{}
				extra.Format = elItem.Material.Extra.Format
				extra.Height = strconv.FormatInt(elItem.Material.Extra.Height, 10)
				extra.Name = elItem.Material.Extra.Name
				extra.Resolution = elItem.Material.Extra.Resolution
				extra.Size = strconv.FormatInt(elItem.Material.Extra.Size, 10)
				extra.TargetHeight = strconv.FormatInt(elItem.Material.Extra.TargetHeight, 10)
				extra.TargetWidth = strconv.FormatInt(elItem.Material.Extra.TargetWidth, 10)
				extra.URIVa = elItem.Material.Extra.URIVa
				extra.VDuration = strconv.FormatInt(elItem.Material.Extra.VDuration, 10)
				extra.Width = strconv.FormatInt(elItem.Material.Extra.Width, 10)

				recognitionRes := lo.Map(elItem.Material.RecognitionRes, func(el M.RecognitionRe, _ int) M.CheckProductRecognition {
					checkProductRecognition := M.CheckProductRecognition{}
					checkProductRecognition.Actions = el.Actions
					checkProductRecognition.PicRecID = strconv.FormatInt(el.PicRecID, 10)
					checkProductRecognition.RecTimeMS = strconv.FormatInt(el.RecTimeMS, 10)
					checkProductRecognition.RecognitionAlgorithm = el.RecognitionAlgorithm
					checkProductRecognition.Score = el.Score
					checkProductRecognition.Status = el.Status
					checkProductRecognition.Type = el.Type
					return checkProductRecognition
				})

				material := M.CheckProductMaterial{
					Extra:          extra,
					ID:             strconv.FormatInt(elItem.Material.ID, 10),
					MaterialStatus: elItem.Material.MaterialStatus,
					MaterialType:   elItem.Material.MaterialType,
					Name:           elItem.Material.Name,
					ParentID:       strconv.FormatInt(elItem.Material.ParentID, 10),
					RecognitionRes: recognitionRes,
					SellerID:       strconv.FormatInt(elItem.Material.SellerID, 10),
					ShopID:         strconv.FormatInt(elItem.Material.ShopID, 10),
					Size:           strconv.FormatInt(elItem.Material.Size, 10),
					URI:            elItem.Material.URI,
					URLMap:         elItem.Material.URLMap,
					Vid:            elItem.Material.Vid,
				}

				checkProductPicture.Material = material
				return checkProductPicture
			})
			mediaInfo.PicType = 2 // TODO PicType 收集不到此值，默认为2
			checkProductSkuDetail.MediaInfo = mediaInfo

			checkProductSkuDetail.SalePropertyList = lo.Map(el.SalePropertyList, func(elItem M.SalePropertyValueInfo, _ int) M.CheckProductSaleProperty {
				checkProductSaleProperty := M.CheckProductSaleProperty{}
				checkProductSaleProperty.PropertyValueID = strconv.FormatInt(elItem.PropertyValueId, 10)
				valueId := strconv.FormatInt(elItem.TTSPropertyValueId, 10)
				checkProductSaleProperty.TTSPropertyValueID = &valueId
				return checkProductSaleProperty
			})
			checkProductSkuDetail.PackageLongestLength = strconv.FormatInt(el.PackageLongestLength, 10)
			checkProductSkuDetail.PackageShortestLength = strconv.FormatInt(el.PackageShortestLength, 10)
			checkProductSkuDetail.PackageMiddleLength = strconv.FormatInt(el.PackageMiddleLength, 10)
			checkProductSkuDetail.PackageWeight = strconv.FormatInt(el.PackageWeight, 10)
			checkProductSkuDetail.ArticleNumber = el.ArticleNumber
			checkProductSkuDetail.Price = strconv.FormatInt(el.Price, 10)
			checkProductSkuDetail.ProductStatus = true // TODO ProductStatus 收集不到此值，默认为true
			checkProductSkuDetail.Stock = strconv.FormatInt(el.Stock, 10)
			checkProductSkuDetail.SupplyPriceCurrencyType = el.PriceCurrencyType
			checkProductSkuDetail.GoodsInStock = true // TODO GoodsInStock 收集不到此值，默认为true
			checkProductSkuDetail.StockMode = el.StockMode
			return checkProductSkuDetail
		})

		checkProductSkcDetail.ArticleNumber = item.ArticleNumber
		checkProductSkcDetail.StockMode = item.SkuDetails[0].StockMode

		return checkProductSkcDetail
	})

	salePropertyValueList := lo.Map(productDesc.SalePropertyList, func(item M.SalePropertyList, _ int) []M.SalePropertyValueList {
		return lo.Map(item.PropertyValues, func(el M.SalePropertyValue, _ int) M.SalePropertyValueList {
			salePropertyValueList := M.SalePropertyValueList{}
			salePropertyValueList.PlmPropertyValueID = strconv.FormatInt(el.PropertyValueID, 10)
			valueId := strconv.FormatInt(el.TTSPropertyValueID, 10)
			salePropertyValueList.PlmTTSAttributeValueID = &valueId
			return salePropertyValueList
		})
	})

	request.ProductInfo.ProductName = productDesc.ProductName
	request.ProductInfo.ProductNameEn = productDesc.ProductNameEn
	request.ProductInfo.CategoryID = strconv.FormatInt(productDesc.CategoryID, 10)
	request.ProductInfo.BrandID = nil
	request.ProductInfo.PropertiesV2 = propertiesV2
	request.ProductInfo.SecurityWarningInfo = M.CheckProductSecurityWarningInfo(productDesc.SecurityWarningInfo)
	request.ProductInfo.SalePropertyIDList = salePropertyIDList
	request.ProductInfo.VideoList = []any{}
	request.ProductInfo.MediaInfo = mediaInfo
	request.ProductInfo.Grading = struct{}{}
	request.ProductInfo.ProductDescEn = productDesc.ProductDescEn
	request.ProductInfo.Certifications = []any{}
	request.ProductInfo.ExcludeRegionCodes = productDesc.ExcludeRegionCodes
	request.ProductInfo.ManufacturerIDS = manufacturerIDS
	request.ProductInfo.RpIDS = rpIDS
	request.ProductInfo.SkcDetails = skcDetails
	request.ProductInfo.SalePropertyValueList = salePropertyValueList
	request.ProductInfo.TicketCode = productDesc.TicketCode
	request.ProductInfo.SpuCode = productDesc.SpuCode

	var response M.CheckProductResponse
	err := http.Request("POST", "/check/check_product", sessionId, ctx, &request, &response)
	if err != nil {
		panic(err)
	}

	return response.PictureCheckResult.CheckResultMap, response.PictureCheckResult.UriToCheckResultMap
}
