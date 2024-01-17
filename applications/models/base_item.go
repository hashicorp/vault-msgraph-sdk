// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseItem 
type BaseItem struct {
    Entity
    // The createdBy property
    createdBy IdentitySetable
    // The createdByUser property
    createdByUser Userable
    // Date and time of item creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Provides a user-visible description of the item. Optional.
    description *string
    // ETag for the item. Read-only.
    eTag *string
    // The lastModifiedBy property
    lastModifiedBy IdentitySetable
    // The lastModifiedByUser property
    lastModifiedByUser Userable
    // Date and time the item was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the item. Read-write.
    name *string
    // The parentReference property
    parentReference ItemReferenceable
    // URL that either displays the resource in the browser (for Office file formats), or is a direct link to the file (for other formats). Read-only.
    webUrl *string
}
// NewBaseItem instantiates a new baseItem and sets the default values.
func NewBaseItem()(*BaseItem) {
    m := &BaseItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBaseItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBaseItem(), nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *BaseItem) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedByUser gets the createdByUser property value. The createdByUser property
func (m *BaseItem) GetCreatedByUser()(Userable) {
    return m.createdByUser
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *BaseItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Provides a user-visible description of the item. Optional.
func (m *BaseItem) GetDescription()(*string) {
    return m.description
}
// GetETag gets the eTag property value. ETag for the item. Read-only.
func (m *BaseItem) GetETag()(*string) {
    return m.eTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUser(val.(Userable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["eTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetETag(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedByUser(val.(Userable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["parentReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentReference(val.(ItemReferenceable))
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *BaseItem) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedByUser gets the lastModifiedByUser property value. The lastModifiedByUser property
func (m *BaseItem) GetLastModifiedByUser()(Userable) {
    return m.lastModifiedByUser
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *BaseItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. The name of the item. Read-write.
func (m *BaseItem) GetName()(*string) {
    return m.name
}
// GetParentReference gets the parentReference property value. The parentReference property
func (m *BaseItem) GetParentReference()(ItemReferenceable) {
    return m.parentReference
}
// GetWebUrl gets the webUrl property value. URL that either displays the resource in the browser (for Office file formats), or is a direct link to the file (for other formats). Read-only.
func (m *BaseItem) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *BaseItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdByUser", m.GetCreatedByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eTag", m.GetETag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedByUser", m.GetLastModifiedByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentReference", m.GetParentReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *BaseItem) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedByUser sets the createdByUser property value. The createdByUser property
func (m *BaseItem) SetCreatedByUser(value Userable)() {
    m.createdByUser = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *BaseItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Provides a user-visible description of the item. Optional.
func (m *BaseItem) SetDescription(value *string)() {
    m.description = value
}
// SetETag sets the eTag property value. ETag for the item. Read-only.
func (m *BaseItem) SetETag(value *string)() {
    m.eTag = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *BaseItem) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedByUser sets the lastModifiedByUser property value. The lastModifiedByUser property
func (m *BaseItem) SetLastModifiedByUser(value Userable)() {
    m.lastModifiedByUser = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *BaseItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. The name of the item. Read-write.
func (m *BaseItem) SetName(value *string)() {
    m.name = value
}
// SetParentReference sets the parentReference property value. The parentReference property
func (m *BaseItem) SetParentReference(value ItemReferenceable)() {
    m.parentReference = value
}
// SetWebUrl sets the webUrl property value. URL that either displays the resource in the browser (for Office file formats), or is a direct link to the file (for other formats). Read-only.
func (m *BaseItem) SetWebUrl(value *string)() {
    m.webUrl = value
}
// BaseItemable 
type BaseItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedByUser()(Userable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetETag()(*string)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedByUser()(Userable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetParentReference()(ItemReferenceable)
    GetWebUrl()(*string)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedByUser(value Userable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetETag(value *string)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedByUser(value Userable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetParentReference(value ItemReferenceable)()
    SetWebUrl(value *string)()
}
