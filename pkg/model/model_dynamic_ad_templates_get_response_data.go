/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DynamicAdTemplatesGetResponseData struct {
	List     *[]DynamicAdTemplatesGetListStruct `json:"list,omitempty"`
	PageInfo *PageInfo                          `json:"page_info,omitempty"`
}