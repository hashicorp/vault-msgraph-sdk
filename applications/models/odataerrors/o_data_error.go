// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package odataerrors

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ODataError 
type ODataError struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // The error property
    errorEscaped MainErrorable
}
// NewODataError instantiates a new ODataError and sets the default values.
func NewODataError()(*ODataError) {
    m := &ODataError{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    return m
}
// CreateODataErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateODataErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewODataError(), nil
}
// Error the primary error message.
func (m *ODataError) Error()(string) {
    return *(m.GetErrorEscaped().GetMessage())
}
// GetErrorEscaped gets the error property value. The error property
func (m *ODataError) GetErrorEscaped()(MainErrorable) {
    return m.errorEscaped
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ODataError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMainErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorEscaped(val.(MainErrorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ODataError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetErrorEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorEscaped sets the error property value. The error property
func (m *ODataError) SetErrorEscaped(value MainErrorable)() {
    m.errorEscaped = value
}
// ODataErrorable 
type ODataErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorEscaped()(MainErrorable)
    SetErrorEscaped(value MainErrorable)()
}
