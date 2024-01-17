// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationError 
type OperationError struct {
    // Operation error code.
    code *string
    // Operation error message.
    message *string
}
// NewOperationError instantiates a new operationError and sets the default values.
func NewOperationError()(*OperationError) {
    m := &OperationError{
    }
    return m
}
// CreateOperationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationError(), nil
}
// GetCode gets the code property value. Operation error code.
func (m *OperationError) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OperationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Operation error message.
func (m *OperationError) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *OperationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCode sets the code property value. Operation error code.
func (m *OperationError) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. Operation error message.
func (m *OperationError) SetMessage(value *string)() {
    m.message = value
}
// OperationErrorable 
type OperationErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetMessage()(*string)
    SetCode(value *string)()
    SetMessage(value *string)()
}
