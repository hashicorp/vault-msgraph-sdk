// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelSummary 
type ChannelSummary struct {
    // Count of guests in a channel.
    guestsCount *int32
    // Indicates whether external members are included on the channel.
    hasMembersFromOtherTenants *bool
    // Count of members in a channel.
    membersCount *int32
    // Count of owners in a channel.
    ownersCount *int32
}
// NewChannelSummary instantiates a new channelSummary and sets the default values.
func NewChannelSummary()(*ChannelSummary) {
    m := &ChannelSummary{
    }
    return m
}
// CreateChannelSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["guestsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestsCount(val)
        }
        return nil
    }
    res["hasMembersFromOtherTenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMembersFromOtherTenants(val)
        }
        return nil
    }
    res["membersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembersCount(val)
        }
        return nil
    }
    res["ownersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnersCount(val)
        }
        return nil
    }
    return res
}
// GetGuestsCount gets the guestsCount property value. Count of guests in a channel.
func (m *ChannelSummary) GetGuestsCount()(*int32) {
    return m.guestsCount
}
// GetHasMembersFromOtherTenants gets the hasMembersFromOtherTenants property value. Indicates whether external members are included on the channel.
func (m *ChannelSummary) GetHasMembersFromOtherTenants()(*bool) {
    return m.hasMembersFromOtherTenants
}
// GetMembersCount gets the membersCount property value. Count of members in a channel.
func (m *ChannelSummary) GetMembersCount()(*int32) {
    return m.membersCount
}
// GetOwnersCount gets the ownersCount property value. Count of owners in a channel.
func (m *ChannelSummary) GetOwnersCount()(*int32) {
    return m.ownersCount
}
// Serialize serializes information the current object
func (m *ChannelSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("guestsCount", m.GetGuestsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasMembersFromOtherTenants", m.GetHasMembersFromOtherTenants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("membersCount", m.GetMembersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ownersCount", m.GetOwnersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGuestsCount sets the guestsCount property value. Count of guests in a channel.
func (m *ChannelSummary) SetGuestsCount(value *int32)() {
    m.guestsCount = value
}
// SetHasMembersFromOtherTenants sets the hasMembersFromOtherTenants property value. Indicates whether external members are included on the channel.
func (m *ChannelSummary) SetHasMembersFromOtherTenants(value *bool)() {
    m.hasMembersFromOtherTenants = value
}
// SetMembersCount sets the membersCount property value. Count of members in a channel.
func (m *ChannelSummary) SetMembersCount(value *int32)() {
    m.membersCount = value
}
// SetOwnersCount sets the ownersCount property value. Count of owners in a channel.
func (m *ChannelSummary) SetOwnersCount(value *int32)() {
    m.ownersCount = value
}
// ChannelSummaryable 
type ChannelSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGuestsCount()(*int32)
    GetHasMembersFromOtherTenants()(*bool)
    GetMembersCount()(*int32)
    GetOwnersCount()(*int32)
    SetGuestsCount(value *int32)()
    SetHasMembersFromOtherTenants(value *bool)()
    SetMembersCount(value *int32)()
    SetOwnersCount(value *int32)()
}
