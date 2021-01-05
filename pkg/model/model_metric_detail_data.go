/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 报表效果数据
type MetricDetailData struct {
	Cost                  int64 `json:"cost,omitempty"`
	ExpPv                 int64 `json:"exp_pv,omitempty"`
	CanvasCpnCouponsGetPv int64 `json:"canvas_cpn_coupons_get_pv,omitempty"`
	CanvasCpnCouponsCost  int64 `json:"canvas_cpn_coupons_cost,omitempty"`
	CanvasCpnCouponsUsePv int64 `json:"canvas_cpn_coupons_use_pv,omitempty"`
	Purchase              int64 `json:"purchase,omitempty"`
	PromotionClaimOfferPv int64 `json:"promotion_claim_offer_pv,omitempty"`
	PageVisitStorePv      int64 `json:"page_visit_store_pv,omitempty"`
}