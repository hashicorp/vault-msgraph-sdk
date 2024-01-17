// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package odataerrors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InnerError the structure of this object is service-specific
type InnerError struct {
}
// NewInnerError instantiates a new InnerError and sets the default values.
func NewInnerError()(*InnerError) {
    m := &InnerError{
    }
    return m
}
// CreateInnerErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInnerErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInnerError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InnerError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *InnerError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// InnerErrorable 
type InnerErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
