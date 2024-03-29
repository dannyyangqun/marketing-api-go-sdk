/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type QualificationsDeleteRequest struct {
	AccountId           *int64            `json:"account_id,omitempty"`
	QualificationType   QualificationType `json:"qualification_type,omitempty"`
	QualificationId     *int64            `json:"qualification_id,omitempty"`
	QualificationIdList *[]int64          `json:"qualification_id_list,omitempty"`
}
