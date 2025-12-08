// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Phone 
type Phone struct {
    // The language property
    language *string
    // The phone number.
    number *string
    // The region property
    region *string
    // The type property
    typeEscaped *PhoneType
}
// NewPhone instantiates a new phone and sets the default values.
func NewPhone()(*Phone) {
    m := &Phone{
    }
    return m
}
// CreatePhoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhone(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Phone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePhoneType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*PhoneType))
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The language property
func (m *Phone) GetLanguage()(*string) {
    return m.language
}
// GetNumber gets the number property value. The phone number.
func (m *Phone) GetNumber()(*string) {
    return m.number
}
// GetRegion gets the region property value. The region property
func (m *Phone) GetRegion()(*string) {
    return m.region
}
// GetTypeEscaped gets the type property value. The type property
func (m *Phone) GetTypeEscaped()(*PhoneType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Phone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLanguage sets the language property value. The language property
func (m *Phone) SetLanguage(value *string)() {
    m.language = value
}
// SetNumber sets the number property value. The phone number.
func (m *Phone) SetNumber(value *string)() {
    m.number = value
}
// SetRegion sets the region property value. The region property
func (m *Phone) SetRegion(value *string)() {
    m.region = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Phone) SetTypeEscaped(value *PhoneType)() {
    m.typeEscaped = value
}
// Phoneable 
type Phoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLanguage()(*string)
    GetNumber()(*string)
    GetRegion()(*string)
    GetTypeEscaped()(*PhoneType)
    SetLanguage(value *string)()
    SetNumber(value *string)()
    SetRegion(value *string)()
    SetTypeEscaped(value *PhoneType)()
}
