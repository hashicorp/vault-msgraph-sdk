package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Image 
type Image struct {
    // Optional. Height of the image, in pixels. Read-only.
    height *int32
    // Optional. Width of the image, in pixels. Read-only.
    width *int32
}
// NewImage instantiates a new image and sets the default values.
func NewImage()(*Image) {
    m := &Image{
    }
    return m
}
// CreateImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Image) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetHeight gets the height property value. Optional. Height of the image, in pixels. Read-only.
func (m *Image) GetHeight()(*int32) {
    return m.height
}
// GetWidth gets the width property value. Optional. Width of the image, in pixels. Read-only.
func (m *Image) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *Image) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHeight sets the height property value. Optional. Height of the image, in pixels. Read-only.
func (m *Image) SetHeight(value *int32)() {
    m.height = value
}
// SetWidth sets the width property value. Optional. Width of the image, in pixels. Read-only.
func (m *Image) SetWidth(value *int32)() {
    m.width = value
}
// Imageable 
type Imageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeight()(*int32)
    GetWidth()(*int32)
    SetHeight(value *int32)()
    SetWidth(value *int32)()
}
