// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VisualInfo 
type VisualInfo struct {
    // The attribution property
    attribution ImageInfoable
    // Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
    backgroundColor *string
    // The content property
    content Jsonable
    // Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
    description *string
    // Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
    displayText *string
}
// NewVisualInfo instantiates a new visualInfo and sets the default values.
func NewVisualInfo()(*VisualInfo) {
    m := &VisualInfo{
    }
    return m
}
// CreateVisualInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVisualInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVisualInfo(), nil
}
// GetAttribution gets the attribution property value. The attribution property
func (m *VisualInfo) GetAttribution()(ImageInfoable) {
    return m.attribution
}
// GetBackgroundColor gets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) GetBackgroundColor()(*string) {
    return m.backgroundColor
}
// GetContent gets the content property value. The content property
func (m *VisualInfo) GetContent()(Jsonable) {
    return m.content
}
// GetDescription gets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) GetDescription()(*string) {
    return m.description
}
// GetDisplayText gets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) GetDisplayText()(*string) {
    return m.displayText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VisualInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attribution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImageInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribution(val.(ImageInfoable))
        }
        return nil
    }
    res["backgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(Jsonable))
        }
        return nil
    }
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
    res["displayText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayText(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *VisualInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attribution", m.GetAttribution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayText", m.GetDisplayText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttribution sets the attribution property value. The attribution property
func (m *VisualInfo) SetAttribution(value ImageInfoable)() {
    m.attribution = value
}
// SetBackgroundColor sets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// SetContent sets the content property value. The content property
func (m *VisualInfo) SetContent(value Jsonable)() {
    m.content = value
}
// SetDescription sets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayText sets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) SetDisplayText(value *string)() {
    m.displayText = value
}
// VisualInfoable 
type VisualInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttribution()(ImageInfoable)
    GetBackgroundColor()(*string)
    GetContent()(Jsonable)
    GetDescription()(*string)
    GetDisplayText()(*string)
    SetAttribution(value ImageInfoable)()
    SetBackgroundColor(value *string)()
    SetContent(value Jsonable)()
    SetDescription(value *string)()
    SetDisplayText(value *string)()
}
