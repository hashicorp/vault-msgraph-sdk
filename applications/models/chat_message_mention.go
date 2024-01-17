// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageMention 
type ChatMessageMention struct {
    // Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
    id *int32
    // The mentioned property
    mentioned ChatMessageMentionedIdentitySetable
    // String used to represent the mention. For example, a user's display name, a team name.
    mentionText *string
}
// NewChatMessageMention instantiates a new chatMessageMention and sets the default values.
func NewChatMessageMention()(*ChatMessageMention) {
    m := &ChatMessageMention{
    }
    return m
}
// CreateChatMessageMentionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageMentionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageMention(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageMention) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["mentionText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionText(val)
        }
        return nil
    }
    res["mentioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageMentionedIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentioned(val.(ChatMessageMentionedIdentitySetable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
func (m *ChatMessageMention) GetId()(*int32) {
    return m.id
}
// GetMentioned gets the mentioned property value. The mentioned property
func (m *ChatMessageMention) GetMentioned()(ChatMessageMentionedIdentitySetable) {
    return m.mentioned
}
// GetMentionText gets the mentionText property value. String used to represent the mention. For example, a user's display name, a team name.
func (m *ChatMessageMention) GetMentionText()(*string) {
    return m.mentionText
}
// Serialize serializes information the current object
func (m *ChatMessageMention) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mentioned", m.GetMentioned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mentionText", m.GetMentionText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetId sets the id property value. Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
func (m *ChatMessageMention) SetId(value *int32)() {
    m.id = value
}
// SetMentioned sets the mentioned property value. The mentioned property
func (m *ChatMessageMention) SetMentioned(value ChatMessageMentionedIdentitySetable)() {
    m.mentioned = value
}
// SetMentionText sets the mentionText property value. String used to represent the mention. For example, a user's display name, a team name.
func (m *ChatMessageMention) SetMentionText(value *string)() {
    m.mentionText = value
}
// ChatMessageMentionable 
type ChatMessageMentionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetMentioned()(ChatMessageMentionedIdentitySetable)
    GetMentionText()(*string)
    SetId(value *int32)()
    SetMentioned(value ChatMessageMentionedIdentitySetable)()
    SetMentionText(value *string)()
}
