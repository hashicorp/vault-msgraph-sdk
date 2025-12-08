// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAxis 
type WorkbookChartAxis struct {
    Entity
    // The format property
    format WorkbookChartAxisFormatable
    // The majorGridlines property
    majorGridlines WorkbookChartGridlinesable
    // The majorUnit property
    majorUnit Jsonable
    // The maximum property
    maximum Jsonable
    // The minimum property
    minimum Jsonable
    // The minorGridlines property
    minorGridlines WorkbookChartGridlinesable
    // The minorUnit property
    minorUnit Jsonable
    // The title property
    title WorkbookChartAxisTitleable
}
// NewWorkbookChartAxis instantiates a new workbookChartAxis and sets the default values.
func NewWorkbookChartAxis()(*WorkbookChartAxis) {
    m := &WorkbookChartAxis{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAxis(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartAxisFormatable))
        }
        return nil
    }
    res["majorGridlines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorGridlines(val.(WorkbookChartGridlinesable))
        }
        return nil
    }
    res["majorUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMajorUnit(val.(Jsonable))
        }
        return nil
    }
    res["maximum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximum(val.(Jsonable))
        }
        return nil
    }
    res["minimum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimum(val.(Jsonable))
        }
        return nil
    }
    res["minorGridlines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartGridlinesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorGridlines(val.(WorkbookChartGridlinesable))
        }
        return nil
    }
    res["minorUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinorUnit(val.(Jsonable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisTitleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(WorkbookChartAxisTitleable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The format property
func (m *WorkbookChartAxis) GetFormat()(WorkbookChartAxisFormatable) {
    return m.format
}
// GetMajorGridlines gets the majorGridlines property value. The majorGridlines property
func (m *WorkbookChartAxis) GetMajorGridlines()(WorkbookChartGridlinesable) {
    return m.majorGridlines
}
// GetMajorUnit gets the majorUnit property value. The majorUnit property
func (m *WorkbookChartAxis) GetMajorUnit()(Jsonable) {
    return m.majorUnit
}
// GetMaximum gets the maximum property value. The maximum property
func (m *WorkbookChartAxis) GetMaximum()(Jsonable) {
    return m.maximum
}
// GetMinimum gets the minimum property value. The minimum property
func (m *WorkbookChartAxis) GetMinimum()(Jsonable) {
    return m.minimum
}
// GetMinorGridlines gets the minorGridlines property value. The minorGridlines property
func (m *WorkbookChartAxis) GetMinorGridlines()(WorkbookChartGridlinesable) {
    return m.minorGridlines
}
// GetMinorUnit gets the minorUnit property value. The minorUnit property
func (m *WorkbookChartAxis) GetMinorUnit()(Jsonable) {
    return m.minorUnit
}
// GetTitle gets the title property value. The title property
func (m *WorkbookChartAxis) GetTitle()(WorkbookChartAxisTitleable) {
    return m.title
}
// Serialize serializes information the current object
func (m *WorkbookChartAxis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorGridlines", m.GetMajorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("majorUnit", m.GetMajorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimum", m.GetMinimum())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorGridlines", m.GetMinorGridlines())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minorUnit", m.GetMinorUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. The format property
func (m *WorkbookChartAxis) SetFormat(value WorkbookChartAxisFormatable)() {
    m.format = value
}
// SetMajorGridlines sets the majorGridlines property value. The majorGridlines property
func (m *WorkbookChartAxis) SetMajorGridlines(value WorkbookChartGridlinesable)() {
    m.majorGridlines = value
}
// SetMajorUnit sets the majorUnit property value. The majorUnit property
func (m *WorkbookChartAxis) SetMajorUnit(value Jsonable)() {
    m.majorUnit = value
}
// SetMaximum sets the maximum property value. The maximum property
func (m *WorkbookChartAxis) SetMaximum(value Jsonable)() {
    m.maximum = value
}
// SetMinimum sets the minimum property value. The minimum property
func (m *WorkbookChartAxis) SetMinimum(value Jsonable)() {
    m.minimum = value
}
// SetMinorGridlines sets the minorGridlines property value. The minorGridlines property
func (m *WorkbookChartAxis) SetMinorGridlines(value WorkbookChartGridlinesable)() {
    m.minorGridlines = value
}
// SetMinorUnit sets the minorUnit property value. The minorUnit property
func (m *WorkbookChartAxis) SetMinorUnit(value Jsonable)() {
    m.minorUnit = value
}
// SetTitle sets the title property value. The title property
func (m *WorkbookChartAxis) SetTitle(value WorkbookChartAxisTitleable)() {
    m.title = value
}
// WorkbookChartAxisable 
type WorkbookChartAxisable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(WorkbookChartAxisFormatable)
    GetMajorGridlines()(WorkbookChartGridlinesable)
    GetMajorUnit()(Jsonable)
    GetMaximum()(Jsonable)
    GetMinimum()(Jsonable)
    GetMinorGridlines()(WorkbookChartGridlinesable)
    GetMinorUnit()(Jsonable)
    GetTitle()(WorkbookChartAxisTitleable)
    SetFormat(value WorkbookChartAxisFormatable)()
    SetMajorGridlines(value WorkbookChartGridlinesable)()
    SetMajorUnit(value Jsonable)()
    SetMaximum(value Jsonable)()
    SetMinimum(value Jsonable)()
    SetMinorGridlines(value WorkbookChartGridlinesable)()
    SetMinorUnit(value Jsonable)()
    SetTitle(value WorkbookChartAxisTitleable)()
}
