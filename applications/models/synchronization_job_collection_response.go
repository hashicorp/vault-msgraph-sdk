package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobCollectionResponse 
type SynchronizationJobCollectionResponse struct {
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []SynchronizationJobable
}
// NewSynchronizationJobCollectionResponse instantiates a new synchronizationJobCollectionResponse and sets the default values.
func NewSynchronizationJobCollectionResponse()(*SynchronizationJobCollectionResponse) {
    m := &SynchronizationJobCollectionResponse{
    }
    return m
}
// CreateSynchronizationJobCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJobCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SynchronizationJobable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *SynchronizationJobCollectionResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *SynchronizationJobCollectionResponse) GetValue()([]SynchronizationJobable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SynchronizationJobCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SynchronizationJobCollectionResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *SynchronizationJobCollectionResponse) SetValue(value []SynchronizationJobable)() {
    m.value = value
}
// SynchronizationJobCollectionResponseable 
type SynchronizationJobCollectionResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataNextLink()(*string)
    GetValue()([]SynchronizationJobable)
    SetOdataNextLink(value *string)()
    SetValue(value []SynchronizationJobable)()
}
