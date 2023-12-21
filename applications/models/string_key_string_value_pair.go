package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StringKeyStringValuePair 
type StringKeyStringValuePair struct {
    // Key.
    key *string
    // Value.
    value *string
}
// NewStringKeyStringValuePair instantiates a new stringKeyStringValuePair and sets the default values.
func NewStringKeyStringValuePair()(*StringKeyStringValuePair) {
    m := &StringKeyStringValuePair{
    }
    return m
}
// CreateStringKeyStringValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyStringValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStringKeyStringValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyStringValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. Key.
func (m *StringKeyStringValuePair) GetKey()(*string) {
    return m.key
}
// GetValue gets the value property value. Value.
func (m *StringKeyStringValuePair) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *StringKeyStringValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. Key.
func (m *StringKeyStringValuePair) SetKey(value *string)() {
    m.key = value
}
// SetValue sets the value property value. Value.
func (m *StringKeyStringValuePair) SetValue(value *string)() {
    m.value = value
}
// StringKeyStringValuePairable 
type StringKeyStringValuePairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    GetValue()(*string)
    SetKey(value *string)()
    SetValue(value *string)()
}
