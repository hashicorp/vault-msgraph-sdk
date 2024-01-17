// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AutomaticRepliesSetting 
type AutomaticRepliesSetting struct {
    // The externalAudience property
    externalAudience *ExternalAudienceScope
    // The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
    externalReplyMessage *string
    // The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
    internalReplyMessage *string
    // The scheduledEndDateTime property
    scheduledEndDateTime DateTimeTimeZoneable
    // The scheduledStartDateTime property
    scheduledStartDateTime DateTimeTimeZoneable
    // The status property
    status *AutomaticRepliesStatus
}
// NewAutomaticRepliesSetting instantiates a new automaticRepliesSetting and sets the default values.
func NewAutomaticRepliesSetting()(*AutomaticRepliesSetting) {
    m := &AutomaticRepliesSetting{
    }
    return m
}
// CreateAutomaticRepliesSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAutomaticRepliesSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutomaticRepliesSetting(), nil
}
// GetExternalAudience gets the externalAudience property value. The externalAudience property
func (m *AutomaticRepliesSetting) GetExternalAudience()(*ExternalAudienceScope) {
    return m.externalAudience
}
// GetExternalReplyMessage gets the externalReplyMessage property value. The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) GetExternalReplyMessage()(*string) {
    return m.externalReplyMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AutomaticRepliesSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalAudience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAudienceScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAudience(val.(*ExternalAudienceScope))
        }
        return nil
    }
    res["externalReplyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalReplyMessage(val)
        }
        return nil
    }
    res["internalReplyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalReplyMessage(val)
        }
        return nil
    }
    res["scheduledEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["scheduledStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutomaticRepliesStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AutomaticRepliesStatus))
        }
        return nil
    }
    return res
}
// GetInternalReplyMessage gets the internalReplyMessage property value. The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) GetInternalReplyMessage()(*string) {
    return m.internalReplyMessage
}
// GetScheduledEndDateTime gets the scheduledEndDateTime property value. The scheduledEndDateTime property
func (m *AutomaticRepliesSetting) GetScheduledEndDateTime()(DateTimeTimeZoneable) {
    return m.scheduledEndDateTime
}
// GetScheduledStartDateTime gets the scheduledStartDateTime property value. The scheduledStartDateTime property
func (m *AutomaticRepliesSetting) GetScheduledStartDateTime()(DateTimeTimeZoneable) {
    return m.scheduledStartDateTime
}
// GetStatus gets the status property value. The status property
func (m *AutomaticRepliesSetting) GetStatus()(*AutomaticRepliesStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *AutomaticRepliesSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExternalAudience() != nil {
        cast := (*m.GetExternalAudience()).String()
        err := writer.WriteStringValue("externalAudience", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalReplyMessage", m.GetExternalReplyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalReplyMessage", m.GetInternalReplyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledEndDateTime", m.GetScheduledEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledStartDateTime", m.GetScheduledStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExternalAudience sets the externalAudience property value. The externalAudience property
func (m *AutomaticRepliesSetting) SetExternalAudience(value *ExternalAudienceScope)() {
    m.externalAudience = value
}
// SetExternalReplyMessage sets the externalReplyMessage property value. The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) SetExternalReplyMessage(value *string)() {
    m.externalReplyMessage = value
}
// SetInternalReplyMessage sets the internalReplyMessage property value. The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) SetInternalReplyMessage(value *string)() {
    m.internalReplyMessage = value
}
// SetScheduledEndDateTime sets the scheduledEndDateTime property value. The scheduledEndDateTime property
func (m *AutomaticRepliesSetting) SetScheduledEndDateTime(value DateTimeTimeZoneable)() {
    m.scheduledEndDateTime = value
}
// SetScheduledStartDateTime sets the scheduledStartDateTime property value. The scheduledStartDateTime property
func (m *AutomaticRepliesSetting) SetScheduledStartDateTime(value DateTimeTimeZoneable)() {
    m.scheduledStartDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *AutomaticRepliesSetting) SetStatus(value *AutomaticRepliesStatus)() {
    m.status = value
}
// AutomaticRepliesSettingable 
type AutomaticRepliesSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExternalAudience()(*ExternalAudienceScope)
    GetExternalReplyMessage()(*string)
    GetInternalReplyMessage()(*string)
    GetScheduledEndDateTime()(DateTimeTimeZoneable)
    GetScheduledStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*AutomaticRepliesStatus)
    SetExternalAudience(value *ExternalAudienceScope)()
    SetExternalReplyMessage(value *string)()
    SetInternalReplyMessage(value *string)()
    SetScheduledEndDateTime(value DateTimeTimeZoneable)()
    SetScheduledStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *AutomaticRepliesStatus)()
}
