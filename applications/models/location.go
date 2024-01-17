// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Location 
type Location struct {
    // The address property
    address PhysicalAddressable
    // The coordinates property
    coordinates OutlookGeoCoordinatesable
    // The name associated with the location.
    displayName *string
    // Optional email address of the location.
    locationEmailAddress *string
    // The locationType property
    locationType *LocationType
    // Optional URI representing the location.
    locationUri *string
    // For internal use only.
    uniqueId *string
    // The uniqueIdType property
    uniqueIdType *LocationUniqueIdType
}
// NewLocation instantiates a new location and sets the default values.
func NewLocation()(*Location) {
    m := &Location{
    }
    return m
}
// CreateLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocation(), nil
}
// GetAddress gets the address property value. The address property
func (m *Location) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetCoordinates gets the coordinates property value. The coordinates property
func (m *Location) GetCoordinates()(OutlookGeoCoordinatesable) {
    return m.coordinates
}
// GetDisplayName gets the displayName property value. The name associated with the location.
func (m *Location) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Location) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["coordinates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoordinates(val.(OutlookGeoCoordinatesable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["locationEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationEmailAddress(val)
        }
        return nil
    }
    res["locationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationType(val.(*LocationType))
        }
        return nil
    }
    res["locationUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationUri(val)
        }
        return nil
    }
    res["uniqueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueId(val)
        }
        return nil
    }
    res["uniqueIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationUniqueIdType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueIdType(val.(*LocationUniqueIdType))
        }
        return nil
    }
    return res
}
// GetLocationEmailAddress gets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) GetLocationEmailAddress()(*string) {
    return m.locationEmailAddress
}
// GetLocationType gets the locationType property value. The locationType property
func (m *Location) GetLocationType()(*LocationType) {
    return m.locationType
}
// GetLocationUri gets the locationUri property value. Optional URI representing the location.
func (m *Location) GetLocationUri()(*string) {
    return m.locationUri
}
// GetUniqueId gets the uniqueId property value. For internal use only.
func (m *Location) GetUniqueId()(*string) {
    return m.uniqueId
}
// GetUniqueIdType gets the uniqueIdType property value. The uniqueIdType property
func (m *Location) GetUniqueIdType()(*LocationUniqueIdType) {
    return m.uniqueIdType
}
// Serialize serializes information the current object
func (m *Location) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("coordinates", m.GetCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locationEmailAddress", m.GetLocationEmailAddress())
        if err != nil {
            return err
        }
    }
    if m.GetLocationType() != nil {
        cast := (*m.GetLocationType()).String()
        err := writer.WriteStringValue("locationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locationUri", m.GetLocationUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uniqueId", m.GetUniqueId())
        if err != nil {
            return err
        }
    }
    if m.GetUniqueIdType() != nil {
        cast := (*m.GetUniqueIdType()).String()
        err := writer.WriteStringValue("uniqueIdType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The address property
func (m *Location) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetCoordinates sets the coordinates property value. The coordinates property
func (m *Location) SetCoordinates(value OutlookGeoCoordinatesable)() {
    m.coordinates = value
}
// SetDisplayName sets the displayName property value. The name associated with the location.
func (m *Location) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLocationEmailAddress sets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) SetLocationEmailAddress(value *string)() {
    m.locationEmailAddress = value
}
// SetLocationType sets the locationType property value. The locationType property
func (m *Location) SetLocationType(value *LocationType)() {
    m.locationType = value
}
// SetLocationUri sets the locationUri property value. Optional URI representing the location.
func (m *Location) SetLocationUri(value *string)() {
    m.locationUri = value
}
// SetUniqueId sets the uniqueId property value. For internal use only.
func (m *Location) SetUniqueId(value *string)() {
    m.uniqueId = value
}
// SetUniqueIdType sets the uniqueIdType property value. The uniqueIdType property
func (m *Location) SetUniqueIdType(value *LocationUniqueIdType)() {
    m.uniqueIdType = value
}
// Locationable 
type Locationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PhysicalAddressable)
    GetCoordinates()(OutlookGeoCoordinatesable)
    GetDisplayName()(*string)
    GetLocationEmailAddress()(*string)
    GetLocationType()(*LocationType)
    GetLocationUri()(*string)
    GetUniqueId()(*string)
    GetUniqueIdType()(*LocationUniqueIdType)
    SetAddress(value PhysicalAddressable)()
    SetCoordinates(value OutlookGeoCoordinatesable)()
    SetDisplayName(value *string)()
    SetLocationEmailAddress(value *string)()
    SetLocationType(value *LocationType)()
    SetLocationUri(value *string)()
    SetUniqueId(value *string)()
    SetUniqueIdType(value *LocationUniqueIdType)()
}
