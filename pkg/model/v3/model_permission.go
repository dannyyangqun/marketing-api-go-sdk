/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 数据源权限
type Permission struct {
	CanCreateAudience       *bool `json:"can_create_audience,omitempty"`
	CanExactConversionClaim *bool `json:"can_exact_conversion_claim,omitempty"`
	CanAsServing            *bool `json:"can_as_serving,omitempty"`
}
