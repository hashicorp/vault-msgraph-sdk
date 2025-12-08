// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RequestSignatureVerification 
type RequestSignatureVerification struct {
    // The allowedWeakAlgorithms property
    allowedWeakAlgorithms *WeakAlgorithms
    // Specifies whether signed authentication requests for this application should be required.
    isSignedRequestRequired *bool
}
// NewRequestSignatureVerification instantiates a new requestSignatureVerification and sets the default values.
func NewRequestSignatureVerification()(*RequestSignatureVerification) {
    m := &RequestSignatureVerification{
    }
    return m
}
// CreateRequestSignatureVerificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRequestSignatureVerificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestSignatureVerification(), nil
}
// GetAllowedWeakAlgorithms gets the allowedWeakAlgorithms property value. The allowedWeakAlgorithms property
func (m *RequestSignatureVerification) GetAllowedWeakAlgorithms()(*WeakAlgorithms) {
    return m.allowedWeakAlgorithms
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestSignatureVerification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedWeakAlgorithms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeakAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedWeakAlgorithms(val.(*WeakAlgorithms))
        }
        return nil
    }
    res["isSignedRequestRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSignedRequestRequired(val)
        }
        return nil
    }
    return res
}
// GetIsSignedRequestRequired gets the isSignedRequestRequired property value. Specifies whether signed authentication requests for this application should be required.
func (m *RequestSignatureVerification) GetIsSignedRequestRequired()(*bool) {
    return m.isSignedRequestRequired
}
// Serialize serializes information the current object
func (m *RequestSignatureVerification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedWeakAlgorithms() != nil {
        cast := (*m.GetAllowedWeakAlgorithms()).String()
        err := writer.WriteStringValue("allowedWeakAlgorithms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSignedRequestRequired", m.GetIsSignedRequestRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedWeakAlgorithms sets the allowedWeakAlgorithms property value. The allowedWeakAlgorithms property
func (m *RequestSignatureVerification) SetAllowedWeakAlgorithms(value *WeakAlgorithms)() {
    m.allowedWeakAlgorithms = value
}
// SetIsSignedRequestRequired sets the isSignedRequestRequired property value. Specifies whether signed authentication requests for this application should be required.
func (m *RequestSignatureVerification) SetIsSignedRequestRequired(value *bool)() {
    m.isSignedRequestRequired = value
}
// RequestSignatureVerificationable 
type RequestSignatureVerificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedWeakAlgorithms()(*WeakAlgorithms)
    GetIsSignedRequestRequired()(*bool)
    SetAllowedWeakAlgorithms(value *WeakAlgorithms)()
    SetIsSignedRequestRequired(value *bool)()
}
