/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件信息
type ComponentInfo struct {
	ComponentId   *int64             `json:"component_id,omitempty"`
	ComponentType ComponentType      `json:"component_type,omitempty"`
	ReviewStatus  ReviewResultStatus `json:"review_status,omitempty"`
}
