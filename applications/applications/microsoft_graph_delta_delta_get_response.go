// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// MicrosoftGraphDeltaDeltaGetResponse 
type MicrosoftGraphDeltaDeltaGetResponse struct {
    // The OdataDeltaLink property
    odataDeltaLink *string
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable
}
// NewMicrosoftGraphDeltaDeltaGetResponse instantiates a new MicrosoftGraphDeltaDeltaGetResponse and sets the default values.
func NewMicrosoftGraphDeltaDeltaGetResponse()(*MicrosoftGraphDeltaDeltaGetResponse) {
    m := &MicrosoftGraphDeltaDeltaGetResponse{
    }
    return m
}
// CreateMicrosoftGraphDeltaDeltaGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphDeltaDeltaGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphDeltaDeltaGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphDeltaDeltaGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.deltaLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataDeltaLink(val)
        }
        return nil
    }
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
        val, err := n.GetCollectionOfObjectValues(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdataDeltaLink gets the @odata.deltaLink property value. The OdataDeltaLink property
func (m *MicrosoftGraphDeltaDeltaGetResponse) GetOdataDeltaLink()(*string) {
    return m.odataDeltaLink
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *MicrosoftGraphDeltaDeltaGetResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *MicrosoftGraphDeltaDeltaGetResponse) GetValue()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MicrosoftGraphDeltaDeltaGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.deltaLink", m.GetOdataDeltaLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataDeltaLink sets the @odata.deltaLink property value. The OdataDeltaLink property
func (m *MicrosoftGraphDeltaDeltaGetResponse) SetOdataDeltaLink(value *string)() {
    m.odataDeltaLink = value
}
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *MicrosoftGraphDeltaDeltaGetResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *MicrosoftGraphDeltaDeltaGetResponse) SetValue(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable)() {
    m.value = value
}
// MicrosoftGraphDeltaDeltaGetResponseable 
type MicrosoftGraphDeltaDeltaGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataDeltaLink()(*string)
    GetOdataNextLink()(*string)
    GetValue()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable)
    SetOdataDeltaLink(value *string)()
    SetOdataNextLink(value *string)()
    SetValue(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Applicationable)()
}
