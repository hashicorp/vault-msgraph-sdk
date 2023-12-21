package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeDefinitionMetadataEntry 
type AttributeDefinitionMetadataEntry struct {
    // The key property
    key *AttributeDefinitionMetadata
    // Value of the metadata property.
    value *string
}
// NewAttributeDefinitionMetadataEntry instantiates a new attributeDefinitionMetadataEntry and sets the default values.
func NewAttributeDefinitionMetadataEntry()(*AttributeDefinitionMetadataEntry) {
    m := &AttributeDefinitionMetadataEntry{
    }
    return m
}
// CreateAttributeDefinitionMetadataEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeDefinitionMetadataEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeDefinitionMetadataEntry(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeDefinitionMetadataEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeDefinitionMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val.(*AttributeDefinitionMetadata))
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
func (m *AttributeDefinitionMetadataEntry) GetKey()(*AttributeDefinitionMetadata) {
    return m.key
}
// GetValue gets the value property value. Value of the metadata property.
func (m *AttributeDefinitionMetadataEntry) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *AttributeDefinitionMetadataEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AttributeDefinitionMetadataEntry) SetKey(value *AttributeDefinitionMetadata)() {
    m.key = value
}
// SetValue sets the value property value. Value of the metadata property.
func (m *AttributeDefinitionMetadataEntry) SetValue(value *string)() {
    m.value = value
}
// AttributeDefinitionMetadataEntryable 
type AttributeDefinitionMetadataEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*AttributeDefinitionMetadata)
    GetValue()(*string)
    SetKey(value *AttributeDefinitionMetadata)()
    SetValue(value *string)()
}
