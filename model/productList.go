package model

// 请求数据
type ProductListRequest struct {
	Filter   Filter          `json:"filter"`
	PageInfo PageInfoRequest `json:"page_info"`
}

type Filter struct {
	ArticleNumber string `json:"article_number"`
	ReverseStatus int    `json:"reverse_status"`
	ReverseType   []int  `json:"reverse_type"`
}

// 响应数据
type ProductListResponse struct {
	BaseResp BaseResp         `json:"base_resp"`
	PageInfo PageInfoResponse `json:"page_info"`
	Products []Product        `json:"product_reverse_items"`
}

type Product struct {
	ArticleNumber string        `json:"article_number"`
	SpuCode       string        `json:"spu_code"`
	ReverseItem   []ReverseItem `json:"reverse_item"`
}

type ReverseItem struct {
	AppealOrderInfo *any `json:"appeal_order_info,omitempty"`
}
