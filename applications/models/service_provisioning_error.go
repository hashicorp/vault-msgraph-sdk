package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceProvisioningError 
type ServiceProvisioningError struct {
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The isResolved property
    isResolved *bool
    // The serviceInstance property
    serviceInstance *string
}
// NewServiceProvisioningError instantiates a new serviceProvisioningError and sets the default values.
func NewServiceProvisioningError()(*ServiceProvisioningError) {
    m := &ServiceProvisioningError{
    }
    return m
}
// CreateServiceProvisioningErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceProvisioningErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceProvisioningError(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ServiceProvisioningError) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceProvisioningError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isResolved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsResolved(val)
        }
        return nil
    }
    res["serviceInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceInstance(val)
        }
        return nil
    }
    return res
}
// GetIsResolved gets the isResolved property value. The isResolved property
func (m *ServiceProvisioningError) GetIsResolved()(*bool) {
    return m.isResolved
}
// GetServiceInstance gets the serviceInstance property value. The serviceInstance property
func (m *ServiceProvisioningError) GetServiceInstance()(*string) {
    return m.serviceInstance
}
// Serialize serializes information the current object
func (m *ServiceProvisioningError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isResolved", m.GetIsResolved())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serviceInstance", m.GetServiceInstance())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ServiceProvisioningError) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetIsResolved sets the isResolved property value. The isResolved property
func (m *ServiceProvisioningError) SetIsResolved(value *bool)() {
    m.isResolved = value
}
// SetServiceInstance sets the serviceInstance property value. The serviceInstance property
func (m *ServiceProvisioningError) SetServiceInstance(value *string)() {
    m.serviceInstance = value
}
// ServiceProvisioningErrorable 
type ServiceProvisioningErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsResolved()(*bool)
    GetServiceInstance()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsResolved(value *bool)()
    SetServiceInstance(value *string)()
}
