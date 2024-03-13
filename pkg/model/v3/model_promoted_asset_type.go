/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// PromotedAssetType : 推广类容类型
type PromotedAssetType string

// List of PromotedAssetType
const (
	PromotedAssetType_UNKNOWN                          PromotedAssetType = "PROMOTED_ASSET_TYPE_UNKNOWN"
	PromotedAssetType_APP_ANDROID                      PromotedAssetType = "PROMOTED_ASSET_TYPE_APP_ANDROID"
	PromotedAssetType_APP_IOS                          PromotedAssetType = "PROMOTED_ASSET_TYPE_APP_IOS"
	PromotedAssetType_WECHAT_OFFICIAL_ACCOUNT          PromotedAssetType = "PROMOTED_ASSET_TYPE_WECHAT_OFFICIAL_ACCOUNT"
	PromotedAssetType_EDUCATION                        PromotedAssetType = "PROMOTED_ASSET_TYPE_EDUCATION"
	PromotedAssetType_TRAFFIC                          PromotedAssetType = "PROMOTED_ASSET_TYPE_TRAFFIC"
	PromotedAssetType_HOUSE_PROPERTY                   PromotedAssetType = "PROMOTED_ASSET_TYPE_HOUSE_PROPERTY"
	PromotedAssetType_LOCAL_STORE                      PromotedAssetType = "PROMOTED_ASSET_TYPE_LOCAL_STORE"
	PromotedAssetType_MINIGAME                         PromotedAssetType = "PROMOTED_ASSET_TYPE_MINIGAME"
	PromotedAssetType_CONSUMER_PRODUCT                 PromotedAssetType = "PROMOTED_ASSET_TYPE_CONSUMER_PRODUCT"
	PromotedAssetType_WECHAT_CHANNELS                  PromotedAssetType = "PROMOTED_ASSET_TYPE_WECHAT_CHANNELS"
	PromotedAssetType_WECHAT_CHANNELS_LIVE             PromotedAssetType = "PROMOTED_ASSET_TYPE_WECHAT_CHANNELS_LIVE"
	PromotedAssetType_WECHAT_CHANNELS_LIVE_RESERVATION PromotedAssetType = "PROMOTED_ASSET_TYPE_WECHAT_CHANNELS_LIVE_RESERVATION"
	PromotedAssetType_MINI_PROGRAM_WECHAT              PromotedAssetType = "PROMOTED_ASSET_TYPE_MINI_PROGRAM_WECHAT"
	PromotedAssetType_APP_QUICK_APP                    PromotedAssetType = "PROMOTED_ASSET_TYPE_APP_QUICK_APP"
	PromotedAssetType_CONSUME_MEDICAL                  PromotedAssetType = "PROMOTED_ASSET_TYPE_CONSUME_MEDICAL"
	PromotedAssetType_COMPREHENSIVE_HOUSEKEEPING       PromotedAssetType = "PROMOTED_ASSET_TYPE_COMPREHENSIVE_HOUSEKEEPING"
	PromotedAssetType_FICTION                          PromotedAssetType = "PROMOTED_ASSET_TYPE_FICTION"
	PromotedAssetType_SHORT_DRAMA                      PromotedAssetType = "PROMOTED_ASSET_TYPE_SHORT_DRAMA"
	PromotedAssetType_AUDIOVISUAL_ENTERTAINMENT        PromotedAssetType = "PROMOTED_ASSET_TYPE_AUDIOVISUAL_ENTERTAINMENT"
	PromotedAssetType_BEAUTY_AND_PERSONAL_CARE         PromotedAssetType = "PROMOTED_ASSET_TYPE_BEAUTY_AND_PERSONAL_CARE"
	PromotedAssetType_WEDDING_AND_PORTRAIT_PHOTOGRAPHY PromotedAssetType = "PROMOTED_ASSET_TYPE_WEDDING_AND_PORTRAIT_PHOTOGRAPHY"
	PromotedAssetType_FRANCHISE_BRAND                  PromotedAssetType = "PROMOTED_ASSET_TYPE_FRANCHISE_BRAND"
	PromotedAssetType_ENTERPRISE_SERVICES              PromotedAssetType = "PROMOTED_ASSET_TYPE_ENTERPRISE_SERVICES"
	PromotedAssetType_EXHIBITION_BOOTH_DESIGN          PromotedAssetType = "PROMOTED_ASSET_TYPE_EXHIBITION_BOOTH_DESIGN"
	PromotedAssetType_INSURANCE                        PromotedAssetType = "PROMOTED_ASSET_TYPE_INSURANCE"
	PromotedAssetType_BANK                             PromotedAssetType = "PROMOTED_ASSET_TYPE_BANK"
	PromotedAssetType_CREDIT                           PromotedAssetType = "PROMOTED_ASSET_TYPE_CREDIT"
	PromotedAssetType_INVESTMENT_CONSULTING            PromotedAssetType = "PROMOTED_ASSET_TYPE_INVESTMENT_CONSULTING"
	PromotedAssetType_REAL_ESTATE                      PromotedAssetType = "PROMOTED_ASSET_TYPE_REAL_ESTATE"
	PromotedAssetType_TELECOMMUNICATIONS_OPERATOR      PromotedAssetType = "PROMOTED_ASSET_TYPE_TELECOMMUNICATIONS_OPERATOR"
	PromotedAssetType_TOURIST_ATTRACTIONS_TICKETS      PromotedAssetType = "PROMOTED_ASSET_TYPE_TOURIST_ATTRACTIONS_TICKETS"
	PromotedAssetType_RENOVATION_SERVICES              PromotedAssetType = "PROMOTED_ASSET_TYPE_RENOVATION_SERVICES"
	PromotedAssetType_FURNITURE_AND_BUILDING_MATERIALS PromotedAssetType = "PROMOTED_ASSET_TYPE_FURNITURE_AND_BUILDING_MATERIALS"
	PromotedAssetType_EXHIBITION_SALES                 PromotedAssetType = "PROMOTED_ASSET_TYPE_EXHIBITION_SALES"
	PromotedAssetType_MEDICINE_INDUSTRY_COMMERCIAL     PromotedAssetType = "PROMOTED_ASSET_TYPE_MEDICINE_INDUSTRY_COMMERCIAL"
	PromotedAssetType_FINANCE                          PromotedAssetType = "PROMOTED_ASSET_TYPE_FINANCE"
	PromotedAssetType_LOCAL_STORE_PACKAGE              PromotedAssetType = "PROMOTED_ASSET_TYPE_LOCAL_STORE_PACKAGE"
	PromotedAssetType_CATERING_AND_LEISURE             PromotedAssetType = "PROMOTED_ASSET_TYPE_CATERING_AND_LEISURE"
	PromotedAssetType_CHAIN_RESTAURANT                 PromotedAssetType = "PROMOTED_ASSET_TYPE_CHAIN_RESTAURANT"
	PromotedAssetType_COMMODITY_SET                    PromotedAssetType = "PROMOTED_ASSET_TYPE_COMMODITY_SET"
	PromotedAssetType_TOURIST_TRAVEL_ROUTE             PromotedAssetType = "PROMOTED_ASSET_TYPE_TOURIST_TRAVEL_ROUTE"
	PromotedAssetType_TOURIST_CRUISE_LINE              PromotedAssetType = "PROMOTED_ASSET_TYPE_TOURIST_CRUISE_LINE"
	PromotedAssetType_TOURIST_HOTEL_SERVICE            PromotedAssetType = "PROMOTED_ASSET_TYPE_TOURIST_HOTEL_SERVICE"
	PromotedAssetType_TOURIST_AIRLINE_TICKETS          PromotedAssetType = "PROMOTED_ASSET_TYPE_TOURIST_AIRLINE_TICKETS"
	PromotedAssetType_LOCAL_STORE_COMBINE_WITH_PRODUCT PromotedAssetType = "PROMOTED_ASSET_TYPE_LOCAL_STORE_COMBINE_WITH_PRODUCT"
	PromotedAssetType_ACTIVITY                         PromotedAssetType = "PROMOTED_ASSET_TYPE_ACTIVITY"
	PromotedAssetType_STORE                            PromotedAssetType = "PROMOTED_ASSET_TYPE_STORE"
	PromotedAssetType_MINI_GAME_QQ                     PromotedAssetType = "PROMOTED_ASSET_TYPE_MINI_GAME_QQ"
	PromotedAssetType_APP_GAME_ANDROID                 PromotedAssetType = "PROMOTED_ASSET_TYPE_APP_GAME_ANDROID"
	PromotedAssetType_APP_GAME_IOS                     PromotedAssetType = "PROMOTED_ASSET_TYPE_APP_GAME_IOS"
	PromotedAssetType_PC_GAME                          PromotedAssetType = "PROMOTED_ASSET_TYPE_PC_GAME"
	PromotedAssetType_WECHAT_WORK                      PromotedAssetType = "PROMOTED_ASSET_TYPE_WECHAT_WORK"
	PromotedAssetType_LIVE_STREAM_ROOM                 PromotedAssetType = "PROMOTED_ASSET_TYPE_LIVE_STREAM_ROOM"
)
