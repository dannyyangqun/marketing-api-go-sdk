/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 拒绝原因明细
type RejectReasonDetailStruct struct {
	RejectReasonId      *string                              `json:"reject_reason_id,omitempty"`
	RejectReasonContent *string                              `json:"reject_reason_content,omitempty"`
	CaseDoc             *string                              `json:"case_doc,omitempty"`
	CaseContent         *string                              `json:"case_content,omitempty"`
	RejectInfoLocations *[]PrereviewRejectInfoLocationStruct `json:"reject_info_locations,omitempty"`
}
