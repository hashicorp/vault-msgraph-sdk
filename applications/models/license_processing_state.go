package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LicenseProcessingState 
type LicenseProcessingState struct {
    // The state property
    state *string
}
// NewLicenseProcessingState instantiates a new licenseProcessingState and sets the default values.
func NewLicenseProcessingState()(*LicenseProcessingState) {
    m := &LicenseProcessingState{
    }
    return m
}
// CreateLicenseProcessingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLicenseProcessingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLicenseProcessingState(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LicenseProcessingState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state property
func (m *LicenseProcessingState) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *LicenseProcessingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetState sets the state property value. The state property
func (m *LicenseProcessingState) SetState(value *string)() {
    m.state = value
}
// LicenseProcessingStateable 
type LicenseProcessingStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetState()(*string)
    SetState(value *string)()
}
