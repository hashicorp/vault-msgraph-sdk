// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignedPlan 
type AssignedPlan struct {
    // The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value.
    capabilityStatus *string
    // The name of the service; for example, exchange.
    service *string
    // A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names, see Product names and service plan identifiers for licensing.
    servicePlanId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewAssignedPlan instantiates a new assignedPlan and sets the default values.
func NewAssignedPlan()(*AssignedPlan) {
    m := &AssignedPlan{
    }
    return m
}
// CreateAssignedPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignedPlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedPlan(), nil
}
// GetAssignedDateTime gets the assignedDateTime property value. The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AssignedPlan) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignedDateTime
}
// GetCapabilityStatus gets the capabilityStatus property value. Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value.
func (m *AssignedPlan) GetCapabilityStatus()(*string) {
    return m.capabilityStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedPlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedDateTime(val)
        }
        return nil
    }
    res["capabilityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityStatus(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    return res
}
// GetService gets the service property value. The name of the service; for example, exchange.
func (m *AssignedPlan) GetService()(*string) {
    return m.service
}
// GetServicePlanId gets the servicePlanId property value. A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names, see Product names and service plan identifiers for licensing.
func (m *AssignedPlan) GetServicePlanId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.servicePlanId
}
// Serialize serializes information the current object
func (m *AssignedPlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("assignedDateTime", m.GetAssignedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedDateTime sets the assignedDateTime property value. The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AssignedPlan) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignedDateTime = value
}
// SetCapabilityStatus sets the capabilityStatus property value. Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value.
func (m *AssignedPlan) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
// SetService sets the service property value. The name of the service; for example, exchange.
func (m *AssignedPlan) SetService(value *string)() {
    m.service = value
}
// SetServicePlanId sets the servicePlanId property value. A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names, see Product names and service plan identifiers for licensing.
func (m *AssignedPlan) SetServicePlanId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.servicePlanId = value
}
// AssignedPlanable 
type AssignedPlanable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCapabilityStatus()(*string)
    GetService()(*string)
    GetServicePlanId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCapabilityStatus(value *string)()
    SetService(value *string)()
    SetServicePlanId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
