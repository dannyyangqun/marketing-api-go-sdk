/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ads

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
)

func (c *SDKClient) AdDiagnosis() *api.AdDiagnosisApiService {
	return c.Client.AdDiagnosisApi
}

func (c *SDKClient) AdLabel() *api.AdLabelApiService {
	return c.Client.AdLabelApi
}

func (c *SDKClient) AdParam() *api.AdParamApiService {
	return c.Client.AdParamApi
}

func (c *SDKClient) AdcreativePreviews() *api.AdcreativePreviewsApiService {
	return c.Client.AdcreativePreviewsApi
}

func (c *SDKClient) AdcreativePreviewsQrcode() *api.AdcreativePreviewsQrcodeApiService {
	return c.Client.AdcreativePreviewsQrcodeApi
}

func (c *SDKClient) AdcreativeTemplate() *api.AdcreativeTemplateApiService {
	return c.Client.AdcreativeTemplateApi
}

func (c *SDKClient) AdcreativeTemplateDetail() *api.AdcreativeTemplateDetailApiService {
	return c.Client.AdcreativeTemplateDetailApi
}

func (c *SDKClient) AdcreativeTemplateList() *api.AdcreativeTemplateListApiService {
	return c.Client.AdcreativeTemplateListApi
}

func (c *SDKClient) AdcreativeTemplatePreview() *api.AdcreativeTemplatePreviewApiService {
	return c.Client.AdcreativeTemplatePreviewApi
}

func (c *SDKClient) AdcreativeTemplatePreviews() *api.AdcreativeTemplatePreviewsApiService {
	return c.Client.AdcreativeTemplatePreviewsApi
}

func (c *SDKClient) AdcreativeTemplates() *api.AdcreativeTemplatesApiService {
	return c.Client.AdcreativeTemplatesApi
}

func (c *SDKClient) Adcreatives() *api.AdcreativesApiService {
	return c.Client.AdcreativesApi
}

func (c *SDKClient) AdcreativesRelatedCapability() *api.AdcreativesRelatedCapabilityApiService {
	return c.Client.AdcreativesRelatedCapabilityApi
}

func (c *SDKClient) Adgroups() *api.AdgroupsApiService {
	return c.Client.AdgroupsApi
}

func (c *SDKClient) Ads() *api.AdsApiService {
	return c.Client.AdsApi
}

func (c *SDKClient) Advertiser() *api.AdvertiserApiService {
	return c.Client.AdvertiserApi
}

func (c *SDKClient) AgencyInnerTransfer() *api.AgencyInnerTransferApiService {
	return c.Client.AgencyInnerTransferApi
}

func (c *SDKClient) AgencyRealtimeCost() *api.AgencyRealtimeCostApiService {
	return c.Client.AgencyRealtimeCostApi
}

func (c *SDKClient) AndroidChannelPackages() *api.AndroidChannelPackagesApiService {
	return c.Client.AndroidChannelPackagesApi
}

func (c *SDKClient) AndroidUnionChannelPackages() *api.AndroidUnionChannelPackagesApiService {
	return c.Client.AndroidUnionChannelPackagesApi
}

func (c *SDKClient) AppAndroidChannelPackages() *api.AppAndroidChannelPackagesApiService {
	return c.Client.AppAndroidChannelPackagesApi
}

func (c *SDKClient) AssetPermissions() *api.AssetPermissionsApiService {
	return c.Client.AssetPermissionsApi
}

func (c *SDKClient) AssetPrePermissions() *api.AssetPrePermissionsApiService {
	return c.Client.AssetPrePermissionsApi
}

func (c *SDKClient) AsyncReportFiles() *api.AsyncReportFilesApiService {
	return c.Client.AsyncReportFilesApi
}

func (c *SDKClient) AsyncReports() *api.AsyncReportsApiService {
	return c.Client.AsyncReportsApi
}

func (c *SDKClient) AsyncTaskFiles() *api.AsyncTaskFilesApiService {
	return c.Client.AsyncTaskFilesApi
}

func (c *SDKClient) AsyncTasks() *api.AsyncTasksApiService {
	return c.Client.AsyncTasksApi
}

func (c *SDKClient) AudienceGrantRelations() *api.AudienceGrantRelationsApiService {
	return c.Client.AudienceGrantRelationsApi
}

func (c *SDKClient) Authorization() *api.AuthorizationApiService {
	return c.Client.AuthorizationApi
}

func (c *SDKClient) Barrage() *api.BarrageApiService {
	return c.Client.BarrageApi
}

func (c *SDKClient) BarrageRecommend() *api.BarrageRecommendApiService {
	return c.Client.BarrageRecommendApi
}

func (c *SDKClient) BatchAsyncRequestSpecification() *api.BatchAsyncRequestSpecificationApiService {
	return c.Client.BatchAsyncRequestSpecificationApi
}

func (c *SDKClient) BatchAsyncRequests() *api.BatchAsyncRequestsApiService {
	return c.Client.BatchAsyncRequestsApi
}

func (c *SDKClient) BatchOperation() *api.BatchOperationApiService {
	return c.Client.BatchOperationApi
}

func (c *SDKClient) BatchRequests() *api.BatchRequestsApiService {
	return c.Client.BatchRequestsApi
}

func (c *SDKClient) BidSimulation() *api.BidSimulationApiService {
	return c.Client.BidSimulationApi
}

func (c *SDKClient) Brand() *api.BrandApiService {
	return c.Client.BrandApi
}

func (c *SDKClient) BusinessManagerRelations() *api.BusinessManagerRelationsApiService {
	return c.Client.BusinessManagerRelationsApi
}

func (c *SDKClient) BusinessMdmAccountRelations() *api.BusinessMdmAccountRelationsApiService {
	return c.Client.BusinessMdmAccountRelationsApi
}

func (c *SDKClient) Campaigns() *api.CampaignsApiService {
	return c.Client.CampaignsApi
}

func (c *SDKClient) Capabilities() *api.CapabilitiesApiService {
	return c.Client.CapabilitiesApi
}

func (c *SDKClient) Channels() *api.ChannelsApiService {
	return c.Client.ChannelsApi
}

func (c *SDKClient) ComplianceValidation() *api.ComplianceValidationApiService {
	return c.Client.ComplianceValidationApi
}

func (c *SDKClient) Conversions() *api.ConversionsApiService {
	return c.Client.ConversionsApi
}

func (c *SDKClient) CreativetoolsText() *api.CreativetoolsTextApiService {
	return c.Client.CreativetoolsTextApi
}

func (c *SDKClient) CustomAudienceEstimations() *api.CustomAudienceEstimationsApiService {
	return c.Client.CustomAudienceEstimationsApi
}

func (c *SDKClient) CustomAudienceFiles() *api.CustomAudienceFilesApiService {
	return c.Client.CustomAudienceFilesApi
}

func (c *SDKClient) CustomAudienceInsights() *api.CustomAudienceInsightsApiService {
	return c.Client.CustomAudienceInsightsApi
}

func (c *SDKClient) CustomAudienceReports() *api.CustomAudienceReportsApiService {
	return c.Client.CustomAudienceReportsApi
}

func (c *SDKClient) CustomAudiences() *api.CustomAudiencesApiService {
	return c.Client.CustomAudiencesApi
}

func (c *SDKClient) CustomFeatures() *api.CustomFeaturesApiService {
	return c.Client.CustomFeaturesApi
}

func (c *SDKClient) CustomTagFiles() *api.CustomTagFilesApiService {
	return c.Client.CustomTagFilesApi
}

func (c *SDKClient) CustomTags() *api.CustomTagsApiService {
	return c.Client.CustomTagsApi
}

func (c *SDKClient) DailyBalanceReport() *api.DailyBalanceReportApiService {
	return c.Client.DailyBalanceReportApi
}

func (c *SDKClient) DailyCost() *api.DailyCostApiService {
	return c.Client.DailyCostApi
}

func (c *SDKClient) DailyReports() *api.DailyReportsApiService {
	return c.Client.DailyReportsApi
}

func (c *SDKClient) Diagnosis() *api.DiagnosisApiService {
	return c.Client.DiagnosisApi
}

func (c *SDKClient) DplabelAdLabel() *api.DplabelAdLabelApiService {
	return c.Client.DplabelAdLabelApi
}

func (c *SDKClient) DynamicAdImageTemplates() *api.DynamicAdImageTemplatesApiService {
	return c.Client.DynamicAdImageTemplatesApi
}

func (c *SDKClient) DynamicAdImages() *api.DynamicAdImagesApiService {
	return c.Client.DynamicAdImagesApi
}

func (c *SDKClient) DynamicAdTemplates() *api.DynamicAdTemplatesApiService {
	return c.Client.DynamicAdTemplatesApi
}

func (c *SDKClient) DynamicAdVideo() *api.DynamicAdVideoApiService {
	return c.Client.DynamicAdVideoApi
}

func (c *SDKClient) DynamicAdVideoTemplates() *api.DynamicAdVideoTemplatesApiService {
	return c.Client.DynamicAdVideoTemplatesApi
}

func (c *SDKClient) DynamicCreatives() *api.DynamicCreativesApiService {
	return c.Client.DynamicCreativesApi
}

func (c *SDKClient) EcommerceOrder() *api.EcommerceOrderApiService {
	return c.Client.EcommerceOrderApi
}

func (c *SDKClient) Estimation() *api.EstimationApiService {
	return c.Client.EstimationApi
}

func (c *SDKClient) FundStatementsDaily() *api.FundStatementsDailyApiService {
	return c.Client.FundStatementsDailyApi
}

func (c *SDKClient) FundStatementsDetailed() *api.FundStatementsDetailedApiService {
	return c.Client.FundStatementsDetailedApi
}

func (c *SDKClient) FundTransfer() *api.FundTransferApiService {
	return c.Client.FundTransferApi
}

func (c *SDKClient) Funds() *api.FundsApiService {
	return c.Client.FundsApi
}

func (c *SDKClient) HourlyReports() *api.HourlyReportsApiService {
	return c.Client.HourlyReportsApi
}

func (c *SDKClient) ImageProcessing() *api.ImageProcessingApiService {
	return c.Client.ImageProcessingApi
}

func (c *SDKClient) Images() *api.ImagesApiService {
	return c.Client.ImagesApi
}

func (c *SDKClient) LabelAudiences() *api.LabelAudiencesApiService {
	return c.Client.LabelAudiencesApi
}

func (c *SDKClient) Labels() *api.LabelsApiService {
	return c.Client.LabelsApi
}

func (c *SDKClient) LeadClues() *api.LeadCluesApiService {
	return c.Client.LeadCluesApi
}

func (c *SDKClient) LeadsForm() *api.LeadsFormApiService {
	return c.Client.LeadsFormApi
}

func (c *SDKClient) LeadsFormList() *api.LeadsFormListApiService {
	return c.Client.LeadsFormListApi
}

func (c *SDKClient) LeadsInvalidPay() *api.LeadsInvalidPayApiService {
	return c.Client.LeadsInvalidPayApi
}

func (c *SDKClient) Local() *api.LocalApiService {
	return c.Client.LocalApi
}

func (c *SDKClient) LocalEndadsmanually() *api.LocalEndadsmanuallyApiService {
	return c.Client.LocalEndadsmanuallyApi
}

func (c *SDKClient) LocalEstimatedamount() *api.LocalEstimatedamountApiService {
	return c.Client.LocalEstimatedamountApi
}

func (c *SDKClient) LocalStores() *api.LocalStoresApiService {
	return c.Client.LocalStoresApi
}

func (c *SDKClient) LocalStoresAddressParsingResult() *api.LocalStoresAddressParsingResultApiService {
	return c.Client.LocalStoresAddressParsingResultApi
}

func (c *SDKClient) LocalStoresCategories() *api.LocalStoresCategoriesApiService {
	return c.Client.LocalStoresCategoriesApi
}

func (c *SDKClient) LocalStoresSearchInfo() *api.LocalStoresSearchInfoApiService {
	return c.Client.LocalStoresSearchInfoApi
}

func (c *SDKClient) LocalStoresWxpayMerchants() *api.LocalStoresWxpayMerchantsApiService {
	return c.Client.LocalStoresWxpayMerchantsApi
}

func (c *SDKClient) MergeFundTypeDailyBalanceReport() *api.MergeFundTypeDailyBalanceReportApiService {
	return c.Client.MergeFundTypeDailyBalanceReportApi
}

func (c *SDKClient) MergeFundTypeFundStatementsDetailed() *api.MergeFundTypeFundStatementsDetailedApiService {
	return c.Client.MergeFundTypeFundStatementsDetailedApi
}

func (c *SDKClient) MergeFundTypeFunds() *api.MergeFundTypeFundsApiService {
	return c.Client.MergeFundTypeFundsApi
}

func (c *SDKClient) MergeFundTypeSubcustomerTransfer() *api.MergeFundTypeSubcustomerTransferApiService {
	return c.Client.MergeFundTypeSubcustomerTransferApi
}

func (c *SDKClient) Oauth() *api.OauthApiService {
	return c.Client.OauthApi
}

func (c *SDKClient) OptimizationGoalPermissions() *api.OptimizationGoalPermissionsApiService {
	return c.Client.OptimizationGoalPermissionsApi
}

func (c *SDKClient) OuterClues() *api.OuterCluesApiService {
	return c.Client.OuterCluesApi
}

func (c *SDKClient) OuterCluesContact() *api.OuterCluesContactApiService {
	return c.Client.OuterCluesContactApi
}

func (c *SDKClient) Pages() *api.PagesApiService {
	return c.Client.PagesApi
}

func (c *SDKClient) PlayablePages() *api.PlayablePagesApiService {
	return c.Client.PlayablePagesApi
}

func (c *SDKClient) ProductCatalogs() *api.ProductCatalogsApiService {
	return c.Client.ProductCatalogsApi
}

func (c *SDKClient) ProductCatalogsReports() *api.ProductCatalogsReportsApiService {
	return c.Client.ProductCatalogsReportsApi
}

func (c *SDKClient) ProductCategoriesList() *api.ProductCategoriesListApiService {
	return c.Client.ProductCategoriesListApi
}

func (c *SDKClient) ProductItems() *api.ProductItemsApiService {
	return c.Client.ProductItemsApi
}

func (c *SDKClient) ProductItemsDetail() *api.ProductItemsDetailApiService {
	return c.Client.ProductItemsDetailApi
}

func (c *SDKClient) ProductItemsVerticals() *api.ProductItemsVerticalsApiService {
	return c.Client.ProductItemsVerticalsApi
}

func (c *SDKClient) ProductSeries() *api.ProductSeriesApiService {
	return c.Client.ProductSeriesApi
}

func (c *SDKClient) ProductsSystemStatus() *api.ProductsSystemStatusApiService {
	return c.Client.ProductsSystemStatusApi
}

func (c *SDKClient) Profiles() *api.ProfilesApiService {
	return c.Client.ProfilesApi
}

func (c *SDKClient) PromotedObjects() *api.PromotedObjectsApiService {
	return c.Client.PromotedObjectsApi
}

func (c *SDKClient) PropertyFileSessions() *api.PropertyFileSessionsApiService {
	return c.Client.PropertyFileSessionsApi
}

func (c *SDKClient) PropertyFiles() *api.PropertyFilesApiService {
	return c.Client.PropertyFilesApi
}

func (c *SDKClient) PropertySetSchemas() *api.PropertySetSchemasApiService {
	return c.Client.PropertySetSchemasApi
}

func (c *SDKClient) PropertySets() *api.PropertySetsApiService {
	return c.Client.PropertySetsApi
}

func (c *SDKClient) Qualifications() *api.QualificationsApiService {
	return c.Client.QualificationsApi
}

func (c *SDKClient) RealtimeCost() *api.RealtimeCostApiService {
	return c.Client.RealtimeCostApi
}

func (c *SDKClient) Report() *api.ReportApiService {
	return c.Client.ReportApi
}

func (c *SDKClient) ReviewElementPrereviewResults() *api.ReviewElementPrereviewResultsApiService {
	return c.Client.ReviewElementPrereviewResultsApi
}

func (c *SDKClient) SceneSpecTags() *api.SceneSpecTagsApiService {
	return c.Client.SceneSpecTagsApi
}

func (c *SDKClient) Shop() *api.ShopApiService {
	return c.Client.ShopApi
}

func (c *SDKClient) SplitTests() *api.SplitTestsApiService {
	return c.Client.SplitTestsApi
}

func (c *SDKClient) SubcustomerTransfer() *api.SubcustomerTransferApiService {
	return c.Client.SubcustomerTransferApi
}

func (c *SDKClient) SystemStatus() *api.SystemStatusApiService {
	return c.Client.SystemStatusApi
}

func (c *SDKClient) TargetingTagReports() *api.TargetingTagReportsApiService {
	return c.Client.TargetingTagReportsApi
}

func (c *SDKClient) TargetingTags() *api.TargetingTagsApiService {
	return c.Client.TargetingTagsApi
}

func (c *SDKClient) TargetingTagsUv() *api.TargetingTagsUvApiService {
	return c.Client.TargetingTagsUvApi
}

func (c *SDKClient) Targetings() *api.TargetingsApiService {
	return c.Client.TargetingsApi
}

func (c *SDKClient) TargetingsShare() *api.TargetingsShareApiService {
	return c.Client.TargetingsShareApi
}

func (c *SDKClient) TrackingReports() *api.TrackingReportsApiService {
	return c.Client.TrackingReportsApi
}

func (c *SDKClient) UnionPositionPackages() *api.UnionPositionPackagesApiService {
	return c.Client.UnionPositionPackagesApi
}

func (c *SDKClient) UserActionSetReports() *api.UserActionSetReportsApiService {
	return c.Client.UserActionSetReportsApi
}

func (c *SDKClient) UserActionSets() *api.UserActionSetsApiService {
	return c.Client.UserActionSetsApi
}

func (c *SDKClient) UserActions() *api.UserActionsApiService {
	return c.Client.UserActionsApi
}

func (c *SDKClient) UserProperties() *api.UserPropertiesApiService {
	return c.Client.UserPropertiesApi
}

func (c *SDKClient) UserPropertySets() *api.UserPropertySetsApiService {
	return c.Client.UserPropertySetsApi
}

func (c *SDKClient) VideomakerAutoadjustments() *api.VideomakerAutoadjustmentsApiService {
	return c.Client.VideomakerAutoadjustmentsApi
}

func (c *SDKClient) VideomakerSubtitles() *api.VideomakerSubtitlesApiService {
	return c.Client.VideomakerSubtitlesApi
}

func (c *SDKClient) VideomakerTasks() *api.VideomakerTasksApiService {
	return c.Client.VideomakerTasksApi
}

func (c *SDKClient) VideomakerVideocaptures() *api.VideomakerVideocapturesApiService {
	return c.Client.VideomakerVideocapturesApi
}

func (c *SDKClient) Videos() *api.VideosApiService {
	return c.Client.VideosApi
}

func (c *SDKClient) WechatAdLabels() *api.WechatAdLabelsApiService {
	return c.Client.WechatAdLabelsApi
}

func (c *SDKClient) WechatAdvertiser() *api.WechatAdvertiserApiService {
	return c.Client.WechatAdvertiserApi
}

func (c *SDKClient) WechatAdvertiserDetail() *api.WechatAdvertiserDetailApiService {
	return c.Client.WechatAdvertiserDetailApi
}

func (c *SDKClient) WechatAdvertiserLocalBusiness() *api.WechatAdvertiserLocalBusinessApiService {
	return c.Client.WechatAdvertiserLocalBusinessApi
}

func (c *SDKClient) WechatAdvertiserSpecification() *api.WechatAdvertiserSpecificationApiService {
	return c.Client.WechatAdvertiserSpecificationApi
}

func (c *SDKClient) WechatAgency() *api.WechatAgencyApiService {
	return c.Client.WechatAgencyApi
}

func (c *SDKClient) WechatDailyCost() *api.WechatDailyCostApiService {
	return c.Client.WechatDailyCostApi
}

func (c *SDKClient) WechatFundStatementsDetailed() *api.WechatFundStatementsDetailedApiService {
	return c.Client.WechatFundStatementsDetailedApi
}

func (c *SDKClient) WechatFundTransfer() *api.WechatFundTransferApiService {
	return c.Client.WechatFundTransferApi
}

func (c *SDKClient) WechatFunds() *api.WechatFundsApiService {
	return c.Client.WechatFundsApi
}

func (c *SDKClient) WechatPages() *api.WechatPagesApiService {
	return c.Client.WechatPagesApi
}

func (c *SDKClient) WechatPagesCsgroupStatus() *api.WechatPagesCsgroupStatusApiService {
	return c.Client.WechatPagesCsgroupStatusApi
}

func (c *SDKClient) WechatPagesCsgroupUser() *api.WechatPagesCsgroupUserApiService {
	return c.Client.WechatPagesCsgroupUserApi
}

func (c *SDKClient) WechatPagesCsgrouplist() *api.WechatPagesCsgrouplistApiService {
	return c.Client.WechatPagesCsgrouplistApi
}

func (c *SDKClient) WechatPagesCustom() *api.WechatPagesCustomApiService {
	return c.Client.WechatPagesCustomApi
}

func (c *SDKClient) WechatPagesGrantinfo() *api.WechatPagesGrantinfoApiService {
	return c.Client.WechatPagesGrantinfoApi
}

func (c *SDKClient) WechatQualifications() *api.WechatQualificationsApiService {
	return c.Client.WechatQualificationsApi
}

func (c *SDKClient) Wildcards() *api.WildcardsApiService {
	return c.Client.WildcardsApi
}

func (c *SDKClient) XijingComplexTemplate() *api.XijingComplexTemplateApiService {
	return c.Client.XijingComplexTemplateApi
}

func (c *SDKClient) XijingPage() *api.XijingPageApiService {
	return c.Client.XijingPageApi
}

func (c *SDKClient) XijingPageByComponents() *api.XijingPageByComponentsApiService {
	return c.Client.XijingPageByComponentsApi
}

func (c *SDKClient) XijingPageInteractive() *api.XijingPageInteractiveApiService {
	return c.Client.XijingPageInteractiveApi
}

func (c *SDKClient) XijingPageList() *api.XijingPageListApiService {
	return c.Client.XijingPageListApi
}

func (c *SDKClient) XijingTemplate() *api.XijingTemplateApiService {
	return c.Client.XijingTemplateApi
}

func (c *SDKClient) XijingTemplateList() *api.XijingTemplateListApiService {
	return c.Client.XijingTemplateListApi
}
