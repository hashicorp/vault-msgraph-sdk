// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicError 
type PublicError struct {
    // Represents the error code.
    code *string
    // Details of the error.
    details []PublicErrorDetailable
    // The innerError property
    innerError PublicInnerErrorable
    // A non-localized message for the developer.
    message *string
    // The target of the error.
    target *string
}
// NewPublicError instantiates a new publicError and sets the default values.
func NewPublicError()(*PublicError) {
    m := &PublicError{
    }
    return m
}
// CreatePublicErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicError(), nil
}
// GetCode gets the code property value. Represents the error code.
func (m *PublicError) GetCode()(*string) {
    return m.code
}
// GetDetails gets the details property value. Details of the error.
func (m *PublicError) GetDetails()([]PublicErrorDetailable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublicErrorDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublicErrorDetailable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PublicErrorDetailable)
                }
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["innerError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicInnerErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(PublicInnerErrorable))
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
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetInnerError gets the innerError property value. The innerError property
func (m *PublicError) GetInnerError()(PublicInnerErrorable) {
    return m.innerError
}
// GetMessage gets the message property value. A non-localized message for the developer.
func (m *PublicError) GetMessage()(*string) {
    return m.message
}
// GetTarget gets the target property value. The target of the error.
func (m *PublicError) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *PublicError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    if m.GetDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("details", cast)
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
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCode sets the code property value. Represents the error code.
func (m *PublicError) SetCode(value *string)() {
    m.code = value
}
// SetDetails sets the details property value. Details of the error.
func (m *PublicError) SetDetails(value []PublicErrorDetailable)() {
    m.details = value
}
// SetInnerError sets the innerError property value. The innerError property
func (m *PublicError) SetInnerError(value PublicInnerErrorable)() {
    m.innerError = value
}
// SetMessage sets the message property value. A non-localized message for the developer.
func (m *PublicError) SetMessage(value *string)() {
    m.message = value
}
// SetTarget sets the target property value. The target of the error.
func (m *PublicError) SetTarget(value *string)() {
    m.target = value
}
// PublicErrorable 
type PublicErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetDetails()([]PublicErrorDetailable)
    GetInnerError()(PublicInnerErrorable)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetDetails(value []PublicErrorDetailable)()
    SetInnerError(value PublicInnerErrorable)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}
