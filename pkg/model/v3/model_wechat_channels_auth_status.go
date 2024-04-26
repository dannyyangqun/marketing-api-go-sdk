/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WechatChannelsAuthStatus : 授权状态
type WechatChannelsAuthStatus string

// List of WechatChannelsAuthStatus
const (
	WechatChannelsAuthStatus_PENDING       WechatChannelsAuthStatus = "PENDING"
	WechatChannelsAuthStatus_AUTHORIZED    WechatChannelsAuthStatus = "AUTHORIZED"
	WechatChannelsAuthStatus_CANCELLED     WechatChannelsAuthStatus = "CANCELLED"
	WechatChannelsAuthStatus_REFUSED       WechatChannelsAuthStatus = "REFUSED"
	WechatChannelsAuthStatus_EXPIRED       WechatChannelsAuthStatus = "EXPIRED"
	WechatChannelsAuthStatus_AUDIT_PENDING WechatChannelsAuthStatus = "AUDIT_PENDING"
	WechatChannelsAuthStatus_AUDIT_REFUSED WechatChannelsAuthStatus = "AUDIT_REFUSED"
)