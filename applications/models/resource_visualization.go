package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResourceVisualization 
type ResourceVisualization struct {
    // A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
    containerDisplayName *string
    // Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
    containerType *string
    // A path leading to the folder in which the item is stored.
    containerWebUrl *string
    // The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Not all Media Mime Types are supported.
    mediaType *string
    // A URL leading to the preview image for the item.
    previewImageUrl *string
    // A preview text for the item.
    previewText *string
    // The item's title text.
    title *string
    // The item's media type. Can be used for filtering for a specific file based on a specific type. See the section Type property values for supported types.
    typeEscaped *string
}
// NewResourceVisualization instantiates a new resourceVisualization and sets the default values.
func NewResourceVisualization()(*ResourceVisualization) {
    m := &ResourceVisualization{
    }
    return m
}
// CreateResourceVisualizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceVisualizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResourceVisualization(), nil
}
// GetContainerDisplayName gets the containerDisplayName property value. A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
func (m *ResourceVisualization) GetContainerDisplayName()(*string) {
    return m.containerDisplayName
}
// GetContainerType gets the containerType property value. Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
func (m *ResourceVisualization) GetContainerType()(*string) {
    return m.containerType
}
// GetContainerWebUrl gets the containerWebUrl property value. A path leading to the folder in which the item is stored.
func (m *ResourceVisualization) GetContainerWebUrl()(*string) {
    return m.containerWebUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceVisualization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["containerDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerDisplayName(val)
        }
        return nil
    }
    res["containerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerType(val)
        }
        return nil
    }
    res["containerWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerWebUrl(val)
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["previewImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewImageUrl(val)
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetMediaType gets the mediaType property value. The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Not all Media Mime Types are supported.
func (m *ResourceVisualization) GetMediaType()(*string) {
    return m.mediaType
}
// GetPreviewImageUrl gets the previewImageUrl property value. A URL leading to the preview image for the item.
func (m *ResourceVisualization) GetPreviewImageUrl()(*string) {
    return m.previewImageUrl
}
// GetPreviewText gets the previewText property value. A preview text for the item.
func (m *ResourceVisualization) GetPreviewText()(*string) {
    return m.previewText
}
// GetTitle gets the title property value. The item's title text.
func (m *ResourceVisualization) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. The item's media type. Can be used for filtering for a specific file based on a specific type. See the section Type property values for supported types.
func (m *ResourceVisualization) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ResourceVisualization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerDisplayName", m.GetContainerDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("containerType", m.GetContainerType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("containerWebUrl", m.GetContainerWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewImageUrl", m.GetPreviewImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainerDisplayName sets the containerDisplayName property value. A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
func (m *ResourceVisualization) SetContainerDisplayName(value *string)() {
    m.containerDisplayName = value
}
// SetContainerType sets the containerType property value. Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
func (m *ResourceVisualization) SetContainerType(value *string)() {
    m.containerType = value
}
// SetContainerWebUrl sets the containerWebUrl property value. A path leading to the folder in which the item is stored.
func (m *ResourceVisualization) SetContainerWebUrl(value *string)() {
    m.containerWebUrl = value
}
// SetMediaType sets the mediaType property value. The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Not all Media Mime Types are supported.
func (m *ResourceVisualization) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetPreviewImageUrl sets the previewImageUrl property value. A URL leading to the preview image for the item.
func (m *ResourceVisualization) SetPreviewImageUrl(value *string)() {
    m.previewImageUrl = value
}
// SetPreviewText sets the previewText property value. A preview text for the item.
func (m *ResourceVisualization) SetPreviewText(value *string)() {
    m.previewText = value
}
// SetTitle sets the title property value. The item's title text.
func (m *ResourceVisualization) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. The item's media type. Can be used for filtering for a specific file based on a specific type. See the section Type property values for supported types.
func (m *ResourceVisualization) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// ResourceVisualizationable 
type ResourceVisualizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainerDisplayName()(*string)
    GetContainerType()(*string)
    GetContainerWebUrl()(*string)
    GetMediaType()(*string)
    GetPreviewImageUrl()(*string)
    GetPreviewText()(*string)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    SetContainerDisplayName(value *string)()
    SetContainerType(value *string)()
    SetContainerWebUrl(value *string)()
    SetMediaType(value *string)()
    SetPreviewImageUrl(value *string)()
    SetPreviewText(value *string)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
}
