/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// oCPC/oCPM投放能力
type WechatOcpa struct {
	PromotedObjectType   PromotedObjectType `json:"promoted_object_type,omitempty"`
	PromotedObjectId     string             `json:"promoted_object_id,omitempty"`
	AdcreativeTemplateId int64              `json:"adcreative_template_id,omitempty"`
}