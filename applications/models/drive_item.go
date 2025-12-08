// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItem 
type DriveItem struct {
    BaseItem
    // The analytics property
    analytics ItemAnalyticsable
    // The audio property
    audio Audioable
    // The bundle property
    bundle Bundleable
    // Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
    children []DriveItemable
    // The content stream, if the item represents a file.
    content []byte
    // An eTag for the content of the item. This eTag isn't changed if only the metadata is changed. Note This property isn't returned if the item is a folder. Read-only.
    cTag *string
    // The deleted property
    deleted Deletedable
    // The file property
    file Fileable
    // The fileSystemInfo property
    fileSystemInfo FileSystemInfoable
    // The folder property
    folder Folderable
    // The image property
    image Imageable
    // The listItem property
    listItem ListItemable
    // The location property
    location GeoCoordinatesable
    // The malware property
    malware Malwareable
    // The package property
    packageEscaped PackageEscapedable
    // The pendingOperations property
    pendingOperations PendingOperationsable
    // The set of permissions for the item. Read-only. Nullable.
    permissions []Permissionable
    // The photo property
    photo Photoable
    // The publication property
    publication PublicationFacetable
    // The remoteItem property
    remoteItem RemoteItemable
    // The retentionLabel property
    retentionLabel ItemRetentionLabelable
    // The root property
    root Rootable
    // The searchResult property
    searchResult SearchResultable
    // The shared property
    shared Sharedable
    // The sharepointIds property
    sharepointIds SharepointIdsable
    // Size of the item in bytes. Read-only.
    size *int64
    // The specialFolder property
    specialFolder SpecialFolderable
    // The set of subscriptions on the item. Only supported on the root of a drive.
    subscriptions []Subscriptionable
    // Collection of [thumbnailSet][] objects associated with the item. For more information, see [getting thumbnails][]. Read-only. Nullable.
    thumbnails []ThumbnailSetable
    // The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
    versions []DriveItemVersionable
    // The video property
    video Videoable
    // WebDAV compatible URL for the item.
    webDavUrl *string
    // The workbook property
    workbook Workbookable
}
// NewDriveItem instantiates a new driveItem and sets the default values.
func NewDriveItem()(*DriveItem) {
    m := &DriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateDriveItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveItem(), nil
}
// GetAnalytics gets the analytics property value. The analytics property
func (m *DriveItem) GetAnalytics()(ItemAnalyticsable) {
    return m.analytics
}
// GetAudio gets the audio property value. The audio property
func (m *DriveItem) GetAudio()(Audioable) {
    return m.audio
}
// GetBundle gets the bundle property value. The bundle property
func (m *DriveItem) GetBundle()(Bundleable) {
    return m.bundle
}
// GetChildren gets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) GetChildren()([]DriveItemable) {
    return m.children
}
// GetContent gets the content property value. The content stream, if the item represents a file.
func (m *DriveItem) GetContent()([]byte) {
    return m.content
}
// GetCTag gets the cTag property value. An eTag for the content of the item. This eTag isn't changed if only the metadata is changed. Note This property isn't returned if the item is a folder. Read-only.
func (m *DriveItem) GetCTag()(*string) {
    return m.cTag
}
// GetDeleted gets the deleted property value. The deleted property
func (m *DriveItem) GetDeleted()(Deletedable) {
    return m.deleted
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["audio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAudioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudio(val.(Audioable))
        }
        return nil
    }
    res["bundle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBundleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundle(val.(Bundleable))
        }
        return nil
    }
    res["cTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCTag(val)
        }
        return nil
    }
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DriveItemable)
                }
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeletedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val.(Deletedable))
        }
        return nil
    }
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(Fileable))
        }
        return nil
    }
    res["fileSystemInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSystemInfo(val.(FileSystemInfoable))
        }
        return nil
    }
    res["folder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val.(Folderable))
        }
        return nil
    }
    res["image"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImage(val.(Imageable))
        }
        return nil
    }
    res["listItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(ListItemable))
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(GeoCoordinatesable))
        }
        return nil
    }
    res["malware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMalwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMalware(val.(Malwareable))
        }
        return nil
    }
    res["package"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageEscapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageEscaped(val.(PackageEscapedable))
        }
        return nil
    }
    res["pendingOperations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingOperationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOperations(val.(PendingOperationsable))
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Permissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Permissionable)
                }
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["photo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(Photoable))
        }
        return nil
    }
    res["publication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicationFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublication(val.(PublicationFacetable))
        }
        return nil
    }
    res["remoteItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemoteItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteItem(val.(RemoteItemable))
        }
        return nil
    }
    res["retentionLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemRetentionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetentionLabel(val.(ItemRetentionLabelable))
        }
        return nil
    }
    res["root"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(Rootable))
        }
        return nil
    }
    res["searchResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchResult(val.(SearchResultable))
        }
        return nil
    }
    res["shared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShared(val.(Sharedable))
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["specialFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSpecialFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecialFolder(val.(SpecialFolderable))
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubscriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Subscriptionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Subscriptionable)
                }
            }
            m.SetSubscriptions(res)
        }
        return nil
    }
    res["thumbnails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThumbnailSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThumbnailSetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ThumbnailSetable)
                }
            }
            m.SetThumbnails(res)
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DriveItemVersionable)
                }
            }
            m.SetVersions(res)
        }
        return nil
    }
    res["video"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVideoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val.(Videoable))
        }
        return nil
    }
    res["webDavUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebDavUrl(val)
        }
        return nil
    }
    res["workbook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbook(val.(Workbookable))
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. The file property
func (m *DriveItem) GetFile()(Fileable) {
    return m.file
}
// GetFileSystemInfo gets the fileSystemInfo property value. The fileSystemInfo property
func (m *DriveItem) GetFileSystemInfo()(FileSystemInfoable) {
    return m.fileSystemInfo
}
// GetFolder gets the folder property value. The folder property
func (m *DriveItem) GetFolder()(Folderable) {
    return m.folder
}
// GetImage gets the image property value. The image property
func (m *DriveItem) GetImage()(Imageable) {
    return m.image
}
// GetListItem gets the listItem property value. The listItem property
func (m *DriveItem) GetListItem()(ListItemable) {
    return m.listItem
}
// GetLocation gets the location property value. The location property
func (m *DriveItem) GetLocation()(GeoCoordinatesable) {
    return m.location
}
// GetMalware gets the malware property value. The malware property
func (m *DriveItem) GetMalware()(Malwareable) {
    return m.malware
}
// GetPackageEscaped gets the package property value. The package property
func (m *DriveItem) GetPackageEscaped()(PackageEscapedable) {
    return m.packageEscaped
}
// GetPendingOperations gets the pendingOperations property value. The pendingOperations property
func (m *DriveItem) GetPendingOperations()(PendingOperationsable) {
    return m.pendingOperations
}
// GetPermissions gets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) GetPermissions()([]Permissionable) {
    return m.permissions
}
// GetPhoto gets the photo property value. The photo property
func (m *DriveItem) GetPhoto()(Photoable) {
    return m.photo
}
// GetPublication gets the publication property value. The publication property
func (m *DriveItem) GetPublication()(PublicationFacetable) {
    return m.publication
}
// GetRemoteItem gets the remoteItem property value. The remoteItem property
func (m *DriveItem) GetRemoteItem()(RemoteItemable) {
    return m.remoteItem
}
// GetRetentionLabel gets the retentionLabel property value. The retentionLabel property
func (m *DriveItem) GetRetentionLabel()(ItemRetentionLabelable) {
    return m.retentionLabel
}
// GetRoot gets the root property value. The root property
func (m *DriveItem) GetRoot()(Rootable) {
    return m.root
}
// GetSearchResult gets the searchResult property value. The searchResult property
func (m *DriveItem) GetSearchResult()(SearchResultable) {
    return m.searchResult
}
// GetShared gets the shared property value. The shared property
func (m *DriveItem) GetShared()(Sharedable) {
    return m.shared
}
// GetSharepointIds gets the sharepointIds property value. The sharepointIds property
func (m *DriveItem) GetSharepointIds()(SharepointIdsable) {
    return m.sharepointIds
}
// GetSize gets the size property value. Size of the item in bytes. Read-only.
func (m *DriveItem) GetSize()(*int64) {
    return m.size
}
// GetSpecialFolder gets the specialFolder property value. The specialFolder property
func (m *DriveItem) GetSpecialFolder()(SpecialFolderable) {
    return m.specialFolder
}
// GetSubscriptions gets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) GetSubscriptions()([]Subscriptionable) {
    return m.subscriptions
}
// GetThumbnails gets the thumbnails property value. Collection of [thumbnailSet][] objects associated with the item. For more information, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) GetThumbnails()([]ThumbnailSetable) {
    return m.thumbnails
}
// GetVersions gets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) GetVersions()([]DriveItemVersionable) {
    return m.versions
}
// GetVideo gets the video property value. The video property
func (m *DriveItem) GetVideo()(Videoable) {
    return m.video
}
// GetWebDavUrl gets the webDavUrl property value. WebDAV compatible URL for the item.
func (m *DriveItem) GetWebDavUrl()(*string) {
    return m.webDavUrl
}
// GetWorkbook gets the workbook property value. The workbook property
func (m *DriveItem) GetWorkbook()(Workbookable) {
    return m.workbook
}
// Serialize serializes information the current object
func (m *DriveItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("audio", m.GetAudio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bundle", m.GetBundle())
        if err != nil {
            return err
        }
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cTag", m.GetCTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deleted", m.GetDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fileSystemInfo", m.GetFileSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("folder", m.GetFolder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("image", m.GetImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("listItem", m.GetListItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("malware", m.GetMalware())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("package", m.GetPackageEscaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("pendingOperations", m.GetPendingOperations())
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissions()))
        for i, v := range m.GetPermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("permissions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publication", m.GetPublication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("remoteItem", m.GetRemoteItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("retentionLabel", m.GetRetentionLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("root", m.GetRoot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("searchResult", m.GetSearchResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shared", m.GetShared())
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
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("specialFolder", m.GetSpecialFolder())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("subscriptions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetThumbnails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetThumbnails()))
        for i, v := range m.GetThumbnails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("thumbnails", cast)
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
    {
        err = writer.WriteObjectValue("video", m.GetVideo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webDavUrl", m.GetWebDavUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workbook", m.GetWorkbook())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnalytics sets the analytics property value. The analytics property
func (m *DriveItem) SetAnalytics(value ItemAnalyticsable)() {
    m.analytics = value
}
// SetAudio sets the audio property value. The audio property
func (m *DriveItem) SetAudio(value Audioable)() {
    m.audio = value
}
// SetBundle sets the bundle property value. The bundle property
func (m *DriveItem) SetBundle(value Bundleable)() {
    m.bundle = value
}
// SetChildren sets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) SetChildren(value []DriveItemable)() {
    m.children = value
}
// SetContent sets the content property value. The content stream, if the item represents a file.
func (m *DriveItem) SetContent(value []byte)() {
    m.content = value
}
// SetCTag sets the cTag property value. An eTag for the content of the item. This eTag isn't changed if only the metadata is changed. Note This property isn't returned if the item is a folder. Read-only.
func (m *DriveItem) SetCTag(value *string)() {
    m.cTag = value
}
// SetDeleted sets the deleted property value. The deleted property
func (m *DriveItem) SetDeleted(value Deletedable)() {
    m.deleted = value
}
// SetFile sets the file property value. The file property
func (m *DriveItem) SetFile(value Fileable)() {
    m.file = value
}
// SetFileSystemInfo sets the fileSystemInfo property value. The fileSystemInfo property
func (m *DriveItem) SetFileSystemInfo(value FileSystemInfoable)() {
    m.fileSystemInfo = value
}
// SetFolder sets the folder property value. The folder property
func (m *DriveItem) SetFolder(value Folderable)() {
    m.folder = value
}
// SetImage sets the image property value. The image property
func (m *DriveItem) SetImage(value Imageable)() {
    m.image = value
}
// SetListItem sets the listItem property value. The listItem property
func (m *DriveItem) SetListItem(value ListItemable)() {
    m.listItem = value
}
// SetLocation sets the location property value. The location property
func (m *DriveItem) SetLocation(value GeoCoordinatesable)() {
    m.location = value
}
// SetMalware sets the malware property value. The malware property
func (m *DriveItem) SetMalware(value Malwareable)() {
    m.malware = value
}
// SetPackageEscaped sets the package property value. The package property
func (m *DriveItem) SetPackageEscaped(value PackageEscapedable)() {
    m.packageEscaped = value
}
// SetPendingOperations sets the pendingOperations property value. The pendingOperations property
func (m *DriveItem) SetPendingOperations(value PendingOperationsable)() {
    m.pendingOperations = value
}
// SetPermissions sets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) SetPermissions(value []Permissionable)() {
    m.permissions = value
}
// SetPhoto sets the photo property value. The photo property
func (m *DriveItem) SetPhoto(value Photoable)() {
    m.photo = value
}
// SetPublication sets the publication property value. The publication property
func (m *DriveItem) SetPublication(value PublicationFacetable)() {
    m.publication = value
}
// SetRemoteItem sets the remoteItem property value. The remoteItem property
func (m *DriveItem) SetRemoteItem(value RemoteItemable)() {
    m.remoteItem = value
}
// SetRetentionLabel sets the retentionLabel property value. The retentionLabel property
func (m *DriveItem) SetRetentionLabel(value ItemRetentionLabelable)() {
    m.retentionLabel = value
}
// SetRoot sets the root property value. The root property
func (m *DriveItem) SetRoot(value Rootable)() {
    m.root = value
}
// SetSearchResult sets the searchResult property value. The searchResult property
func (m *DriveItem) SetSearchResult(value SearchResultable)() {
    m.searchResult = value
}
// SetShared sets the shared property value. The shared property
func (m *DriveItem) SetShared(value Sharedable)() {
    m.shared = value
}
// SetSharepointIds sets the sharepointIds property value. The sharepointIds property
func (m *DriveItem) SetSharepointIds(value SharepointIdsable)() {
    m.sharepointIds = value
}
// SetSize sets the size property value. Size of the item in bytes. Read-only.
func (m *DriveItem) SetSize(value *int64)() {
    m.size = value
}
// SetSpecialFolder sets the specialFolder property value. The specialFolder property
func (m *DriveItem) SetSpecialFolder(value SpecialFolderable)() {
    m.specialFolder = value
}
// SetSubscriptions sets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) SetSubscriptions(value []Subscriptionable)() {
    m.subscriptions = value
}
// SetThumbnails sets the thumbnails property value. Collection of [thumbnailSet][] objects associated with the item. For more information, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) SetThumbnails(value []ThumbnailSetable)() {
    m.thumbnails = value
}
// SetVersions sets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) SetVersions(value []DriveItemVersionable)() {
    m.versions = value
}
// SetVideo sets the video property value. The video property
func (m *DriveItem) SetVideo(value Videoable)() {
    m.video = value
}
// SetWebDavUrl sets the webDavUrl property value. WebDAV compatible URL for the item.
func (m *DriveItem) SetWebDavUrl(value *string)() {
    m.webDavUrl = value
}
// SetWorkbook sets the workbook property value. The workbook property
func (m *DriveItem) SetWorkbook(value Workbookable)() {
    m.workbook = value
}
// DriveItemable 
type DriveItemable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetAudio()(Audioable)
    GetBundle()(Bundleable)
    GetChildren()([]DriveItemable)
    GetContent()([]byte)
    GetCTag()(*string)
    GetDeleted()(Deletedable)
    GetFile()(Fileable)
    GetFileSystemInfo()(FileSystemInfoable)
    GetFolder()(Folderable)
    GetImage()(Imageable)
    GetListItem()(ListItemable)
    GetLocation()(GeoCoordinatesable)
    GetMalware()(Malwareable)
    GetPackageEscaped()(PackageEscapedable)
    GetPendingOperations()(PendingOperationsable)
    GetPermissions()([]Permissionable)
    GetPhoto()(Photoable)
    GetPublication()(PublicationFacetable)
    GetRemoteItem()(RemoteItemable)
    GetRetentionLabel()(ItemRetentionLabelable)
    GetRoot()(Rootable)
    GetSearchResult()(SearchResultable)
    GetShared()(Sharedable)
    GetSharepointIds()(SharepointIdsable)
    GetSize()(*int64)
    GetSpecialFolder()(SpecialFolderable)
    GetSubscriptions()([]Subscriptionable)
    GetThumbnails()([]ThumbnailSetable)
    GetVersions()([]DriveItemVersionable)
    GetVideo()(Videoable)
    GetWebDavUrl()(*string)
    GetWorkbook()(Workbookable)
    SetAnalytics(value ItemAnalyticsable)()
    SetAudio(value Audioable)()
    SetBundle(value Bundleable)()
    SetChildren(value []DriveItemable)()
    SetContent(value []byte)()
    SetCTag(value *string)()
    SetDeleted(value Deletedable)()
    SetFile(value Fileable)()
    SetFileSystemInfo(value FileSystemInfoable)()
    SetFolder(value Folderable)()
    SetImage(value Imageable)()
    SetListItem(value ListItemable)()
    SetLocation(value GeoCoordinatesable)()
    SetMalware(value Malwareable)()
    SetPackageEscaped(value PackageEscapedable)()
    SetPendingOperations(value PendingOperationsable)()
    SetPermissions(value []Permissionable)()
    SetPhoto(value Photoable)()
    SetPublication(value PublicationFacetable)()
    SetRemoteItem(value RemoteItemable)()
    SetRetentionLabel(value ItemRetentionLabelable)()
    SetRoot(value Rootable)()
    SetSearchResult(value SearchResultable)()
    SetShared(value Sharedable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSize(value *int64)()
    SetSpecialFolder(value SpecialFolderable)()
    SetSubscriptions(value []Subscriptionable)()
    SetThumbnails(value []ThumbnailSetable)()
    SetVersions(value []DriveItemVersionable)()
    SetVideo(value Videoable)()
    SetWebDavUrl(value *string)()
    SetWorkbook(value Workbookable)()
}
