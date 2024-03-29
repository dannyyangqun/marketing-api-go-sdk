/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ads

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
)

func (c *SDKClient) AdDiagnosis() *api.AdDiagnosisApiService {
	return c.Client.AdDiagnosisApi
}

func (c *SDKClient) AdParam() *api.AdParamApiService {
	return c.Client.AdParamApi
}

func (c *SDKClient) AdUnionReports() *api.AdUnionReportsApiService {
	return c.Client.AdUnionReportsApi
}

func (c *SDKClient) AdcreativePreviews() *api.AdcreativePreviewsApiService {
	return c.Client.AdcreativePreviewsApi
}

func (c *SDKClient) AdcreativePreviewsQrcode() *api.AdcreativePreviewsQrcodeApiService {
	return c.Client.AdcreativePreviewsQrcodeApi
}

func (c *SDKClient) AdgroupNegativewords() *api.AdgroupNegativewordsApiService {
	return c.Client.AdgroupNegativewordsApi
}

func (c *SDKClient) Adgroups() *api.AdgroupsApiService {
	return c.Client.AdgroupsApi
}

func (c *SDKClient) Advertiser() *api.AdvertiserApiService {
	return c.Client.AdvertiserApi
}

func (c *SDKClient) AdvertiserDailyBudget() *api.AdvertiserDailyBudgetApiService {
	return c.Client.AdvertiserDailyBudgetApi
}

func (c *SDKClient) Agency() *api.AgencyApiService {
	return c.Client.AgencyApi
}

func (c *SDKClient) AgencyRealtimeCost() *api.AgencyRealtimeCostApiService {
	return c.Client.AgencyRealtimeCostApi
}

func (c *SDKClient) AndroidChannel() *api.AndroidChannelApiService {
	return c.Client.AndroidChannelApi
}

func (c *SDKClient) AsyncReportFiles() *api.AsyncReportFilesApiService {
	return c.Client.AsyncReportFilesApi
}

func (c *SDKClient) AsyncReports() *api.AsyncReportsApiService {
	return c.Client.AsyncReportsApi
}

func (c *SDKClient) AsyncTasks() *api.AsyncTasksApiService {
	return c.Client.AsyncTasksApi
}

func (c *SDKClient) AudienceGrantRelations() *api.AudienceGrantRelationsApiService {
	return c.Client.AudienceGrantRelationsApi
}

func (c *SDKClient) BatchAsyncRequestSpecification() *api.BatchAsyncRequestSpecificationApiService {
	return c.Client.BatchAsyncRequestSpecificationApi
}

func (c *SDKClient) BatchAsyncRequests() *api.BatchAsyncRequestsApiService {
	return c.Client.BatchAsyncRequestsApi
}

func (c *SDKClient) BatchRequests() *api.BatchRequestsApiService {
	return c.Client.BatchRequestsApi
}

func (c *SDKClient) BidSimulation() *api.BidSimulationApiService {
	return c.Client.BidSimulationApi
}

func (c *SDKClient) Bidword() *api.BidwordApiService {
	return c.Client.BidwordApi
}

func (c *SDKClient) BidwordFlow() *api.BidwordFlowApiService {
	return c.Client.BidwordFlowApi
}

func (c *SDKClient) Brand() *api.BrandApiService {
	return c.Client.BrandApi
}

func (c *SDKClient) BusinessPoint() *api.BusinessPointApiService {
	return c.Client.BusinessPointApi
}

func (c *SDKClient) ComponentReviewResults() *api.ComponentReviewResultsApiService {
	return c.Client.ComponentReviewResultsApi
}

func (c *SDKClient) Conversions() *api.ConversionsApiService {
	return c.Client.ConversionsApi
}

func (c *SDKClient) CreativeTemplate() *api.CreativeTemplateApiService {
	return c.Client.CreativeTemplateApi
}

func (c *SDKClient) CreativeTemplateList() *api.CreativeTemplateListApiService {
	return c.Client.CreativeTemplateListApi
}

func (c *SDKClient) CreativeTemplatePreviews() *api.CreativeTemplatePreviewsApiService {
	return c.Client.CreativeTemplatePreviewsApi
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

func (c *SDKClient) CustomAudiences() *api.CustomAudiencesApiService {
	return c.Client.CustomAudiencesApi
}

func (c *SDKClient) DailyBalanceReport() *api.DailyBalanceReportApiService {
	return c.Client.DailyBalanceReportApi
}

func (c *SDKClient) DailyReports() *api.DailyReportsApiService {
	return c.Client.DailyReportsApi
}

func (c *SDKClient) DataSourceDispatch() *api.DataSourceDispatchApiService {
	return c.Client.DataSourceDispatchApi
}

func (c *SDKClient) DynamicAdImageTemplates() *api.DynamicAdImageTemplatesApiService {
	return c.Client.DynamicAdImageTemplatesApi
}

func (c *SDKClient) DynamicAdImages() *api.DynamicAdImagesApiService {
	return c.Client.DynamicAdImagesApi
}

func (c *SDKClient) DynamicAdVideo() *api.DynamicAdVideoApiService {
	return c.Client.DynamicAdVideoApi
}

func (c *SDKClient) DynamicAdVideoTemplates() *api.DynamicAdVideoTemplatesApiService {
	return c.Client.DynamicAdVideoTemplatesApi
}

func (c *SDKClient) DynamicCreativeReviewResults() *api.DynamicCreativeReviewResultsApiService {
	return c.Client.DynamicCreativeReviewResultsApi
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

func (c *SDKClient) ExtendPackage() *api.ExtendPackageApiService {
	return c.Client.ExtendPackageApi
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

func (c *SDKClient) GameFeature() *api.GameFeatureApiService {
	return c.Client.GameFeatureApi
}

func (c *SDKClient) GameFeatureTags() *api.GameFeatureTagsApiService {
	return c.Client.GameFeatureTagsApi
}

func (c *SDKClient) HourlyReports() *api.HourlyReportsApiService {
	return c.Client.HourlyReportsApi
}

func (c *SDKClient) Images() *api.ImagesApiService {
	return c.Client.ImagesApi
}

func (c *SDKClient) KeywordRecommend() *api.KeywordRecommendApiService {
	return c.Client.KeywordRecommendApi
}

func (c *SDKClient) Labels() *api.LabelsApiService {
	return c.Client.LabelsApi
}

func (c *SDKClient) LandingPageSellStrategy() *api.LandingPageSellStrategyApiService {
	return c.Client.LandingPageSellStrategyApi
}

func (c *SDKClient) Leads() *api.LeadsApiService {
	return c.Client.LeadsApi
}

func (c *SDKClient) LeadsActionTypeReport() *api.LeadsActionTypeReportApiService {
	return c.Client.LeadsActionTypeReportApi
}

func (c *SDKClient) LeadsCallRecord() *api.LeadsCallRecordApiService {
	return c.Client.LeadsCallRecordApi
}

func (c *SDKClient) LeadsCallRecords() *api.LeadsCallRecordsApiService {
	return c.Client.LeadsCallRecordsApi
}

func (c *SDKClient) LeadsCallVirtualNumber() *api.LeadsCallVirtualNumberApiService {
	return c.Client.LeadsCallVirtualNumberApi
}

func (c *SDKClient) LeadsClaim() *api.LeadsClaimApiService {
	return c.Client.LeadsClaimApi
}

func (c *SDKClient) LeadsInvalidPay() *api.LeadsInvalidPayApiService {
	return c.Client.LeadsInvalidPayApi
}

func (c *SDKClient) LeadsList() *api.LeadsListApiService {
	return c.Client.LeadsListApi
}

func (c *SDKClient) LeadsStatus() *api.LeadsStatusApiService {
	return c.Client.LeadsStatusApi
}

func (c *SDKClient) LeadsVoipCall() *api.LeadsVoipCallApiService {
	return c.Client.LeadsVoipCallApi
}

func (c *SDKClient) LeadsVoipCallToken() *api.LeadsVoipCallTokenApiService {
	return c.Client.LeadsVoipCallTokenApi
}

func (c *SDKClient) LiveRoomComponentStatus() *api.LiveRoomComponentStatusApiService {
	return c.Client.LiveRoomComponentStatusApi
}

func (c *SDKClient) LiveRoomComponents() *api.LiveRoomComponentsApiService {
	return c.Client.LiveRoomComponentsApi
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

func (c *SDKClient) MarketingTargetAssetCategories() *api.MarketingTargetAssetCategoriesApiService {
	return c.Client.MarketingTargetAssetCategoriesApi
}

func (c *SDKClient) MarketingTargetAssetDetail() *api.MarketingTargetAssetDetailApiService {
	return c.Client.MarketingTargetAssetDetailApi
}

func (c *SDKClient) MarketingTargetAssetProperties() *api.MarketingTargetAssetPropertiesApiService {
	return c.Client.MarketingTargetAssetPropertiesApi
}

func (c *SDKClient) MarketingTargetAssetPropertyValues() *api.MarketingTargetAssetPropertyValuesApiService {
	return c.Client.MarketingTargetAssetPropertyValuesApi
}

func (c *SDKClient) MarketingTargetAssets() *api.MarketingTargetAssetsApiService {
	return c.Client.MarketingTargetAssetsApi
}

func (c *SDKClient) MarketingTargetTypes() *api.MarketingTargetTypesApiService {
	return c.Client.MarketingTargetTypesApi
}

func (c *SDKClient) MaterialDcatag() *api.MaterialDcatagApiService {
	return c.Client.MaterialDcatagApi
}

func (c *SDKClient) MaterialLabels() *api.MaterialLabelsApiService {
	return c.Client.MaterialLabelsApi
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

func (c *SDKClient) OrganizationAccountRelation() *api.OrganizationAccountRelationApiService {
	return c.Client.OrganizationAccountRelationApi
}

func (c *SDKClient) Pages() *api.PagesApiService {
	return c.Client.PagesApi
}

func (c *SDKClient) ProductCatalogs() *api.ProductCatalogsApiService {
	return c.Client.ProductCatalogsApi
}

func (c *SDKClient) ProductItems() *api.ProductItemsApiService {
	return c.Client.ProductItemsApi
}

func (c *SDKClient) ProductSeries() *api.ProductSeriesApiService {
	return c.Client.ProductSeriesApi
}

func (c *SDKClient) Profiles() *api.ProfilesApiService {
	return c.Client.ProfilesApi
}

func (c *SDKClient) Programmed() *api.ProgrammedApiService {
	return c.Client.ProgrammedApi
}

func (c *SDKClient) ProgrammedTemplate() *api.ProgrammedTemplateApiService {
	return c.Client.ProgrammedTemplateApi
}

func (c *SDKClient) Qualifications() *api.QualificationsApiService {
	return c.Client.QualificationsApi
}

func (c *SDKClient) RealtimeCost() *api.RealtimeCostApiService {
	return c.Client.RealtimeCostApi
}

func (c *SDKClient) ReviewElementPrereviewResults() *api.ReviewElementPrereviewResultsApiService {
	return c.Client.ReviewElementPrereviewResultsApi
}

func (c *SDKClient) SceneSpecTags() *api.SceneSpecTagsApiService {
	return c.Client.SceneSpecTagsApi
}

func (c *SDKClient) SubcustomerTransfer() *api.SubcustomerTransferApiService {
	return c.Client.SubcustomerTransferApi
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

func (c *SDKClient) VideoChannelDealerData() *api.VideoChannelDealerDataApiService {
	return c.Client.VideoChannelDealerDataApi
}

func (c *SDKClient) VideoChannelFansData() *api.VideoChannelFansDataApiService {
	return c.Client.VideoChannelFansDataApi
}

func (c *SDKClient) VideoChannelLeadsData() *api.VideoChannelLeadsDataApiService {
	return c.Client.VideoChannelLeadsDataApi
}

func (c *SDKClient) VideoChannelLiveData() *api.VideoChannelLiveDataApiService {
	return c.Client.VideoChannelLiveDataApi
}

func (c *SDKClient) Videos() *api.VideosApiService {
	return c.Client.VideosApi
}

func (c *SDKClient) WechatPages() *api.WechatPagesApiService {
	return c.Client.WechatPagesApi
}

func (c *SDKClient) WechatPagesCustom() *api.WechatPagesCustomApiService {
	return c.Client.WechatPagesCustomApi
}

func (c *SDKClient) WechatPagesGrantinfo() *api.WechatPagesGrantinfoApiService {
	return c.Client.WechatPagesGrantinfoApi
}

func (c *SDKClient) Wildcards() *api.WildcardsApiService {
	return c.Client.WildcardsApi
}

func (c *SDKClient) WxPackageAccount() *api.WxPackageAccountApiService {
	return c.Client.WxPackageAccountApi
}

func (c *SDKClient) WxPackagePackage() *api.WxPackagePackageApiService {
	return c.Client.WxPackagePackageApi
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

func (c *SDKClient) XijingPageList() *api.XijingPageListApiService {
	return c.Client.XijingPageListApi
}

func (c *SDKClient) XijingTemplate() *api.XijingTemplateApiService {
	return c.Client.XijingTemplateApi
}

func (c *SDKClient) XijingTemplateList() *api.XijingTemplateListApiService {
	return c.Client.XijingTemplateListApi
}
