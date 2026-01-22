// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntegerRange 
type IntegerRange struct {
    // The inclusive upper bound of the integer range.
    end *int64
    // The inclusive lower bound of the integer range.
    start *int64
}
// NewIntegerRange instantiates a new integerRange and sets the default values.
func NewIntegerRange()(*IntegerRange) {
    m := &IntegerRange{
    }
    return m
}
// CreateIntegerRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntegerRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegerRange(), nil
}
// GetEnd gets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) GetEnd()(*int64) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntegerRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    return res
}
// GetStart gets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) GetStart()(*int64) {
    return m.start
}
// Serialize serializes information the current object
func (m *IntegerRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnd sets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) SetEnd(value *int64)() {
    m.end = value
}
// SetStart sets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) SetStart(value *int64)() {
    m.start = value
}
// IntegerRangeable 
type IntegerRangeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnd()(*int64)
    GetStart()(*int64)
    SetEnd(value *int64)()
    SetStart(value *int64)()
}
