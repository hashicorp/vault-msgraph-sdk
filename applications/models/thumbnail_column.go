package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThumbnailColumn 
type ThumbnailColumn struct {
}
// NewThumbnailColumn instantiates a new thumbnailColumn and sets the default values.
func NewThumbnailColumn()(*ThumbnailColumn) {
    m := &ThumbnailColumn{
    }
    return m
}
// CreateThumbnailColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThumbnailColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThumbnailColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThumbnailColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ThumbnailColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// ThumbnailColumnable 
type ThumbnailColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
