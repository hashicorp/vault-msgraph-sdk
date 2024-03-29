// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssociatedTeamInfo 
type AssociatedTeamInfo struct {
    TeamInfo
}
// NewAssociatedTeamInfo instantiates a new associatedTeamInfo and sets the default values.
func NewAssociatedTeamInfo()(*AssociatedTeamInfo) {
    m := &AssociatedTeamInfo{
        TeamInfo: *NewTeamInfo(),
    }
    return m
}
// CreateAssociatedTeamInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssociatedTeamInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssociatedTeamInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssociatedTeamInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamInfo.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AssociatedTeamInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamInfo.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AssociatedTeamInfoable 
type AssociatedTeamInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeamInfoable
}
