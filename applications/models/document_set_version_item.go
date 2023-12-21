package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetVersionItem 
type DocumentSetVersionItem struct {
    // The unique identifier for the item.
    itemId *string
    // The title of the item.
    title *string
    // The version ID of the item.
    versionId *string
}
// NewDocumentSetVersionItem instantiates a new documentSetVersionItem and sets the default values.
func NewDocumentSetVersionItem()(*DocumentSetVersionItem) {
    m := &DocumentSetVersionItem{
    }
    return m
}
// CreateDocumentSetVersionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetVersionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentSetVersionItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSetVersionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
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
    res["versionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionId(val)
        }
        return nil
    }
    return res
}
// GetItemId gets the itemId property value. The unique identifier for the item.
func (m *DocumentSetVersionItem) GetItemId()(*string) {
    return m.itemId
}
// GetTitle gets the title property value. The title of the item.
func (m *DocumentSetVersionItem) GetTitle()(*string) {
    return m.title
}
// GetVersionId gets the versionId property value. The version ID of the item.
func (m *DocumentSetVersionItem) GetVersionId()(*string) {
    return m.versionId
}
// Serialize serializes information the current object
func (m *DocumentSetVersionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("itemId", m.GetItemId())
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
        err := writer.WriteStringValue("versionId", m.GetVersionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemId sets the itemId property value. The unique identifier for the item.
func (m *DocumentSetVersionItem) SetItemId(value *string)() {
    m.itemId = value
}
// SetTitle sets the title property value. The title of the item.
func (m *DocumentSetVersionItem) SetTitle(value *string)() {
    m.title = value
}
// SetVersionId sets the versionId property value. The version ID of the item.
func (m *DocumentSetVersionItem) SetVersionId(value *string)() {
    m.versionId = value
}
// DocumentSetVersionItemable 
type DocumentSetVersionItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemId()(*string)
    GetTitle()(*string)
    GetVersionId()(*string)
    SetItemId(value *string)()
    SetTitle(value *string)()
    SetVersionId(value *string)()
}
