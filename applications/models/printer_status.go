// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterStatus 
type PrinterStatus struct {
    // A human-readable description of the printer's current processing state. Read-only.
    description *string
    // The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
    details []PrinterProcessingStateDetail
    // The state property
    state *PrinterProcessingState
}
// NewPrinterStatus instantiates a new printerStatus and sets the default values.
func NewPrinterStatus()(*PrinterStatus) {
    m := &PrinterStatus{
    }
    return m
}
// CreatePrinterStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterStatus(), nil
}
// GetDescription gets the description property value. A human-readable description of the printer's current processing state. Read-only.
func (m *PrinterStatus) GetDescription()(*string) {
    return m.description
}
// GetDetails gets the details property value. The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) GetDetails()([]PrinterProcessingStateDetail) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterProcessingStateDetail)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterProcessingStateDetail, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*PrinterProcessingStateDetail))
                }
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PrinterProcessingState))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state property
func (m *PrinterStatus) GetState()(*PrinterProcessingState) {
    return m.state
}
// Serialize serializes information the current object
func (m *PrinterStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDetails() != nil {
        err := writer.WriteCollectionOfStringValues("details", SerializePrinterProcessingStateDetail(m.GetDetails()))
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. A human-readable description of the printer's current processing state. Read-only.
func (m *PrinterStatus) SetDescription(value *string)() {
    m.description = value
}
// SetDetails sets the details property value. The list of details describing why the printer is in the current state. Valid values are described in the following table. Read-only.
func (m *PrinterStatus) SetDetails(value []PrinterProcessingStateDetail)() {
    m.details = value
}
// SetState sets the state property value. The state property
func (m *PrinterStatus) SetState(value *PrinterProcessingState)() {
    m.state = value
}
// PrinterStatusable 
type PrinterStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDetails()([]PrinterProcessingStateDetail)
    GetState()(*PrinterProcessingState)
    SetDescription(value *string)()
    SetDetails(value []PrinterProcessingStateDetail)()
    SetState(value *PrinterProcessingState)()
}
