// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applicationtemplates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphInstantiateInstantiatePostRequestBody 
type ItemMicrosoftGraphInstantiateInstantiatePostRequestBody struct {
    // The displayName property
    displayName *string
}
// NewItemMicrosoftGraphInstantiateInstantiatePostRequestBody instantiates a new ItemMicrosoftGraphInstantiateInstantiatePostRequestBody and sets the default values.
func NewItemMicrosoftGraphInstantiateInstantiatePostRequestBody()(*ItemMicrosoftGraphInstantiateInstantiatePostRequestBody) {
    m := &ItemMicrosoftGraphInstantiateInstantiatePostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphInstantiateInstantiatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphInstantiateInstantiatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphInstantiateInstantiatePostRequestBody(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ItemMicrosoftGraphInstantiateInstantiatePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphInstantiateInstantiatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphInstantiateInstantiatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ItemMicrosoftGraphInstantiateInstantiatePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// ItemMicrosoftGraphInstantiateInstantiatePostRequestBodyable 
type ItemMicrosoftGraphInstantiateInstantiatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    SetDisplayName(value *string)()
}
