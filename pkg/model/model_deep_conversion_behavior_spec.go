/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// oCPC/oCPM 优化转化行为配置
type DeepConversionBehaviorSpec struct {
	Goal      DeepConversionBehaviorGoal `json:"goal,omitempty"`
	BidAmount int64                      `json:"bid_amount,omitempty"`
}