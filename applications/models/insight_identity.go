// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InsightIdentity 
type InsightIdentity struct {
    // The email address of the user who shared the item.
    address *string
    // The display name of the user who shared the item.
    displayName *string
    // The ID of the user who shared the item.
    id *string
}
// NewInsightIdentity instantiates a new insightIdentity and sets the default values.
func NewInsightIdentity()(*InsightIdentity) {
    m := &InsightIdentity{
    }
    return m
}
// CreateInsightIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInsightIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInsightIdentity(), nil
}
// GetAddress gets the address property value. The email address of the user who shared the item.
func (m *InsightIdentity) GetAddress()(*string) {
    return m.address
}
// GetDisplayName gets the displayName property value. The display name of the user who shared the item.
func (m *InsightIdentity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the user who shared the item.
func (m *InsightIdentity) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *InsightIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The email address of the user who shared the item.
func (m *InsightIdentity) SetAddress(value *string)() {
    m.address = value
}
// SetDisplayName sets the displayName property value. The display name of the user who shared the item.
func (m *InsightIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The ID of the user who shared the item.
func (m *InsightIdentity) SetId(value *string)() {
    m.id = value
}
// InsightIdentityable 
type InsightIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    SetAddress(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
}
