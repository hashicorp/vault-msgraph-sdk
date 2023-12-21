package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageLinks 
type PageLinks struct {
    // The oneNoteClientUrl property
    oneNoteClientUrl ExternalLinkable
    // The oneNoteWebUrl property
    oneNoteWebUrl ExternalLinkable
}
// NewPageLinks instantiates a new pageLinks and sets the default values.
func NewPageLinks()(*PageLinks) {
    m := &PageLinks{
    }
    return m
}
// CreatePageLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePageLinksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPageLinks(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PageLinks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["oneNoteClientUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteClientUrl(val.(ExternalLinkable))
        }
        return nil
    }
    res["oneNoteWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteWebUrl(val.(ExternalLinkable))
        }
        return nil
    }
    return res
}
// GetOneNoteClientUrl gets the oneNoteClientUrl property value. The oneNoteClientUrl property
func (m *PageLinks) GetOneNoteClientUrl()(ExternalLinkable) {
    return m.oneNoteClientUrl
}
// GetOneNoteWebUrl gets the oneNoteWebUrl property value. The oneNoteWebUrl property
func (m *PageLinks) GetOneNoteWebUrl()(ExternalLinkable) {
    return m.oneNoteWebUrl
}
// Serialize serializes information the current object
func (m *PageLinks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("oneNoteClientUrl", m.GetOneNoteClientUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oneNoteWebUrl", m.GetOneNoteWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOneNoteClientUrl sets the oneNoteClientUrl property value. The oneNoteClientUrl property
func (m *PageLinks) SetOneNoteClientUrl(value ExternalLinkable)() {
    m.oneNoteClientUrl = value
}
// SetOneNoteWebUrl sets the oneNoteWebUrl property value. The oneNoteWebUrl property
func (m *PageLinks) SetOneNoteWebUrl(value ExternalLinkable)() {
    m.oneNoteWebUrl = value
}
// PageLinksable 
type PageLinksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOneNoteClientUrl()(ExternalLinkable)
    GetOneNoteWebUrl()(ExternalLinkable)
    SetOneNoteClientUrl(value ExternalLinkable)()
    SetOneNoteWebUrl(value ExternalLinkable)()
}
