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
type AdvertiserGetListStruct struct {
	AccountId               *int64                      `json:"account_id,omitempty"`
	DailyBudget             *int64                      `json:"daily_budget,omitempty"`
	SystemStatus            CustomerSystemStatus        `json:"system_status,omitempty"`
	RejectMessage           *string                     `json:"reject_message,omitempty"`
	CorporationName         *string                     `json:"corporation_name,omitempty"`
	CorporationLicence      *string                     `json:"corporation_licence,omitempty"`
	CertificationImageId    *string                     `json:"certification_image_id,omitempty"`
	CertificationImage      *string                     `json:"certification_image,omitempty"`
	IdentityNumber          *string                     `json:"identity_number,omitempty"`
	IndividualQualification *IndividualQualification    `json:"individual_qualification,omitempty"`
	SystemIndustryId        *int64                      `json:"system_industry_id,omitempty"`
	CustomizedIndustry      *string                     `json:"customized_industry,omitempty"`
	IntroductionUrl         *string                     `json:"introduction_url,omitempty"`
	CorporateBrandName      *string                     `json:"corporate_brand_name,omitempty"`
	ContactPerson           *string                     `json:"contact_person,omitempty"`
	ContactPersonEmail      *string                     `json:"contact_person_email,omitempty"`
	ContactPersonTelephone  *string                     `json:"contact_person_telephone,omitempty"`
	ContactPersonMobile     *string                     `json:"contact_person_mobile,omitempty"`
	Websites                *[]WebsiteReadStruct        `json:"websites,omitempty"`
	MdmId                   *int64                      `json:"mdm_id,omitempty"`
	MdmName                 *string                     `json:"mdm_name,omitempty"`
	AgencyAccountId         *int64                      `json:"agency_account_id,omitempty"`
	Operators               *[]AdvertiserOperatorStruct `json:"operators,omitempty"`
	Memo                    *string                     `json:"memo,omitempty"`
	AreaCode                *int64                      `json:"area_code,omitempty"`
}
