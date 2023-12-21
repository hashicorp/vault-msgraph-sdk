package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppIdentifier the identifier for a mobile app.
type MobileAppIdentifier struct {
}
// NewMobileAppIdentifier instantiates a new mobileAppIdentifier and sets the default values.
func NewMobileAppIdentifier()(*MobileAppIdentifier) {
    m := &MobileAppIdentifier{
    }
    return m
}
// CreateMobileAppIdentifierFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppIdentifierFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppIdentifier(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIdentifier) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *MobileAppIdentifier) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// MobileAppIdentifierable 
type MobileAppIdentifierable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
