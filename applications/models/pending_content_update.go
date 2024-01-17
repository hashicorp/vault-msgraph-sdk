// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingContentUpdate 
type PendingContentUpdate struct {
    // Date and time the pending binary operation was queued in UTC time. Read-only.
    queuedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewPendingContentUpdate instantiates a new pendingContentUpdate and sets the default values.
func NewPendingContentUpdate()(*PendingContentUpdate) {
    m := &PendingContentUpdate{
    }
    return m
}
// CreatePendingContentUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePendingContentUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPendingContentUpdate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PendingContentUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["queuedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueuedDateTime(val)
        }
        return nil
    }
    return res
}
// GetQueuedDateTime gets the queuedDateTime property value. Date and time the pending binary operation was queued in UTC time. Read-only.
func (m *PendingContentUpdate) GetQueuedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.queuedDateTime
}
// Serialize serializes information the current object
func (m *PendingContentUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("queuedDateTime", m.GetQueuedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetQueuedDateTime sets the queuedDateTime property value. Date and time the pending binary operation was queued in UTC time. Read-only.
func (m *PendingContentUpdate) SetQueuedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.queuedDateTime = value
}
// PendingContentUpdateable 
type PendingContentUpdateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetQueuedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetQueuedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
