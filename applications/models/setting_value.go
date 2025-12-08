// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SettingValue 
type SettingValue struct {
    // Name of the setting (as defined by the groupSettingTemplate).
    name *string
    // Value of the setting.
    value *string
}
// NewSettingValue instantiates a new settingValue and sets the default values.
func NewSettingValue()(*SettingValue) {
    m := &SettingValue{
    }
    return m
}
// CreateSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the setting (as defined by the groupSettingTemplate).
func (m *SettingValue) GetName()(*string) {
    return m.name
}
// GetValue gets the value property value. Value of the setting.
func (m *SettingValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *SettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name of the setting (as defined by the groupSettingTemplate).
func (m *SettingValue) SetName(value *string)() {
    m.name = value
}
// SetValue sets the value property value. Value of the setting.
func (m *SettingValue) SetValue(value *string)() {
    m.value = value
}
// SettingValueable 
type SettingValueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetValue()(*string)
    SetName(value *string)()
    SetValue(value *string)()
}
