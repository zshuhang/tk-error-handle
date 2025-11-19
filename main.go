package main

import "encoding/json"

// import (
// 	"context"
// 	"fmt"
// 	"github.com/carlmjohnson/requests"
// )

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

// type Res struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// 	Result  struct {
// 		From string `json:"from"`
// 		Name string `json:"nae"`
// 	} `json:"result"`
// }

func main() {
	// var result Res

	// err := requests.
	// 	URL("https://api.apiopen.top/api/sentences").
	// 	ToJSON(&result).
	// 	Fetch(context.Background())

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%+v", result)
}

func UnmarshalProductListResponse(data []byte) (ProductListResponse, error) {
	var r ProductListResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ProductListResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ProductListResponse struct {
	BaseResp            BaseResp             `json:"base_resp"`
	PageInfo            PageInfo             `json:"page_info"`
	ProductReverseItems []ProductReverseItem `json:"product_reverse_items"`
}

type ProductReverseItem struct {
	ArticleNumber            string                              `json:"article_number"`
	FullProductInfo          FullProductInfo                     `json:"fullProductInfo"`
	IsKeyProduct             int64                               `json:"is_key_product"`
	KeyAttributeOne          ProductReverseItemKeyAttributeOne   `json:"key_attribute_one"`
	KeyAttributeOneValueMap  map[string]KeyAttributeOneValueMap  `json:"key_attribute_one_value_map"`
	ProductImageURL          string                              `json:"product_image_url"`
	RelativeTaskIDS          []string                            `json:"relative_task_ids"`
	ReverseItem              []ReverseItem                       `json:"reverse_item"`
	SaleAttributeOne         SaleAttributeOne                    `json:"sale_attribute_one"`
	SaleAttributeOneValueMap map[string]SaleAttributeOneValueMap `json:"sale_attribute_one_value_map"`
	SpuCode                  string                              `json:"spu_code"`
	TagList                  []TagList                           `json:"tag_list"`
}

type FullProductInfo struct {
	BaseSkuInfos []BaseSkuInfo `json:"base_sku_infos"`
	SkcInfos     []SkcInfo     `json:"skc_infos"`
	SkuInfos     []SkuInfo     `json:"sku_infos"`
	SpuInfo      SpuInfo       `json:"spu_info"`
}

type BaseSkuInfo struct {
	AddSpecsType           int64              `json:"add_specs_type"`
	ArticleNum             string             `json:"article_num"`
	Extra                  BaseSkuInfoExtra   `json:"extra"`
	IsActual               int64              `json:"is_actual"`
	KeyAttributeOneID      string             `json:"key_attribute_one_id"`
	KeyAttributeOneValueID string             `json:"key_attribute_one_value_id"`
	KeyAttributePairs      []KeyAttributePair `json:"key_attribute_pairs"`
	KeyAttributeTwoID      string             `json:"key_attribute_two_id"`
	KeyAttributeTwoValueID string             `json:"key_attribute_two_value_id"`
	SkcCode                string             `json:"skc_code"`
	SkuCode                string             `json:"sku_code"`
	SkuID                  string             `json:"sku_id"`
	SourceID               string             `json:"source_id"`
	SourceType             string             `json:"source_type"`
	SpuCode                string             `json:"spu_code"`
	Status                 int64              `json:"status"`
	StatusV2               int64              `json:"status_v2"`
	StockCount             string             `json:"stock_count"`
	StockMode              string             `json:"stock_mode"`
	SupplierID             string             `json:"supplier_id"`
}

type BaseSkuInfoExtra struct {
	ArticleNum               string `json:"article_num"`
	GsProductAutoFlowProduct string `json:"gs_product_auto_flow_product"`
	GsProductNewFlow         string `json:"gs_product_new_flow"`
	GsProductNewProduct      string `json:"gs_product_new_product"`
	GsSkuStatusFromSkuInfo   string `json:"gs_sku_status_from_sku_info"`
	IsActual                 string `json:"is_actual"`
	MmoGray                  string `json:"mmo_gray"`
	MmoGrayMaterial          string `json:"mmo_gray_material"`
	StockCount               string `json:"stock_count"`
	StockMode                string `json:"stock_mode"`
	SupplyPriceCent          string `json:"supply_price_cent"`
	SupplyPriceCurrencyType  string `json:"supply_price_currency_type"`
}

type KeyAttributePair struct {
	KeyPropertyID       string `json:"key_property_id"`
	KeyPropertyValueID  string `json:"key_property_value_id"`
	TTSAttributeID      string `json:"tts_attribute_id"`
	TTSAttributeValueID string `json:"tts_attribute_value_id"`
}

type SkcInfo struct {
	AddSpecsType           int64                  `json:"add_specs_type"`
	Extra                  ColorBlockClass        `json:"extra"`
	KeyAttributeOne        SkcInfoKeyAttributeOne `json:"key_attribute_one"`
	KeyAttributeOneID      string                 `json:"key_attribute_one_id"`
	KeyAttributeOneValueID string                 `json:"key_attribute_one_value_id"`
	KeyAttributeTwo        KeyAttributeTwo        `json:"key_attribute_two"`
	KeyAttributeTwoID      string                 `json:"key_attribute_two_id"`
	ProductType            int64                  `json:"product_type"`
	PushPhoto              PushPhoto              `json:"push_photo"`
	SiteMainSkc            SiteMainSkc            `json:"site_main_skc"`
	SkcCode                string                 `json:"skc_code"`
	SkcID                  string                 `json:"skc_id"`
	SkcSupplierCode        string                 `json:"skc_supplier_code"`
	SpuCode                string                 `json:"spu_code"`
	SupplierID             string                 `json:"supplier_id"`
}

type ColorBlockClass struct {
}

type SkcInfoKeyAttributeOne struct {
	KeyPropertyID      string `json:"key_property_id"`
	KeyPropertyValueID string `json:"key_property_value_id"`
}

type KeyAttributeTwo struct {
	KeyPropertyID string `json:"key_property_id"`
}

type PushPhoto struct {
	PicSupplyWay int64 `json:"pic_supply_way"`
}

type SiteMainSkc struct {
	ColorBlock          ColorBlockClass `json:"color_block"`
	DevelopTypeID       string          `json:"develop_type_id"`
	DeveloperID         string          `json:"developer_id"`
	DeveloperName       string          `json:"developer_name"`
	DisplayImagesetCode string          `json:"display_imageset_code"`
	FirstKeyProperty    Property        `json:"first_key_property"`
	ImageSetInfo        ColorBlockClass `json:"image_set_info"`
	ImageSets           []interface{}   `json:"image_sets"`
	Images              []Image         `json:"images"`
	Index               int64           `json:"index"`
	Is3D                int64           `json:"is_3d"`
	IsCombo             int64           `json:"is_combo"`
	IsPopular           int64           `json:"is_popular"`
	Models              []interface{}   `json:"models"`
	SkcCode             string          `json:"skc_code"`
	SpuCode             string          `json:"spu_code"`
}

type Property struct {
	InputType      int64             `json:"input_type"`
	PropertyValues []PropertyValue   `json:"property_values"`
	SPropertyCode  SPropertyCodeEnum `json:"s_property_code"`
	SPropertyID    string            `json:"s_property_id"`
	SPropertyName  ItemNameEnEnum    `json:"s_property_name"`
	TTSAttributeID string            `json:"tts_attribute_id"`
}

type PropertyValue struct {
	PlmPropertyValueID string        `json:"plm_property_value_id"`
	SPropertyValueCode string        `json:"s_property_value_code"`
	SPropertyValueID   string        `json:"s_property_value_id"`
	SPropertyValueName ProductNameEn `json:"s_property_value_name"`
	SkcCode            *string       `json:"skc_code,omitempty"`
}

type Image struct {
	Height int64  `json:"height"`
	URI    string `json:"uri"`
	Width  int64  `json:"width"`
}

type SkuInfo struct {
	BaseSkuInfo               BaseSkuInfo                 `json:"base_sku_info"`
	BizSystemStatus           BizSystemStatus             `json:"biz_system_status"`
	BusinessType              int64                       `json:"business_type"`
	ComboList                 []interface{}               `json:"combo_list"`
	ComboMode                 int64                       `json:"combo_mode"`
	ComboType                 int64                       `json:"combo_type"`
	CreateTime                string                      `json:"create_time"`
	CreateUid                 string                      `json:"create_uid"`
	Description               string                      `json:"description"`
	PackageInfo               PackageInfo                 `json:"package_info"`
	ProductSourceType         int64                       `json:"product_source_type"`
	ProductType               int64                       `json:"product_type"`
	SiteMainSku               SiteMainSku                 `json:"site_main_sku"`
	SkuMainStatusTimeInfoList []SkuMainStatusTimeInfoList `json:"skuMainStatusTimeInfoList"`
	UpdateTime                string                      `json:"update_time"`
	UpdateUid                 string                      `json:"update_uid"`
}

type BizSystemStatus struct {
	FirstOrderStatus  string `json:"first_order_status"`
	PriceVerifyStatus string `json:"price_verify_status"`
	SelectPoolStatus  string `json:"select_pool_status"`
}

type PackageInfo struct {
	PackageLongestLength  string `json:"package_longest_length"`
	PackageMiddleLength   string `json:"package_middle_length"`
	PackageShortestLength string `json:"package_shortest_length"`
	PackageWeight         string `json:"package_weight"`
}

type SiteMainSku struct {
	ChildSkuStatusDatas []interface{}  `json:"child_sku_status_datas"`
	ComboSubSkus        []interface{}  `json:"combo_sub_skus"`
	Index               int64          `json:"index"`
	IsCombo             int64          `json:"is_combo"`
	IsVisible           int64          `json:"is_visible"`
	PropertyPairs       []PropertyPair `json:"property_pairs"`
	SkcCode             string         `json:"skc_code"`
	SkuCode             string         `json:"sku_code"`
	SpuCode             string         `json:"spu_code"`
}

type PropertyPair struct {
	SPropertyID      string `json:"s_property_id"`
	SPropertyValueID string `json:"s_property_value_id"`
}

type SkuMainStatusTimeInfoList struct {
	MainStatus int64  `json:"main_status"`
	TimeStamp  string `json:"time_stamp"`
}

type SpuInfo struct {
	BaseSpuInfo                BaseSpuInfo              `json:"base_spu_info"`
	ComboType                  int64                    `json:"combo_type"`
	ManufacturerInfoIDList     []ManufacturerInfoIDList `json:"manufacturer_info_id_list"`
	ProductExtendInfos         []interface{}            `json:"productExtendInfos"`
	ProductType                int64                    `json:"product_type"`
	Properties                 []PropertyElement        `json:"properties"`
	PropertiesV2               []PropertiesV2           `json:"properties_v2"`
	RpIDList                   []RpIDList               `json:"rp_id_list"`
	SecurityWarning            SecurityWarning          `json:"security_warning"`
	SiteProduct                SiteProduct              `json:"site_product"`
	SupplierExcludeRegionInfos []interface{}            `json:"supplier_exclude_region_infos"`
}

type BaseSpuInfo struct {
	BrandID           string           `json:"brand_id"`
	BusinessType      int64            `json:"business_type"`
	CategoryID        string           `json:"category_id"`
	CreateTime        string           `json:"create_time"`
	CreateUid         string           `json:"create_uid"`
	DevelopGroupID    string           `json:"develop_group_id"`
	DevelopTypeID     string           `json:"develop_type_id"`
	Developer         string           `json:"developer"`
	DeveloperID       string           `json:"developer_id"`
	Extra             BaseSpuInfoExtra `json:"extra"`
	ImageSetCode      string           `json:"image_set_code"`
	IsTest            bool             `json:"is_test"`
	KeyAttributeOneID string           `json:"key_attribute_one_id"`
	KeyAttributeTwoID string           `json:"key_attribute_two_id"`
	KeyPropertyIDS    []KeyPropertyID  `json:"key_property_ids"`
	ProductDescEn     string           `json:"product_desc_en"`
	ProductEn         string           `json:"product_en"`
	ProductName       string           `json:"product_name"`
	SpuCode           string           `json:"spu_code"`
	SpuID             string           `json:"spu_id"`
	SupplierID        string           `json:"supplier_id"`
	UpdateTime        string           `json:"update_time"`
	UpdateUid         string           `json:"update_uid"`
}

type BaseSpuInfoExtra struct {
	AlgoGoodsTag             string `json:"algo_goods_tag"`
	GsProductAutoFlowProduct string `json:"gs_product_auto_flow_product"`
	GsProductNewFlow         string `json:"gs_product_new_flow"`
	GsProductNewProduct      string `json:"gs_product_new_product"`
	ProductSource            string `json:"product_source"`
}

type KeyPropertyID struct {
	KeyPropertyID  string `json:"key_property_id"`
	TTSAttributeID string `json:"tts_attribute_id"`
}

type PropertyElement struct {
	PropertyCode      PropertyPropertyCode `json:"property_code"`
	PropertyID        string               `json:"property_id"`
	PropertyValueList []PropertyValueList  `json:"property_value_list"`
}

type PropertyValueList struct {
	PropertyValueCode    PropertyValueListPropertyValueCode `json:"property_value_code"`
	PropertyValueContent string                             `json:"property_value_content"`
	PropertyValueID      string                             `json:"property_value_id"`
}

type PropertiesV2 struct {
	Property PropertyElement `json:"property"`
	Region   Region          `json:"region"`
}

type SecurityWarning struct {
	SecurityWarningLanguages []interface{} `json:"security_warning_languages"`
}

type SiteProduct struct {
	CategoryID                string          `json:"category_id"`
	CreateTimeMS              string          `json:"create_time_ms"`
	Desc                      string          `json:"desc"`
	DisplayImagesetCode       string          `json:"display_imageset_code"`
	DisplayVideosetCode       string          `json:"display_videoset_code"`
	ImageSetMeta              ImageSetMeta    `json:"image_set_meta"`
	Images                    []interface{}   `json:"images"`
	InstructionFile           InstructionFile `json:"instruction_file"`
	InstructionVideo          Video           `json:"instruction_video"`
	IsDescSupplementaryImages int64           `json:"is_desc_supplementary_images"`
	IsSpuImageset             bool            `json:"is_spu_imageset"`
	OecSellerCancelReasonList []interface{}   `json:"oec_seller_cancel_reason_list"`
	OecSellerIDS              []string        `json:"oec_seller_ids"`
	RpID                      string          `json:"rp_id"`
	SaleProperties            []Property      `json:"sale_properties"`
	SkuList                   []interface{}   `json:"sku_list"`
	SpuCode                   string          `json:"spu_code"`
	SpuImageSets              []interface{}   `json:"spu_image_sets"`
	SupplierTitle             string          `json:"supplier_title"`
	SupplierTitleEn           string          `json:"supplier_title_en"`
	Title                     string          `json:"title"`
	UpdateTimeMS              string          `json:"update_time_ms"`
	Video                     Video           `json:"video"`
	VideoSets                 []interface{}   `json:"video_sets"`
}

type ImageSetMeta struct {
	ImagesetCodes []string `json:"imageset_codes"`
}

type InstructionFile struct {
	FileName  string        `json:"file_name"`
	FileSize  string        `json:"file_size"`
	FileType  string        `json:"file_type"`
	Languages []interface{} `json:"languages"`
	OID       string        `json:"o_id"`
}

type Video struct {
	CoverURI string `json:"cover_uri"`
	CoverURL string `json:"cover_url"`
	Duration int64  `json:"duration"`
	Vid      string `json:"vid"`
}

type ProductReverseItemKeyAttributeOne struct {
	CreateTime    string            `json:"create_time"`
	ItemNameCN    ItemNameCNEnum    `json:"item_name_cn"`
	ItemNameEn    ItemNameEnEnum    `json:"item_name_en"`
	ModifyTime    string            `json:"modify_time"`
	ProductNameCN ItemNameCNEnum    `json:"product_name_cn"`
	ProductNameEn ItemNameEnEnum    `json:"product_name_en"`
	PropertyCode  SPropertyCodeEnum `json:"property_code"`
	PropertyID    string            `json:"property_id"`
	PropertyType  int64             `json:"property_type"`
}

type KeyAttributeOneValueMap struct {
	CreateTime        string                                   `json:"create_time"`
	IsPrivateValue    bool                                     `json:"is_private_value"`
	ModifyTime        string                                   `json:"modify_time"`
	ProductNameCN     KeyAttributeOneValueMapProductNameCN     `json:"product_name_cn"`
	ProductNameEn     ProductNameEn                            `json:"product_name_en"`
	PropertyCode      string                                   `json:"property_code"`
	PropertyValueCode KeyAttributeOneValueMapPropertyValueCode `json:"property_value_code"`
	PropertyValueID   string                                   `json:"property_value_id"`
}

type ReverseItem struct {
	CreateTime     string        `json:"create_time"`
	Deadline       string        `json:"deadline"`
	FinishTime     string        `json:"finish_time"`
	ImageUrls      []string      `json:"image_urls"`
	ReverseProcess int64         `json:"reverse_process"`
	ReverseReason  ReverseReason `json:"reverse_reason"`
	TaskID         string        `json:"task_id"`
	Type           int64         `json:"type"`
}

type SaleAttributeOne struct {
	CreateTime    string            `json:"create_time"`
	ModifyTime    string            `json:"modify_time"`
	NameCN        ItemNameCNEnum    `json:"name_cn"`
	NameEn        ItemNameEnEnum    `json:"name_en"`
	PropertyCode  SPropertyCodeEnum `json:"property_code"`
	PropertyID    string            `json:"property_id"`
	PropertyType  int64             `json:"property_type"`
	TTSPropertyID string            `json:"tts_property_id"`
}

type SaleAttributeOneValueMap struct {
	CreateTime         string                                   `json:"create_time"`
	IsPrivateValue     bool                                     `json:"is_private_value"`
	ModifyTime         string                                   `json:"modify_time"`
	NameCN             KeyAttributeOneValueMapProductNameCN     `json:"name_cn"`
	NameEn             ProductNameEn                            `json:"name_en"`
	PropertyValueCode  KeyAttributeOneValueMapPropertyValueCode `json:"property_value_code"`
	PropertyValueID    string                                   `json:"property_value_id"`
	TTSPropertyValueID string                                   `json:"tts_property_value_id"`
}

type TagList struct {
	TagCode TagCode `json:"tag_code"`
	Text    string  `json:"text"`
}

type ProductNameEn string

const (
	Multicolor ProductNameEn = "Multicolor"
	OneSize    ProductNameEn = "one-size"
)

type SPropertyCodeEnum string

const (
	Colour       SPropertyCodeEnum = "COLOUR"
	PropertyCode SPropertyCodeEnum = ""
	Size         SPropertyCodeEnum = "SIZE"
)

type ItemNameEnEnum string

const (
	Color          ItemNameEnEnum = "Color"
	ItemNameEn     ItemNameEnEnum = ""
	ItemNameEnSize ItemNameEnEnum = "Size"
)

type ManufacturerInfoIDList string

const (
	The67D28Dc5E61B905Cdd399839 ManufacturerInfoIDList = "67d28dc5e61b905cdd399839"
)

type PropertyPropertyCode string

const (
	CAProp65_CA                        PropertyPropertyCode = "CA_PROP_65_CA"
	CAProp65_Re                        PropertyPropertyCode = "CA_PROP_65_RE"
	ConfirmAnyRegulatoryMarkingOrLabel PropertyPropertyCode = "CONFIRM_ANY_REGULATORY_MARKING_OR_LABEL"
	FdaProductCode                     PropertyPropertyCode = "FDA_PRODUCT_CODE"
	Material                           PropertyPropertyCode = "MATERIAL"
	PlugType                           PropertyPropertyCode = "PLUG_TYPE"
	ProductContainWoodenMaterials      PropertyPropertyCode = "PRODUCT_CONTAIN_WOODEN_MATERIALS"
	Quantity                           PropertyPropertyCode = "QUANTITY"
	RegionOfOrigin                     PropertyPropertyCode = "REGION_OF_ORIGIN"
	SensitiveGoodsType                 PropertyPropertyCode = "SENSITIVE_GOODS_TYPE"
)

type PropertyValueListPropertyValueCode string

const (
	CANo                                  PropertyValueListPropertyValueCode = "CA_NO"
	ConfirmAnyRegulatoryMarkingOrLabelYes PropertyValueListPropertyValueCode = "CONFIRM_ANY_REGULATORY_MARKING_OR_LABEL_YES"
	MaterialCeramic                       PropertyValueListPropertyValueCode = "MATERIAL_CERAMIC"
	NoPlug                                PropertyValueListPropertyValueCode = "NO_PLUG"
	NonSensitiveGood                      PropertyValueListPropertyValueCode = "NON_SENSITIVE_GOOD"
	ReNo                                  PropertyValueListPropertyValueCode = "RE_NO"
	RegionOfOriginChina                   PropertyValueListPropertyValueCode = "REGION_OF_ORIGIN_CHINA"
	The1_Piece                            PropertyValueListPropertyValueCode = "1_PIECE"
	The7517194954022094599                PropertyValueListPropertyValueCode = "7517194954022094599"
	Wsdfghngfdsxzasedr                    PropertyValueListPropertyValueCode = "WSDFGHNGFDSXZASEDR"
)

type Region string

const (
	De Region = "DE"
	Es Region = "ES"
	Fr Region = "FR"
	GB Region = "GB"
	It Region = "IT"
	Jp Region = "JP"
	MX Region = "MX"
	Sa Region = "SA"
	Us Region = "US"
)

type RpIDList string

const (
	The67Ca66Bf396Fd5D6B80309Cc RpIDList = "67ca66bf396fd5d6b80309cc"
)

type ItemNameCNEnum string

const (
	尺码 ItemNameCNEnum = "尺码"
)

type KeyAttributeOneValueMapProductNameCN string

const (
	均码 KeyAttributeOneValueMapProductNameCN = "均码"
)

type KeyAttributeOneValueMapPropertyValueCode string

const (
	PropertyValueCodeONESIZE KeyAttributeOneValueMapPropertyValueCode = "ONE-SIZE"
)

type ReverseReason string

const (
	尺寸图可能不符规范尺寸图可能不符规范 ReverseReason = "尺寸图可能不符规范; 尺寸图可能不符规范"
)

type TagCode string

const (
	AutoProduct                   TagCode = "AUTO_PRODUCT"
	AveSupplyPriceThan8           TagCode = "AVE_SUPPLY_PRICE_THAN_8"
	RealtimeAmazon70Supplyprice   TagCode = "REALTIME_AMAZON70_SUPPLYPRICE"
	RealtimeHotstockSearchProduct TagCode = "REALTIME_HOTSTOCK_SEARCH_PRODUCT"
)
