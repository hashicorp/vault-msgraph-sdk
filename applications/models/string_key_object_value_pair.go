package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StringKeyObjectValuePair 
type StringKeyObjectValuePair struct {
    // Key.
    key *string
}
// NewStringKeyObjectValuePair instantiates a new stringKeyObjectValuePair and sets the default values.
func NewStringKeyObjectValuePair()(*StringKeyObjectValuePair) {
    m := &StringKeyObjectValuePair{
    }
    return m
}
// CreateStringKeyObjectValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStringKeyObjectValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStringKeyObjectValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StringKeyObjectValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetKey gets the key property value. Key.
func (m *StringKeyObjectValuePair) GetKey()(*string) {
    return m.key
}
// Serialize serializes information the current object
func (m *StringKeyObjectValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKey sets the key property value. Key.
func (m *StringKeyObjectValuePair) SetKey(value *string)() {
    m.key = value
}
// StringKeyObjectValuePairable 
type StringKeyObjectValuePairable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKey()(*string)
    SetKey(value *string)()
}
