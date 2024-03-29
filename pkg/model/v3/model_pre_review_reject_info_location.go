/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 标注位置信息
type PreReviewRejectInfoLocation struct {
	X              *int64   `json:"x,omitempty"`
	Y              *int64   `json:"y,omitempty"`
	Width          *int64   `json:"width,omitempty"`
	Height         *int64   `json:"height,omitempty"`
	TimeSecond     *float64 `json:"time_second,omitempty"`
	LocationImgUrl *string  `json:"location_img_url,omitempty"`
	ImgUrl         *string  `json:"img_url,omitempty"`
	RelatedImgUrl  *string  `json:"related_img_url,omitempty"`
}
