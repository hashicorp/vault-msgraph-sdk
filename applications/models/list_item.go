package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListItem 
type ListItem struct {
    BaseItem
    // The analytics property
    analytics ItemAnalyticsable
    // The contentType property
    contentType ContentTypeInfoable
    // Version information for a document set version created by a user.
    documentSetVersions []DocumentSetVersionable
    // The driveItem property
    driveItem DriveItemable
    // The fields property
    fields FieldValueSetable
    // The sharepointIds property
    sharepointIds SharepointIdsable
    // The list of previous versions of the list item.
    versions []ListItemVersionable
}
// NewListItem instantiates a new listItem and sets the default values.
func NewListItem()(*ListItem) {
    m := &ListItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateListItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListItem(), nil
}
// GetAnalytics gets the analytics property value. The analytics property
func (m *ListItem) GetAnalytics()(ItemAnalyticsable) {
    return m.analytics
}
// GetContentType gets the contentType property value. The contentType property
func (m *ListItem) GetContentType()(ContentTypeInfoable) {
    return m.contentType
}
// GetDocumentSetVersions gets the documentSetVersions property value. Version information for a document set version created by a user.
func (m *ListItem) GetDocumentSetVersions()([]DocumentSetVersionable) {
    return m.documentSetVersions
}
// GetDriveItem gets the driveItem property value. The driveItem property
func (m *ListItem) GetDriveItem()(DriveItemable) {
    return m.driveItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalytics(val.(ItemAnalyticsable))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
        }
        return nil
    }
    res["documentSetVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDocumentSetVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentSetVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DocumentSetVersionable)
                }
            }
            m.SetDocumentSetVersions(res)
        }
        return nil
    }
    res["driveItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFieldValueSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(FieldValueSetable))
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
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListItemVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItemVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ListItemVersionable)
                }
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The fields property
func (m *ListItem) GetFields()(FieldValueSetable) {
    return m.fields
}
// GetSharepointIds gets the sharepointIds property value. The sharepointIds property
func (m *ListItem) GetSharepointIds()(SharepointIdsable) {
    return m.sharepointIds
}
// GetVersions gets the versions property value. The list of previous versions of the list item.
func (m *ListItem) GetVersions()([]ListItemVersionable) {
    return m.versions
}
// Serialize serializes information the current object
func (m *ListItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("analytics", m.GetAnalytics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    if m.GetDocumentSetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDocumentSetVersions()))
        for i, v := range m.GetDocumentSetVersions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("documentSetVersions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnalytics sets the analytics property value. The analytics property
func (m *ListItem) SetAnalytics(value ItemAnalyticsable)() {
    m.analytics = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *ListItem) SetContentType(value ContentTypeInfoable)() {
    m.contentType = value
}
// SetDocumentSetVersions sets the documentSetVersions property value. Version information for a document set version created by a user.
func (m *ListItem) SetDocumentSetVersions(value []DocumentSetVersionable)() {
    m.documentSetVersions = value
}
// SetDriveItem sets the driveItem property value. The driveItem property
func (m *ListItem) SetDriveItem(value DriveItemable)() {
    m.driveItem = value
}
// SetFields sets the fields property value. The fields property
func (m *ListItem) SetFields(value FieldValueSetable)() {
    m.fields = value
}
// SetSharepointIds sets the sharepointIds property value. The sharepointIds property
func (m *ListItem) SetSharepointIds(value SharepointIdsable)() {
    m.sharepointIds = value
}
// SetVersions sets the versions property value. The list of previous versions of the list item.
func (m *ListItem) SetVersions(value []ListItemVersionable)() {
    m.versions = value
}
// ListItemable 
type ListItemable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetContentType()(ContentTypeInfoable)
    GetDocumentSetVersions()([]DocumentSetVersionable)
    GetDriveItem()(DriveItemable)
    GetFields()(FieldValueSetable)
    GetSharepointIds()(SharepointIdsable)
    GetVersions()([]ListItemVersionable)
    SetAnalytics(value ItemAnalyticsable)()
    SetContentType(value ContentTypeInfoable)()
    SetDocumentSetVersions(value []DocumentSetVersionable)()
    SetDriveItem(value DriveItemable)()
    SetFields(value FieldValueSetable)()
    SetSharepointIds(value SharepointIdsable)()
    SetVersions(value []ListItemVersionable)()
}
