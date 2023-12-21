package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SettingSource 
type SettingSource struct {
    // Not yet documented
    displayName *string
    // Not yet documented
    id *string
    // The sourceType property
    sourceType *SettingSourceType
}
// NewSettingSource instantiates a new settingSource and sets the default values.
func NewSettingSource()(*SettingSource) {
    m := &SettingSource{
    }
    return m
}
// CreateSettingSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettingSource(), nil
}
// GetDisplayName gets the displayName property value. Not yet documented
func (m *SettingSource) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["sourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSettingSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceType(val.(*SettingSourceType))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Not yet documented
func (m *SettingSource) GetId()(*string) {
    return m.id
}
// GetSourceType gets the sourceType property value. The sourceType property
func (m *SettingSource) GetSourceType()(*SettingSourceType) {
    return m.sourceType
}
// Serialize serializes information the current object
func (m *SettingSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := (*m.GetSourceType()).String()
        err := writer.WriteStringValue("sourceType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Not yet documented
func (m *SettingSource) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. Not yet documented
func (m *SettingSource) SetId(value *string)() {
    m.id = value
}
// SetSourceType sets the sourceType property value. The sourceType property
func (m *SettingSource) SetSourceType(value *SettingSourceType)() {
    m.sourceType = value
}
// SettingSourceable 
type SettingSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetId()(*string)
    GetSourceType()(*SettingSourceType)
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetSourceType(value *SettingSourceType)()
}
