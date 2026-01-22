// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookOperationError 
type WorkbookOperationError struct {
    // The error code.
    code *string
    // The innerError property
    innerError WorkbookOperationErrorable
    // The error message.
    message *string
}
// NewWorkbookOperationError instantiates a new workbookOperationError and sets the default values.
func NewWorkbookOperationError()(*WorkbookOperationError) {
    m := &WorkbookOperationError{
    }
    return m
}
// CreateWorkbookOperationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookOperationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookOperationError(), nil
}
// GetCode gets the code property value. The error code.
func (m *WorkbookOperationError) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookOperationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["innerError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(WorkbookOperationErrorable))
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
// GetInnerError gets the innerError property value. The innerError property
func (m *WorkbookOperationError) GetInnerError()(WorkbookOperationErrorable) {
    return m.innerError
}
// GetMessage gets the message property value. The error message.
func (m *WorkbookOperationError) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *WorkbookOperationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("innerError", m.GetInnerError())
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
// SetCode sets the code property value. The error code.
func (m *WorkbookOperationError) SetCode(value *string)() {
    m.code = value
}
// SetInnerError sets the innerError property value. The innerError property
func (m *WorkbookOperationError) SetInnerError(value WorkbookOperationErrorable)() {
    m.innerError = value
}
// SetMessage sets the message property value. The error message.
func (m *WorkbookOperationError) SetMessage(value *string)() {
    m.message = value
}
// WorkbookOperationErrorable 
type WorkbookOperationErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetInnerError()(WorkbookOperationErrorable)
    GetMessage()(*string)
    SetCode(value *string)()
    SetInnerError(value WorkbookOperationErrorable)()
    SetMessage(value *string)()
}
