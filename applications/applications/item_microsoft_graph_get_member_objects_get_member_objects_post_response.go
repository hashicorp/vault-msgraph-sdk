// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse 
type ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse struct {
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []string
}
// NewItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse instantiates a new ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse and sets the default values.
func NewItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse()(*ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) {
    m := &ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse{
    }
    return m
}
// CreateItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) GetValue()([]string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponse) SetValue(value []string)() {
    m.value = value
}
// ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponseable 
type ItemMicrosoftGraphGetMemberObjectsGetMemberObjectsPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataNextLink()(*string)
    GetValue()([]string)
    SetOdataNextLink(value *string)()
    SetValue(value []string)()
}
