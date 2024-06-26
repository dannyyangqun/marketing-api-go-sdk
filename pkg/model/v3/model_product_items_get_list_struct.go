/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type ProductItemsGetListStruct struct {
	ProductOuterId     *string  `json:"product_outer_id,omitempty"`
	ProductName        *string  `json:"product_name,omitempty"`
	ProductImageUrl    *string  `json:"product_image_url,omitempty"`
	ProductShortName   *string  `json:"product_short_name,omitempty"`
	Price              *float64 `json:"price,omitempty"`
	FirstCategoryId    *int64   `json:"first_category_id,omitempty"`
	FirstCategoryName  *string  `json:"first_category_name,omitempty"`
	SecondCategoryId   *int64   `json:"second_category_id,omitempty"`
	SecondCategoryName *string  `json:"second_category_name,omitempty"`
	ThirdCategoryId    *int64   `json:"third_category_id,omitempty"`
	ThirdCategoryName  *string  `json:"third_category_name,omitempty"`
	FourthCategoryId   *int64   `json:"fourth_category_id,omitempty"`
	FourthCategoryName *string  `json:"fourth_category_name,omitempty"`
	BrandName          *string  `json:"brand_name,omitempty"`
	BrandId            *int64   `json:"brand_id,omitempty"`
	Description        *string  `json:"description,omitempty"`
	CustomData         *string  `json:"custom_data,omitempty"`
	IsVideo            *bool    `json:"is_video,omitempty"`
	DataSource         *string  `json:"data_source,omitempty"`
	ProductSelectScore *float64 `json:"product_select_score,omitempty"`
}
