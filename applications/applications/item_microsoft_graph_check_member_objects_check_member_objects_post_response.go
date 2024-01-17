// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse 
type ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse struct {
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []string
}
// NewItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse instantiates a new ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse and sets the default values.
func NewItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse()(*ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) {
    m := &ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse{
    }
    return m
}
// CreateItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.nextLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataNextLink(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        err := writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponse) SetValue(value []string)() {
    m.value = value
}
// ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponseable 
type ItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataNextLink()(*string)
    GetValue()([]string)
    SetOdataNextLink(value *string)()
    SetValue(value []string)()
}
