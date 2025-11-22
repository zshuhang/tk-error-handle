package model

// 请求数据
type AppealOrederRequest struct {
	SpuDetail         ProductInfo         `json:"spu_detail"`
	Scene             int64               `json:"scene"`
	PicIssues         map[string][]string `json:"pic_issues"`
	AppealPictures    []AppealPicture     `json:"appeal_pictures"`
	AppealSceneParams AppealSceneParams   `json:"appeal_scene_params"`
}

type AppealPicture struct {
	Uri     string   `json:"uri"`
	Issues  []string `json:"issues"`
	PicType int64    `json:"pic_type"`
}

type AppealSceneParams struct {
	FeedbackContent any      `json:"feedback_content"`
	ReverseTaskIds  []string `json:"reverse_task_ids"`
}

// 响应数据
type AppealOrederResponse struct {
	BaseResp BaseResp `json:"base_resp"`
	ID       int64   `json:"id"`
}
