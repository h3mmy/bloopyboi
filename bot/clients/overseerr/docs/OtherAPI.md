# \OtherAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeywordKeywordIdGet**](OtherAPI.md#KeywordKeywordIdGet) | **Get** /keyword/{keywordId} | Get keyword
[**WatchprovidersMoviesGet**](OtherAPI.md#WatchprovidersMoviesGet) | **Get** /watchproviders/movies | Get watch provider movies
[**WatchprovidersRegionsGet**](OtherAPI.md#WatchprovidersRegionsGet) | **Get** /watchproviders/regions | Get watch provider regions
[**WatchprovidersTvGet**](OtherAPI.md#WatchprovidersTvGet) | **Get** /watchproviders/tv | Get watch provider series



## KeywordKeywordIdGet

> Keyword KeywordKeywordIdGet(ctx, keywordId).Execute()

Get keyword



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    keywordId := float32(1) // float32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherAPI.KeywordKeywordIdGet(context.Background(), keywordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.KeywordKeywordIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `KeywordKeywordIdGet`: Keyword
    fmt.Fprintf(os.Stdout, "Response from `OtherAPI.KeywordKeywordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywordId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeywordKeywordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Keyword**](Keyword.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchprovidersMoviesGet

> []WatchProviderDetails WatchprovidersMoviesGet(ctx).WatchRegion(watchRegion).Execute()

Get watch provider movies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    watchRegion := "US" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherAPI.WatchprovidersMoviesGet(context.Background()).WatchRegion(watchRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.WatchprovidersMoviesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchprovidersMoviesGet`: []WatchProviderDetails
    fmt.Fprintf(os.Stdout, "Response from `OtherAPI.WatchprovidersMoviesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchprovidersMoviesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchRegion** | **string** |  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchprovidersRegionsGet

> []WatchProviderRegion WatchprovidersRegionsGet(ctx).Execute()

Get watch provider regions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherAPI.WatchprovidersRegionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.WatchprovidersRegionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchprovidersRegionsGet`: []WatchProviderRegion
    fmt.Fprintf(os.Stdout, "Response from `OtherAPI.WatchprovidersRegionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWatchprovidersRegionsGetRequest struct via the builder pattern


### Return type

[**[]WatchProviderRegion**](WatchProviderRegion.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchprovidersTvGet

> []WatchProviderDetails WatchprovidersTvGet(ctx).WatchRegion(watchRegion).Execute()

Get watch provider series



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    watchRegion := "US" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OtherAPI.WatchprovidersTvGet(context.Background()).WatchRegion(watchRegion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.WatchprovidersTvGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchprovidersTvGet`: []WatchProviderDetails
    fmt.Fprintf(os.Stdout, "Response from `OtherAPI.WatchprovidersTvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchprovidersTvGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchRegion** | **string** |  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

