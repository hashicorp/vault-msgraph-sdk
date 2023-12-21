package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceSpecificPermission 
type ResourceSpecificPermission struct {
    // Describes the level of access that the resource-specific permission represents.
    description *string
    // The display name for the resource-specific permission.
    displayName *string
    // The unique identifier for the resource-specific application permission.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Indicates whether the permission is enabled.
    isEnabled *bool
    // The value of the permission.
    value *string
}
// NewResourceSpecificPermission instantiates a new resourceSpecificPermission and sets the default values.
func NewResourceSpecificPermission()(*ResourceSpecificPermission) {
    m := &ResourceSpecificPermission{
    }
    return m
}
// CreateResourceSpecificPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceSpecificPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceSpecificPermission(), nil
}
// GetDescription gets the description property value. Describes the level of access that the resource-specific permission represents.
func (m *ResourceSpecificPermission) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the resource-specific permission.
func (m *ResourceSpecificPermission) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceSpecificPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier for the resource-specific application permission.
func (m *ResourceSpecificPermission) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the permission is enabled.
func (m *ResourceSpecificPermission) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetValue gets the value property value. The value of the permission.
func (m *ResourceSpecificPermission) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ResourceSpecificPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Describes the level of access that the resource-specific permission represents.
func (m *ResourceSpecificPermission) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the resource-specific permission.
func (m *ResourceSpecificPermission) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. The unique identifier for the resource-specific application permission.
func (m *ResourceSpecificPermission) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the permission is enabled.
func (m *ResourceSpecificPermission) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetValue sets the value property value. The value of the permission.
func (m *ResourceSpecificPermission) SetValue(value *string)() {
    m.value = value
}
// ResourceSpecificPermissionable 
type ResourceSpecificPermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIsEnabled()(*bool)
    GetValue()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIsEnabled(value *bool)()
    SetValue(value *string)()
}
