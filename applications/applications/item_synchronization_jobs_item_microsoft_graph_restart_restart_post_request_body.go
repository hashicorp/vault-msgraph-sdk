// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody 
type ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody struct {
    // The criteria property
    criteria i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobRestartCriteriaable
}
// NewItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody instantiates a new ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody()(*ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody) {
    m := &ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody{
    }
    return m
}
// CreateItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody(), nil
}
// GetCriteria gets the criteria property value. The criteria property
func (m *ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody) GetCriteria()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobRestartCriteriaable) {
    return m.criteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateSynchronizationJobRestartCriteriaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria(val.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobRestartCriteriaable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("criteria", m.GetCriteria())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCriteria sets the criteria property value. The criteria property
func (m *ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBody) SetCriteria(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobRestartCriteriaable)() {
    m.criteria = value
}
// ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBodyable 
type ItemSynchronizationJobsItemMicrosoftGraphRestartRestartPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCriteria()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobRestartCriteriaable)
    SetCriteria(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.SynchronizationJobRestartCriteriaable)()
}
