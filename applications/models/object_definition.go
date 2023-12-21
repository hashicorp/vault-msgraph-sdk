package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectDefinition 
type ObjectDefinition struct {
    // Defines attributes of the object.
    attributes []AttributeDefinitionable
    // Metadata for the given object.
    metadata []ObjectDefinitionMetadataEntryable
    // Name of the object. Must be unique within a directory definition. Not nullable.
    name *string
    // The API that the provisioning service queries to retrieve data for synchronization.
    supportedApis []string
}
// NewObjectDefinition instantiates a new objectDefinition and sets the default values.
func NewObjectDefinition()(*ObjectDefinition) {
    m := &ObjectDefinition{
    }
    return m
}
// CreateObjectDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectDefinition(), nil
}
// GetAttributes gets the attributes property value. Defines attributes of the object.
func (m *ObjectDefinition) GetAttributes()([]AttributeDefinitionable) {
    return m.attributes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AttributeDefinitionable)
                }
            }
            m.SetAttributes(res)
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateObjectDefinitionMetadataEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ObjectDefinitionMetadataEntryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ObjectDefinitionMetadataEntryable)
                }
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["supportedApis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportedApis(res)
        }
        return nil
    }
    return res
}
// GetMetadata gets the metadata property value. Metadata for the given object.
func (m *ObjectDefinition) GetMetadata()([]ObjectDefinitionMetadataEntryable) {
    return m.metadata
}
// GetName gets the name property value. Name of the object. Must be unique within a directory definition. Not nullable.
func (m *ObjectDefinition) GetName()(*string) {
    return m.name
}
// GetSupportedApis gets the supportedApis property value. The API that the provisioning service queries to retrieve data for synchronization.
func (m *ObjectDefinition) GetSupportedApis()([]string) {
    return m.supportedApis
}
// Serialize serializes information the current object
func (m *ObjectDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("attributes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedApis() != nil {
        err := writer.WriteCollectionOfStringValues("supportedApis", m.GetSupportedApis())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttributes sets the attributes property value. Defines attributes of the object.
func (m *ObjectDefinition) SetAttributes(value []AttributeDefinitionable)() {
    m.attributes = value
}
// SetMetadata sets the metadata property value. Metadata for the given object.
func (m *ObjectDefinition) SetMetadata(value []ObjectDefinitionMetadataEntryable)() {
    m.metadata = value
}
// SetName sets the name property value. Name of the object. Must be unique within a directory definition. Not nullable.
func (m *ObjectDefinition) SetName(value *string)() {
    m.name = value
}
// SetSupportedApis sets the supportedApis property value. The API that the provisioning service queries to retrieve data for synchronization.
func (m *ObjectDefinition) SetSupportedApis(value []string)() {
    m.supportedApis = value
}
// ObjectDefinitionable 
type ObjectDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributes()([]AttributeDefinitionable)
    GetMetadata()([]ObjectDefinitionMetadataEntryable)
    GetName()(*string)
    GetSupportedApis()([]string)
    SetAttributes(value []AttributeDefinitionable)()
    SetMetadata(value []ObjectDefinitionMetadataEntryable)()
    SetName(value *string)()
    SetSupportedApis(value []string)()
}
