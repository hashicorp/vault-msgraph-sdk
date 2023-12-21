package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SpecialFolder 
type SpecialFolder struct {
    // The unique identifier for this item in the /drive/special collection
    name *string
}
// NewSpecialFolder instantiates a new specialFolder and sets the default values.
func NewSpecialFolder()(*SpecialFolder) {
    m := &SpecialFolder{
    }
    return m
}
// CreateSpecialFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSpecialFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSpecialFolder(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SpecialFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The unique identifier for this item in the /drive/special collection
func (m *SpecialFolder) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *SpecialFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The unique identifier for this item in the /drive/special collection
func (m *SpecialFolder) SetName(value *string)() {
    m.name = value
}
// SpecialFolderable 
type SpecialFolderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    SetName(value *string)()
}
