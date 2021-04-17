/*
 * lighthouse
 *
 * lighthouse  20210117
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// LighthouseApiService LighthouseApi service
type LighthouseApiService service

type ApiSearchAutoCompleteRequest struct {
	ctx _context.Context
	ApiService *LighthouseApiService
	s *string
	size *int32
	from *int32
	nsfw *bool
}

func (r ApiSearchAutoCompleteRequest) S(s string) ApiSearchAutoCompleteRequest {
	r.s = &s
	return r
}
func (r ApiSearchAutoCompleteRequest) Size(size int32) ApiSearchAutoCompleteRequest {
	r.size = &size
	return r
}
func (r ApiSearchAutoCompleteRequest) From(from int32) ApiSearchAutoCompleteRequest {
	r.from = &from
	return r
}
func (r ApiSearchAutoCompleteRequest) Nsfw(nsfw bool) ApiSearchAutoCompleteRequest {
	r.nsfw = &nsfw
	return r
}

func (r ApiSearchAutoCompleteRequest) Execute() ([]string, *_nethttp.Response, error) {
	return r.ApiService.SearchAutoCompleteExecute(r)
}

/*
 * SearchAutoComplete 搜尋自動完成建議
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchAutoCompleteRequest
 */
func (a *LighthouseApiService) SearchAutoComplete(ctx _context.Context) ApiSearchAutoCompleteRequest {
	return ApiSearchAutoCompleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []string
 */
func (a *LighthouseApiService) SearchAutoCompleteExecute(r ApiSearchAutoCompleteRequest) ([]string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LighthouseApiService.SearchAutoComplete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/autocomplete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.s == nil {
		return localVarReturnValue, nil, reportError("s is required and must be specified")
	}

	localVarQueryParams.Add("s", parameterToString(*r.s, ""))
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.nsfw != nil {
		localVarQueryParams.Add("nsfw", parameterToString(*r.nsfw, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"array"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchByStringRequest struct {
	ctx _context.Context
	ApiService *LighthouseApiService
	s *string
	size *int32
	from *int32
	channel *string
	channelId *string
	relatedTo *string
	sortBy *string
	include *string
	nsfw *bool
	freeOnly *bool
	resolve *bool
}

func (r ApiSearchByStringRequest) S(s string) ApiSearchByStringRequest {
	r.s = &s
	return r
}
func (r ApiSearchByStringRequest) Size(size int32) ApiSearchByStringRequest {
	r.size = &size
	return r
}
func (r ApiSearchByStringRequest) From(from int32) ApiSearchByStringRequest {
	r.from = &from
	return r
}
func (r ApiSearchByStringRequest) Channel(channel string) ApiSearchByStringRequest {
	r.channel = &channel
	return r
}
func (r ApiSearchByStringRequest) ChannelId(channelId string) ApiSearchByStringRequest {
	r.channelId = &channelId
	return r
}
func (r ApiSearchByStringRequest) RelatedTo(relatedTo string) ApiSearchByStringRequest {
	r.relatedTo = &relatedTo
	return r
}
func (r ApiSearchByStringRequest) SortBy(sortBy string) ApiSearchByStringRequest {
	r.sortBy = &sortBy
	return r
}
func (r ApiSearchByStringRequest) Include(include string) ApiSearchByStringRequest {
	r.include = &include
	return r
}
func (r ApiSearchByStringRequest) Nsfw(nsfw bool) ApiSearchByStringRequest {
	r.nsfw = &nsfw
	return r
}
func (r ApiSearchByStringRequest) FreeOnly(freeOnly bool) ApiSearchByStringRequest {
	r.freeOnly = &freeOnly
	return r
}
func (r ApiSearchByStringRequest) Resolve(resolve bool) ApiSearchByStringRequest {
	r.resolve = &resolve
	return r
}

func (r ApiSearchByStringRequest) Execute() ([]SearchResponse, *_nethttp.Response, error) {
	return r.ApiService.SearchByStringExecute(r)
}

/*
 * SearchByString 字串搜尋
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchByStringRequest
 */
func (a *LighthouseApiService) SearchByString(ctx _context.Context) ApiSearchByStringRequest {
	return ApiSearchByStringRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []SearchResponse
 */
func (a *LighthouseApiService) SearchByStringExecute(r ApiSearchByStringRequest) ([]SearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LighthouseApiService.SearchByString")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.s == nil {
		return localVarReturnValue, nil, reportError("s is required and must be specified")
	}

	localVarQueryParams.Add("s", parameterToString(*r.s, ""))
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.channel != nil {
		localVarQueryParams.Add("channel", parameterToString(*r.channel, ""))
	}
	if r.channelId != nil {
		localVarQueryParams.Add("channel_id", parameterToString(*r.channelId, ""))
	}
	if r.relatedTo != nil {
		localVarQueryParams.Add("related_to", parameterToString(*r.relatedTo, ""))
	}
	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.nsfw != nil {
		localVarQueryParams.Add("nsfw", parameterToString(*r.nsfw, ""))
	}
	if r.freeOnly != nil {
		localVarQueryParams.Add("free_only", parameterToString(*r.freeOnly, ""))
	}
	if r.resolve != nil {
		localVarQueryParams.Add("resolve", parameterToString(*r.resolve, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
