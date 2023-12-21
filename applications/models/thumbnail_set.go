package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThumbnailSet 
type ThumbnailSet struct {
    Entity
    // The large property
    large Thumbnailable
    // The medium property
    medium Thumbnailable
    // The small property
    small Thumbnailable
    // The source property
    source Thumbnailable
}
// NewThumbnailSet instantiates a new thumbnailSet and sets the default values.
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateThumbnailSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThumbnailSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThumbnailSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThumbnailSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["large"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLarge(val.(Thumbnailable))
        }
        return nil
    }
    res["medium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedium(val.(Thumbnailable))
        }
        return nil
    }
    res["small"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmall(val.(Thumbnailable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(Thumbnailable))
        }
        return nil
    }
    return res
}
// GetLarge gets the large property value. The large property
func (m *ThumbnailSet) GetLarge()(Thumbnailable) {
    return m.large
}
// GetMedium gets the medium property value. The medium property
func (m *ThumbnailSet) GetMedium()(Thumbnailable) {
    return m.medium
}
// GetSmall gets the small property value. The small property
func (m *ThumbnailSet) GetSmall()(Thumbnailable) {
    return m.small
}
// GetSource gets the source property value. The source property
func (m *ThumbnailSet) GetSource()(Thumbnailable) {
    return m.source
}
// Serialize serializes information the current object
func (m *ThumbnailSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("large", m.GetLarge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("medium", m.GetMedium())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("small", m.GetSmall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLarge sets the large property value. The large property
func (m *ThumbnailSet) SetLarge(value Thumbnailable)() {
    m.large = value
}
// SetMedium sets the medium property value. The medium property
func (m *ThumbnailSet) SetMedium(value Thumbnailable)() {
    m.medium = value
}
// SetSmall sets the small property value. The small property
func (m *ThumbnailSet) SetSmall(value Thumbnailable)() {
    m.small = value
}
// SetSource sets the source property value. The source property
func (m *ThumbnailSet) SetSource(value Thumbnailable)() {
    m.source = value
}
// ThumbnailSetable 
type ThumbnailSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLarge()(Thumbnailable)
    GetMedium()(Thumbnailable)
    GetSmall()(Thumbnailable)
    GetSource()(Thumbnailable)
    SetLarge(value Thumbnailable)()
    SetMedium(value Thumbnailable)()
    SetSmall(value Thumbnailable)()
    SetSource(value Thumbnailable)()
}
