/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Aggregation : 聚合维度，是否将结果按照指定类型细分，可选值'DOMAIN', 'ACTION_TYPE'
type Aggregation string

// List of Aggregation
const (
	Aggregation_DOMAIN      Aggregation = "DOMAIN"
	Aggregation_ACTION_TYPE Aggregation = "ACTION_TYPE"
)