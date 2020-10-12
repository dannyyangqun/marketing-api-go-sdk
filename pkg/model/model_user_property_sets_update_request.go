/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type UserPropertySetsUpdateRequest struct {
	AccountId         int64  `json:"account_id,omitempty"`
	UserPropertySetId int64  `json:"user_property_set_id,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
}