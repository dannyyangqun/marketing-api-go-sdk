/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 文生图参数信息
type Text2imgStruct struct {
	Prompt        *string             `json:"prompt,omitempty"`
	DimensionSize MuseAiDimensionSize `json:"dimension_size,omitempty"`
}