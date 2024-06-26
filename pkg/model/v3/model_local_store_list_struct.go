/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 门店列表项
type LocalStoreListStruct struct {
	PoiId          *string `json:"poi_id,omitempty"`
	LocalStoreName *string `json:"local_store_name,omitempty"`
	OwnerAccountId *int64  `json:"owner_account_id,omitempty"`
}
