/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 任务所需条件
type ApiTaskSpec struct {
	TaskTypeCreateAndroidChannelPackageSpec *TaskTypeCreateAndroidChannelPackageSpec `json:"task_type_create_android_channel_package_spec,omitempty"`
	TaskTypeUpdateAndroidChannelPackageSpec *TaskTypeUpdateAndroidChannelPackageSpec `json:"task_type_update_android_channel_package_spec,omitempty"`
}