// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BroadcastMeetingCaptionSettings 
type BroadcastMeetingCaptionSettings struct {
    // Indicates whether captions are enabled for this Teams live event.
    isCaptionEnabled *bool
    // The spoken language.
    spokenLanguage *string
    // The translation languages (choose up to 6).
    translationLanguages []string
}
// NewBroadcastMeetingCaptionSettings instantiates a new broadcastMeetingCaptionSettings and sets the default values.
func NewBroadcastMeetingCaptionSettings()(*BroadcastMeetingCaptionSettings) {
    m := &BroadcastMeetingCaptionSettings{
    }
    return m
}
// CreateBroadcastMeetingCaptionSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBroadcastMeetingCaptionSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBroadcastMeetingCaptionSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BroadcastMeetingCaptionSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isCaptionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCaptionEnabled(val)
        }
        return nil
    }
    res["spokenLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpokenLanguage(val)
        }
        return nil
    }
    res["translationLanguages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTranslationLanguages(res)
        }
        return nil
    }
    return res
}
// GetIsCaptionEnabled gets the isCaptionEnabled property value. Indicates whether captions are enabled for this Teams live event.
func (m *BroadcastMeetingCaptionSettings) GetIsCaptionEnabled()(*bool) {
    return m.isCaptionEnabled
}
// GetSpokenLanguage gets the spokenLanguage property value. The spoken language.
func (m *BroadcastMeetingCaptionSettings) GetSpokenLanguage()(*string) {
    return m.spokenLanguage
}
// GetTranslationLanguages gets the translationLanguages property value. The translation languages (choose up to 6).
func (m *BroadcastMeetingCaptionSettings) GetTranslationLanguages()([]string) {
    return m.translationLanguages
}
// Serialize serializes information the current object
func (m *BroadcastMeetingCaptionSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isCaptionEnabled", m.GetIsCaptionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("spokenLanguage", m.GetSpokenLanguage())
        if err != nil {
            return err
        }
    }
    if m.GetTranslationLanguages() != nil {
        err := writer.WriteCollectionOfStringValues("translationLanguages", m.GetTranslationLanguages())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsCaptionEnabled sets the isCaptionEnabled property value. Indicates whether captions are enabled for this Teams live event.
func (m *BroadcastMeetingCaptionSettings) SetIsCaptionEnabled(value *bool)() {
    m.isCaptionEnabled = value
}
// SetSpokenLanguage sets the spokenLanguage property value. The spoken language.
func (m *BroadcastMeetingCaptionSettings) SetSpokenLanguage(value *string)() {
    m.spokenLanguage = value
}
// SetTranslationLanguages sets the translationLanguages property value. The translation languages (choose up to 6).
func (m *BroadcastMeetingCaptionSettings) SetTranslationLanguages(value []string)() {
    m.translationLanguages = value
}
// BroadcastMeetingCaptionSettingsable 
type BroadcastMeetingCaptionSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsCaptionEnabled()(*bool)
    GetSpokenLanguage()(*string)
    GetTranslationLanguages()([]string)
    SetIsCaptionEnabled(value *bool)()
    SetSpokenLanguage(value *string)()
    SetTranslationLanguages(value []string)()
}
