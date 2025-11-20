package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
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

		fmt.Println(propList, propValueList)
		// data, err := json.MarshalIndent(propList, "", "  ")
		// if err != nil {
		// 	panic(err)
		// }
		// os.WriteFile("propListlll.json", data, 0644)

		// data, err = json.MarshalIndent(propValueList, "", "  ")
		// if err != nil {
		// 	panic(err)
		// }
		// os.WriteFile("propValueListlll.json", data, 0644)

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
