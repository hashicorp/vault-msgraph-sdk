package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentTypeOrder 
type ContentTypeOrder struct {
    // Indicates whether this is the default content type
    defaultEscaped *bool
    // Specifies the position in which the content type appears in the selection UI.
    position *int32
}
// NewContentTypeOrder instantiates a new contentTypeOrder and sets the default values.
func NewContentTypeOrder()(*ContentTypeOrder) {
    m := &ContentTypeOrder{
    }
    return m
}
// CreateContentTypeOrderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentTypeOrderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentTypeOrder(), nil
}
// GetDefaultEscaped gets the default property value. Indicates whether this is the default content type
func (m *ContentTypeOrder) GetDefaultEscaped()(*bool) {
    return m.defaultEscaped
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentTypeOrder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEscaped(val)
        }
        return nil
    }
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetPosition gets the position property value. Specifies the position in which the content type appears in the selection UI.
func (m *ContentTypeOrder) GetPosition()(*int32) {
    return m.position
}
// Serialize serializes information the current object
func (m *ContentTypeOrder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("default", m.GetDefaultEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultEscaped sets the default property value. Indicates whether this is the default content type
func (m *ContentTypeOrder) SetDefaultEscaped(value *bool)() {
    m.defaultEscaped = value
}
// SetPosition sets the position property value. Specifies the position in which the content type appears in the selection UI.
func (m *ContentTypeOrder) SetPosition(value *int32)() {
    m.position = value
}
// ContentTypeOrderable 
type ContentTypeOrderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultEscaped()(*bool)
    GetPosition()(*int32)
    SetDefaultEscaped(value *bool)()
    SetPosition(value *int32)()
}
