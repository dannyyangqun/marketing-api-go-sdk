/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DailyCostGetResponseData struct {
	List     *[]DailyCostGetListStruct `json:"list,omitempty"`
	PageInfo *Conf                     `json:"page_info,omitempty"`
}