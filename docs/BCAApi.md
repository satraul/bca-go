# \BCAApi

All URIs are relative to *https://m.klikbca.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountStatementView**](BCAApi.md#AccountStatementView) | **Post** /accountstmt.do?value(actions)&#x3D;acctstmtview | AccountStatementView
[**BalanceInquiry**](BCAApi.md#BalanceInquiry) | **Post** /balanceinquiry.do | BalanceInquiry
[**Login**](BCAApi.md#Login) | **Post** /authentication.do | Login
[**Logout**](BCAApi.md#Logout) | **Get** /authentication.do?value(actions)&#x3D;logout | Logout
[**Menu**](BCAApi.md#Menu) | **Post** /accountstmt.do?value(actions)&#x3D;menu | Menu



## AccountStatementView

> map[string]interface{} AccountStatementView(ctx).ContentType(contentType).UserAgent(userAgent).Referer(referer).ValueActions(valueActions).Cookie(cookie).ContentLength(contentLength).R1(r1).Value28D129(value28D129).Value28startDt29(value28startDt29).Value28startMt29(value28startMt29).Value28startYr29(value28startYr29).Value28endDt29(value28endDt29).Value28endMt29(value28endMt29).Value28endYr29(value28endYr29).Execute()

AccountStatementView

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
    contentType := "contentType_example" // string | 
    userAgent := "userAgent_example" // string | 
    referer := "referer_example" // string | 
    valueActions := "valueActions_example" // string | 
    cookie := "cookie_example" // string | 
    contentLength := "contentLength_example" // string | 
    r1 := "r1_example" // string | 
    value28D129 := "value28D129_example" // string | 
    value28startDt29 := "value28startDt29_example" // string | 
    value28startMt29 := "value28startMt29_example" // string | 
    value28startYr29 := "value28startYr29_example" // string | 
    value28endDt29 := "value28endDt29_example" // string | 
    value28endMt29 := "value28endMt29_example" // string | 
    value28endYr29 := "value28endYr29_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BCAApi.AccountStatementView(context.Background(), contentType, userAgent, referer, valueActions, cookie, contentLength, r1, value28D129, value28startDt29, value28startMt29, value28startYr29, value28endDt29, value28endMt29, value28endYr29).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BCAApi.AccountStatementView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountStatementView`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BCAApi.AccountStatementView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountStatementViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **userAgent** | **string** |  | 
 **referer** | **string** |  | 
 **valueActions** | **string** |  | 
 **cookie** | **string** |  | 
 **contentLength** | **string** |  | 
 **r1** | **string** |  | 
 **value28D129** | **string** |  | 
 **value28startDt29** | **string** |  | 
 **value28startMt29** | **string** |  | 
 **value28startYr29** | **string** |  | 
 **value28endDt29** | **string** |  | 
 **value28endMt29** | **string** |  | 
 **value28endYr29** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BalanceInquiry

> map[string]interface{} BalanceInquiry(ctx).ContentType(contentType).UserAgent(userAgent).Referer(referer).Cookie(cookie).ContentLength(contentLength).Execute()

BalanceInquiry

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
    contentType := "contentType_example" // string | 
    userAgent := "userAgent_example" // string | 
    referer := "referer_example" // string | 
    cookie := "cookie_example" // string | 
    contentLength := "contentLength_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BCAApi.BalanceInquiry(context.Background(), contentType, userAgent, referer, cookie, contentLength).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BCAApi.BalanceInquiry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BalanceInquiry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BCAApi.BalanceInquiry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBalanceInquiryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **userAgent** | **string** |  | 
 **referer** | **string** |  | 
 **cookie** | **string** |  | 
 **contentLength** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> map[string]interface{} Login(ctx).ContentType(contentType).UserAgent(userAgent).Referer(referer).Cookie(cookie).ContentLength(contentLength).Value28userId29(value28userId29).Value28pswd29(value28pswd29).Value28Submit29(value28Submit29).Value28actions29(value28actions29).Value28userIp29(value28userIp29).UserIp(userIp).Value28mobile29(value28mobile29).Value28browserInfo29(value28browserInfo29).Mobile(mobile).Execute()

Login

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
    contentType := "contentType_example" // string | 
    userAgent := "userAgent_example" // string | 
    referer := "referer_example" // string | 
    cookie := "cookie_example" // string | 
    contentLength := 987 // int32 | 
    value28userId29 := "value28userId29_example" // string | 
    value28pswd29 := "value28pswd29_example" // string | 
    value28Submit29 := "value28Submit29_example" // string | 
    value28actions29 := "value28actions29_example" // string | 
    value28userIp29 := "value28userIp29_example" // string | 
    userIp := "userIp_example" // string | 
    value28mobile29 := "value28mobile29_example" // string | 
    value28browserInfo29 := "value28browserInfo29_example" // string | 
    mobile := "mobile_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BCAApi.Login(context.Background(), contentType, userAgent, referer, cookie, contentLength, value28userId29, value28pswd29, value28Submit29, value28actions29, value28userIp29, userIp, value28mobile29, value28browserInfo29, mobile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BCAApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BCAApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **userAgent** | **string** |  | 
 **referer** | **string** |  | 
 **cookie** | **string** |  | 
 **contentLength** | **int32** |  | 
 **value28userId29** | **string** |  | 
 **value28pswd29** | **string** |  | 
 **value28Submit29** | **string** |  | 
 **value28actions29** | **string** |  | 
 **value28userIp29** | **string** |  | 
 **userIp** | **string** |  | 
 **value28mobile29** | **string** |  | 
 **value28browserInfo29** | **string** |  | 
 **mobile** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx).UserAgent(userAgent).Referer(referer).ValueActions(valueActions).Cookie(cookie).Execute()

Logout

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
    userAgent := "userAgent_example" // string | 
    referer := "referer_example" // string | 
    valueActions := "valueActions_example" // string | 
    cookie := "cookie_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BCAApi.Logout(context.Background(), userAgent, referer, valueActions, cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BCAApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BCAApi.Logout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **string** |  | 
 **referer** | **string** |  | 
 **valueActions** | **string** |  | 
 **cookie** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Menu

> map[string]interface{} Menu(ctx).ContentType(contentType).UserAgent(userAgent).Referer(referer).ValueActions(valueActions).Cookie(cookie).ContentLength(contentLength).Execute()

Menu

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
    contentType := "contentType_example" // string | 
    userAgent := "userAgent_example" // string | 
    referer := "referer_example" // string | 
    valueActions := "valueActions_example" // string | 
    cookie := "cookie_example" // string | 
    contentLength := "contentLength_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BCAApi.Menu(context.Background(), contentType, userAgent, referer, valueActions, cookie, contentLength).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BCAApi.Menu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Menu`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BCAApi.Menu`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMenuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **userAgent** | **string** |  | 
 **referer** | **string** |  | 
 **valueActions** | **string** |  | 
 **cookie** | **string** |  | 
 **contentLength** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

