/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type WechatFundsGetExample struct {
	TAds               *ads.SDKClient
	AccessToken        string
	AccountId          int64
	WechatFundsGetOpts *api.WechatFundsGetOpts
}

func (e *WechatFundsGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.WechatFundsGetOpts = &api.WechatFundsGetOpts{

		Fields: optional.NewInterface([]string{"fund_type", "balance", "credit_roll_spec", "miniprogram_spec"}),
	}
}

func (e *WechatFundsGetExample) RunExample() (model.WechatFundsGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.WechatFunds().Get(ctx, e.AccountId, e.WechatFundsGetOpts)
}

func main() {
	e := &WechatFundsGetExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}