/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 查询定向标签的条件
type QuerySpec struct {
	Query                  *string                        `json:"query,omitempty"`
	QueryList              *[]string                      `json:"query_list,omitempty"`
	MaxResultNumber        *int64                         `json:"max_result_number,omitempty"`
	ExcludingTargetingTags *[]string                      `json:"excluding_targeting_tags,omitempty"`
	AdvancedRecommendType  TargetingAdvancedRecommendType `json:"advanced_recommend_type,omitempty"`
}