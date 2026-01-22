// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationServicePrincipal 
type ApplicationServicePrincipal struct {
    // The application property
    application Applicationable
    // The servicePrincipal property
    servicePrincipal ServicePrincipalable
}
// NewApplicationServicePrincipal instantiates a new applicationServicePrincipal and sets the default values.
func NewApplicationServicePrincipal()(*ApplicationServicePrincipal) {
    m := &ApplicationServicePrincipal{
    }
    return m
}
// CreateApplicationServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationServicePrincipal(), nil
}
// GetApplication gets the application property value. The application property
func (m *ApplicationServicePrincipal) GetApplication()(Applicationable) {
    return m.application
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Applicationable))
        }
        return nil
    }
    res["servicePrincipal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipal(val.(ServicePrincipalable))
        }
        return nil
    }
    return res
}
// GetServicePrincipal gets the servicePrincipal property value. The servicePrincipal property
func (m *ApplicationServicePrincipal) GetServicePrincipal()(ServicePrincipalable) {
    return m.servicePrincipal
}
// Serialize serializes information the current object
func (m *ApplicationServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("servicePrincipal", m.GetServicePrincipal())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The application property
func (m *ApplicationServicePrincipal) SetApplication(value Applicationable)() {
    m.application = value
}
// SetServicePrincipal sets the servicePrincipal property value. The servicePrincipal property
func (m *ApplicationServicePrincipal) SetServicePrincipal(value ServicePrincipalable)() {
    m.servicePrincipal = value
}
// ApplicationServicePrincipalable 
type ApplicationServicePrincipalable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(Applicationable)
    GetServicePrincipal()(ServicePrincipalable)
    SetApplication(value Applicationable)()
    SetServicePrincipal(value ServicePrincipalable)()
}
