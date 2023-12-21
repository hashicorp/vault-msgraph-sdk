package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Bundle 
type Bundle struct {
    // The album property
    album Albumable
    // Number of children contained immediately within this container.
    childCount *int32
}
// NewBundle instantiates a new bundle and sets the default values.
func NewBundle()(*Bundle) {
    m := &Bundle{
    }
    return m
}
// CreateBundleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBundleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBundle(), nil
}
// GetAlbum gets the album property value. The album property
func (m *Bundle) GetAlbum()(Albumable) {
    return m.album
}
// GetChildCount gets the childCount property value. Number of children contained immediately within this container.
func (m *Bundle) GetChildCount()(*int32) {
    return m.childCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bundle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["album"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlbumFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbum(val.(Albumable))
        }
        return nil
    }
    res["childCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildCount(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Bundle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("album", m.GetAlbum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("childCount", m.GetChildCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlbum sets the album property value. The album property
func (m *Bundle) SetAlbum(value Albumable)() {
    m.album = value
}
// SetChildCount sets the childCount property value. Number of children contained immediately within this container.
func (m *Bundle) SetChildCount(value *int32)() {
    m.childCount = value
}
// Bundleable 
type Bundleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlbum()(Albumable)
    GetChildCount()(*int32)
    SetAlbum(value Albumable)()
    SetChildCount(value *int32)()
}
