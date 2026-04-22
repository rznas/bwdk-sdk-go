# BusinessAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Address** | **string** |  | 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**CityName** | **string** |  | 
**StateName** | **string** |  | 
**DistrictId** | Pointer to **NullableInt32** |  | [optional] 
**DistrictName** | Pointer to **NullableString** |  | [optional] 
**Longitude** | Pointer to **NullableFloat64** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** |  | [optional] 
**BuildingNumber** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**ReceiverName** | Pointer to **NullableString** |  | [optional] 
**ReceiverPhone** | Pointer to **NullableString** |  | [optional] 
**IsAccurate** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**CreatedAt** | **time.Time** |  | [readonly] 
**ModifiedAt** | **time.Time** |  | [readonly] 

## Methods

### NewBusinessAddress

`func NewBusinessAddress(id int32, address string, cityName string, stateName string, createdAt time.Time, modifiedAt time.Time, ) *BusinessAddress`

NewBusinessAddress instantiates a new BusinessAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessAddressWithDefaults

`func NewBusinessAddressWithDefaults() *BusinessAddress`

NewBusinessAddressWithDefaults instantiates a new BusinessAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BusinessAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *BusinessAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BusinessAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BusinessAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPostalCode

`func (o *BusinessAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *BusinessAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *BusinessAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *BusinessAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *BusinessAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *BusinessAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCityName

`func (o *BusinessAddress) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *BusinessAddress) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *BusinessAddress) SetCityName(v string)`

SetCityName sets CityName field to given value.


### GetStateName

`func (o *BusinessAddress) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *BusinessAddress) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *BusinessAddress) SetStateName(v string)`

SetStateName sets StateName field to given value.


### GetDistrictId

`func (o *BusinessAddress) GetDistrictId() int32`

GetDistrictId returns the DistrictId field if non-nil, zero value otherwise.

### GetDistrictIdOk

`func (o *BusinessAddress) GetDistrictIdOk() (*int32, bool)`

GetDistrictIdOk returns a tuple with the DistrictId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictId

`func (o *BusinessAddress) SetDistrictId(v int32)`

SetDistrictId sets DistrictId field to given value.

### HasDistrictId

`func (o *BusinessAddress) HasDistrictId() bool`

HasDistrictId returns a boolean if a field has been set.

### SetDistrictIdNil

`func (o *BusinessAddress) SetDistrictIdNil(b bool)`

 SetDistrictIdNil sets the value for DistrictId to be an explicit nil

### UnsetDistrictId
`func (o *BusinessAddress) UnsetDistrictId()`

UnsetDistrictId ensures that no value is present for DistrictId, not even an explicit nil
### GetDistrictName

`func (o *BusinessAddress) GetDistrictName() string`

GetDistrictName returns the DistrictName field if non-nil, zero value otherwise.

### GetDistrictNameOk

`func (o *BusinessAddress) GetDistrictNameOk() (*string, bool)`

GetDistrictNameOk returns a tuple with the DistrictName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrictName

`func (o *BusinessAddress) SetDistrictName(v string)`

SetDistrictName sets DistrictName field to given value.

### HasDistrictName

`func (o *BusinessAddress) HasDistrictName() bool`

HasDistrictName returns a boolean if a field has been set.

### SetDistrictNameNil

`func (o *BusinessAddress) SetDistrictNameNil(b bool)`

 SetDistrictNameNil sets the value for DistrictName to be an explicit nil

### UnsetDistrictName
`func (o *BusinessAddress) UnsetDistrictName()`

UnsetDistrictName ensures that no value is present for DistrictName, not even an explicit nil
### GetLongitude

`func (o *BusinessAddress) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *BusinessAddress) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *BusinessAddress) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *BusinessAddress) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *BusinessAddress) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *BusinessAddress) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetLatitude

`func (o *BusinessAddress) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *BusinessAddress) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *BusinessAddress) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *BusinessAddress) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *BusinessAddress) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *BusinessAddress) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetBuildingNumber

`func (o *BusinessAddress) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *BusinessAddress) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *BusinessAddress) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *BusinessAddress) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### SetBuildingNumberNil

`func (o *BusinessAddress) SetBuildingNumberNil(b bool)`

 SetBuildingNumberNil sets the value for BuildingNumber to be an explicit nil

### UnsetBuildingNumber
`func (o *BusinessAddress) UnsetBuildingNumber()`

UnsetBuildingNumber ensures that no value is present for BuildingNumber, not even an explicit nil
### GetUnit

`func (o *BusinessAddress) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BusinessAddress) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BusinessAddress) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BusinessAddress) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *BusinessAddress) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *BusinessAddress) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetReceiverName

`func (o *BusinessAddress) GetReceiverName() string`

GetReceiverName returns the ReceiverName field if non-nil, zero value otherwise.

### GetReceiverNameOk

`func (o *BusinessAddress) GetReceiverNameOk() (*string, bool)`

GetReceiverNameOk returns a tuple with the ReceiverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverName

`func (o *BusinessAddress) SetReceiverName(v string)`

SetReceiverName sets ReceiverName field to given value.

### HasReceiverName

`func (o *BusinessAddress) HasReceiverName() bool`

HasReceiverName returns a boolean if a field has been set.

### SetReceiverNameNil

`func (o *BusinessAddress) SetReceiverNameNil(b bool)`

 SetReceiverNameNil sets the value for ReceiverName to be an explicit nil

### UnsetReceiverName
`func (o *BusinessAddress) UnsetReceiverName()`

UnsetReceiverName ensures that no value is present for ReceiverName, not even an explicit nil
### GetReceiverPhone

`func (o *BusinessAddress) GetReceiverPhone() string`

GetReceiverPhone returns the ReceiverPhone field if non-nil, zero value otherwise.

### GetReceiverPhoneOk

`func (o *BusinessAddress) GetReceiverPhoneOk() (*string, bool)`

GetReceiverPhoneOk returns a tuple with the ReceiverPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverPhone

`func (o *BusinessAddress) SetReceiverPhone(v string)`

SetReceiverPhone sets ReceiverPhone field to given value.

### HasReceiverPhone

`func (o *BusinessAddress) HasReceiverPhone() bool`

HasReceiverPhone returns a boolean if a field has been set.

### SetReceiverPhoneNil

`func (o *BusinessAddress) SetReceiverPhoneNil(b bool)`

 SetReceiverPhoneNil sets the value for ReceiverPhone to be an explicit nil

### UnsetReceiverPhone
`func (o *BusinessAddress) UnsetReceiverPhone()`

UnsetReceiverPhone ensures that no value is present for ReceiverPhone, not even an explicit nil
### GetIsAccurate

`func (o *BusinessAddress) GetIsAccurate() bool`

GetIsAccurate returns the IsAccurate field if non-nil, zero value otherwise.

### GetIsAccurateOk

`func (o *BusinessAddress) GetIsAccurateOk() (*bool, bool)`

GetIsAccurateOk returns a tuple with the IsAccurate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccurate

`func (o *BusinessAddress) SetIsAccurate(v bool)`

SetIsAccurate sets IsAccurate field to given value.

### HasIsAccurate

`func (o *BusinessAddress) HasIsAccurate() bool`

HasIsAccurate returns a boolean if a field has been set.

### GetIsActive

`func (o *BusinessAddress) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BusinessAddress) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BusinessAddress) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BusinessAddress) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BusinessAddress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BusinessAddress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BusinessAddress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *BusinessAddress) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BusinessAddress) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BusinessAddress) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


