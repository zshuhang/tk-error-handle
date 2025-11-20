package model

type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PageInfo struct {
	HasMore    int `json:"has_more"`
	PageNo     int `json:"page_no"`
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`
}
