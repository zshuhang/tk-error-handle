package model

type CategoryRelationRequest struct {
	CategoryIds []string `json:"category_ids"`
	RegionList  []string `json:"region_list"`
}

type CategoryRelationResponse struct {
	BaseResp      BaseResp            `json:"base_resp"`
	IdRelationMap map[string]Relation `json:"id_relation_map"`
}

type Relation struct {
	PropList      []Prop      `json:"prop_list"`
	PropValueList []PropValue `json:"prop_value_list"`
}

type Prop struct {
	CreateTime    int64  `json:"create_time"`
	ModifyTime    int64  `json:"modify_time"`
	NameCn        string `json:"name_cn"`
	NameEn        string `json:"name_en"`
	PropertyCode  string `json:"property_code"`
	PropertyId    int64  `json:"property_id"`
	PropertyType  int64  `json:"property_type"`
	TtsPropertyId int64  `json:"tts_property_id"`
}

type PropValue struct {
	CreateTime         int64  `json:"create_time"`
	IsPrivateValue     bool   `json:"is_private_value"`
	ModifyTime         int64  `json:"modify_time"`
	NameCn             string `json:"name_cn"`
	NameEn             string `json:"name_en"`
	PropertyValueCode  string `json:"property_value_code"`
	PropertyValueId    int64  `json:"property_value_id"`
	TtsPropertyValueId int64  `json:"tts_property_value_id"`
}
