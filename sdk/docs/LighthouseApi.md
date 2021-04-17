# \LighthouseApi

All URIs are relative to *https://lighthouse.lbry.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAutoComplete**](LighthouseApi.md#SearchAutoComplete) | **Get** /autocomplete | 搜尋自動完成建議
[**SearchByString**](LighthouseApi.md#SearchByString) | **Get** /search | 字串搜尋



## SearchAutoComplete

> []string SearchAutoComplete(ctx).S(s).Size(size).From(from).Nsfw(nsfw).Execute()

搜尋自動完成建議

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    s := "s_example" // string | 關鍵字
    size := int32(56) // int32 | array size (回傳數量) (optional)
    from := int32(56) // int32 | 第幾頁 (optional) (default to 0)
    nsfw := true // bool | 成人模式 false:close (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LighthouseApi.SearchAutoComplete(context.Background()).S(s).Size(size).From(from).Nsfw(nsfw).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LighthouseApi.SearchAutoComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAutoComplete`: []string
    fmt.Fprintf(os.Stdout, "Response from `LighthouseApi.SearchAutoComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAutoCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string** | 關鍵字 | 
 **size** | **int32** | array size (回傳數量) | 
 **from** | **int32** | 第幾頁 | [default to 0]
 **nsfw** | **bool** | 成人模式 false:close | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: array

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchByString

> []SearchResponse SearchByString(ctx).S(s).Size(size).From(from).Channel(channel).ChannelId(channelId).RelatedTo(relatedTo).SortBy(sortBy).Include(include).Nsfw(nsfw).FreeOnly(freeOnly).Resolve(resolve).Execute()

字串搜尋

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    s := "s_example" // string | 關鍵字
    size := int32(56) // int32 | array size (回傳數量) (optional)
    from := int32(56) // int32 | 第幾頁 (optional) (default to 0)
    channel := "channel_example" // string | 搜尋包含這個關鍵字的頻道中的內容 (optional)
    channelId := "channelId_example" // string | 測試只會回傳空集合 20210117 (optional)
    relatedTo := "relatedTo_example" // string | 與...影片相關 參數為claimId (optional)
    sortBy := "sortBy_example" // string | (開發中)參數release_time https://github.com/lbryio/lighthouse/issues/12 (optional)
    include := "include_example" // string | https://github.com/lbryio/lighthouse/issues/12 (optional)
    nsfw := true // bool | 成人模式 false:close (optional) (default to false)
    freeOnly := true // bool | 只搜尋免費內容 (optional) (default to true)
    resolve := true // bool | 顯示完整資訊 (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LighthouseApi.SearchByString(context.Background()).S(s).Size(size).From(from).Channel(channel).ChannelId(channelId).RelatedTo(relatedTo).SortBy(sortBy).Include(include).Nsfw(nsfw).FreeOnly(freeOnly).Resolve(resolve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LighthouseApi.SearchByString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchByString`: []SearchResponse
    fmt.Fprintf(os.Stdout, "Response from `LighthouseApi.SearchByString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchByStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s** | **string** | 關鍵字 | 
 **size** | **int32** | array size (回傳數量) | 
 **from** | **int32** | 第幾頁 | [default to 0]
 **channel** | **string** | 搜尋包含這個關鍵字的頻道中的內容 | 
 **channelId** | **string** | 測試只會回傳空集合 20210117 | 
 **relatedTo** | **string** | 與...影片相關 參數為claimId | 
 **sortBy** | **string** | (開發中)參數release_time https://github.com/lbryio/lighthouse/issues/12 | 
 **include** | **string** | https://github.com/lbryio/lighthouse/issues/12 | 
 **nsfw** | **bool** | 成人模式 false:close | [default to false]
 **freeOnly** | **bool** | 只搜尋免費內容 | [default to true]
 **resolve** | **bool** | 顯示完整資訊 | [default to false]

### Return type

[**[]SearchResponse**](SearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

