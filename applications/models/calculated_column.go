// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalculatedColumn 
type CalculatedColumn struct {
    // For dateTime output types, the format of the value. Possible values are: dateOnly or dateTime.
    format *string
    // The formula used to compute the value for this column.
    formula *string
    // The output type used to format values in this column. Possible values are: boolean, currency, dateTime, number, or text.
    outputType *string
}
// NewCalculatedColumn instantiates a new calculatedColumn and sets the default values.
func NewCalculatedColumn()(*CalculatedColumn) {
    m := &CalculatedColumn{
    }
    return m
}
// CreateCalculatedColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalculatedColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalculatedColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalculatedColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["formula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormula(val)
        }
        return nil
    }
    res["outputType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputType(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. For dateTime output types, the format of the value. Possible values are: dateOnly or dateTime.
func (m *CalculatedColumn) GetFormat()(*string) {
    return m.format
}
// GetFormula gets the formula property value. The formula used to compute the value for this column.
func (m *CalculatedColumn) GetFormula()(*string) {
    return m.formula
}
// GetOutputType gets the outputType property value. The output type used to format values in this column. Possible values are: boolean, currency, dateTime, number, or text.
func (m *CalculatedColumn) GetOutputType()(*string) {
    return m.outputType
}
// Serialize serializes information the current object
func (m *CalculatedColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputType", m.GetOutputType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. For dateTime output types, the format of the value. Possible values are: dateOnly or dateTime.
func (m *CalculatedColumn) SetFormat(value *string)() {
    m.format = value
}
// SetFormula sets the formula property value. The formula used to compute the value for this column.
func (m *CalculatedColumn) SetFormula(value *string)() {
    m.formula = value
}
// SetOutputType sets the outputType property value. The output type used to format values in this column. Possible values are: boolean, currency, dateTime, number, or text.
func (m *CalculatedColumn) SetOutputType(value *string)() {
    m.outputType = value
}
// CalculatedColumnable 
type CalculatedColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(*string)
    GetFormula()(*string)
    GetOutputType()(*string)
    SetFormat(value *string)()
    SetFormula(value *string)()
    SetOutputType(value *string)()
}
