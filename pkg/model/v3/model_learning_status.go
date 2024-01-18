/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// LearningStatus : 学习状态
type LearningStatus string

// List of LearningStatus
const (
	LearningStatus_UNKNOWN  LearningStatus = "LEARNING_STATUS_UNKNOWN"
	LearningStatus_WIP      LearningStatus = "LEARNING_STATUS_WIP"
	LearningStatus_FINISHED LearningStatus = "LEARNING_STATUS_FINISHED"
	LearningStatus_FAILED   LearningStatus = "LEARNING_STATUS_FAILED"
)