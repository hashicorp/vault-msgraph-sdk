// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApplication 
type WebApplication struct {
    // Home page or landing page of the application.
    homePageUrl *string
    // The implicitGrantSettings property
    implicitGrantSettings ImplicitGrantSettingsable
    // Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel, back-channel or SAML logout protocols.
    logoutUrl *string
    // Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
    redirectUris []string
    // The redirectUriSettings property
    redirectUriSettings []RedirectUriSettingsable
}
// NewWebApplication instantiates a new webApplication and sets the default values.
func NewWebApplication()(*WebApplication) {
    m := &WebApplication{
    }
    return m
}
// CreateWebApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebApplication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["homePageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomePageUrl(val)
        }
        return nil
    }
    res["implicitGrantSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImplicitGrantSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplicitGrantSettings(val.(ImplicitGrantSettingsable))
        }
        return nil
    }
    res["logoutUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoutUrl(val)
        }
        return nil
    }
    res["redirectUriSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRedirectUriSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RedirectUriSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RedirectUriSettingsable)
                }
            }
            m.SetRedirectUriSettings(res)
        }
        return nil
    }
    res["redirectUris"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRedirectUris(res)
        }
        return nil
    }
    return res
}
// GetHomePageUrl gets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) GetHomePageUrl()(*string) {
    return m.homePageUrl
}
// GetImplicitGrantSettings gets the implicitGrantSettings property value. The implicitGrantSettings property
func (m *WebApplication) GetImplicitGrantSettings()(ImplicitGrantSettingsable) {
    return m.implicitGrantSettings
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) GetLogoutUrl()(*string) {
    return m.logoutUrl
}
// GetRedirectUris gets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) GetRedirectUris()([]string) {
    return m.redirectUris
}
// GetRedirectUriSettings gets the redirectUriSettings property value. The redirectUriSettings property
func (m *WebApplication) GetRedirectUriSettings()([]RedirectUriSettingsable) {
    return m.redirectUriSettings
}
// Serialize serializes information the current object
func (m *WebApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("homePageUrl", m.GetHomePageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("implicitGrantSettings", m.GetImplicitGrantSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRedirectUris() != nil {
        err := writer.WriteCollectionOfStringValues("redirectUris", m.GetRedirectUris())
        if err != nil {
            return err
        }
    }
    if m.GetRedirectUriSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRedirectUriSettings()))
        for i, v := range m.GetRedirectUriSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("redirectUriSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHomePageUrl sets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) SetHomePageUrl(value *string)() {
    m.homePageUrl = value
}
// SetImplicitGrantSettings sets the implicitGrantSettings property value. The implicitGrantSettings property
func (m *WebApplication) SetImplicitGrantSettings(value ImplicitGrantSettingsable)() {
    m.implicitGrantSettings = value
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that is used by Microsoft's authorization service to log out a user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// SetRedirectUris sets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) SetRedirectUris(value []string)() {
    m.redirectUris = value
}
// SetRedirectUriSettings sets the redirectUriSettings property value. The redirectUriSettings property
func (m *WebApplication) SetRedirectUriSettings(value []RedirectUriSettingsable)() {
    m.redirectUriSettings = value
}
// WebApplicationable 
type WebApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHomePageUrl()(*string)
    GetImplicitGrantSettings()(ImplicitGrantSettingsable)
    GetLogoutUrl()(*string)
    GetRedirectUris()([]string)
    GetRedirectUriSettings()([]RedirectUriSettingsable)
    SetHomePageUrl(value *string)()
    SetImplicitGrantSettings(value ImplicitGrantSettingsable)()
    SetLogoutUrl(value *string)()
    SetRedirectUris(value []string)()
    SetRedirectUriSettings(value []RedirectUriSettingsable)()
}
