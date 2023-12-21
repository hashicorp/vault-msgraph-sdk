package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SamlSingleSignOnSettings 
type SamlSingleSignOnSettings struct {
    // The relative URI the service provider would redirect to after completion of the single sign-on flow.
    relayState *string
}
// NewSamlSingleSignOnSettings instantiates a new samlSingleSignOnSettings and sets the default values.
func NewSamlSingleSignOnSettings()(*SamlSingleSignOnSettings) {
    m := &SamlSingleSignOnSettings{
    }
    return m
}
// CreateSamlSingleSignOnSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSamlSingleSignOnSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSamlSingleSignOnSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SamlSingleSignOnSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["relayState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelayState(val)
        }
        return nil
    }
    return res
}
// GetRelayState gets the relayState property value. The relative URI the service provider would redirect to after completion of the single sign-on flow.
func (m *SamlSingleSignOnSettings) GetRelayState()(*string) {
    return m.relayState
}
// Serialize serializes information the current object
func (m *SamlSingleSignOnSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("relayState", m.GetRelayState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRelayState sets the relayState property value. The relative URI the service provider would redirect to after completion of the single sign-on flow.
func (m *SamlSingleSignOnSettings) SetRelayState(value *string)() {
    m.relayState = value
}
// SamlSingleSignOnSettingsable 
type SamlSingleSignOnSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRelayState()(*string)
    SetRelayState(value *string)()
}
