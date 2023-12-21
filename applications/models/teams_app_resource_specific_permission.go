package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppResourceSpecificPermission 
type TeamsAppResourceSpecificPermission struct {
    // The permissionType property
    permissionType *TeamsAppResourceSpecificPermissionType
    // The name of the resource-specific permission.
    permissionValue *string
}
// NewTeamsAppResourceSpecificPermission instantiates a new teamsAppResourceSpecificPermission and sets the default values.
func NewTeamsAppResourceSpecificPermission()(*TeamsAppResourceSpecificPermission) {
    m := &TeamsAppResourceSpecificPermission{
    }
    return m
}
// CreateTeamsAppResourceSpecificPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppResourceSpecificPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppResourceSpecificPermission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppResourceSpecificPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["permissionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppResourceSpecificPermissionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionType(val.(*TeamsAppResourceSpecificPermissionType))
        }
        return nil
    }
    res["permissionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionValue(val)
        }
        return nil
    }
    return res
}
// GetPermissionType gets the permissionType property value. The permissionType property
func (m *TeamsAppResourceSpecificPermission) GetPermissionType()(*TeamsAppResourceSpecificPermissionType) {
    return m.permissionType
}
// GetPermissionValue gets the permissionValue property value. The name of the resource-specific permission.
func (m *TeamsAppResourceSpecificPermission) GetPermissionValue()(*string) {
    return m.permissionValue
}
// Serialize serializes information the current object
func (m *TeamsAppResourceSpecificPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPermissionType() != nil {
        cast := (*m.GetPermissionType()).String()
        err := writer.WriteStringValue("permissionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("permissionValue", m.GetPermissionValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPermissionType sets the permissionType property value. The permissionType property
func (m *TeamsAppResourceSpecificPermission) SetPermissionType(value *TeamsAppResourceSpecificPermissionType)() {
    m.permissionType = value
}
// SetPermissionValue sets the permissionValue property value. The name of the resource-specific permission.
func (m *TeamsAppResourceSpecificPermission) SetPermissionValue(value *string)() {
    m.permissionValue = value
}
// TeamsAppResourceSpecificPermissionable 
type TeamsAppResourceSpecificPermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPermissionType()(*TeamsAppResourceSpecificPermissionType)
    GetPermissionValue()(*string)
    SetPermissionType(value *TeamsAppResourceSpecificPermissionType)()
    SetPermissionValue(value *string)()
}
