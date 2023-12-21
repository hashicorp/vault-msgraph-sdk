package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalizedLabel 
type LocalizedLabel struct {
    // Indicates whether the label is the default label.
    isDefault *bool
    // The language tag for the label.
    languageTag *string
    // The name of the label.
    name *string
}
// NewLocalizedLabel instantiates a new localizedLabel and sets the default values.
func NewLocalizedLabel()(*LocalizedLabel) {
    m := &LocalizedLabel{
    }
    return m
}
// CreateLocalizedLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocalizedLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalizedLabel(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocalizedLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates whether the label is the default label.
func (m *LocalizedLabel) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetLanguageTag gets the languageTag property value. The language tag for the label.
func (m *LocalizedLabel) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetName gets the name property value. The name of the label.
func (m *LocalizedLabel) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *LocalizedLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
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
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. Indicates whether the label is the default label.
func (m *LocalizedLabel) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetLanguageTag sets the languageTag property value. The language tag for the label.
func (m *LocalizedLabel) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetName sets the name property value. The name of the label.
func (m *LocalizedLabel) SetName(value *string)() {
    m.name = value
}
// LocalizedLabelable 
type LocalizedLabelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDefault()(*bool)
    GetLanguageTag()(*string)
    GetName()(*string)
    SetIsDefault(value *bool)()
    SetLanguageTag(value *string)()
    SetName(value *string)()
}
