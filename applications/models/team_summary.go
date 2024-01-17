// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamSummary 
type TeamSummary struct {
    // Count of guests in a team.
    guestsCount *int32
    // Count of members in a team.
    membersCount *int32
    // Count of owners in a team.
    ownersCount *int32
}
// NewTeamSummary instantiates a new teamSummary and sets the default values.
func NewTeamSummary()(*TeamSummary) {
    m := &TeamSummary{
    }
    return m
}
// CreateTeamSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetGuestsCount gets the guestsCount property value. Count of guests in a team.
func (m *TeamSummary) GetGuestsCount()(*int32) {
    return m.guestsCount
}
// GetMembersCount gets the membersCount property value. Count of members in a team.
func (m *TeamSummary) GetMembersCount()(*int32) {
    return m.membersCount
}
// GetOwnersCount gets the ownersCount property value. Count of owners in a team.
func (m *TeamSummary) GetOwnersCount()(*int32) {
    return m.ownersCount
}
// Serialize serializes information the current object
func (m *TeamSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("guestsCount", m.GetGuestsCount())
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
// SetGuestsCount sets the guestsCount property value. Count of guests in a team.
func (m *TeamSummary) SetGuestsCount(value *int32)() {
    m.guestsCount = value
}
// SetMembersCount sets the membersCount property value. Count of members in a team.
func (m *TeamSummary) SetMembersCount(value *int32)() {
    m.membersCount = value
}
// SetOwnersCount sets the ownersCount property value. Count of owners in a team.
func (m *TeamSummary) SetOwnersCount(value *int32)() {
    m.ownersCount = value
}
// TeamSummaryable 
type TeamSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGuestsCount()(*int32)
    GetMembersCount()(*int32)
    GetOwnersCount()(*int32)
    SetGuestsCount(value *int32)()
    SetMembersCount(value *int32)()
    SetOwnersCount(value *int32)()
}
