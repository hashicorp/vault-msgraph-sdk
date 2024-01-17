// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectDefinitionMetadataEntry 
type ObjectDefinitionMetadataEntry struct {
    // The key property
    key *ObjectDefinitionMetadata
    // Value of the metadata property.
    value *string
}
// NewObjectDefinitionMetadataEntry instantiates a new objectDefinitionMetadataEntry and sets the default values.
func NewObjectDefinitionMetadataEntry()(*ObjectDefinitionMetadataEntry) {
    m := &ObjectDefinitionMetadataEntry{
    }
    return m
}
// CreateObjectDefinitionMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectDefinitionMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectDefinitionMetadataEntry(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectDefinitionMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectDefinitionMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*ObjectDefinitionMetadata))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The key property
func (m *ObjectDefinitionMetadataEntry) GetKey()(*ObjectDefinitionMetadata) {
    return m.key
}
// GetValue gets the value property value. Value of the metadata property.
func (m *ObjectDefinitionMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ObjectDefinitionMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKey() != nil {
        cast := (*m.GetKey()).String()
        err := writer.WriteStringValue("key", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. The key property
func (m *ObjectDefinitionMetadataEntry) SetKey(value *ObjectDefinitionMetadata)() {
    m.key = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *ObjectDefinitionMetadataEntry) SetValue(value *string)() {
    m.value = value
}
// ObjectDefinitionMetadataEntryable 
type ObjectDefinitionMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*ObjectDefinitionMetadata)
    GetValue()(*string)
    SetKey(value *ObjectDefinitionMetadata)()
    SetValue(value *string)()
}
