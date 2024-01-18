/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 关键词信息
type BidwordDataStructs struct {
	Bidword         *string  `json:"bidword,omitempty"`
	MonthQueryCount *int64   `json:"month_query_count,omitempty"`
	ClickCount      *int64   `json:"click_count,omitempty"`
	Price           *float64 `json:"price,omitempty"`
	TotalAccts      *int64   `json:"total_accts,omitempty"`
	RecommendReason *string  `json:"recommend_reason,omitempty"`
}