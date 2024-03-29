// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package applications

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody 
type ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody struct {
    // The keyId property
    keyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The proof property
    proof *string
}
// NewItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody instantiates a new ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody and sets the default values.
func NewItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody()(*ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) {
    m := &ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["proof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The keyId property
func (m *ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) GetKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.keyId
}
// GetProof gets the proof property value. The proof property
func (m *ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeyId sets the keyId property value. The keyId property
func (m *ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.keyId = value
}
// SetProof sets the proof property value. The proof property
func (m *ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
// ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBodyable 
type ItemMicrosoftGraphRemoveKeyRemoveKeyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetProof()(*string)
    SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetProof(value *string)()
}
