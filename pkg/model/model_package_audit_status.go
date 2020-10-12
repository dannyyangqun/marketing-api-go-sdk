/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// PackageAuditStatus : 渠道包审核结果状态
type PackageAuditStatus string

// List of PackageAuditStatus
const (
	PackageAuditStatus_DIRTY_WORD            PackageAuditStatus = "AUDIT_STATUS_DIRTY_WORD"
	PackageAuditStatus_PIRATED_SIGNATURE     PackageAuditStatus = "AUDIT_STATUS_PIRATED_SIGNATURE"
	PackageAuditStatus_VIRUS                 PackageAuditStatus = "AUDIT_STATUS_VIRUS"
	PackageAuditStatus_ADMINISTRATOR_DELETE  PackageAuditStatus = "AUDIT_STATUS_ADMINISTRATOR_DELETE"
	PackageAuditStatus_MISS_RESOURCE         PackageAuditStatus = "AUDIT_STATUS_MISS_RESOURCE"
	PackageAuditStatus_REJECT                PackageAuditStatus = "AUDIT_STATUS_REJECT"
	PackageAuditStatus_GO_ILLEGAL            PackageAuditStatus = "AUDIT_STATUS_GO_ILLEGAL"
	PackageAuditStatus_OFFLINE               PackageAuditStatus = "AUDIT_STATUS_OFFLINE"
	PackageAuditStatus_INTERNAL_SERVER_ERROR PackageAuditStatus = "AUDIT_STATUS_INTERNAL_SERVER_ERROR"
	PackageAuditStatus_ONLINE                PackageAuditStatus = "AUDIT_STATUS_ONLINE"
	PackageAuditStatus_NEW_VERSION           PackageAuditStatus = "AUDIT_STATUS_NEW_VERSION"
)