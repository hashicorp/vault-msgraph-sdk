// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicClientApplication 
type PublicClientApplication struct {
    // Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
    redirectUris []string
}
// NewPublicClientApplication instantiates a new publicClientApplication and sets the default values.
func NewPublicClientApplication()(*PublicClientApplication) {
    m := &PublicClientApplication{
    }
    return m
}
// CreatePublicClientApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicClientApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicClientApplication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicClientApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetRedirectUris gets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *PublicClientApplication) GetRedirectUris()([]string) {
    return m.redirectUris
}
// Serialize serializes information the current object
func (m *PublicClientApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRedirectUris() != nil {
        err := writer.WriteCollectionOfStringValues("redirectUris", m.GetRedirectUris())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRedirectUris sets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *PublicClientApplication) SetRedirectUris(value []string)() {
    m.redirectUris = value
}
// PublicClientApplicationable 
type PublicClientApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRedirectUris()([]string)
    SetRedirectUris(value []string)()
}
