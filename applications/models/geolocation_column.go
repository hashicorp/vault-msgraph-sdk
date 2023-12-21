package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GeolocationColumn 
type GeolocationColumn struct {
}
// NewGeolocationColumn instantiates a new geolocationColumn and sets the default values.
func NewGeolocationColumn()(*GeolocationColumn) {
    m := &GeolocationColumn{
    }
    return m
}
// CreateGeolocationColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGeolocationColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeolocationColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeolocationColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *GeolocationColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// GeolocationColumnable 
type GeolocationColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
