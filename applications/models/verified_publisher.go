// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifiedPublisher 
type VerifiedPublisher struct {
    // The timestamp when the verified publisher was first added or most recently updated.
    addedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The verified publisher name from the app publisher's Partner Center account.
    displayName *string
    // The ID of the verified publisher from the app publisher's Partner Center account.
    verifiedPublisherId *string
}
// NewVerifiedPublisher instantiates a new verifiedPublisher and sets the default values.
func NewVerifiedPublisher()(*VerifiedPublisher) {
    m := &VerifiedPublisher{
    }
    return m
}
// CreateVerifiedPublisherFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifiedPublisherFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiedPublisher(), nil
}
// GetAddedDateTime gets the addedDateTime property value. The timestamp when the verified publisher was first added or most recently updated.
func (m *VerifiedPublisher) GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.addedDateTime
}
// GetDisplayName gets the displayName property value. The verified publisher name from the app publisher's Partner Center account.
func (m *VerifiedPublisher) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifiedPublisher) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedDateTime(val)
        }
        return nil
    }
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
// GetVerifiedPublisherId gets the verifiedPublisherId property value. The ID of the verified publisher from the app publisher's Partner Center account.
func (m *VerifiedPublisher) GetVerifiedPublisherId()(*string) {
    return m.verifiedPublisherId
}
// Serialize serializes information the current object
func (m *VerifiedPublisher) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("addedDateTime", m.GetAddedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("verifiedPublisherId", m.GetVerifiedPublisherId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedDateTime sets the addedDateTime property value. The timestamp when the verified publisher was first added or most recently updated.
func (m *VerifiedPublisher) SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.addedDateTime = value
}
// SetDisplayName sets the displayName property value. The verified publisher name from the app publisher's Partner Center account.
func (m *VerifiedPublisher) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetVerifiedPublisherId sets the verifiedPublisherId property value. The ID of the verified publisher from the app publisher's Partner Center account.
func (m *VerifiedPublisher) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
// VerifiedPublisherable 
type VerifiedPublisherable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetVerifiedPublisherId()(*string)
    SetAddedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetVerifiedPublisherId(value *string)()
}
