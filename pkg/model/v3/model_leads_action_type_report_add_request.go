/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LeadsActionTypeReportAddRequest struct {
	AccountId                 *int64                             `json:"account_id,omitempty"`
	MatchType                 LeadsMatchType                     `json:"match_type,omitempty"`
	LeadsActionTypeReportList *[]LeadsActionTypeReportListStruct `json:"leads_action_type_report_list,omitempty"`
}
