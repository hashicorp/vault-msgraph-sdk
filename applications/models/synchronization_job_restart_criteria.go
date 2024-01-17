// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobRestartCriteria 
type SynchronizationJobRestartCriteria struct {
    // The resetScope property
    resetScope *SynchronizationJobRestartScope
}
// NewSynchronizationJobRestartCriteria instantiates a new synchronizationJobRestartCriteria and sets the default values.
func NewSynchronizationJobRestartCriteria()(*SynchronizationJobRestartCriteria) {
    m := &SynchronizationJobRestartCriteria{
    }
    return m
}
// CreateSynchronizationJobRestartCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobRestartCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJobRestartCriteria(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobRestartCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resetScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationJobRestartScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetScope(val.(*SynchronizationJobRestartScope))
        }
        return nil
    }
    return res
}
// GetResetScope gets the resetScope property value. The resetScope property
func (m *SynchronizationJobRestartCriteria) GetResetScope()(*SynchronizationJobRestartScope) {
    return m.resetScope
}
// Serialize serializes information the current object
func (m *SynchronizationJobRestartCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetResetScope() != nil {
        cast := (*m.GetResetScope()).String()
        err := writer.WriteStringValue("resetScope", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResetScope sets the resetScope property value. The resetScope property
func (m *SynchronizationJobRestartCriteria) SetResetScope(value *SynchronizationJobRestartScope)() {
    m.resetScope = value
}
// SynchronizationJobRestartCriteriaable 
type SynchronizationJobRestartCriteriaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResetScope()(*SynchronizationJobRestartScope)
    SetResetScope(value *SynchronizationJobRestartScope)()
}
