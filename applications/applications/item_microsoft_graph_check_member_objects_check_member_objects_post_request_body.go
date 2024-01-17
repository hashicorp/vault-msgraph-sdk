// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody 
type ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody struct {
    // The ids property
    ids []string
}
// NewItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody instantiates a new ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody and sets the default values.
func NewItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody()(*ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) {
    m := &ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIds(res)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) GetIds()([]string) {
    return m.ids
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIds sets the ids property value. The ids property
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyable 
type ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIds()([]string)
    SetIds(value []string)()
}
