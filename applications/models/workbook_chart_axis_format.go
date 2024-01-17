// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxisFormat 
type WorkbookChartAxisFormat struct {
    Entity
    // The font property
    font WorkbookChartFontable
    // The line property
    line WorkbookChartLineFormatable
}
// NewWorkbookChartAxisFormat instantiates a new workbookChartAxisFormat and sets the default values.
func NewWorkbookChartAxisFormat()(*WorkbookChartAxisFormat) {
    m := &WorkbookChartAxisFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAxisFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxisFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["font"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFontFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(WorkbookChartFontable))
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
// GetFont gets the font property value. The font property
func (m *WorkbookChartAxisFormat) GetFont()(WorkbookChartFontable) {
    return m.font
}
// GetLine gets the line property value. The line property
func (m *WorkbookChartAxisFormat) GetLine()(WorkbookChartLineFormatable) {
    return m.line
}
// Serialize serializes information the current object
func (m *WorkbookChartAxisFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
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
// SetFont sets the font property value. The font property
func (m *WorkbookChartAxisFormat) SetFont(value WorkbookChartFontable)() {
    m.font = value
}
// SetLine sets the line property value. The line property
func (m *WorkbookChartAxisFormat) SetLine(value WorkbookChartLineFormatable)() {
    m.line = value
}
// WorkbookChartAxisFormatable 
type WorkbookChartAxisFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFont()(WorkbookChartFontable)
    GetLine()(WorkbookChartLineFormatable)
    SetFont(value WorkbookChartFontable)()
    SetLine(value WorkbookChartLineFormatable)()
}
