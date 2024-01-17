// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StringKeyLongValuePair 
type StringKeyLongValuePair struct {
    // The mapping of the user type from the source system to the target system. For example:User to User - For Microsoft Entra ID to Microsoft Entra ID synchronization worker to user - For Workday to Microsoft Entra synchronization.
    key *string
    // Total number of synchronized objects.
    value *int64
}
// NewStringKeyLongValuePair instantiates a new stringKeyLongValuePair and sets the default values.
func NewStringKeyLongValuePair()(*StringKeyLongValuePair) {
    m := &StringKeyLongValuePair{
    }
    return m
}
// CreateStringKeyLongValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyLongValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStringKeyLongValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyLongValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
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
// GetKey gets the key property value. The mapping of the user type from the source system to the target system. For example:User to User - For Microsoft Entra ID to Microsoft Entra ID synchronization worker to user - For Workday to Microsoft Entra synchronization.
func (m *StringKeyLongValuePair) GetKey()(*string) {
    return m.key
}
// GetValue gets the value property value. Total number of synchronized objects.
func (m *StringKeyLongValuePair) GetValue()(*int64) {
    return m.value
}
// Serialize serializes information the current object
func (m *StringKeyLongValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. The mapping of the user type from the source system to the target system. For example:User to User - For Microsoft Entra ID to Microsoft Entra ID synchronization worker to user - For Workday to Microsoft Entra synchronization.
func (m *StringKeyLongValuePair) SetKey(value *string)() {
    m.key = value
}
// SetValue sets the value property value. Total number of synchronized objects.
func (m *StringKeyLongValuePair) SetValue(value *int64)() {
    m.value = value
}
// StringKeyLongValuePairable 
type StringKeyLongValuePairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetValue()(*int64)
    SetKey(value *string)()
    SetValue(value *int64)()
}
