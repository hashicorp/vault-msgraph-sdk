// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectMappingMetadataEntry 
type ObjectMappingMetadataEntry struct {
    // The key property
    key *ObjectMappingMetadata
    // Value of the metadata property.
    value *string
}
// NewObjectMappingMetadataEntry instantiates a new objectMappingMetadataEntry and sets the default values.
func NewObjectMappingMetadataEntry()(*ObjectMappingMetadataEntry) {
    m := &ObjectMappingMetadataEntry{
    }
    return m
}
// CreateObjectMappingMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectMappingMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectMappingMetadataEntry(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectMappingMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectMappingMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*ObjectMappingMetadata))
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
func (m *ObjectMappingMetadataEntry) GetKey()(*ObjectMappingMetadata) {
    return m.key
}
// GetValue gets the value property value. Value of the metadata property.
func (m *ObjectMappingMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ObjectMappingMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ObjectMappingMetadataEntry) SetKey(value *ObjectMappingMetadata)() {
    m.key = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *ObjectMappingMetadataEntry) SetValue(value *string)() {
    m.value = value
}
// ObjectMappingMetadataEntryable 
type ObjectMappingMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*ObjectMappingMetadata)
    GetValue()(*string)
    SetKey(value *ObjectMappingMetadata)()
    SetValue(value *string)()
}
