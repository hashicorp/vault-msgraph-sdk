// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody 
type ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody struct {
    // The credentials property
    credentials []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable
}
// NewItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody instantiates a new ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody and sets the default values.
func NewItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody()(*ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody) {
    m := &ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody{
    }
    return m
}
// CreateItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody(), nil
}
// GetCredentials gets the credentials property value. The credentials property
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody) GetCredentials()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateSynchronizationSecretKeyStringValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)
                }
            }
            m.SetCredentials(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredentials()))
        for i, v := range m.GetCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("credentials", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCredentials sets the credentials property value. The credentials property
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBody) SetCredentials(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)() {
    m.credentials = value
}
// ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBodyable 
type ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCredentials()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)
    SetCredentials(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)()
}
