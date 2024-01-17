// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharingLink 
type SharingLink struct {
    // The application property
    application Identityable
    // If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
    preventsDownload *bool
    // The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
    scope *string
    // The type of the link created.
    typeEscaped *string
    // For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
    webHtml *string
    // A URL that opens the item in the browser on the OneDrive website.
    webUrl *string
}
// NewSharingLink instantiates a new sharingLink and sets the default values.
func NewSharingLink()(*SharingLink) {
    m := &SharingLink{
    }
    return m
}
// CreateSharingLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharingLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharingLink(), nil
}
// GetApplication gets the application property value. The application property
func (m *SharingLink) GetApplication()(Identityable) {
    return m.application
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharingLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Identityable))
        }
        return nil
    }
    res["preventsDownload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreventsDownload(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
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
    res["webHtml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebHtml(val)
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
// GetPreventsDownload gets the preventsDownload property value. If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
func (m *SharingLink) GetPreventsDownload()(*bool) {
    return m.preventsDownload
}
// GetScope gets the scope property value. The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
func (m *SharingLink) GetScope()(*string) {
    return m.scope
}
// GetTypeEscaped gets the type property value. The type of the link created.
func (m *SharingLink) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetWebHtml gets the webHtml property value. For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
func (m *SharingLink) GetWebHtml()(*string) {
    return m.webHtml
}
// GetWebUrl gets the webUrl property value. A URL that opens the item in the browser on the OneDrive website.
func (m *SharingLink) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *SharingLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preventsDownload", m.GetPreventsDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
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
    {
        err := writer.WriteStringValue("webHtml", m.GetWebHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The application property
func (m *SharingLink) SetApplication(value Identityable)() {
    m.application = value
}
// SetPreventsDownload sets the preventsDownload property value. If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
func (m *SharingLink) SetPreventsDownload(value *bool)() {
    m.preventsDownload = value
}
// SetScope sets the scope property value. The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
func (m *SharingLink) SetScope(value *string)() {
    m.scope = value
}
// SetTypeEscaped sets the type property value. The type of the link created.
func (m *SharingLink) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetWebHtml sets the webHtml property value. For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
func (m *SharingLink) SetWebHtml(value *string)() {
    m.webHtml = value
}
// SetWebUrl sets the webUrl property value. A URL that opens the item in the browser on the OneDrive website.
func (m *SharingLink) SetWebUrl(value *string)() {
    m.webUrl = value
}
// SharingLinkable 
type SharingLinkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(Identityable)
    GetPreventsDownload()(*bool)
    GetScope()(*string)
    GetTypeEscaped()(*string)
    GetWebHtml()(*string)
    GetWebUrl()(*string)
    SetApplication(value Identityable)()
    SetPreventsDownload(value *bool)()
    SetScope(value *string)()
    SetTypeEscaped(value *string)()
    SetWebHtml(value *string)()
    SetWebUrl(value *string)()
}
