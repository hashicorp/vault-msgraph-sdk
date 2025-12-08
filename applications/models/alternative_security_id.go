// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlternativeSecurityId 
type AlternativeSecurityId struct {
    // For internal use only.
    identityProvider *string
    // For internal use only.
    key []byte
    // For internal use only.
    typeEscaped *int32
}
// NewAlternativeSecurityId instantiates a new alternativeSecurityId and sets the default values.
func NewAlternativeSecurityId()(*AlternativeSecurityId) {
    m := &AlternativeSecurityId{
    }
    return m
}
// CreateAlternativeSecurityIdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlternativeSecurityIdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlternativeSecurityId(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlternativeSecurityId) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["identityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityProvider(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetIdentityProvider gets the identityProvider property value. For internal use only.
func (m *AlternativeSecurityId) GetIdentityProvider()(*string) {
    return m.identityProvider
}
// GetKey gets the key property value. For internal use only.
func (m *AlternativeSecurityId) GetKey()([]byte) {
    return m.key
}
// GetTypeEscaped gets the type property value. For internal use only.
func (m *AlternativeSecurityId) GetTypeEscaped()(*int32) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AlternativeSecurityId) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("identityProvider", m.GetIdentityProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentityProvider sets the identityProvider property value. For internal use only.
func (m *AlternativeSecurityId) SetIdentityProvider(value *string)() {
    m.identityProvider = value
}
// SetKey sets the key property value. For internal use only.
func (m *AlternativeSecurityId) SetKey(value []byte)() {
    m.key = value
}
// SetTypeEscaped sets the type property value. For internal use only.
func (m *AlternativeSecurityId) SetTypeEscaped(value *int32)() {
    m.typeEscaped = value
}
// AlternativeSecurityIdable 
type AlternativeSecurityIdable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentityProvider()(*string)
    GetKey()([]byte)
    GetTypeEscaped()(*int32)
    SetIdentityProvider(value *string)()
    SetKey(value []byte)()
    SetTypeEscaped(value *int32)()
}
