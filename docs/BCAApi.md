# \BCAApi

All URIs are relative to *https://m.klikbca.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountStatementView**](BCAApi.md#AccountStatementView) | **Post** /accountstmt.do?value(actions)&#x3D;acctstmtview | AccountStatementView
[**BalanceInquiry**](BCAApi.md#BalanceInquiry) | **Post** /balanceinquiry.do | BalanceInquiry
[**Login**](BCAApi.md#Login) | **Post** /authentication.do | Login
[**Logout**](BCAApi.md#Logout) | **Get** /authentication.do?value(actions)&#x3D;logout | Logout



## AccountStatementView

> map[string]interface{} AccountStatementView(ctx, contentType, userAgent, referer, valueActions, cookie, contentLength, r1, value28D129, value28startDt29, value28startMt29, value28startYr29, value28endDt29, value28endMt29, value28endYr29)

AccountStatementView

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentType** | **string**|  | 
**userAgent** | **string**|  | 
**referer** | **string**|  | 
**valueActions** | **string**|  | 
**cookie** | **string**|  | 
**contentLength** | **string**|  | 
**r1** | **string**|  | 
**value28D129** | **string**|  | 
**value28startDt29** | **string**|  | 
**value28startMt29** | **string**|  | 
**value28startYr29** | **string**|  | 
**value28endDt29** | **string**|  | 
**value28endMt29** | **string**|  | 
**value28endYr29** | **string**|  | 

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

> map[string]interface{} BalanceInquiry(ctx, contentType, userAgent, referer, cookie, contentLength)

BalanceInquiry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentType** | **string**|  | 
**userAgent** | **string**|  | 
**referer** | **string**|  | 
**cookie** | **string**|  | 
**contentLength** | **string**|  | 

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

> map[string]interface{} Login(ctx, contentType, userAgent, referer, cookie, contentLength, value28userId29, value28pswd29, value28Submit29, value28actions29, value28userIp29, userIp, value28mobile29, value28browserInfo29, mobile)

Login

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentType** | **string**|  | 
**userAgent** | **string**|  | 
**referer** | **string**|  | 
**cookie** | **string**|  | 
**contentLength** | **int32**|  | 
**value28userId29** | **string**|  | 
**value28pswd29** | **string**|  | 
**value28Submit29** | **string**|  | 
**value28actions29** | **string**|  | 
**value28userIp29** | **string**|  | 
**userIp** | **string**|  | 
**value28mobile29** | **string**|  | 
**value28browserInfo29** | **string**|  | 
**mobile** | **string**|  | 

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

> map[string]interface{} Logout(ctx, userAgent, referer, valueActions, cookie)

Logout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAgent** | **string**|  | 
**referer** | **string**|  | 
**valueActions** | **string**|  | 
**cookie** | **string**|  | 

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

