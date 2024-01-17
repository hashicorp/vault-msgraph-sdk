// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListItemVersion 
type ListItemVersion struct {
    BaseItemVersion
    // The fields property
    fields FieldValueSetable
}
// NewListItemVersion instantiates a new listItemVersion and sets the default values.
func NewListItemVersion()(*ListItemVersion) {
    m := &ListItemVersion{
        BaseItemVersion: *NewBaseItemVersion(),
    }
    return m
}
// CreateListItemVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListItemVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListItemVersion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListItemVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItemVersion.GetFieldDeserializers()
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFieldValueSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(FieldValueSetable))
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The fields property
func (m *ListItemVersion) GetFields()(FieldValueSetable) {
    return m.fields
}
// Serialize serializes information the current object
func (m *ListItemVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItemVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFields sets the fields property value. The fields property
func (m *ListItemVersion) SetFields(value FieldValueSetable)() {
    m.fields = value
}
// ListItemVersionable 
type ListItemVersionable interface {
    BaseItemVersionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFields()(FieldValueSetable)
    SetFields(value FieldValueSetable)()
}
