/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组屏蔽优量汇流量包信息
type UpdateExcludeUnionPositionPackageItem struct {
	AdgroupId                   *int64   `json:"adgroup_id,omitempty"`
	ExcludeUnionPositionPackage *[]int64 `json:"exclude_union_position_package,omitempty"`
}
