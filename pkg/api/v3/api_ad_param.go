/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
	. "github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

// Linger please
var (
	_ context.Context
)

type AdParamApiService service

/*
AdParamApiService 获取词包
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param marketingGoal
 * @param creativeTemplateId
 * @param siteSet
 * @param optional nil or *AdParamGetOpts - Optional Parameters:
     * @param "MarketingSubGoal" (optional.String) -
     * @param "MarketingCarrierType" (optional.String) -
     * @param "MarketingTargetType" (optional.String) -
     * @param "ProductCatalogId" (optional.Int64) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return AdParamGetResponse
*/

type AdParamGetOpts struct {
	MarketingSubGoal     optional.String
	MarketingCarrierType optional.String
	MarketingTargetType  optional.String
	ProductCatalogId     optional.Int64
	Fields               optional.Interface
}

func (a *AdParamApiService) Get(ctx context.Context, accountId int64, marketingGoal string, creativeTemplateId int64, siteSet []string, localVarOptionals *AdParamGetOpts) (AdParamGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue AdParamGetResponseData
		localVarResponse    AdParamGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/ad_param/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	localVarQueryParams.Add("marketing_goal", parameterToString(marketingGoal, ""))
	if localVarOptionals != nil && localVarOptionals.MarketingSubGoal.IsSet() {
		localVarQueryParams.Add("marketing_sub_goal", parameterToString(localVarOptionals.MarketingSubGoal.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MarketingCarrierType.IsSet() {
		localVarQueryParams.Add("marketing_carrier_type", parameterToString(localVarOptionals.MarketingCarrierType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MarketingTargetType.IsSet() {
		localVarQueryParams.Add("marketing_target_type", parameterToString(localVarOptionals.MarketingTargetType.Value(), ""))
	}
	localVarQueryParams.Add("creative_template_id", parameterToString(creativeTemplateId, ""))
	localVarQueryParams.Add("site_set", parameterToString(siteSet, "multi"))
	if localVarOptionals != nil && localVarOptionals.ProductCatalogId.IsSet() {
		localVarQueryParams.Add("product_catalog_id", parameterToString(localVarOptionals.ProductCatalogId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *localVarResponse.Code != 0 {
				var localVarResponseErrors []model.ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			if localVarResponse.Data == nil {
				return localVarReturnValue, localVarHttpResponse.Header, err
			} else {
				return *localVarResponse.Data, localVarHttpResponse.Header, err
			}
		} else {
			return localVarReturnValue, localVarHttpResponse.Header, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v AdParamGetResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse.Header, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse.Header, newErr
		}

		return localVarReturnValue, localVarHttpResponse.Header, newErr
	}

	return localVarReturnValue, localVarHttpResponse.Header, nil
}
