// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteOperationError 
type OnenoteOperationError struct {
    // The error code.
    code *string
    // The error message.
    message *string
}
// NewOnenoteOperationError instantiates a new onenoteOperationError and sets the default values.
func NewOnenoteOperationError()(*OnenoteOperationError) {
    m := &OnenoteOperationError{
    }
    return m
}
// CreateOnenoteOperationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteOperationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteOperationError(), nil
}
// GetCode gets the code property value. The error code.
func (m *OnenoteOperationError) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteOperationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetMessage gets the message property value. The error message.
func (m *OnenoteOperationError) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *OnenoteOperationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetCode sets the code property value. The error code.
func (m *OnenoteOperationError) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. The error message.
func (m *OnenoteOperationError) SetMessage(value *string)() {
    m.message = value
}
// OnenoteOperationErrorable 
type OnenoteOperationErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetMessage()(*string)
    SetCode(value *string)()
    SetMessage(value *string)()
}
