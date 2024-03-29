// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody 
type ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody struct {
    // The applicationIdentifier property
    applicationIdentifier *string
    // The credentials property
    credentials []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable
    // The templateId property
    templateId *string
    // The useSavedCredentials property
    useSavedCredentials *bool
}
// NewItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody instantiates a new ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody and sets the default values.
func NewItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody()(*ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) {
    m := &ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody{
    }
    return m
}
// CreateItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody(), nil
}
// GetApplicationIdentifier gets the applicationIdentifier property value. The applicationIdentifier property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) GetApplicationIdentifier()(*string) {
    return m.applicationIdentifier
}
// GetCredentials gets the credentials property value. The credentials property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) GetCredentials()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationIdentifier(val)
        }
        return nil
    }
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
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["useSavedCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSavedCredentials(val)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. The templateId property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) GetTemplateId()(*string) {
    return m.templateId
}
// GetUseSavedCredentials gets the useSavedCredentials property value. The useSavedCredentials property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) GetUseSavedCredentials()(*bool) {
    return m.useSavedCredentials
}
// Serialize serializes information the current object
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationIdentifier", m.GetApplicationIdentifier())
        if err != nil {
            return err
        }
    }
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
    {
        err := writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("useSavedCredentials", m.GetUseSavedCredentials())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationIdentifier sets the applicationIdentifier property value. The applicationIdentifier property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) SetApplicationIdentifier(value *string)() {
    m.applicationIdentifier = value
}
// SetCredentials sets the credentials property value. The credentials property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) SetCredentials(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)() {
    m.credentials = value
}
// SetTemplateId sets the templateId property value. The templateId property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) SetTemplateId(value *string)() {
    m.templateId = value
}
// SetUseSavedCredentials sets the useSavedCredentials property value. The useSavedCredentials property
func (m *ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBody) SetUseSavedCredentials(value *bool)() {
    m.useSavedCredentials = value
}
// ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBodyable 
type ItemSynchronizationJobsMicrosoftGraphValidateCredentialsValidateCredentialsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationIdentifier()(*string)
    GetCredentials()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)
    GetTemplateId()(*string)
    GetUseSavedCredentials()(*bool)
    SetApplicationIdentifier(value *string)()
    SetCredentials(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationSecretKeyStringValuePairable)()
    SetTemplateId(value *string)()
    SetUseSavedCredentials(value *bool)()
}
