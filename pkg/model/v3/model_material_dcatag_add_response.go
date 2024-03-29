/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type MaterialDcatagAddResponse struct {
	Code      *int64                         `json:"code,omitempty"`
	Message   *string                        `json:"message,omitempty"`
	MessageCn *string                        `json:"message_cn,omitempty"`
	Errors    *[]model.ApiErrorStruct        `json:"errors,omitempty"`
	Data      *MaterialDcatagAddResponseData `json:"data,omitempty"`
}
