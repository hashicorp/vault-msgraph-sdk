// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PreAuthorizedApplication 
type PreAuthorizedApplication struct {
    // The unique identifier for the application.
    appId *string
    // The unique identifier for the oauth2PermissionScopes the application requires.
    delegatedPermissionIds []string
}
// NewPreAuthorizedApplication instantiates a new preAuthorizedApplication and sets the default values.
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    return m
}
// CreatePreAuthorizedApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePreAuthorizedApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPreAuthorizedApplication(), nil
}
// GetAppId gets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    return m.appId
}
// GetDelegatedPermissionIds gets the delegatedPermissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) GetDelegatedPermissionIds()([]string) {
    return m.delegatedPermissionIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["delegatedPermissionIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDelegatedPermissionIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PreAuthorizedApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedPermissionIds() != nil {
        err := writer.WriteCollectionOfStringValues("delegatedPermissionIds", m.GetDelegatedPermissionIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    m.appId = value
}
// SetDelegatedPermissionIds sets the delegatedPermissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) SetDelegatedPermissionIds(value []string)() {
    m.delegatedPermissionIds = value
}
// PreAuthorizedApplicationable 
type PreAuthorizedApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetDelegatedPermissionIds()([]string)
    SetAppId(value *string)()
    SetDelegatedPermissionIds(value []string)()
}
