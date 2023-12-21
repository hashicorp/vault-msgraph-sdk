package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionLabelSettings 
type RetentionLabelSettings struct {
    // Specifies whether updates to document content are allowed. Read-only.
    isContentUpdateAllowed *bool
    // Specifies whether the document deletion is allowed. Read-only.
    isDeleteAllowed *bool
    // Specifies whether you're allowed to change the retention label on the document. Read-only.
    isLabelUpdateAllowed *bool
    // Specifies whether updates to the item metadata (for example, the Title field) are blocked. Read-only.
    isMetadataUpdateAllowed *bool
    // Specifies whether the item is locked. Read-write.
    isRecordLocked *bool
}
// NewRetentionLabelSettings instantiates a new retentionLabelSettings and sets the default values.
func NewRetentionLabelSettings()(*RetentionLabelSettings) {
    m := &RetentionLabelSettings{
    }
    return m
}
// CreateRetentionLabelSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionLabelSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionLabelSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionLabelSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isContentUpdateAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentUpdateAllowed(val)
        }
        return nil
    }
    res["isDeleteAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleteAllowed(val)
        }
        return nil
    }
    res["isLabelUpdateAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLabelUpdateAllowed(val)
        }
        return nil
    }
    res["isMetadataUpdateAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMetadataUpdateAllowed(val)
        }
        return nil
    }
    res["isRecordLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRecordLocked(val)
        }
        return nil
    }
    return res
}
// GetIsContentUpdateAllowed gets the isContentUpdateAllowed property value. Specifies whether updates to document content are allowed. Read-only.
func (m *RetentionLabelSettings) GetIsContentUpdateAllowed()(*bool) {
    return m.isContentUpdateAllowed
}
// GetIsDeleteAllowed gets the isDeleteAllowed property value. Specifies whether the document deletion is allowed. Read-only.
func (m *RetentionLabelSettings) GetIsDeleteAllowed()(*bool) {
    return m.isDeleteAllowed
}
// GetIsLabelUpdateAllowed gets the isLabelUpdateAllowed property value. Specifies whether you're allowed to change the retention label on the document. Read-only.
func (m *RetentionLabelSettings) GetIsLabelUpdateAllowed()(*bool) {
    return m.isLabelUpdateAllowed
}
// GetIsMetadataUpdateAllowed gets the isMetadataUpdateAllowed property value. Specifies whether updates to the item metadata (for example, the Title field) are blocked. Read-only.
func (m *RetentionLabelSettings) GetIsMetadataUpdateAllowed()(*bool) {
    return m.isMetadataUpdateAllowed
}
// GetIsRecordLocked gets the isRecordLocked property value. Specifies whether the item is locked. Read-write.
func (m *RetentionLabelSettings) GetIsRecordLocked()(*bool) {
    return m.isRecordLocked
}
// Serialize serializes information the current object
func (m *RetentionLabelSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isContentUpdateAllowed", m.GetIsContentUpdateAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDeleteAllowed", m.GetIsDeleteAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isLabelUpdateAllowed", m.GetIsLabelUpdateAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMetadataUpdateAllowed", m.GetIsMetadataUpdateAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRecordLocked", m.GetIsRecordLocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsContentUpdateAllowed sets the isContentUpdateAllowed property value. Specifies whether updates to document content are allowed. Read-only.
func (m *RetentionLabelSettings) SetIsContentUpdateAllowed(value *bool)() {
    m.isContentUpdateAllowed = value
}
// SetIsDeleteAllowed sets the isDeleteAllowed property value. Specifies whether the document deletion is allowed. Read-only.
func (m *RetentionLabelSettings) SetIsDeleteAllowed(value *bool)() {
    m.isDeleteAllowed = value
}
// SetIsLabelUpdateAllowed sets the isLabelUpdateAllowed property value. Specifies whether you're allowed to change the retention label on the document. Read-only.
func (m *RetentionLabelSettings) SetIsLabelUpdateAllowed(value *bool)() {
    m.isLabelUpdateAllowed = value
}
// SetIsMetadataUpdateAllowed sets the isMetadataUpdateAllowed property value. Specifies whether updates to the item metadata (for example, the Title field) are blocked. Read-only.
func (m *RetentionLabelSettings) SetIsMetadataUpdateAllowed(value *bool)() {
    m.isMetadataUpdateAllowed = value
}
// SetIsRecordLocked sets the isRecordLocked property value. Specifies whether the item is locked. Read-write.
func (m *RetentionLabelSettings) SetIsRecordLocked(value *bool)() {
    m.isRecordLocked = value
}
// RetentionLabelSettingsable 
type RetentionLabelSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsContentUpdateAllowed()(*bool)
    GetIsDeleteAllowed()(*bool)
    GetIsLabelUpdateAllowed()(*bool)
    GetIsMetadataUpdateAllowed()(*bool)
    GetIsRecordLocked()(*bool)
    SetIsContentUpdateAllowed(value *bool)()
    SetIsDeleteAllowed(value *bool)()
    SetIsLabelUpdateAllowed(value *bool)()
    SetIsMetadataUpdateAllowed(value *bool)()
    SetIsRecordLocked(value *bool)()
}
