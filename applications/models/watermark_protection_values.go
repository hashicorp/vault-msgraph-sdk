// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WatermarkProtectionValues 
type WatermarkProtectionValues struct {
    // Indicates whether to apply a watermark to any shared content.
    isEnabledForContentSharing *bool
    // Indicates whether to apply a watermark to everyone's video feed.
    isEnabledForVideo *bool
}
// NewWatermarkProtectionValues instantiates a new watermarkProtectionValues and sets the default values.
func NewWatermarkProtectionValues()(*WatermarkProtectionValues) {
    m := &WatermarkProtectionValues{
    }
    return m
}
// CreateWatermarkProtectionValuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWatermarkProtectionValuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWatermarkProtectionValues(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WatermarkProtectionValues) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isEnabledForContentSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForContentSharing(val)
        }
        return nil
    }
    res["isEnabledForVideo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForVideo(val)
        }
        return nil
    }
    return res
}
// GetIsEnabledForContentSharing gets the isEnabledForContentSharing property value. Indicates whether to apply a watermark to any shared content.
func (m *WatermarkProtectionValues) GetIsEnabledForContentSharing()(*bool) {
    return m.isEnabledForContentSharing
}
// GetIsEnabledForVideo gets the isEnabledForVideo property value. Indicates whether to apply a watermark to everyone's video feed.
func (m *WatermarkProtectionValues) GetIsEnabledForVideo()(*bool) {
    return m.isEnabledForVideo
}
// Serialize serializes information the current object
func (m *WatermarkProtectionValues) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabledForContentSharing", m.GetIsEnabledForContentSharing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabledForVideo", m.GetIsEnabledForVideo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsEnabledForContentSharing sets the isEnabledForContentSharing property value. Indicates whether to apply a watermark to any shared content.
func (m *WatermarkProtectionValues) SetIsEnabledForContentSharing(value *bool)() {
    m.isEnabledForContentSharing = value
}
// SetIsEnabledForVideo sets the isEnabledForVideo property value. Indicates whether to apply a watermark to everyone's video feed.
func (m *WatermarkProtectionValues) SetIsEnabledForVideo(value *bool)() {
    m.isEnabledForVideo = value
}
// WatermarkProtectionValuesable 
type WatermarkProtectionValuesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEnabledForContentSharing()(*bool)
    GetIsEnabledForVideo()(*bool)
    SetIsEnabledForContentSharing(value *bool)()
    SetIsEnabledForVideo(value *bool)()
}
