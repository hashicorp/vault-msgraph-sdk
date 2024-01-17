// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BroadcastMeetingSettings 
type BroadcastMeetingSettings struct {
    // The allowedAudience property
    allowedAudience *BroadcastMeetingAudience
    // The captions property
    captions BroadcastMeetingCaptionSettingsable
    // Indicates whether attendee report is enabled for this Teams live event. Default value is false.
    isAttendeeReportEnabled *bool
    // Indicates whether Q&A is enabled for this Teams live event. Default value is false.
    isQuestionAndAnswerEnabled *bool
    // Indicates whether recording is enabled for this Teams live event. Default value is false.
    isRecordingEnabled *bool
    // Indicates whether video on demand is enabled for this Teams live event. Default value is false.
    isVideoOnDemandEnabled *bool
}
// NewBroadcastMeetingSettings instantiates a new broadcastMeetingSettings and sets the default values.
func NewBroadcastMeetingSettings()(*BroadcastMeetingSettings) {
    m := &BroadcastMeetingSettings{
    }
    return m
}
// CreateBroadcastMeetingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBroadcastMeetingSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBroadcastMeetingSettings(), nil
}
// GetAllowedAudience gets the allowedAudience property value. The allowedAudience property
func (m *BroadcastMeetingSettings) GetAllowedAudience()(*BroadcastMeetingAudience) {
    return m.allowedAudience
}
// GetCaptions gets the captions property value. The captions property
func (m *BroadcastMeetingSettings) GetCaptions()(BroadcastMeetingCaptionSettingsable) {
    return m.captions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BroadcastMeetingSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedAudience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBroadcastMeetingAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAudience(val.(*BroadcastMeetingAudience))
        }
        return nil
    }
    res["captions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBroadcastMeetingCaptionSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaptions(val.(BroadcastMeetingCaptionSettingsable))
        }
        return nil
    }
    res["isAttendeeReportEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAttendeeReportEnabled(val)
        }
        return nil
    }
    res["isQuestionAndAnswerEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsQuestionAndAnswerEnabled(val)
        }
        return nil
    }
    res["isRecordingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRecordingEnabled(val)
        }
        return nil
    }
    res["isVideoOnDemandEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVideoOnDemandEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsAttendeeReportEnabled gets the isAttendeeReportEnabled property value. Indicates whether attendee report is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsAttendeeReportEnabled()(*bool) {
    return m.isAttendeeReportEnabled
}
// GetIsQuestionAndAnswerEnabled gets the isQuestionAndAnswerEnabled property value. Indicates whether Q&A is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsQuestionAndAnswerEnabled()(*bool) {
    return m.isQuestionAndAnswerEnabled
}
// GetIsRecordingEnabled gets the isRecordingEnabled property value. Indicates whether recording is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsRecordingEnabled()(*bool) {
    return m.isRecordingEnabled
}
// GetIsVideoOnDemandEnabled gets the isVideoOnDemandEnabled property value. Indicates whether video on demand is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsVideoOnDemandEnabled()(*bool) {
    return m.isVideoOnDemandEnabled
}
// Serialize serializes information the current object
func (m *BroadcastMeetingSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedAudience() != nil {
        cast := (*m.GetAllowedAudience()).String()
        err := writer.WriteStringValue("allowedAudience", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("captions", m.GetCaptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAttendeeReportEnabled", m.GetIsAttendeeReportEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQuestionAndAnswerEnabled", m.GetIsQuestionAndAnswerEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRecordingEnabled", m.GetIsRecordingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVideoOnDemandEnabled", m.GetIsVideoOnDemandEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedAudience sets the allowedAudience property value. The allowedAudience property
func (m *BroadcastMeetingSettings) SetAllowedAudience(value *BroadcastMeetingAudience)() {
    m.allowedAudience = value
}
// SetCaptions sets the captions property value. The captions property
func (m *BroadcastMeetingSettings) SetCaptions(value BroadcastMeetingCaptionSettingsable)() {
    m.captions = value
}
// SetIsAttendeeReportEnabled sets the isAttendeeReportEnabled property value. Indicates whether attendee report is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsAttendeeReportEnabled(value *bool)() {
    m.isAttendeeReportEnabled = value
}
// SetIsQuestionAndAnswerEnabled sets the isQuestionAndAnswerEnabled property value. Indicates whether Q&A is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsQuestionAndAnswerEnabled(value *bool)() {
    m.isQuestionAndAnswerEnabled = value
}
// SetIsRecordingEnabled sets the isRecordingEnabled property value. Indicates whether recording is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsRecordingEnabled(value *bool)() {
    m.isRecordingEnabled = value
}
// SetIsVideoOnDemandEnabled sets the isVideoOnDemandEnabled property value. Indicates whether video on demand is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsVideoOnDemandEnabled(value *bool)() {
    m.isVideoOnDemandEnabled = value
}
// BroadcastMeetingSettingsable 
type BroadcastMeetingSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedAudience()(*BroadcastMeetingAudience)
    GetCaptions()(BroadcastMeetingCaptionSettingsable)
    GetIsAttendeeReportEnabled()(*bool)
    GetIsQuestionAndAnswerEnabled()(*bool)
    GetIsRecordingEnabled()(*bool)
    GetIsVideoOnDemandEnabled()(*bool)
    SetAllowedAudience(value *BroadcastMeetingAudience)()
    SetCaptions(value BroadcastMeetingCaptionSettingsable)()
    SetIsAttendeeReportEnabled(value *bool)()
    SetIsQuestionAndAnswerEnabled(value *bool)()
    SetIsRecordingEnabled(value *bool)()
    SetIsVideoOnDemandEnabled(value *bool)()
}
