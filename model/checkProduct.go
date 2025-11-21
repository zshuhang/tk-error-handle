package model

// 请求数据
type CheckProductRequest struct {
	CheckOption CheckOption `json:"check_option"`
	ProductInfo ProductInfo `json:"product_info"`
}

type CheckOption struct {
	CheckPrice          bool `json:"check_price"`
	CheckCertification  bool `json:"check_certification"`
	CheckPackage        bool `json:"check_package"`
	CheckPic            bool `json:"check_pic"`
	CheckProductDescPic bool `json:"check_product_desc_pic"`
}

type ProductInfo struct {
	ProductName           string                          `json:"product_name"`
	ProductNameEn         string                          `json:"product_name_en"`
	CategoryID            string                          `json:"category_id"`
	BrandID               any                             `json:"brand_id"`
	PropertiesV2          []CheckProductPropertiesV2      `json:"properties_v2"`
	SecurityWarningInfo   CheckProductSecurityWarningInfo `json:"security_warning_info"`
	SalePropertyIDList    []SalePropertyIDList            `json:"sale_property_id_list"`
	VideoList             []any                           `json:"video_list"`
	MediaInfo             CheckProductMediaInfo           `json:"media_info"`
	Grading               any                             `json:"grading"`
	ProductDescEn         string                          `json:"product_desc_en"`
	Certifications        []any                           `json:"certifications"`
	ExcludeRegionCodes    []any                           `json:"exclude_region_codes"`
	ManufacturerIDS       []string                        `json:"manufacturer_ids"`
	RpIDS                 []string                        `json:"rp_ids"`
	SkcDetails            []CheckProductSkcDetail         `json:"skc_details"`
	SalePropertyValueList [][]SalePropertyValueList       `json:"sale_property_value_list"`
	TicketCode            string                          `json:"ticket_code"`
	SpuCode               string                          `json:"spu_code"`
}

type CheckProductPropertiesV2 struct {
	Region   string               `json:"region"`
	Property CheckProductProperty `json:"property"`
}

type CheckProductProperty struct {
	PropertyValueList []CheckProductPropertyValue `json:"property_value_list"`
	PropertyID        string                      `json:"property_id"`
	PropertyCode      string                      `json:"property_code"`
	TTSAttributeID    string                      `json:"tts_attribute_id"`
}

type CheckProductPropertyValue struct {
	PropertyValueID     string `json:"property_value_id"`
	PropertyValueCode   string `json:"property_value_code"`
	PropertyValueCN     string `json:"property_value_cn"`
	PropertyValueEn     string `json:"property_value_en"`
	TTSAttributeValueID string `json:"tts_attribute_value_id,omitempty"`
}

type CheckProductSecurityWarningInfo struct {
	SecurityWarningLanguages []any `json:"security_warning_languages"`
}

type SalePropertyIDList struct {
	PropertyID    string `json:"property_id"`
	TTSPropertyID string `json:"tts_property_id"`
}

type CheckProductMediaInfo struct {
	PictureList []CheckProductPicture `json:"picture_list"`
	PicType     int64                 `json:"pic_type"`
	PicSetType  int64                 `json:"pic_set_type"`
}

type CheckProductPicture struct {
	ID                  string               `json:"id"`
	LinkType            int64                `json:"link_type"`
	MaterialShowType    int64                `json:"material_show_type"`
	MaterialUseTypeCode *string              `json:"material_use_type_code,omitempty"`
	OrderNum            string               `json:"order_num"`
	Material            CheckProductMaterial `json:"material"`
}

type CheckProductMaterial struct {
	Extra          CheckProductMaterialExtra `json:"extra"`
	ID             string                    `json:"id"`
	MaterialStatus int64                     `json:"material_status"`
	MaterialType   int64                     `json:"material_type"`
	Name           string                    `json:"name"`
	ParentID       string                    `json:"parent_id"`
	RecognitionRes []CheckProductRecognition `json:"recognition_res"`
	SellerID       string                    `json:"seller_id"`
	ShopID         string                    `json:"shop_id"`
	Size           string                    `json:"size"`
	URI            string                    `json:"uri"`
	URLMap         map[string]string         `json:"urlMap"`
	Vid            string                    `json:"vid"`
}

type CheckProductMaterialExtra struct {
	Format       string `json:"format"`
	Height       string `json:"height"`
	Name         string `json:"name"`
	Resolution   string `json:"resolution"`
	Size         string `json:"size"`
	TargetHeight string `json:"target_height"`
	TargetWidth  string `json:"target_width"`
	URIVa        string `json:"uri_va"`
	VDuration    string `json:"v_duration"`
	Width        string `json:"width"`
}

type CheckProductRecognition struct {
	Actions              []int64 `json:"actions"`
	PicRecID             string  `json:"pic_rec_id"`
	RecTimeMS            string  `json:"rec_time_ms"`
	RecognitionAlgorithm int64   `json:"recognitionAlgorithm"`
	Score                int64   `json:"score"`
	Status               int64   `json:"status"`
	Type                 int64   `json:"type"`
}

type CheckProductSkcDetail struct {
	Index         string                         `json:"index"`
	SkcCode       string                         `json:"skc_code"`
	SaleProperty  CheckProductSaleProperty       `json:"sale_property"`
	MediaInfo     CheckProductSkcDetailMediaInfo `json:"media_info"`
	SkuDetails    []CheckProductSkuDetail        `json:"sku_details"`
	ArticleNumber string                         `json:"article_number"`
	StockMode     int64                          `json:"stock_mode"`
}

type CheckProductSaleProperty struct {
	PropertyValueID    string  `json:"property_value_id"`
	TTSPropertyValueID *string `json:"tts_property_value_id,omitempty"`
}

type CheckProductSkcDetailMediaInfo struct {
	PictureList []CheckProductPicture `json:"picture_list"`
	PicType     int64                 `json:"pic_type"`
}

type CheckProductSkuDetail struct {
	SkuCode                 string                         `json:"sku_code"`
	MediaInfo               CheckProductSkcDetailMediaInfo `json:"media_info"`
	SalePropertyList        []CheckProductSaleProperty     `json:"sale_property_list"`
	PackageLongestLength    string                         `json:"package_longest_length"`
	PackageShortestLength   string                         `json:"package_shortest_length"`
	PackageMiddleLength     string                         `json:"package_middle_length"`
	PackageWeight           string                         `json:"package_weight"`
	ArticleNumber           string                         `json:"article_number"`
	Price                   string                         `json:"price"`
	ProductStatus           bool                           `json:"product_status"`
	Stock                   string                         `json:"stock"`
	SupplyPriceCurrencyType int64                          `json:"supply_price_currency_type"`
	GoodsInStock            bool                           `json:"goods_in_stock"`
	StockMode               int64                          `json:"stock_mode"`
}

type SalePropertyValueList struct {
	PlmPropertyValueID     string  `json:"plm_property_value_id"`
	PlmTTSAttributeValueID *string `json:"plm_tts_attribute_value_id,omitempty"`
}

// 响应数据
type CheckProductResponse struct {
	BaseResp           BaseResp           `json:"base_resp"`
	PictureCheckResult PictureCheckResult `json:"picture_check_result"`
}

type 	PictureCheckResult struct {
	CheckResultMap      []CheckResult               `json:"check_result_map"`
	UriToCheckResultMap map[string]UriToCheckResult `json:"uri_to_check_result_map"`
}

type CheckResult struct {
	CheckType              int64                   `json:"check_type"`
	PictureUri             string                  `json:"picture_uri"`
	PictureUrl             string                  `json:"picture_url"`
	RecognitionResultItems []RecognitionResultItem `json:"recognition_result_items"`
}

type RecognitionResultItem struct {
	Actions              []int64 `json:"actions"`
	PicRecId             int64  `json:"pic_rec_id"`
	RecTimeMs            int64  `json:"rec_time_ms"`
	RecognitionAlgorithm int64   `json:"recognitionAlgorithm"`
	Score                int64   `json:"score"`
	Status               int64   `json:"status"`
	Type                 int64   `json:"type"`
	Value                string  `json:"value"`
}

type UriToCheckResult struct {
	CheckType              int64                   `json:"check_type"`
	PictureTags            []PictureTag            `json:"picture_tags"`
	PictureUri             string                  `json:"picture_uri"`
	PictureUrl             string                  `json:"picture_url"`
	RecognitionResultItems []RecognitionResultItem `json:"recognition_result_items"`
}

type PictureTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
