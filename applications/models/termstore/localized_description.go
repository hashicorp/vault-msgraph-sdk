package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalizedDescription 
type LocalizedDescription struct {
    // The description in the localized language.
    description *string
    // The language tag for the label.
    languageTag *string
}
// NewLocalizedDescription instantiates a new localizedDescription and sets the default values.
func NewLocalizedDescription()(*LocalizedDescription) {
    m := &LocalizedDescription{
    }
    return m
}
// CreateLocalizedDescriptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocalizedDescriptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalizedDescription(), nil
}
// GetDescription gets the description property value. The description in the localized language.
func (m *LocalizedDescription) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocalizedDescription) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. The language tag for the label.
func (m *LocalizedDescription) GetLanguageTag()(*string) {
    return m.languageTag
}
// Serialize serializes information the current object
func (m *LocalizedDescription) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description in the localized language.
func (m *LocalizedDescription) SetDescription(value *string)() {
    m.description = value
}
// SetLanguageTag sets the languageTag property value. The language tag for the label.
func (m *LocalizedDescription) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// LocalizedDescriptionable 
type LocalizedDescriptionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetLanguageTag()(*string)
    SetDescription(value *string)()
    SetLanguageTag(value *string)()
}
