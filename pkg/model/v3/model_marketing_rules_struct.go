/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 过滤条件
type MarketingRulesStruct struct {
	MarketingGoal        *string `json:"marketing_goal,omitempty"`
	MarketingSubGoal     *string `json:"marketing_sub_goal,omitempty"`
	MarketingTargetType  *string `json:"marketing_target_type,omitempty"`
	MarketingCarrierType *string `json:"marketing_carrier_type,omitempty"`
}