# AuthStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthenticated** | **bool** | وضعیت لاگین بودن | 

## Methods

### NewAuthStatusResponse

`func NewAuthStatusResponse(isAuthenticated bool, ) *AuthStatusResponse`

NewAuthStatusResponse instantiates a new AuthStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStatusResponseWithDefaults

`func NewAuthStatusResponseWithDefaults() *AuthStatusResponse`

NewAuthStatusResponseWithDefaults instantiates a new AuthStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthenticated

`func (o *AuthStatusResponse) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *AuthStatusResponse) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *AuthStatusResponse) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


