/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WechatAccountTradeStatus : 账户交易状态
type WechatAccountTradeStatus string

// List of WechatAccountTradeStatus
const (
	WechatAccountTradeStatus_PROCESSING  WechatAccountTradeStatus = "TRADE_STATUS_PROCESSING"
	WechatAccountTradeStatus_TRANSFERRED WechatAccountTradeStatus = "TRADE_STATUS_TRANSFERRED"
	WechatAccountTradeStatus_DISAPPROVED WechatAccountTradeStatus = "TRADE_STATUS_DISAPPROVED"
	WechatAccountTradeStatus_UNSUPPORTED WechatAccountTradeStatus = "TRADE_STATUS_UNSUPPORTED"
)