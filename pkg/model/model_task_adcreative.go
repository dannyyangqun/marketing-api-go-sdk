/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告创意
type TaskAdcreative struct {
	PageType DestinationType `json:"page_type,omitempty"`
	PageSpec *TaskPageSpec   `json:"page_spec,omitempty"`
}