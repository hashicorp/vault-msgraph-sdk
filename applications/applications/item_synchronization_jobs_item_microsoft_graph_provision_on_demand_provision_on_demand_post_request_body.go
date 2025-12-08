// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody 
type ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody struct {
    // The parameters property
    parameters []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable
}
// NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody instantiates a new ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody()(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) {
    m := &ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody{
    }
    return m
}
// CreateItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateSynchronizationJobApplicationParametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable)
                }
            }
            m.SetParameters(res)
        }
        return nil
    }
    return res
}
// GetParameters gets the parameters property value. The parameters property
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) GetParameters()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable) {
    return m.parameters
}
// Serialize serializes information the current object
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParameters sets the parameters property value. The parameters property
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBody) SetParameters(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable)() {
    m.parameters = value
}
// ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyable 
type ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetParameters()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable)
    SetParameters(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobApplicationParametersable)()
}
