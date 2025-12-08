// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody 
type ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody struct {
    // The groupIds property
    groupIds []string
}
// NewItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody instantiates a new ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody and sets the default values.
func NewItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody()(*ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) {
    m := &ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGroupIds(res)
        }
        return nil
    }
    return res
}
// GetGroupIds gets the groupIds property value. The groupIds property
func (m *ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) GetGroupIds()([]string) {
    return m.groupIds
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupIds() != nil {
        err := writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupIds sets the groupIds property value. The groupIds property
func (m *ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBody) SetGroupIds(value []string)() {
    m.groupIds = value
}
// ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBodyable 
type ItemMicrosoftGraphCheckMemberGroupsCheckMemberGroupsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupIds()([]string)
    SetGroupIds(value []string)()
}
