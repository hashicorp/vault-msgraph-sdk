// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppPermissionSet 
type TeamsAppPermissionSet struct {
    // A collection of resource-specific permissions.
    resourceSpecificPermissions []TeamsAppResourceSpecificPermissionable
}
// NewTeamsAppPermissionSet instantiates a new teamsAppPermissionSet and sets the default values.
func NewTeamsAppPermissionSet()(*TeamsAppPermissionSet) {
    m := &TeamsAppPermissionSet{
    }
    return m
}
// CreateTeamsAppPermissionSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppPermissionSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppPermissionSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppPermissionSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["resourceSpecificPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsAppResourceSpecificPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppResourceSpecificPermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamsAppResourceSpecificPermissionable)
                }
            }
            m.SetResourceSpecificPermissions(res)
        }
        return nil
    }
    return res
}
// GetResourceSpecificPermissions gets the resourceSpecificPermissions property value. A collection of resource-specific permissions.
func (m *TeamsAppPermissionSet) GetResourceSpecificPermissions()([]TeamsAppResourceSpecificPermissionable) {
    return m.resourceSpecificPermissions
}
// Serialize serializes information the current object
func (m *TeamsAppPermissionSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetResourceSpecificPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceSpecificPermissions()))
        for i, v := range m.GetResourceSpecificPermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("resourceSpecificPermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceSpecificPermissions sets the resourceSpecificPermissions property value. A collection of resource-specific permissions.
func (m *TeamsAppPermissionSet) SetResourceSpecificPermissions(value []TeamsAppResourceSpecificPermissionable)() {
    m.resourceSpecificPermissions = value
}
// TeamsAppPermissionSetable 
type TeamsAppPermissionSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceSpecificPermissions()([]TeamsAppResourceSpecificPermissionable)
    SetResourceSpecificPermissions(value []TeamsAppResourceSpecificPermissionable)()
}
