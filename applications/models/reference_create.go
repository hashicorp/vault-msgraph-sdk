package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferenceCreate 
type ReferenceCreate struct {
    // The OdataId property
    odataId *string
}
// NewReferenceCreate instantiates a new ReferenceCreate and sets the default values.
func NewReferenceCreate()(*ReferenceCreate) {
    m := &ReferenceCreate{
    }
    return m
}
// CreateReferenceCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReferenceCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferenceCreate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReferenceCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataId(val)
        }
        return nil
    }
    return res
}
// GetOdataId gets the @odata.id property value. The OdataId property
func (m *ReferenceCreate) GetOdataId()(*string) {
    return m.odataId
}
// Serialize serializes information the current object
func (m *ReferenceCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.id", m.GetOdataId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataId sets the @odata.id property value. The OdataId property
func (m *ReferenceCreate) SetOdataId(value *string)() {
    m.odataId = value
}
// ReferenceCreateable 
type ReferenceCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataId()(*string)
    SetOdataId(value *string)()
}
