// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxSettings 
type MailboxSettings struct {
    // Folder ID of an archive folder for the user.
    archiveFolder *string
    // The automaticRepliesSetting property
    automaticRepliesSetting AutomaticRepliesSettingable
    // The date format for the user's mailbox.
    dateFormat *string
    // The delegateMeetingMessageDeliveryOptions property
    delegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions
    // The language property
    language LocaleInfoable
    // The time format for the user's mailbox.
    timeFormat *string
    // The default time zone for the user's mailbox.
    timeZone *string
    // The userPurpose property
    userPurpose *UserPurpose
    // The workingHours property
    workingHours WorkingHoursable
}
// NewMailboxSettings instantiates a new mailboxSettings and sets the default values.
func NewMailboxSettings()(*MailboxSettings) {
    m := &MailboxSettings{
    }
    return m
}
// CreateMailboxSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailboxSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxSettings(), nil
}
// GetArchiveFolder gets the archiveFolder property value. Folder ID of an archive folder for the user.
func (m *MailboxSettings) GetArchiveFolder()(*string) {
    return m.archiveFolder
}
// GetAutomaticRepliesSetting gets the automaticRepliesSetting property value. The automaticRepliesSetting property
func (m *MailboxSettings) GetAutomaticRepliesSetting()(AutomaticRepliesSettingable) {
    return m.automaticRepliesSetting
}
// GetDateFormat gets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) GetDateFormat()(*string) {
    return m.dateFormat
}
// GetDelegateMeetingMessageDeliveryOptions gets the delegateMeetingMessageDeliveryOptions property value. The delegateMeetingMessageDeliveryOptions property
func (m *MailboxSettings) GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions) {
    return m.delegateMeetingMessageDeliveryOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailboxSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archiveFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveFolder(val)
        }
        return nil
    }
    res["automaticRepliesSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutomaticRepliesSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticRepliesSetting(val.(AutomaticRepliesSettingable))
        }
        return nil
    }
    res["dateFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateFormat(val)
        }
        return nil
    }
    res["delegateMeetingMessageDeliveryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegateMeetingMessageDeliveryOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegateMeetingMessageDeliveryOptions(val.(*DelegateMeetingMessageDeliveryOptions))
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["timeFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["userPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPurpose(val.(*UserPurpose))
        }
        return nil
    }
    res["workingHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkingHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkingHours(val.(WorkingHoursable))
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The language property
func (m *MailboxSettings) GetLanguage()(LocaleInfoable) {
    return m.language
}
// GetTimeFormat gets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) GetTimeFormat()(*string) {
    return m.timeFormat
}
// GetTimeZone gets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) GetTimeZone()(*string) {
    return m.timeZone
}
// GetUserPurpose gets the userPurpose property value. The userPurpose property
func (m *MailboxSettings) GetUserPurpose()(*UserPurpose) {
    return m.userPurpose
}
// GetWorkingHours gets the workingHours property value. The workingHours property
func (m *MailboxSettings) GetWorkingHours()(WorkingHoursable) {
    return m.workingHours
}
// Serialize serializes information the current object
func (m *MailboxSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archiveFolder", m.GetArchiveFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("automaticRepliesSetting", m.GetAutomaticRepliesSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dateFormat", m.GetDateFormat())
        if err != nil {
            return err
        }
    }
    if m.GetDelegateMeetingMessageDeliveryOptions() != nil {
        cast := (*m.GetDelegateMeetingMessageDeliveryOptions()).String()
        err := writer.WriteStringValue("delegateMeetingMessageDeliveryOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeFormat", m.GetTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    if m.GetUserPurpose() != nil {
        cast := (*m.GetUserPurpose()).String()
        err := writer.WriteStringValue("userPurpose", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("workingHours", m.GetWorkingHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetArchiveFolder sets the archiveFolder property value. Folder ID of an archive folder for the user.
func (m *MailboxSettings) SetArchiveFolder(value *string)() {
    m.archiveFolder = value
}
// SetAutomaticRepliesSetting sets the automaticRepliesSetting property value. The automaticRepliesSetting property
func (m *MailboxSettings) SetAutomaticRepliesSetting(value AutomaticRepliesSettingable)() {
    m.automaticRepliesSetting = value
}
// SetDateFormat sets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) SetDateFormat(value *string)() {
    m.dateFormat = value
}
// SetDelegateMeetingMessageDeliveryOptions sets the delegateMeetingMessageDeliveryOptions property value. The delegateMeetingMessageDeliveryOptions property
func (m *MailboxSettings) SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)() {
    m.delegateMeetingMessageDeliveryOptions = value
}
// SetLanguage sets the language property value. The language property
func (m *MailboxSettings) SetLanguage(value LocaleInfoable)() {
    m.language = value
}
// SetTimeFormat sets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) SetTimeFormat(value *string)() {
    m.timeFormat = value
}
// SetTimeZone sets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) SetTimeZone(value *string)() {
    m.timeZone = value
}
// SetUserPurpose sets the userPurpose property value. The userPurpose property
func (m *MailboxSettings) SetUserPurpose(value *UserPurpose)() {
    m.userPurpose = value
}
// SetWorkingHours sets the workingHours property value. The workingHours property
func (m *MailboxSettings) SetWorkingHours(value WorkingHoursable)() {
    m.workingHours = value
}
// MailboxSettingsable 
type MailboxSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveFolder()(*string)
    GetAutomaticRepliesSetting()(AutomaticRepliesSettingable)
    GetDateFormat()(*string)
    GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions)
    GetLanguage()(LocaleInfoable)
    GetTimeFormat()(*string)
    GetTimeZone()(*string)
    GetUserPurpose()(*UserPurpose)
    GetWorkingHours()(WorkingHoursable)
    SetArchiveFolder(value *string)()
    SetAutomaticRepliesSetting(value AutomaticRepliesSettingable)()
    SetDateFormat(value *string)()
    SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)()
    SetLanguage(value LocaleInfoable)()
    SetTimeFormat(value *string)()
    SetTimeZone(value *string)()
    SetUserPurpose(value *UserPurpose)()
    SetWorkingHours(value WorkingHoursable)()
}
