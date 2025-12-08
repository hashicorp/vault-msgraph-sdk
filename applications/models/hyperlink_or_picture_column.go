// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HyperlinkOrPictureColumn 
type HyperlinkOrPictureColumn struct {
    // Specifies whether the display format used for URL columns is an image or a hyperlink.
    isPicture *bool
}
// NewHyperlinkOrPictureColumn instantiates a new hyperlinkOrPictureColumn and sets the default values.
func NewHyperlinkOrPictureColumn()(*HyperlinkOrPictureColumn) {
    m := &HyperlinkOrPictureColumn{
    }
    return m
}
// CreateHyperlinkOrPictureColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHyperlinkOrPictureColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHyperlinkOrPictureColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HyperlinkOrPictureColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isPicture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPicture(val)
        }
        return nil
    }
    return res
}
// GetIsPicture gets the isPicture property value. Specifies whether the display format used for URL columns is an image or a hyperlink.
func (m *HyperlinkOrPictureColumn) GetIsPicture()(*bool) {
    return m.isPicture
}
// Serialize serializes information the current object
func (m *HyperlinkOrPictureColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isPicture", m.GetIsPicture())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsPicture sets the isPicture property value. Specifies whether the display format used for URL columns is an image or a hyperlink.
func (m *HyperlinkOrPictureColumn) SetIsPicture(value *bool)() {
    m.isPicture = value
}
// HyperlinkOrPictureColumnable 
type HyperlinkOrPictureColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsPicture()(*bool)
    SetIsPicture(value *bool)()
}
