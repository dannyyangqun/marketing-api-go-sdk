/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告组信息最后更新时间范围，一次最多允许获取7天的数据，最远可获取2018.01.01的数据
type LastModifiedTimeRange struct {
	BeginTime int64 `json:"begin_time,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
}