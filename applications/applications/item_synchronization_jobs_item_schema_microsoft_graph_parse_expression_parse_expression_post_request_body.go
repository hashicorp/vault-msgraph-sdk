// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody 
type ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody struct {
    // The expression property
    expression *string
    // The targetAttributeDefinition property
    targetAttributeDefinition i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeDefinitionable
    // The testInputObject property
    testInputObject i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.ExpressionInputObjectable
}
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody instantiates a new ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody()(*ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) {
    m := &ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody{
    }
    return m
}
// CreateItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody(), nil
}
// GetExpression gets the expression property value. The expression property
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) GetExpression()(*string) {
    return m.expression
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["targetAttributeDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateAttributeDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetAttributeDefinition(val.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeDefinitionable))
        }
        return nil
    }
    res["testInputObject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateExpressionInputObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestInputObject(val.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.ExpressionInputObjectable))
        }
        return nil
    }
    return res
}
// GetTargetAttributeDefinition gets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) GetTargetAttributeDefinition()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeDefinitionable) {
    return m.targetAttributeDefinition
}
// GetTestInputObject gets the testInputObject property value. The testInputObject property
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) GetTestInputObject()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.ExpressionInputObjectable) {
    return m.testInputObject
}
// Serialize serializes information the current object
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("expression", m.GetExpression())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("targetAttributeDefinition", m.GetTargetAttributeDefinition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("testInputObject", m.GetTestInputObject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpression sets the expression property value. The expression property
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) SetExpression(value *string)() {
    m.expression = value
}
// SetTargetAttributeDefinition sets the targetAttributeDefinition property value. The targetAttributeDefinition property
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) SetTargetAttributeDefinition(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeDefinitionable)() {
    m.targetAttributeDefinition = value
}
// SetTestInputObject sets the testInputObject property value. The testInputObject property
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBody) SetTestInputObject(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.ExpressionInputObjectable)() {
    m.testInputObject = value
}
// ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBodyable 
type ItemSynchronizationJobsItemSchemaMicrosoftGraphParseExpressionParseExpressionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpression()(*string)
    GetTargetAttributeDefinition()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeDefinitionable)
    GetTestInputObject()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.ExpressionInputObjectable)
    SetExpression(value *string)()
    SetTargetAttributeDefinition(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeDefinitionable)()
    SetTestInputObject(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.ExpressionInputObjectable)()
}
