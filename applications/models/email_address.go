// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailAddress 
type EmailAddress struct {
    // The email address of the person or entity.
    address *string
    // The display name of the person or entity.
    name *string
}
// NewEmailAddress instantiates a new emailAddress and sets the default values.
func NewEmailAddress()(*EmailAddress) {
    m := &EmailAddress{
    }
    return m
}
// CreateEmailAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailAddress(), nil
}
// GetAddress gets the address property value. The email address of the person or entity.
func (m *EmailAddress) GetAddress()(*string) {
    return m.address
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The display name of the person or entity.
func (m *EmailAddress) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *EmailAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The email address of the person or entity.
func (m *EmailAddress) SetAddress(value *string)() {
    m.address = value
}
// SetName sets the name property value. The display name of the person or entity.
func (m *EmailAddress) SetName(value *string)() {
    m.name = value
}
// EmailAddressable 
type EmailAddressable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(*string)
    GetName()(*string)
    SetAddress(value *string)()
    SetName(value *string)()
}
