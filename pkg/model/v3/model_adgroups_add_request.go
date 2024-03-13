/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdgroupsAddRequest struct {
	AccountId                         *int64                      `json:"account_id,omitempty"`
	AdgroupName                       *string                     `json:"adgroup_name,omitempty"`
	MarketingGoal                     MarketingGoal               `json:"marketing_goal,omitempty"`
	MarketingSubGoal                  MarketingSubGoal            `json:"marketing_sub_goal,omitempty"`
	MarketingCarrierType              MarketingCarrierType        `json:"marketing_carrier_type,omitempty"`
	MarketingCarrierDetail            *MarketingCarrierDetail     `json:"marketing_carrier_detail,omitempty"`
	BeginDate                         *string                     `json:"begin_date,omitempty"`
	EndDate                           *string                     `json:"end_date,omitempty"`
	FirstDayBeginTime                 *string                     `json:"first_day_begin_time,omitempty"`
	BidAmount                         *int64                      `json:"bid_amount,omitempty"`
	OptimizationGoal                  OptimizationGoal            `json:"optimization_goal,omitempty"`
	TimeSeries                        *string                     `json:"time_series,omitempty"`
	AutomaticSiteEnabled              *bool                       `json:"automatic_site_enabled,omitempty"`
	SiteSet                           *[]string                   `json:"site_set,omitempty"`
	DailyBudget                       *int64                      `json:"daily_budget,omitempty"`
	Targeting                         *WriteTargetingSetting      `json:"targeting,omitempty"`
	SceneSpec                         *SceneTargetingForWrite     `json:"scene_spec,omitempty"`
	UserActionSets                    *[]UserActionSetStruct      `json:"user_action_sets,omitempty"`
	BidStrategy                       BidStrategy                 `json:"bid_strategy,omitempty"`
	DeepConversionSpec                *DeepConversionSpec         `json:"deep_conversion_spec,omitempty"`
	ConversionId                      *int64                      `json:"conversion_id,omitempty"`
	DeepConversionBehaviorBid         *int64                      `json:"deep_conversion_behavior_bid,omitempty"`
	DeepConversionWorthRate           *float64                    `json:"deep_conversion_worth_rate,omitempty"`
	DeepConversionWorthAdvancedRate   *float64                    `json:"deep_conversion_worth_advanced_rate,omitempty"`
	DeepConversionBehaviorAdvancedBid *int64                      `json:"deep_conversion_behavior_advanced_bid,omitempty"`
	BidMode                           BidMode                     `json:"bid_mode,omitempty"`
	AutoAcquisitionEnabled            *bool                       `json:"auto_acquisition_enabled,omitempty"`
	AutoAcquisitionBudget             *int64                      `json:"auto_acquisition_budget,omitempty"`
	SmartBidType                      SmartBidType                `json:"smart_bid_type,omitempty"`
	SmartCostCap                      *int64                      `json:"smart_cost_cap,omitempty"`
	AutoDerivedCreativeEnabled        *bool                       `json:"auto_derived_creative_enabled,omitempty"`
	SearchExpandTargetingSwitch       SearchExpandTargetingSwitch `json:"search_expand_targeting_switch,omitempty"`
	AutoDerivedLandingPageSwitch      *bool                       `json:"auto_derived_landing_page_switch,omitempty"`
	BidScene                          BidScene                    `json:"bid_scene,omitempty"`
	ConfiguredStatus                  ConfiguredStatus            `json:"configured_status,omitempty"`
	FlowOptimizationEnabled           *bool                       `json:"flow_optimization_enabled,omitempty"`
	MaterialPackageId                 *int64                      `json:"material_package_id,omitempty"`
	MarketingAssetId                  *int64                      `json:"marketing_asset_id,omitempty"`
	MarketingAssetOuterSpec           *MarketingAssetOuterSpec    `json:"marketing_asset_outer_spec,omitempty"`
	PoiList                           *[]string                   `json:"poi_list,omitempty"`
	ExplorationStrategy               SiteSetExplorationStrategy  `json:"exploration_strategy,omitempty"`
	PrioritySiteSet                   *[]string                   `json:"priority_site_set,omitempty"`
	EcomPkamSwitch                    EcomPkamSwitch              `json:"ecom_pkam_switch,omitempty"`
	ForwardLinkAssist                 OptimizationGoal            `json:"forward_link_assist,omitempty"`
	MpaSpec                           *MpaSpec                    `json:"mpa_spec,omitempty"`
	CostConstraintScene               CostConstraintScene         `json:"cost_constraint_scene,omitempty"`
	CustomCostCap                     *int64                      `json:"custom_cost_cap,omitempty"`
}
