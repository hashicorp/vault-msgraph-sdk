package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemReference 
type ItemReference struct {
    // Unique identifier of the drive instance that contains the driveItem. Only returned if the item is located in a [drive][]. Read-only.
    driveId *string
    // Identifies the type of drive. Only returned if the item is located in a [drive][]. See [drive][] resource for values.
    driveType *string
    // Unique identifier of the driveItem in the drive or a listItem in a list. Read-only.
    id *string
    // The name of the item being referenced. Read-only.
    name *string
    // Path that can be used to navigate to the item. Read-only.
    path *string
    // A unique identifier for a shared resource that can be accessed via the [Shares][] API.
    shareId *string
    // The sharepointIds property
    sharepointIds SharepointIdsable
    // For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent document library of the driveItem resource or the parent list of the listItem resource. The value is the same as the id property of that [site][] resource. It is an opaque string that consists of three identifiers of the site. For OneDrive, this property is not populated.
    siteId *string
}
// NewItemReference instantiates a new itemReference and sets the default values.
func NewItemReference()(*ItemReference) {
    m := &ItemReference{
    }
    return m
}
// CreateItemReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemReference(), nil
}
// GetDriveId gets the driveId property value. Unique identifier of the drive instance that contains the driveItem. Only returned if the item is located in a [drive][]. Read-only.
func (m *ItemReference) GetDriveId()(*string) {
    return m.driveId
}
// GetDriveType gets the driveType property value. Identifies the type of drive. Only returned if the item is located in a [drive][]. See [drive][] resource for values.
func (m *ItemReference) GetDriveType()(*string) {
    return m.driveType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["driveId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveId(val)
        }
        return nil
    }
    res["driveType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveType(val)
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["shareId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareId(val)
        }
        return nil
    }
    res["sharepointIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["siteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier of the driveItem in the drive or a listItem in a list. Read-only.
func (m *ItemReference) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name of the item being referenced. Read-only.
func (m *ItemReference) GetName()(*string) {
    return m.name
}
// GetPath gets the path property value. Path that can be used to navigate to the item. Read-only.
func (m *ItemReference) GetPath()(*string) {
    return m.path
}
// GetShareId gets the shareId property value. A unique identifier for a shared resource that can be accessed via the [Shares][] API.
func (m *ItemReference) GetShareId()(*string) {
    return m.shareId
}
// GetSharepointIds gets the sharepointIds property value. The sharepointIds property
func (m *ItemReference) GetSharepointIds()(SharepointIdsable) {
    return m.sharepointIds
}
// GetSiteId gets the siteId property value. For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent document library of the driveItem resource or the parent list of the listItem resource. The value is the same as the id property of that [site][] resource. It is an opaque string that consists of three identifiers of the site. For OneDrive, this property is not populated.
func (m *ItemReference) GetSiteId()(*string) {
    return m.siteId
}
// Serialize serializes information the current object
func (m *ItemReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("driveId", m.GetDriveId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("driveType", m.GetDriveType())
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
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("shareId", m.GetShareId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteId", m.GetSiteId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDriveId sets the driveId property value. Unique identifier of the drive instance that contains the driveItem. Only returned if the item is located in a [drive][]. Read-only.
func (m *ItemReference) SetDriveId(value *string)() {
    m.driveId = value
}
// SetDriveType sets the driveType property value. Identifies the type of drive. Only returned if the item is located in a [drive][]. See [drive][] resource for values.
func (m *ItemReference) SetDriveType(value *string)() {
    m.driveType = value
}
// SetId sets the id property value. Unique identifier of the driveItem in the drive or a listItem in a list. Read-only.
func (m *ItemReference) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name of the item being referenced. Read-only.
func (m *ItemReference) SetName(value *string)() {
    m.name = value
}
// SetPath sets the path property value. Path that can be used to navigate to the item. Read-only.
func (m *ItemReference) SetPath(value *string)() {
    m.path = value
}
// SetShareId sets the shareId property value. A unique identifier for a shared resource that can be accessed via the [Shares][] API.
func (m *ItemReference) SetShareId(value *string)() {
    m.shareId = value
}
// SetSharepointIds sets the sharepointIds property value. The sharepointIds property
func (m *ItemReference) SetSharepointIds(value SharepointIdsable)() {
    m.sharepointIds = value
}
// SetSiteId sets the siteId property value. For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent document library of the driveItem resource or the parent list of the listItem resource. The value is the same as the id property of that [site][] resource. It is an opaque string that consists of three identifiers of the site. For OneDrive, this property is not populated.
func (m *ItemReference) SetSiteId(value *string)() {
    m.siteId = value
}
// ItemReferenceable 
type ItemReferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDriveId()(*string)
    GetDriveType()(*string)
    GetId()(*string)
    GetName()(*string)
    GetPath()(*string)
    GetShareId()(*string)
    GetSharepointIds()(SharepointIdsable)
    GetSiteId()(*string)
    SetDriveId(value *string)()
    SetDriveType(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetPath(value *string)()
    SetShareId(value *string)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteId(value *string)()
}
