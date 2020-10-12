/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdvertiserUpdateRequest struct {
	AccountId               int64                    `json:"account_id,omitempty"`
	DailyBudget             *int64                   `json:"daily_budget,omitempty"`
	SystemIndustryId        int64                    `json:"system_industry_id,omitempty"`
	CorporationName         string                   `json:"corporation_name,omitempty"`
	CertificationImageId    string                   `json:"certification_image_id,omitempty"`
	CorporateImageName      string                   `json:"corporate_image_name,omitempty"`
	IndividualQualification *IndividualQualification `json:"individual_qualification,omitempty"`
	IntroductionUrl         string                   `json:"introduction_url,omitempty"`
	ContactPersonTelephone  string                   `json:"contact_person_telephone,omitempty"`
	ContactPersonMobile     string                   `json:"contact_person_mobile,omitempty"`
	WechatSpec              *MpInfoUpdate            `json:"wechat_spec,omitempty"`
	Websites                *[]WebsiteUpdateStruct   `json:"websites,omitempty"`
}