// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationMetadataEntry 
type SynchronizationMetadataEntry struct {
    // The key property
    key *SynchronizationMetadata
    // Value of the metadata property.
    value *string
}
// NewSynchronizationMetadataEntry instantiates a new synchronizationMetadataEntry and sets the default values.
func NewSynchronizationMetadataEntry()(*SynchronizationMetadataEntry) {
    m := &SynchronizationMetadataEntry{
    }
    return m
}
// CreateSynchronizationMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationMetadataEntry(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*SynchronizationMetadata))
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
func (m *SynchronizationMetadataEntry) GetKey()(*SynchronizationMetadata) {
    return m.key
}
// GetValue gets the value property value. Value of the metadata property.
func (m *SynchronizationMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *SynchronizationMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SynchronizationMetadataEntry) SetKey(value *SynchronizationMetadata)() {
    m.key = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *SynchronizationMetadataEntry) SetValue(value *string)() {
    m.value = value
}
// SynchronizationMetadataEntryable 
type SynchronizationMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*SynchronizationMetadata)
    GetValue()(*string)
    SetKey(value *SynchronizationMetadata)()
    SetValue(value *string)()
}
