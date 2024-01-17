// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageHistoryItem 
type ChatMessageHistoryItem struct {
    // The actions property
    actions *ChatMessageActions
    // The date and time when the message was modified.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reaction property
    reaction ChatMessageReactionable
}
// NewChatMessageHistoryItem instantiates a new chatMessageHistoryItem and sets the default values.
func NewChatMessageHistoryItem()(*ChatMessageHistoryItem) {
    m := &ChatMessageHistoryItem{
    }
    return m
}
// CreateChatMessageHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageHistoryItem(), nil
}
// GetActions gets the actions property value. The actions property
func (m *ChatMessageHistoryItem) GetActions()(*ChatMessageActions) {
    return m.actions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActions(val.(*ChatMessageActions))
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["reaction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageReactionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReaction(val.(ChatMessageReactionable))
        }
        return nil
    }
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The date and time when the message was modified.
func (m *ChatMessageHistoryItem) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetReaction gets the reaction property value. The reaction property
func (m *ChatMessageHistoryItem) GetReaction()(ChatMessageReactionable) {
    return m.reaction
}
// Serialize serializes information the current object
func (m *ChatMessageHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := (*m.GetActions()).String()
        err := writer.WriteStringValue("actions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reaction", m.GetReaction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The actions property
func (m *ChatMessageHistoryItem) SetActions(value *ChatMessageActions)() {
    m.actions = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The date and time when the message was modified.
func (m *ChatMessageHistoryItem) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetReaction sets the reaction property value. The reaction property
func (m *ChatMessageHistoryItem) SetReaction(value ChatMessageReactionable)() {
    m.reaction = value
}
// ChatMessageHistoryItemable 
type ChatMessageHistoryItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()(*ChatMessageActions)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReaction()(ChatMessageReactionable)
    SetActions(value *ChatMessageActions)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReaction(value ChatMessageReactionable)()
}
