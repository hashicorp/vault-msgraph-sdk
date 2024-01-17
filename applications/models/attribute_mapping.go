// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMapping 
type AttributeMapping struct {
    // Default value to be used in case the source property was evaluated to null. Optional.
    defaultValue *string
    // For internal use only.
    exportMissingReferences *bool
    // The flowBehavior property
    flowBehavior *AttributeFlowBehavior
    // The flowType property
    flowType *AttributeFlowType
    // If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
    matchingPriority *int32
    // The source property
    source AttributeMappingSourceable
    // Name of the attribute on the target object.
    targetAttributeName *string
}
// NewAttributeMapping instantiates a new attributeMapping and sets the default values.
func NewAttributeMapping()(*AttributeMapping) {
    m := &AttributeMapping{
    }
    return m
}
// CreateAttributeMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMapping(), nil
}
// GetDefaultValue gets the defaultValue property value. Default value to be used in case the source property was evaluated to null. Optional.
func (m *AttributeMapping) GetDefaultValue()(*string) {
    return m.defaultValue
}
// GetExportMissingReferences gets the exportMissingReferences property value. For internal use only.
func (m *AttributeMapping) GetExportMissingReferences()(*bool) {
    return m.exportMissingReferences
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["exportMissingReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportMissingReferences(val)
        }
        return nil
    }
    res["flowBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeFlowBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlowBehavior(val.(*AttributeFlowBehavior))
        }
        return nil
    }
    res["flowType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeFlowType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlowType(val.(*AttributeFlowType))
        }
        return nil
    }
    res["matchingPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingPriority(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttributeMappingSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(AttributeMappingSourceable))
        }
        return nil
    }
    res["targetAttributeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeName(val)
        }
        return nil
    }
    return res
}
// GetFlowBehavior gets the flowBehavior property value. The flowBehavior property
func (m *AttributeMapping) GetFlowBehavior()(*AttributeFlowBehavior) {
    return m.flowBehavior
}
// GetFlowType gets the flowType property value. The flowType property
func (m *AttributeMapping) GetFlowType()(*AttributeFlowType) {
    return m.flowType
}
// GetMatchingPriority gets the matchingPriority property value. If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
func (m *AttributeMapping) GetMatchingPriority()(*int32) {
    return m.matchingPriority
}
// GetSource gets the source property value. The source property
func (m *AttributeMapping) GetSource()(AttributeMappingSourceable) {
    return m.source
}
// GetTargetAttributeName gets the targetAttributeName property value. Name of the attribute on the target object.
func (m *AttributeMapping) GetTargetAttributeName()(*string) {
    return m.targetAttributeName
}
// Serialize serializes information the current object
func (m *AttributeMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("exportMissingReferences", m.GetExportMissingReferences())
        if err != nil {
            return err
        }
    }
    if m.GetFlowBehavior() != nil {
        cast := (*m.GetFlowBehavior()).String()
        err := writer.WriteStringValue("flowBehavior", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFlowType() != nil {
        cast := (*m.GetFlowType()).String()
        err := writer.WriteStringValue("flowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("matchingPriority", m.GetMatchingPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetAttributeName", m.GetTargetAttributeName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. Default value to be used in case the source property was evaluated to null. Optional.
func (m *AttributeMapping) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetExportMissingReferences sets the exportMissingReferences property value. For internal use only.
func (m *AttributeMapping) SetExportMissingReferences(value *bool)() {
    m.exportMissingReferences = value
}
// SetFlowBehavior sets the flowBehavior property value. The flowBehavior property
func (m *AttributeMapping) SetFlowBehavior(value *AttributeFlowBehavior)() {
    m.flowBehavior = value
}
// SetFlowType sets the flowType property value. The flowType property
func (m *AttributeMapping) SetFlowType(value *AttributeFlowType)() {
    m.flowType = value
}
// SetMatchingPriority sets the matchingPriority property value. If higher than 0, this attribute will be used to perform an initial match of the objects between source and target directories. The synchronization engine will try to find the matching object using attribute with lowest value of matching priority first. If not found, the attribute with the next matching priority will be used, and so on a until match is found or no more matching attributes are left. Only attributes that are expected to have unique values, such as email, should be used as matching attributes.
func (m *AttributeMapping) SetMatchingPriority(value *int32)() {
    m.matchingPriority = value
}
// SetSource sets the source property value. The source property
func (m *AttributeMapping) SetSource(value AttributeMappingSourceable)() {
    m.source = value
}
// SetTargetAttributeName sets the targetAttributeName property value. Name of the attribute on the target object.
func (m *AttributeMapping) SetTargetAttributeName(value *string)() {
    m.targetAttributeName = value
}
// AttributeMappingable 
type AttributeMappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetExportMissingReferences()(*bool)
    GetFlowBehavior()(*AttributeFlowBehavior)
    GetFlowType()(*AttributeFlowType)
    GetMatchingPriority()(*int32)
    GetSource()(AttributeMappingSourceable)
    GetTargetAttributeName()(*string)
    SetDefaultValue(value *string)()
    SetExportMissingReferences(value *bool)()
    SetFlowBehavior(value *AttributeFlowBehavior)()
    SetFlowType(value *AttributeFlowType)()
    SetMatchingPriority(value *int32)()
    SetSource(value AttributeMappingSourceable)()
    SetTargetAttributeName(value *string)()
}
