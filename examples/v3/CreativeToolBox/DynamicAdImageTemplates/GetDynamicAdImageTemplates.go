/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type DynamicAdImageTemplatesGetExample struct {
	TAds                           *ads.SDKClient
	AccessToken                    string
	AccountId                      int64
	ProductCatalogId               int64
	ProductMode                    string
	DynamicAdTemplateWidth         int64
	DynamicAdTemplateHeight        int64
	DynamicAdImageTemplatesGetOpts *api.DynamicAdImageTemplatesGetOpts
}

func (e *DynamicAdImageTemplatesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.ProductCatalogId = 789
	e.ProductMode = "productMode_example"
	e.DynamicAdTemplateWidth = 789
	e.DynamicAdTemplateHeight = 789
	e.DynamicAdImageTemplatesGetOpts = &api.DynamicAdImageTemplatesGetOpts{}
}

func (e *DynamicAdImageTemplatesGetExample) RunExample() (model.DynamicAdImageTemplatesGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.DynamicAdImageTemplates().Get(ctx, e.AccountId, e.ProductCatalogId, e.ProductMode, e.DynamicAdTemplateWidth, e.DynamicAdTemplateHeight, e.DynamicAdImageTemplatesGetOpts)
}

func main() {
	e := &DynamicAdImageTemplatesGetExample{}
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
