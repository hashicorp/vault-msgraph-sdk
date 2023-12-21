package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deleted 
type Deleted struct {
    // Represents the state of the deleted item.
    state *string
}
// NewDeleted instantiates a new deleted and sets the default values.
func NewDeleted()(*Deleted) {
    m := &Deleted{
    }
    return m
}
// CreateDeletedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleted(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Deleted) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetState gets the state property value. Represents the state of the deleted item.
func (m *Deleted) GetState()(*string) {
    return m.state
}
// Serialize serializes information the current object
func (m *Deleted) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetState sets the state property value. Represents the state of the deleted item.
func (m *Deleted) SetState(value *string)() {
    m.state = value
}
// Deletedable 
type Deletedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetState()(*string)
    SetState(value *string)()
}
