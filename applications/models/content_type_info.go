package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentTypeInfo 
type ContentTypeInfo struct {
    // The ID of the content type.
    id *string
    // The name of the content type.
    name *string
}
// NewContentTypeInfo instantiates a new contentTypeInfo and sets the default values.
func NewContentTypeInfo()(*ContentTypeInfo) {
    m := &ContentTypeInfo{
    }
    return m
}
// CreateContentTypeInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentTypeInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentTypeInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentTypeInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
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
// GetId gets the id property value. The ID of the content type.
func (m *ContentTypeInfo) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name of the content type.
func (m *ContentTypeInfo) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ContentTypeInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetId sets the id property value. The ID of the content type.
func (m *ContentTypeInfo) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name of the content type.
func (m *ContentTypeInfo) SetName(value *string)() {
    m.name = value
}
// ContentTypeInfoable 
type ContentTypeInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetName()(*string)
    SetId(value *string)()
    SetName(value *string)()
}
