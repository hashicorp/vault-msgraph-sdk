// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody 
type ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody struct {
    // The passwordCredential property
    passwordCredential i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable
}
// NewItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody instantiates a new ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody and sets the default values.
func NewItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody()(*ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) {
    m := &ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphAddPasswordAddPasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphAddPasswordAddPasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["passwordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable))
        }
        return nil
    }
    return res
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) GetPasswordCredential()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable) {
    return m.passwordCredential
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) SetPasswordCredential(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable)() {
    m.passwordCredential = value
}
// ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBodyable 
type ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPasswordCredential()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable)
    SetPasswordCredential(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable)()
}
