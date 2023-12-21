package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse 
type ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse struct {
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable
}
// NewItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse instantiates a new ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse and sets the default values.
func NewItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse()(*ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) {
    m := &ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse{
    }
    return m
}
// CreateItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateAttributeMappingFunctionSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) GetValue()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponse) SetValue(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable)() {
    m.value = value
}
// ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseable 
type ItemSynchronizationTemplatesItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataNextLink()(*string)
    GetValue()([]i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable)
    SetOdataNextLink(value *string)()
    SetValue(value []i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AttributeMappingFunctionSchemaable)()
}
