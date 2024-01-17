// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelIdentity 
type ChannelIdentity struct {
    // The identity of the channel in which the message was posted.
    channelId *string
    // The identity of the team in which the message was posted.
    teamId *string
}
// NewChannelIdentity instantiates a new channelIdentity and sets the default values.
func NewChannelIdentity()(*ChannelIdentity) {
    m := &ChannelIdentity{
    }
    return m
}
// CreateChannelIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelIdentity(), nil
}
// GetChannelId gets the channelId property value. The identity of the channel in which the message was posted.
func (m *ChannelIdentity) GetChannelId()(*string) {
    return m.channelId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["channelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelId(val)
        }
        return nil
    }
    res["teamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
// GetTeamId gets the teamId property value. The identity of the team in which the message was posted.
func (m *ChannelIdentity) GetTeamId()(*string) {
    return m.teamId
}
// Serialize serializes information the current object
func (m *ChannelIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("channelId", m.GetChannelId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannelId sets the channelId property value. The identity of the channel in which the message was posted.
func (m *ChannelIdentity) SetChannelId(value *string)() {
    m.channelId = value
}
// SetTeamId sets the teamId property value. The identity of the team in which the message was posted.
func (m *ChannelIdentity) SetTeamId(value *string)() {
    m.teamId = value
}
// ChannelIdentityable 
type ChannelIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelId()(*string)
    GetTeamId()(*string)
    SetChannelId(value *string)()
    SetTeamId(value *string)()
}
