package model

// 请求数据
type ProductDescRequest struct {
	SpuCode       string `json:"spu_code"`
	ReverseStatus int64  `json:"reverse_status"`
}

// 响应数据
type ProductDescResponse struct {
	BaseResp        BaseResp        `json:"base_resp"`
	Info            ProductDescInfo `json:"info"`
	RelativeTaskIds []int64         `json:"relative_task_ids"`
}

type ProductDescInfo struct {
	SpuDetail ProductDesc `json:"spu_detail"`
}

type ProductDesc struct {
	BrandID             int64               `json:"brand_id"`
	CategoryID          int64               `json:"category_id"`
	ExcludeRegionCodes  []any               `json:"exclude_region_codes"`
	ManufacturerInfos   []ManufacturerInfo  `json:"manufacturer_infos"`
	PictureUrls         []any               `json:"picture_urls"`
	ProductDescEn       string              `json:"product_desc_en"`
	ProductMediaInfo    ProductMediaInfo    `json:"product_media_info"`
	ProductName         string              `json:"product_name"`
	ProductNameEn       string              `json:"product_name_en"`
	PropertiesV2        []PropertiesV2      `json:"properties_v2"`
	RpModels            []RpModel           `json:"rp_models"`
	SalePropertyList    []SalePropertyList  `json:"sale_property_list"`
	SecurityWarningInfo SecurityWarningInfo `json:"security_warning_info"`
	SkcDetails          []SkcDetail         `json:"skc_details"`
	SpuCode             string              `json:"spu_code"`
	TicketCode          string              `json:"ticket_code"`
}

type ManufacturerInfo struct {
	ID string `json:"id"`
}

type ProductMediaInfo struct {
	PicSet PicSet `json:"pic_set"`
}

type PicSet struct {
	Extra          PicSetExtra                    `json:"extra"`
	ObjectMaterial map[string][]ObjectMaterialObj `json:"object_material"`
	SetType        int64                          `json:"set_type"`
	SpuMaterial    []SpuMaterial                  `json:"spu_material"`
}

type PicSetExtra struct {
	SkcSkuMapping map[string][]string `json:"skc_sku_mapping"`
}

type ObjectMaterialObj struct {
	ID               int64    `json:"id"`
	LinkType         int64    `json:"link_type"`
	Material         Material `json:"material"`
	MaterialShowType int64    `json:"material_show_type"`
	OrderNum         int64    `json:"order_num"`
}

type Material struct {
	Extra          MaterialExtra     `json:"extra"`
	ID             int64             `json:"id"`
	MaterialStatus int64             `json:"material_status"`
	MaterialType   int64             `json:"material_type"`
	Name           string            `json:"name"`
	ParentID       int64             `json:"parent_id"`
	RecognitionRes []RecognitionRe   `json:"recognition_res"`
	SellerID       int64             `json:"seller_id"`
	ShopID         int64             `json:"shop_id"`
	Size           int64             `json:"size"`
	URI            string            `json:"uri"`
	URLMap         map[string]string `json:"urlMap"`
	Vid            string            `json:"vid"`
}

type MaterialExtra struct {
	Format       string `json:"format"`
	Height       int64  `json:"height"`
	Name         string `json:"name"`
	Resolution   string `json:"resolution"`
	Size         int64  `json:"size"`
	TargetHeight int64  `json:"target_height"`
	TargetWidth  int64  `json:"target_width"`
	URIVa        string `json:"uri_va"`
	VDuration    int64  `json:"v_duration"`
	Width        int64  `json:"width"`
}

type RecognitionRe struct {
	Actions              []int64 `json:"actions"`
	PicRecID             int64   `json:"pic_rec_id"`
	RecTimeMS            int64   `json:"rec_time_ms"`
	RecognitionAlgorithm int64   `json:"recognitionAlgorithm"`
	Score                int64   `json:"score"`
	Status               int64   `json:"status"`
	Type                 int64   `json:"type"`
}

type SpuMaterial struct {
	Extra                 any      `json:"extra"`
	ID                    int64    `json:"id"`
	IsMain                bool     `json:"is_main"`
	LinkType              int64    `json:"link_type"`
	Material              Material `json:"material"`
	MaterialShowType      int64    `json:"material_show_type"`
	MaterialUseTypeCode   string   `json:"material_use_type_code"`
	MaterialUseTypeName   string   `json:"material_use_type_name"`
	ObjectCode            string   `json:"object_code"`
	OrderNum              int64    `json:"order_num"`
	ProductMaterialStatus int64    `json:"product_material_status"`
	SellerId              int64    `json:"seller_id"`
	ShopId                int64    `json:"shop_id"`
	SpuCode               string   `json:"spu_code"`
	Tags                  []any    `json:"tags"`
}

type PropertiesV2 struct {
	Property PropertiesV2Property `json:"property"`
	Region   string               `json:"region"`
}

type PropertiesV2Property struct {
	PropertyCode      string              `json:"property_code"`
	PropertyID        int64               `json:"property_id"`
	PropertyValueList []PropertyValueList `json:"property_value_list"`
}

type PropertyValueList struct {
	PropertyValueCN   string `json:"property_value_cn"`
	PropertyValueCode string `json:"property_value_code"`
	PropertyValueEn   string `json:"property_value_en"`
	PropertyValueID   int64  `json:"property_value_id"`
}

type RpModel struct {
	ID string `json:"id"`
}

type SalePropertyList struct {
	Property       SaleProperty        `json:"property"`
	PropertyValues []SalePropertyValue `json:"property_values"`
}

type SaleProperty struct {
	PropertyID    int64 `json:"property_id"`
	TTSPropertyID int64 `json:"tts_property_id"`
}

type SalePropertyValue struct {
	PropertyValueID    int64 `json:"property_value_id"`
	TTSPropertyValueID int64 `json:"tts_property_value_id"`
}

type SecurityWarningInfo struct {
	SecurityWarningLanguages []any `json:"security_warning_languages"`
}

type SkcDetail struct {
	ArticleNumber            string                `json:"article_number"`
	Index                    int64                 `json:"index"`
	KeyAttributeOneValueID   int64                 `json:"key_attribute_one_value_id"`
	KeyAttributeOneValueInfo ValueInfo             `json:"key_attribute_one_value_info"`
	PicSetInfo               []any                 `json:"pic_set_info"`
	PicSupplyWay             int64                 `json:"pic_supply_way"`
	PictureUrls              []any                 `json:"picture_urls"`
	PreviewPicList           []any                 `json:"preview_pic_list"`
	SalePropertyValueInfo    SalePropertyValueInfo `json:"sale_property_value_info"`
	SkcCode                  string                `json:"skc_code"`
	SkuDetails               []SkuDetail           `json:"sku_details"`
}

type ValueInfo struct {
	CreateTime              int64  `json:"create_time"`
	EnabledStatus           int64  `json:"enabled_status"`
	IsExternalPropertyValue bool   `json:"is_external_property_value"`
	ItemNameCn              string `json:"item_name_cn"`
	ItemNameEn              string `json:"item_name_en"`
	ModifyTime              int64  `json:"modify_time"`
	ProductNameCn           string `json:"product_name_cn"`
	ProductNameEn           string `json:"product_name_en"`
	PropertyCode            string `json:"property_code"`
	PropertyId              int64  `json:"property_id"`
	PropertyName            string `json:"property_name"`
	PropertyValueCode       string `json:"property_value_code"`
	PropertyValueId         int64  `json:"property_value_id"`
	Sequence                int64  `json:"sequence"`
}

type SalePropertyValueInfo struct {
	CreateTime         int64  `json:"create_time"`
	IsPrivateValue     bool   `json:"is_private_value"`
	ModifyTime         int64  `json:"modify_time"`
	NameCn             string `json:"name_cn"`
	NameEn             string `json:"name_en"`
	PropertyValueCode  string `json:"property_value_code"`
	PropertyValueId    int64  `json:"property_value_id"`
	TTSPropertyValueId int64  `json:"tts_property_value_id"`
}

type SkuDetail struct {
	AddSpecsType             int64                   `json:"add_specs_type"`
	ArticleNumber            string                  `json:"article_number"`
	ComboRelations           []any                   `json:"combo_relations"`
	ComboType                int64                   `json:"combo_type"`
	Description              string                  `json:"description"`
	Index                    int64                   `json:"index"`
	KeyAttributeTwoValueID   int64                   `json:"key_attribute_two_value_id"`
	KeyAttributeTwoValueInfo ValueInfo               `json:"key_attribute_two_value_info"`
	MainVersionSkuStatus     int64                   `json:"main_version_sku_status"`
	PackageLongestLength     int64                   `json:"package_longest_length"`
	PackageMiddleLength      int64                   `json:"package_middle_length"`
	PackageShortestLength    int64                   `json:"package_shortest_length"`
	PackageWeight            int64                   `json:"package_weight"`
	Price                    int64                   `json:"price"`
	PriceCurrencyType        int64                   `json:"price_currency_type"`
	SalePropertyList         []SalePropertyValueInfo `json:"sale_property_list"`
	SkuCode                  string                  `json:"sku_code"`
	SkuStatus                int64                   `json:"sku_status"`
	SkuStatusV2              int64                   `json:"sku_status_v2"`
	Stock                    int64                   `json:"stock"`
	StockMode                int64                   `json:"stock_mode"`
}
