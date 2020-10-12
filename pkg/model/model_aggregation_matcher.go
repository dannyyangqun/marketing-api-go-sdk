/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 匹配规则组
type AggregationMatcher struct {
	AggregationType    AggregationType `json:"aggregation_type,omitempty"`
	CountType          CountType       `json:"count_type,omitempty"`
	ParamName          string          `json:"param_name,omitempty"`
	Comparator         Comparator      `json:"comparator,omitempty"`
	ComparisonValue    int64           `json:"comparison_value,omitempty"`
	ComparisonMinValue int64           `json:"comparison_min_value,omitempty"`
	ComparisonMaxValue int64           `json:"comparison_max_value,omitempty"`
}