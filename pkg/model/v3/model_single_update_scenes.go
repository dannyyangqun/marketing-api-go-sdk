/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 应用场景信息
type SingleUpdateScenes struct {
	Scene    DataNexusScene       `json:"scene,omitempty"`
	AssetIds *[]SingleUpdateAsset `json:"asset_ids,omitempty"`
}
