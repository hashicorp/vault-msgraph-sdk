// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppIdentity 
type AppIdentity struct {
    // Refers to the unique ID representing application in Microsoft Entra ID.
    appId *string
    // Refers to the application name displayed in the Microsoft Entra admin center.
    displayName *string
    // Refers to the unique ID for the service principal in Microsoft Entra ID.
    servicePrincipalId *string
    // Refers to the Service Principal Name is the Application name in the tenant.
    servicePrincipalName *string
}
// NewAppIdentity instantiates a new appIdentity and sets the default values.
func NewAppIdentity()(*AppIdentity) {
    m := &AppIdentity{
    }
    return m
}
// CreateAppIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppIdentity(), nil
}
// GetAppId gets the appId property value. Refers to the unique ID representing application in Microsoft Entra ID.
func (m *AppIdentity) GetAppId()(*string) {
    return m.appId
}
// GetDisplayName gets the displayName property value. Refers to the application name displayed in the Microsoft Entra admin center.
func (m *AppIdentity) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
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
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    return res
}
// GetServicePrincipalId gets the servicePrincipalId property value. Refers to the unique ID for the service principal in Microsoft Entra ID.
func (m *AppIdentity) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// GetServicePrincipalName gets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// Serialize serializes information the current object
func (m *AppIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
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
        err := writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. Refers to the unique ID representing application in Microsoft Entra ID.
func (m *AppIdentity) SetAppId(value *string)() {
    m.appId = value
}
// SetDisplayName sets the displayName property value. Refers to the application name displayed in the Microsoft Entra admin center.
func (m *AppIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. Refers to the unique ID for the service principal in Microsoft Entra ID.
func (m *AppIdentity) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. Refers to the Service Principal Name is the Application name in the tenant.
func (m *AppIdentity) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// AppIdentityable 
type AppIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetDisplayName()(*string)
    GetServicePrincipalId()(*string)
    GetServicePrincipalName()(*string)
    SetAppId(value *string)()
    SetDisplayName(value *string)()
    SetServicePrincipalId(value *string)()
    SetServicePrincipalName(value *string)()
}
