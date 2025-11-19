import fs from "fs";
import axios from "axios";
import JSONBig from "json-bigint";

axios.defaults.transformResponse = [
  function (data) {
    try {
      // 使用 json-bigint 解析原始的响应数据（字符串）
      // { storeAsString: true } 选项会将大数自动转换为字符串
      return JSONBig.parse(data);
    } catch (err) {
      // 如果解析失败，降级为普通 JSON.parse 或返回原始数据
      console.error("JSONBig parse error:", err);
      return JSON.parse(data); // 或者直接 return data;
    }
  },
];

// 全部需要 sessionid

// 获取问题列表
// const data = await axios.request({
//   url: "https://api16-normal-sg.tiktokshopglobalselling.com/api/full-service/product-center/reverse/list",
//   method: "POST",
//   headers: {
//     cookie: "sessionid=bfa765bd3284cededa8eed1da8ad5ea3",
//   },
//   data: {
//     filter: { article_number: "MKB", reverse_status: 10 },
//     page_info: { page_no: 5, page_size: 10 },
//   },
// });
// fs.writeFileSync("productList.json", JSON.stringify(data.data, null, 2));

// 获取商品详情
// const data = await axios.request({
//   url: "https://api16-normal-sg.tiktokshopglobalselling.com/api/full-service/product-center/reverse/get_detail",
//   method: "POST",
//   headers: {
//     cookie: "sessionid=bfa765bd3284cededa8eed1da8ad5ea3",
//   },
//   data: {
//     spu_code: "S251118011105",
//     reverse_status: 10,
//   },
// });
// fs.writeFileSync("productDesc.json", JSON.stringify(data.data, null, 2));

// 获取商品的类目属性关系 （补充properties_v2属性）
// const data = await axios.request({
//   url: "https://api16-normal-sg.tiktokshopglobalselling.com/api/full-service/product-center/category/m_get_category_prop_relation",
//   method: "POST",
//   headers: {
//     cookie: "sessionid=bfa765bd3284cededa8eed1da8ad5ea3",
//   },
//   data: {
//     category_ids: ["7382947714688599815"],
//     region_list: ["SA", "GB", "US", "FR", "DE", "IT", "ES", "MX", "JP"],
//   },
// });
// fs.writeFileSync("categoryPropRelation.json", JSON.stringify(data.data, null, 2));

let productDesc = JSON.parse(fs.readFileSync("./productDesc.json"));
productDesc = productDesc.info.spu_detail;

let categoryPropRelation = JSON.parse(fs.readFileSync("./categoryPropRelation.json"));
categoryPropRelation = categoryPropRelation.id_relation_map[productDesc.category_id];

const propertiesV2 = productDesc.properties_v2.map((item) => {
  let propertyRelation = categoryPropRelation.prop_list.find((el) => el.property_id === item.property.property_id);

  item.property.property_value_list = item.property.property_value_list.map((itemel) => {
    let propertyValueRelation = categoryPropRelation.prop_value_list.find((el) => el.property_value_id === itemel.property_value_id);
    const data = { ...itemel };
    if (propertyValueRelation) data.tts_attribute_value_id = String(propertyValueRelation.tts_property_value_id);
    if (data.tts_attribute_value_id === "0") delete data.tts_attribute_value_id;
    delete data.property_value_content;
    return data;
  });

  item.property.tts_attribute_id = String(propertyRelation.tts_property_id);
  return { ...item };
});

const salePropertyIdList = productDesc.sale_property_list.map((item) => {
  return {
    property_id: item.property.property_id,
    tts_property_id: String(item.property.tts_property_id),
  };
});

const mediaInfo = {
  picture_list: [],
  pic_type: 2 || productDesc.product_media_info.pic_set.set_status, // TODO pic_type 收集不到此值，存疑
  pic_set_type: productDesc.product_media_info.pic_set.set_type,
};

mediaInfo.picture_list = productDesc.product_media_info.pic_set.spu_material.map((item) => {
  for (const key in item.material) {
    if (key === "extra") {
      item.material[key].height = String(item.material[key].height);
      item.material[key].size = String(item.material[key].size);
      item.material[key].target_height = String(item.material[key].target_height);
      item.material[key].target_width = String(item.material[key].target_width);
      item.material[key].v_duration = String(item.material[key].v_duration);
      item.material[key].width = String(item.material[key].width);
    }
    if (key === "id") item.material[key] = String(item.material[key]);
    if (key === "parent_id") item.material[key] = String(item.material[key]);
    for (const el of item.material.recognition_res) {
      el.pic_rec_id = String(el.pic_rec_id);
      el.rec_time_ms = String(el.rec_time_ms);
    }
    if (key === "shop_id") item.material[key] = String(item.material[key]);
    if (key === "size") item.material[key] = String(item.material[key]);
  }

  return {
    id: String(item.id),
    link_type: item.link_type,
    material_show_type: item.material_show_type,
    material_use_type_code: item.material_use_type_code,
    order_num: String(item.order_num),
    material: item.material,
  };
});

const skuList = [];

for (const skc in productDesc.product_media_info.pic_set.extra.skc_sku_mapping) {
  skuList.push(...productDesc.product_media_info.pic_set.extra.skc_sku_mapping[skc]);
}

let skuMediaInfo = productDesc.product_media_info.pic_set.object_material[skuList[0]];

skuMediaInfo = skuMediaInfo.map((item) => {
  for (const key in item.material) {
    if (key === "extra") {
      item.material[key].height = String(item.material[key].height);
      item.material[key].size = String(item.material[key].size);
      item.material[key].target_height = String(item.material[key].target_height);
      item.material[key].target_width = String(item.material[key].target_width);
      item.material[key].v_duration = String(item.material[key].v_duration);
      item.material[key].width = String(item.material[key].width);
    }
    if (key === "id") item.material[key] = String(item.material[key]);
    if (key === "parent_id") item.material[key] = String(item.material[key]);
    for (const el of item.material.recognition_res) {
      el.pic_rec_id = String(el.pic_rec_id);
      el.rec_time_ms = String(el.rec_time_ms);
    }
    if (key === "shop_id") item.material[key] = String(item.material[key]);
    if (key === "size") item.material[key] = String(item.material[key]);
  }

  return {
    id: String(item.id),
    material_show_type: item.material_show_type,
    link_type: item.link_type,
    order_num: String(item.order_num),
    material: item.material,
  };
});

const skcDetails = productDesc.skc_details.map((item) => {
  let data = {};
  data.index = String(item.index);
  data.skc_code = item.skc_code;
  data.sale_property = {
    property_value_id: item.sale_property_value_info.property_value_id,
    tts_property_value_id: String(item.sale_property_value_info.tts_property_value_id),
  };
  data.media_info = {
    picture_list: item.picture_urls,
    pic_type: 2, // TODO pic_type 收集不到此值，存疑
  };
  data.sku_details = item.sku_details.map((elItem) => {
    return {
      sku_code: elItem.sku_code,
      media_info: {
        picture_list: skuMediaInfo,
        pic_type: 2, // TODO pic_type 收集不到此值，存疑
      },
      sale_property_list: elItem.sale_property_list.map((item) => ({
        property_value_id: item.property_value_id,
        tts_property_value_id: String(item.tts_property_value_id),
      })),
      package_longest_length: String(elItem.package_longest_length),
      package_shortest_length: String(elItem.package_shortest_length),
      package_middle_length: String(elItem.package_middle_length),
      package_weight: String(elItem.package_weight),
      article_number: elItem.article_number,
      price: String(elItem.price),
      product_status: true, // TODO product_status 收集不到此值，存疑
      stock: String(elItem.stock),
      supply_price_currency_type: elItem.price_currency_type,
      goods_in_stock: true, // TODO goods_in_stock 收集不到此值，存疑
      stock_mode: elItem.stock_mode,
    };
  });
  data.article_number = item.article_number;
  data.stock_mode = item.sku_details[0].stock_mode;
  return data;
});

const salePropertyValueList = productDesc.sale_property_list.map((item) => {
  return item.property_values.map((el) => ({ plm_property_value_id: el.property_value_id, plm_tts_property_value_id: String(el.tts_property_value_id) }));
});

let checkProductParams = {
  product_name: productDesc.product_name,
  product_name_en: productDesc.product_name_en,
  category_id: productDesc.category_id,
  brand_id: null,
  properties_v2: propertiesV2,
  security_warning_info: productDesc.security_warning_info,
  sale_property_id_list: salePropertyIdList,
  video_list: [],
  media_info: mediaInfo,
  grading: {},
  product_desc_en: productDesc.product_desc_en,
  certifications: [],
  exclude_region_codes: productDesc.exclude_region_codes,
  manufacturer_ids: productDesc.manufacturer_infos.map((item) => item.id),
  rp_ids: productDesc.rp_models.map((item) => item.id),
  skc_details: skcDetails,
  sale_property_value_list: salePropertyValueList,
  ticket_code: productDesc.ticket_code,
  spu_code: productDesc.spu_code,
};

fs.writeFileSync("./test.json", JSON.stringify(checkProductParams, null, 2));

// 检查产品请求（图片异常）
// const data = await axios.request({
//   url: "https://api16-normal-sg.tiktokshopglobalselling.com/api/full-service/product-center/check/check_product",
//   method: "POST",
//   headers: {
//     cookie: "sessionid=bfa765bd3284cededa8eed1da8ad5ea3",
//   },
//   data: {
//     check_option: {
//       check_price: false,
//       check_certification: false,
//       check_package: false,
//       check_pic: true,
//       check_product_desc_pic: false,
//     },
//     product_info: checkProductParams,
//   },
// });

// fs.writeFileSync("./checkProductResult.json", JSON.stringify(data.data, null, 2));

let checkProductResult = JSON.parse(fs.readFileSync("./checkProductResult.json"));

checkProductResult = checkProductResult.picture_check_result;

const picIssues = {};
const issuesKeys = [];
for (const picUrl in checkProductResult.uri_to_check_result_map) {
  checkProductResult.uri_to_check_result_map[picUrl].recognition_result_items.map((item) => {
    if (item.actions && Array.isArray(item.actions) && item.actions.length !== 0) {
      if (picIssues[picUrl] && Array.isArray(picIssues[picUrl])) picIssues[picUrl].push(...item.actions);
      else {
        if (!issuesKeys.includes(picUrl.split("/")[1])) {
          picIssues[picUrl] = JSON.parse(JSON.stringify(item.actions));
          issuesKeys.push(picUrl.split("/")[1]);
        }
      }
    }
  });
  if (picIssues[picUrl] && Array.isArray(picIssues[picUrl])) picIssues[picUrl] = picIssues[picUrl].map(String);
}

const appealPictures = checkProductResult.check_result_map
  .map((item) => {
    const issuesArr = item.recognition_result_items.reduce((last, cur) => {
      if (cur.actions && Array.isArray(cur.actions) && cur.actions.length !== 0) {
        last.push(...cur.actions);
      }
      return last;
    }, []);
    if (issuesArr.length !== 0) return { uri: item.picture_uri, issues: issuesArr.map(String), pic_type: 1 };
    else return null;
  })
  .filter(Boolean);

const productInfo = JSON.parse(fs.readFileSync("./productDesc.json"));

const appealSceneParams = {
  feedback_content: {},
  relative_task_ids: productInfo.relative_task_ids,
};

const appealOrderParams = {
  spu_detail: checkProductParams,
  scene: 3,
  pic_issues: picIssues,
  appeal_pictures: appealPictures,
  appeal_scene_params: appealSceneParams,
};

fs.writeFileSync("./createAppeal.json", JSON.stringify(appealOrderParams, null, 2));

// 发起申诉请求
// const data = await axios.request({
//   url: "https://api16-normal-sg.tiktokshopglobalselling.com/api/full-service/product-center/appeal_order/create",
//   method: "POST",
//   headers: {
//     cookie: "sessionid=bfa765bd3284cededa8eed1da8ad5ea3",
//   },
//   data: appealOrderParams
// });

console.log(data.data);
