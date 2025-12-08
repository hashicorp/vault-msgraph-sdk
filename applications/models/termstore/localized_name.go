// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalizedName 
type LocalizedName struct {
    // The language tag for the label.
    languageTag *string
    // The name in the localized language.
    name *string
}
// NewLocalizedName instantiates a new localizedName and sets the default values.
func NewLocalizedName()(*LocalizedName) {
    m := &LocalizedName{
    }
    return m
}
// CreateLocalizedNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocalizedNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalizedName(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocalizedName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetLanguageTag gets the languageTag property value. The language tag for the label.
func (m *LocalizedName) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetName gets the name property value. The name in the localized language.
func (m *LocalizedName) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *LocalizedName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetLanguageTag sets the languageTag property value. The language tag for the label.
func (m *LocalizedName) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetName sets the name property value. The name in the localized language.
func (m *LocalizedName) SetName(value *string)() {
    m.name = value
}
// LocalizedNameable 
type LocalizedNameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLanguageTag()(*string)
    GetName()(*string)
    SetLanguageTag(value *string)()
    SetName(value *string)()
}
