package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartAreaFormat 
type WorkbookChartAreaFormat struct {
    Entity
    // The fill property
    fill WorkbookChartFillable
    // The font property
    font WorkbookChartFontable
}
// NewWorkbookChartAreaFormat instantiates a new workbookChartAreaFormat and sets the default values.
func NewWorkbookChartAreaFormat()(*WorkbookChartAreaFormat) {
    m := &WorkbookChartAreaFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAreaFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAreaFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartAreaFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAreaFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetFill gets the fill property value. The fill property
func (m *WorkbookChartAreaFormat) GetFill()(WorkbookChartFillable) {
    return m.fill
}
// GetFont gets the font property value. The font property
func (m *WorkbookChartAreaFormat) GetFont()(WorkbookChartFontable) {
    return m.font
}
// Serialize serializes information the current object
func (m *WorkbookChartAreaFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. The fill property
func (m *WorkbookChartAreaFormat) SetFill(value WorkbookChartFillable)() {
    m.fill = value
}
// SetFont sets the font property value. The font property
func (m *WorkbookChartAreaFormat) SetFont(value WorkbookChartFontable)() {
    m.font = value
}
// WorkbookChartAreaFormatable 
type WorkbookChartAreaFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFill()(WorkbookChartFillable)
    GetFont()(WorkbookChartFontable)
    SetFill(value WorkbookChartFillable)()
    SetFont(value WorkbookChartFontable)()
}