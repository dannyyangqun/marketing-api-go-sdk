/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdgroupsUpdateDatetimeRequest struct {
	AccountId          *int64                  `json:"account_id,omitempty"`
	UpdateDatetimeSpec *[]UpdateDatetimeStruct `json:"update_datetime_spec,omitempty"`
}
