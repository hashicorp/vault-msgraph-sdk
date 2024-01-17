// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartSeriesFormat 
type WorkbookChartSeriesFormat struct {
    Entity
    // The fill property
    fill WorkbookChartFillable
    // The line property
    line WorkbookChartLineFormatable
}
// NewWorkbookChartSeriesFormat instantiates a new workbookChartSeriesFormat and sets the default values.
func NewWorkbookChartSeriesFormat()(*WorkbookChartSeriesFormat) {
    m := &WorkbookChartSeriesFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartSeriesFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartSeriesFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartSeriesFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartSeriesFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookChartFillable))
        }
        return nil
    }
    res["line"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartLineFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLine(val.(WorkbookChartLineFormatable))
        }
        return nil
    }
    return res
}
// GetFill gets the fill property value. The fill property
func (m *WorkbookChartSeriesFormat) GetFill()(WorkbookChartFillable) {
    return m.fill
}
// GetLine gets the line property value. The line property
func (m *WorkbookChartSeriesFormat) GetLine()(WorkbookChartLineFormatable) {
    return m.line
}
// Serialize serializes information the current object
func (m *WorkbookChartSeriesFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("line", m.GetLine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. The fill property
func (m *WorkbookChartSeriesFormat) SetFill(value WorkbookChartFillable)() {
    m.fill = value
}
// SetLine sets the line property value. The line property
func (m *WorkbookChartSeriesFormat) SetLine(value WorkbookChartLineFormatable)() {
    m.line = value
}
// WorkbookChartSeriesFormatable 
type WorkbookChartSeriesFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFill()(WorkbookChartFillable)
    GetLine()(WorkbookChartLineFormatable)
    SetFill(value WorkbookChartFillable)()
    SetLine(value WorkbookChartLineFormatable)()
}
