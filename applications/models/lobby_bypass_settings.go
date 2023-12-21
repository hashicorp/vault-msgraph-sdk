package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LobbyBypassSettings 
type LobbyBypassSettings struct {
    // Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
    isDialInBypassEnabled *bool
    // The scope property
    scope *LobbyBypassScope
}
// NewLobbyBypassSettings instantiates a new lobbyBypassSettings and sets the default values.
func NewLobbyBypassSettings()(*LobbyBypassSettings) {
    m := &LobbyBypassSettings{
    }
    return m
}
// CreateLobbyBypassSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLobbyBypassSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLobbyBypassSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LobbyBypassSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isDialInBypassEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDialInBypassEnabled(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLobbyBypassScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*LobbyBypassScope))
        }
        return nil
    }
    return res
}
// GetIsDialInBypassEnabled gets the isDialInBypassEnabled property value. Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
func (m *LobbyBypassSettings) GetIsDialInBypassEnabled()(*bool) {
    return m.isDialInBypassEnabled
}
// GetScope gets the scope property value. The scope property
func (m *LobbyBypassSettings) GetScope()(*LobbyBypassScope) {
    return m.scope
}
// Serialize serializes information the current object
func (m *LobbyBypassSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDialInBypassEnabled", m.GetIsDialInBypassEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err := writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDialInBypassEnabled sets the isDialInBypassEnabled property value. Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
func (m *LobbyBypassSettings) SetIsDialInBypassEnabled(value *bool)() {
    m.isDialInBypassEnabled = value
}
// SetScope sets the scope property value. The scope property
func (m *LobbyBypassSettings) SetScope(value *LobbyBypassScope)() {
    m.scope = value
}
// LobbyBypassSettingsable 
type LobbyBypassSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDialInBypassEnabled()(*bool)
    GetScope()(*LobbyBypassScope)
    SetIsDialInBypassEnabled(value *bool)()
    SetScope(value *LobbyBypassScope)()
}
