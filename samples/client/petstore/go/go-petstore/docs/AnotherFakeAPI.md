# \AnotherFakeAPI

All URIs are relative to *http://petstore.swagger.io:80/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Call123TestSpecialTags**](AnotherFakeAPI.md#Call123TestSpecialTags) | **Patch** /another-fake/dummy | To test special tags



## Call123TestSpecialTags

> Client Call123TestSpecialTags(ctx).UuidTest(uuidTest).Body(body).Execute()

To test special tags



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
    uuidTest := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | to test uuid example value
    body := *openapiclient.NewClient() // Client | client model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnotherFakeAPI.Call123TestSpecialTags(context.Background()).UuidTest(uuidTest).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnotherFakeAPI.Call123TestSpecialTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Call123TestSpecialTags`: Client
    fmt.Fprintf(os.Stdout, "Response from `AnotherFakeAPI.Call123TestSpecialTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCall123TestSpecialTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuidTest** | **string** | to test uuid example value | 
 **body** | [**Client**](Client.md) | client model | 

### Return type

[**Client**](Client.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

