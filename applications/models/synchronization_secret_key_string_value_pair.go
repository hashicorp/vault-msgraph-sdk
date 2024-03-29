// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationSecretKeyStringValuePair 
type SynchronizationSecretKeyStringValuePair struct {
    // The key property
    key *SynchronizationSecret
    // The value of the secret.
    value *string
}
// NewSynchronizationSecretKeyStringValuePair instantiates a new synchronizationSecretKeyStringValuePair and sets the default values.
func NewSynchronizationSecretKeyStringValuePair()(*SynchronizationSecretKeyStringValuePair) {
    m := &SynchronizationSecretKeyStringValuePair{
    }
    return m
}
// CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationSecretKeyStringValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationSecretKeyStringValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationSecret)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*SynchronizationSecret))
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
// GetKey gets the key property value. The key property
func (m *SynchronizationSecretKeyStringValuePair) GetKey()(*SynchronizationSecret) {
    return m.key
}
// GetValue gets the value property value. The value of the secret.
func (m *SynchronizationSecretKeyStringValuePair) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *SynchronizationSecretKeyStringValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKey() != nil {
        cast := (*m.GetKey()).String()
        err := writer.WriteStringValue("key", &cast)
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
// SetKey sets the key property value. The key property
func (m *SynchronizationSecretKeyStringValuePair) SetKey(value *SynchronizationSecret)() {
    m.key = value
}
// SetValue sets the value property value. The value of the secret.
func (m *SynchronizationSecretKeyStringValuePair) SetValue(value *string)() {
    m.value = value
}
// SynchronizationSecretKeyStringValuePairable 
type SynchronizationSecretKeyStringValuePairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*SynchronizationSecret)
    GetValue()(*string)
    SetKey(value *SynchronizationSecret)()
    SetValue(value *string)()
}
