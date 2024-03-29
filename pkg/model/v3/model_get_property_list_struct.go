/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回信息结构
type GetPropertyListStruct struct {
	Properties      *[]DetailPropertyStruct `json:"properties,omitempty"`
	ExtraProperties *[]ExtraProperty        `json:"extra_properties,omitempty"`
}
