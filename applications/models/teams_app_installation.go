// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppInstallation 
type TeamsAppInstallation struct {
    Entity
    // The consentedPermissionSet property
    consentedPermissionSet TeamsAppPermissionSetable
    // The teamsApp property
    teamsApp TeamsAppable
    // The teamsAppDefinition property
    teamsAppDefinition TeamsAppDefinitionable
}
// NewTeamsAppInstallation instantiates a new teamsAppInstallation and sets the default values.
func NewTeamsAppInstallation()(*TeamsAppInstallation) {
    m := &TeamsAppInstallation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppInstallationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppInstallationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppInstallation(), nil
}
// GetConsentedPermissionSet gets the consentedPermissionSet property value. The consentedPermissionSet property
func (m *TeamsAppInstallation) GetConsentedPermissionSet()(TeamsAppPermissionSetable) {
    return m.consentedPermissionSet
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppInstallation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["consentedPermissionSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppPermissionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsentedPermissionSet(val.(TeamsAppPermissionSetable))
        }
        return nil
    }
    res["teamsApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsApp(val.(TeamsAppable))
        }
        return nil
    }
    res["teamsAppDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppDefinition(val.(TeamsAppDefinitionable))
        }
        return nil
    }
    return res
}
// GetTeamsApp gets the teamsApp property value. The teamsApp property
func (m *TeamsAppInstallation) GetTeamsApp()(TeamsAppable) {
    return m.teamsApp
}
// GetTeamsAppDefinition gets the teamsAppDefinition property value. The teamsAppDefinition property
func (m *TeamsAppInstallation) GetTeamsAppDefinition()(TeamsAppDefinitionable) {
    return m.teamsAppDefinition
}
// Serialize serializes information the current object
func (m *TeamsAppInstallation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("consentedPermissionSet", m.GetConsentedPermissionSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsAppDefinition", m.GetTeamsAppDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConsentedPermissionSet sets the consentedPermissionSet property value. The consentedPermissionSet property
func (m *TeamsAppInstallation) SetConsentedPermissionSet(value TeamsAppPermissionSetable)() {
    m.consentedPermissionSet = value
}
// SetTeamsApp sets the teamsApp property value. The teamsApp property
func (m *TeamsAppInstallation) SetTeamsApp(value TeamsAppable)() {
    m.teamsApp = value
}
// SetTeamsAppDefinition sets the teamsAppDefinition property value. The teamsAppDefinition property
func (m *TeamsAppInstallation) SetTeamsAppDefinition(value TeamsAppDefinitionable)() {
    m.teamsAppDefinition = value
}
// TeamsAppInstallationable 
type TeamsAppInstallationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsentedPermissionSet()(TeamsAppPermissionSetable)
    GetTeamsApp()(TeamsAppable)
    GetTeamsAppDefinition()(TeamsAppDefinitionable)
    SetConsentedPermissionSet(value TeamsAppPermissionSetable)()
    SetTeamsApp(value TeamsAppable)()
    SetTeamsAppDefinition(value TeamsAppDefinitionable)()
}
