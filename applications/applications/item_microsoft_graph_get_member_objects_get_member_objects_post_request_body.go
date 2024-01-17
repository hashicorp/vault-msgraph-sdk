// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody 
type ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody struct {
    // The securityEnabledOnly property
    securityEnabledOnly *bool
}
// NewItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody instantiates a new ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody and sets the default values.
func NewItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody()(*ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody) {
    m := &ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["securityEnabledOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabledOnly(val)
        }
        return nil
    }
    return res
}
// GetSecurityEnabledOnly gets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody) GetSecurityEnabledOnly()(*bool) {
    return m.securityEnabledOnly
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("securityEnabledOnly", m.GetSecurityEnabledOnly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSecurityEnabledOnly sets the securityEnabledOnly property value. The securityEnabledOnly property
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBody) SetSecurityEnabledOnly(value *bool)() {
    m.securityEnabledOnly = value
}
// ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBodyable 
type ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecurityEnabledOnly()(*bool)
    SetSecurityEnabledOnly(value *bool)()
}
