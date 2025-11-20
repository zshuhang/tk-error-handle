package main

import (
	"context"
	"errors"
	"fmt"
	"tk-error-handle/http"
	M "tk-error-handle/model"
)

var ctx = context.Background()

var sessionId = "bfa765bd3284cededa8eed1da8ad5ea3"

func main() {
	fmt.Println("此脚本目前只处理图片异常，所以搜索出的异常数据只包含图片异常！！！！！！ ")
	fmt.Println("sessionId如何获取？")
	fmt.Println("登录TK商家中心->按F12->点击应用程序(应用、Application)->双击Cookie->")
	fmt.Println("选Cookie下面的第一个->右边会弹出一个表格->在名称列下找sessionid->复制其对应的值")

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
	fmt.Printf("查询到%d个异常待处理\n", len(products))
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
