/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type WechatDailyCostGetListStruct struct {
	AccountId       int64  `json:"account_id,omitempty"`
	WechatAccountId string `json:"wechat_account_id,omitempty"`
	Date            string `json:"date,omitempty"`
	Cost            int64  `json:"cost,omitempty"`
}