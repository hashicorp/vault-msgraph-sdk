// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppAuthorization 
type TeamsAppAuthorization struct {
    // The requiredPermissionSet property
    requiredPermissionSet TeamsAppPermissionSetable
}
// NewTeamsAppAuthorization instantiates a new teamsAppAuthorization and sets the default values.
func NewTeamsAppAuthorization()(*TeamsAppAuthorization) {
    m := &TeamsAppAuthorization{
    }
    return m
}
// CreateTeamsAppAuthorizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppAuthorizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppAuthorization(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppAuthorization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["requiredPermissionSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppPermissionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredPermissionSet(val.(TeamsAppPermissionSetable))
        }
        return nil
    }
    return res
}
// GetRequiredPermissionSet gets the requiredPermissionSet property value. The requiredPermissionSet property
func (m *TeamsAppAuthorization) GetRequiredPermissionSet()(TeamsAppPermissionSetable) {
    return m.requiredPermissionSet
}
// Serialize serializes information the current object
func (m *TeamsAppAuthorization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("requiredPermissionSet", m.GetRequiredPermissionSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRequiredPermissionSet sets the requiredPermissionSet property value. The requiredPermissionSet property
func (m *TeamsAppAuthorization) SetRequiredPermissionSet(value TeamsAppPermissionSetable)() {
    m.requiredPermissionSet = value
}
// TeamsAppAuthorizationable 
type TeamsAppAuthorizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequiredPermissionSet()(TeamsAppPermissionSetable)
    SetRequiredPermissionSet(value TeamsAppPermissionSetable)()
}
