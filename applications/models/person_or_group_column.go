// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonOrGroupColumn 
type PersonOrGroupColumn struct {
    // Indicates whether multiple values can be selected from the source.
    allowMultipleSelection *bool
    // Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
    chooseFromType *string
    // How to display the information about the person or group chosen. See below.
    displayAs *string
}
// NewPersonOrGroupColumn instantiates a new personOrGroupColumn and sets the default values.
func NewPersonOrGroupColumn()(*PersonOrGroupColumn) {
    m := &PersonOrGroupColumn{
    }
    return m
}
// CreatePersonOrGroupColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonOrGroupColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonOrGroupColumn(), nil
}
// GetAllowMultipleSelection gets the allowMultipleSelection property value. Indicates whether multiple values can be selected from the source.
func (m *PersonOrGroupColumn) GetAllowMultipleSelection()(*bool) {
    return m.allowMultipleSelection
}
// GetChooseFromType gets the chooseFromType property value. Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
func (m *PersonOrGroupColumn) GetChooseFromType()(*string) {
    return m.chooseFromType
}
// GetDisplayAs gets the displayAs property value. How to display the information about the person or group chosen. See below.
func (m *PersonOrGroupColumn) GetDisplayAs()(*string) {
    return m.displayAs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonOrGroupColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowMultipleSelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleSelection(val)
        }
        return nil
    }
    res["chooseFromType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChooseFromType(val)
        }
        return nil
    }
    res["displayAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PersonOrGroupColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleSelection", m.GetAllowMultipleSelection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("chooseFromType", m.GetChooseFromType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowMultipleSelection sets the allowMultipleSelection property value. Indicates whether multiple values can be selected from the source.
func (m *PersonOrGroupColumn) SetAllowMultipleSelection(value *bool)() {
    m.allowMultipleSelection = value
}
// SetChooseFromType sets the chooseFromType property value. Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
func (m *PersonOrGroupColumn) SetChooseFromType(value *string)() {
    m.chooseFromType = value
}
// SetDisplayAs sets the displayAs property value. How to display the information about the person or group chosen. See below.
func (m *PersonOrGroupColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
// PersonOrGroupColumnable 
type PersonOrGroupColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowMultipleSelection()(*bool)
    GetChooseFromType()(*string)
    GetDisplayAs()(*string)
    SetAllowMultipleSelection(value *bool)()
    SetChooseFromType(value *string)()
    SetDisplayAs(value *string)()
}
