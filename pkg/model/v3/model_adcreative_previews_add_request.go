/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdcreativePreviewsAddRequest struct {
	AccountId  *int64       `json:"account_id,omitempty"`
	AdgroupId  *int64       `json:"adgroup_id,omitempty"`
	UserIdType ViewerIdType `json:"user_id_type,omitempty"`
	UserIdList *[]string    `json:"user_id_list,omitempty"`
}