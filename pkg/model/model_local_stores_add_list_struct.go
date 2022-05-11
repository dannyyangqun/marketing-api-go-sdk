/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type LocalStoresAddListStruct struct {
	PoiId                   *string                     `json:"poi_id,omitempty"`
	LocalStoreName          *string                     `json:"local_store_name,omitempty"`
	LocalStoreProvince      *string                     `json:"local_store_province,omitempty"`
	LocalStoreCity          *string                     `json:"local_store_city,omitempty"`
	LocalStoreAddress       *string                     `json:"local_store_address,omitempty"`
	LocalStoreBizInfo       *LocalStoreBizInfoStructRsp `json:"local_store_biz_info,omitempty"`
	WechatEcosystemAccounts *WechatEcosystemAccounts    `json:"wechat_ecosystem_accounts,omitempty"`
	CreatedTime             *int64                      `json:"created_time,omitempty"`
	LastModifiedTime        *int64                      `json:"last_modified_time,omitempty"`
	SystemStatus            AdStatus                    `json:"system_status,omitempty"`
	LocalStoreStreet        *string                     `json:"local_store_street,omitempty"`
	LocalStoreBusinessArea  *string                     `json:"local_store_business_area,omitempty"`
	LocalStoreDistrict      *string                     `json:"local_store_district,omitempty"`
	LocalStoreLocation      *LocalStoreLocation         `json:"local_store_location,omitempty"`
}
