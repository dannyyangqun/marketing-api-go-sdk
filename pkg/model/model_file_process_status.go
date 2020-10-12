/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// FileProcessStatus : 处理状态
type FileProcessStatus string

// List of FileProcessStatus
const (
	FileProcessStatus_PENDING    FileProcessStatus = "PENDING"
	FileProcessStatus_PROCESSING FileProcessStatus = "PROCESSING"
	FileProcessStatus_SUCCESS    FileProcessStatus = "SUCCESS"
	FileProcessStatus_ERROR      FileProcessStatus = "ERROR"
)