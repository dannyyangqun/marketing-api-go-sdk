/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	. "github.com/tencentad/marketing-api-go-sdk/pkg/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type WechatPagesCsgrouplistApiService service

/*
WechatPagesCsgrouplistApiService 增加企业微信组件客服组
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param data

@return WechatPagesCsgrouplistAddResponse
*/
func (a *WechatPagesCsgrouplistApiService) Add(ctx context.Context, data WechatPagesCsgrouplistAddRequest) (interface{}, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue interface{}
		localVarResponse    WechatPagesCsgrouplistAddResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/wechat_pages_csgrouplist/add"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = &data
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
				var localVarResponseErrors []ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			return localVarReturnValue, localVarHttpResponse.Header, err
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
			var v WechatPagesCsgrouplistAddResponse
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

/*
WechatPagesCsgrouplistApiService 获取企业微信客服组列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param optional nil or *WechatPagesCsgrouplistGetOpts - Optional Parameters:
     * @param "Page" (optional.Int64) -
     * @param "PageSize" (optional.Int64) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return WechatPagesCsgrouplistGetResponse
*/

type WechatPagesCsgrouplistGetOpts struct {
	Page     optional.Int64
	PageSize optional.Int64
	Fields   optional.Interface
}

func (a *WechatPagesCsgrouplistApiService) Get(ctx context.Context, accountId int64, localVarOptionals *WechatPagesCsgrouplistGetOpts) (WechatPagesCsgrouplistGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue WechatPagesCsgrouplistGetResponseData
		localVarResponse    WechatPagesCsgrouplistGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/wechat_pages_csgrouplist/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("page_size", parameterToString(localVarOptionals.PageSize.Value(), ""))
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
				var localVarResponseErrors []ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			return *localVarResponse.Data, localVarHttpResponse.Header, err
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
			var v WechatPagesCsgrouplistGetResponse
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

/*
WechatPagesCsgrouplistApiService 更新企业微信组件客服组
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param data

@return WechatPagesCsgrouplistUpdateResponse
*/
func (a *WechatPagesCsgrouplistApiService) Update(ctx context.Context, data WechatPagesCsgrouplistUpdateRequest) (interface{}, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue interface{}
		localVarResponse    WechatPagesCsgrouplistUpdateResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/wechat_pages_csgrouplist/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = &data
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
				var localVarResponseErrors []ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			return localVarReturnValue, localVarHttpResponse.Header, err
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
			var v WechatPagesCsgrouplistUpdateResponse
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
