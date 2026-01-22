// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody 
type ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody struct {
    // The verifiedPublisherId property
    verifiedPublisherId *string
}
// NewItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody instantiates a new ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody and sets the default values.
func NewItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody()(*ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) {
    m := &ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["verifiedPublisherId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisherId(val)
        }
        return nil
    }
    return res
}
// GetVerifiedPublisherId gets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) GetVerifiedPublisherId()(*string) {
    return m.verifiedPublisherId
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("verifiedPublisherId", m.GetVerifiedPublisherId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVerifiedPublisherId sets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBody) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
// ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBodyable 
type ItemMicrosoftGraphSetVerifiedPublisherSetVerifiedPublisherPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVerifiedPublisherId()(*string)
    SetVerifiedPublisherId(value *string)()
}
