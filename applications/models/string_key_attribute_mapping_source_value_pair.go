// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StringKeyAttributeMappingSourceValuePair 
type StringKeyAttributeMappingSourceValuePair struct {
    // The name of the parameter.
    key *string
    // The value property
    value AttributeMappingSourceable
}
// NewStringKeyAttributeMappingSourceValuePair instantiates a new stringKeyAttributeMappingSourceValuePair and sets the default values.
func NewStringKeyAttributeMappingSourceValuePair()(*StringKeyAttributeMappingSourceValuePair) {
    m := &StringKeyAttributeMappingSourceValuePair{
    }
    return m
}
// CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStringKeyAttributeMappingSourceValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyAttributeMappingSourceValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(AttributeMappingSourceable))
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) GetKey()(*string) {
    return m.key
}
// GetValue gets the value property value. The value property
func (m *StringKeyAttributeMappingSourceValuePair) GetValue()(AttributeMappingSourceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *StringKeyAttributeMappingSourceValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. The name of the parameter.
func (m *StringKeyAttributeMappingSourceValuePair) SetKey(value *string)() {
    m.key = value
}
// SetValue sets the value property value. The value property
func (m *StringKeyAttributeMappingSourceValuePair) SetValue(value AttributeMappingSourceable)() {
    m.value = value
}
// StringKeyAttributeMappingSourceValuePairable 
type StringKeyAttributeMappingSourceValuePairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetValue()(AttributeMappingSourceable)
    SetKey(value *string)()
    SetValue(value AttributeMappingSourceable)()
}
