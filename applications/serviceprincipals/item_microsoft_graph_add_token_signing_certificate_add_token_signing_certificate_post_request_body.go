// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody 
type ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody struct {
    // The displayName property
    displayName *string
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody instantiates a new ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody and sets the default values.
func NewItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody()(*ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) {
    m := &ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBodyable 
type ItemMicrosoftGraphAddTokenSigningCertificateAddTokenSigningCertificatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDisplayName(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
