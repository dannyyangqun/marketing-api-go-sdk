/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatPagesCsgroupStatusUpdateRequest struct {
	AccountId *int64  `json:"account_id,omitempty"`
	CorpId    *string `json:"corp_id,omitempty"`
	Userid    *string `json:"userid,omitempty"`
	Status    *int64  `json:"status,omitempty"`
}
