// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeMappingSource 
type AttributeMappingSource struct {
    // Equivalent expression representation of this attributeMappingSource object.
    expression *string
    // Name parameter of the mapping source. Depending on the type property value, this can be the name of the function, the name of the source attribute, or a constant value to be used.
    name *string
    // If this object represents a function, lists function parameters. Parameters consist of attributeMappingSource objects themselves, allowing for complex expressions. If type isn't Function, this property is null/empty array.
    parameters []StringKeyAttributeMappingSourceValuePairable
    // The type property
    typeEscaped *AttributeMappingSourceType
}
// NewAttributeMappingSource instantiates a new attributeMappingSource and sets the default values.
func NewAttributeMappingSource()(*AttributeMappingSource) {
    m := &AttributeMappingSource{
    }
    return m
}
// CreateAttributeMappingSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeMappingSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeMappingSource(), nil
}
// GetExpression gets the expression property value. Equivalent expression representation of this attributeMappingSource object.
func (m *AttributeMappingSource) GetExpression()(*string) {
    return m.expression
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeMappingSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expression"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpression(val)
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
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStringKeyAttributeMappingSourceValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyAttributeMappingSourceValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StringKeyAttributeMappingSourceValuePairable)
                }
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttributeMappingSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*AttributeMappingSourceType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name parameter of the mapping source. Depending on the type property value, this can be the name of the function, the name of the source attribute, or a constant value to be used.
func (m *AttributeMappingSource) GetName()(*string) {
    return m.name
}
// GetParameters gets the parameters property value. If this object represents a function, lists function parameters. Parameters consist of attributeMappingSource objects themselves, allowing for complex expressions. If type isn't Function, this property is null/empty array.
func (m *AttributeMappingSource) GetParameters()([]StringKeyAttributeMappingSourceValuePairable) {
    return m.parameters
}
// GetTypeEscaped gets the type property value. The type property
func (m *AttributeMappingSource) GetTypeEscaped()(*AttributeMappingSourceType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *AttributeMappingSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
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
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpression sets the expression property value. Equivalent expression representation of this attributeMappingSource object.
func (m *AttributeMappingSource) SetExpression(value *string)() {
    m.expression = value
}
// SetName sets the name property value. Name parameter of the mapping source. Depending on the type property value, this can be the name of the function, the name of the source attribute, or a constant value to be used.
func (m *AttributeMappingSource) SetName(value *string)() {
    m.name = value
}
// SetParameters sets the parameters property value. If this object represents a function, lists function parameters. Parameters consist of attributeMappingSource objects themselves, allowing for complex expressions. If type isn't Function, this property is null/empty array.
func (m *AttributeMappingSource) SetParameters(value []StringKeyAttributeMappingSourceValuePairable)() {
    m.parameters = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *AttributeMappingSource) SetTypeEscaped(value *AttributeMappingSourceType)() {
    m.typeEscaped = value
}
// AttributeMappingSourceable 
type AttributeMappingSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpression()(*string)
    GetName()(*string)
    GetParameters()([]StringKeyAttributeMappingSourceValuePairable)
    GetTypeEscaped()(*AttributeMappingSourceType)
    SetExpression(value *string)()
    SetName(value *string)()
    SetParameters(value []StringKeyAttributeMappingSourceValuePairable)()
    SetTypeEscaped(value *AttributeMappingSourceType)()
}
