/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 行为上报数据
type UserAction struct {
	ActionTime    *int64                 `json:"action_time,omitempty"`
	UserId        *UserIdDn              `json:"user_id,omitempty"`
	ActionType    ActionType             `json:"action_type,omitempty"`
	OuterActionId *string                `json:"outer_action_id,omitempty"`
	ActionParam   map[string]interface{} `json:"action_param,omitempty"`
	CustomAction  *string                `json:"custom_action,omitempty"`
	Trace         *Trace                 `json:"trace,omitempty"`
	Url           *string                `json:"url,omitempty"`
	ProductInform *ProductInform         `json:"product_inform,omitempty"`
	Channel       ActionChannelType      `json:"channel,omitempty"`
	ExtInfo       *DeviceInfo            `json:"ext_info,omitempty"`
}