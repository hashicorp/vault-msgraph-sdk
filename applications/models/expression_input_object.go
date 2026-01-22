// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpressionInputObject 
type ExpressionInputObject struct {
    // The definition property
    definition ObjectDefinitionable
    // Property values of the test object.
    properties []StringKeyObjectValuePairable
}
// NewExpressionInputObject instantiates a new expressionInputObject and sets the default values.
func NewExpressionInputObject()(*ExpressionInputObject) {
    m := &ExpressionInputObject{
    }
    return m
}
// CreateExpressionInputObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpressionInputObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExpressionInputObject(), nil
}
// GetDefinition gets the definition property value. The definition property
func (m *ExpressionInputObject) GetDefinition()(ObjectDefinitionable) {
    return m.definition
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpressionInputObject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateObjectDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinition(val.(ObjectDefinitionable))
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStringKeyObjectValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyObjectValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StringKeyObjectValuePairable)
                }
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) GetProperties()([]StringKeyObjectValuePairable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *ExpressionInputObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefinition sets the definition property value. The definition property
func (m *ExpressionInputObject) SetDefinition(value ObjectDefinitionable)() {
    m.definition = value
}
// SetProperties sets the properties property value. Property values of the test object.
func (m *ExpressionInputObject) SetProperties(value []StringKeyObjectValuePairable)() {
    m.properties = value
}
// ExpressionInputObjectable 
type ExpressionInputObjectable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefinition()(ObjectDefinitionable)
    GetProperties()([]StringKeyObjectValuePairable)
    SetDefinition(value ObjectDefinitionable)()
    SetProperties(value []StringKeyObjectValuePairable)()
}
