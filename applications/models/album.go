// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Album 
type Album struct {
    // Unique identifier of the [driveItem][] that is the cover of the album.
    coverImageItemId *string
}
// NewAlbum instantiates a new album and sets the default values.
func NewAlbum()(*Album) {
    m := &Album{
    }
    return m
}
// CreateAlbumFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlbumFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlbum(), nil
}
// GetCoverImageItemId gets the coverImageItemId property value. Unique identifier of the [driveItem][] that is the cover of the album.
func (m *Album) GetCoverImageItemId()(*string) {
    return m.coverImageItemId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Album) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["coverImageItemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoverImageItemId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Album) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("coverImageItemId", m.GetCoverImageItemId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCoverImageItemId sets the coverImageItemId property value. Unique identifier of the [driveItem][] that is the cover of the album.
func (m *Album) SetCoverImageItemId(value *string)() {
    m.coverImageItemId = value
}
// Albumable 
type Albumable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCoverImageItemId()(*string)
    SetCoverImageItemId(value *string)()
}
