package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalLink 
type ExternalLink struct {
    // The URL of the link.
    href *string
}
// NewExternalLink instantiates a new externalLink and sets the default values.
func NewExternalLink()(*ExternalLink) {
    m := &ExternalLink{
    }
    return m
}
// CreateExternalLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalLink(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["href"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHref(val)
        }
        return nil
    }
    return res
}
// GetHref gets the href property value. The URL of the link.
func (m *ExternalLink) GetHref()(*string) {
    return m.href
}
// Serialize serializes information the current object
func (m *ExternalLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("href", m.GetHref())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHref sets the href property value. The URL of the link.
func (m *ExternalLink) SetHref(value *string)() {
    m.href = value
}
// ExternalLinkable 
type ExternalLinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHref()(*string)
    SetHref(value *string)()
}
