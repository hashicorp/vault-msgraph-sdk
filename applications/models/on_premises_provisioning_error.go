package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesProvisioningError 
type OnPremisesProvisioningError struct {
    // Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
    category *string
    // The date and time at which the error occurred.
    occurredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
    propertyCausingError *string
    // Value of the property causing the error.
    value *string
}
// NewOnPremisesProvisioningError instantiates a new onPremisesProvisioningError and sets the default values.
func NewOnPremisesProvisioningError()(*OnPremisesProvisioningError) {
    m := &OnPremisesProvisioningError{
    }
    return m
}
// CreateOnPremisesProvisioningErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesProvisioningErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesProvisioningError(), nil
}
// GetCategory gets the category property value. Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
func (m *OnPremisesProvisioningError) GetCategory()(*string) {
    return m.category
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesProvisioningError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["occurredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOccurredDateTime(val)
        }
        return nil
    }
    res["propertyCausingError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyCausingError(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetOccurredDateTime gets the occurredDateTime property value. The date and time at which the error occurred.
func (m *OnPremisesProvisioningError) GetOccurredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.occurredDateTime
}
// GetPropertyCausingError gets the propertyCausingError property value. Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
func (m *OnPremisesProvisioningError) GetPropertyCausingError()(*string) {
    return m.propertyCausingError
}
// GetValue gets the value property value. Value of the property causing the error.
func (m *OnPremisesProvisioningError) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *OnPremisesProvisioningError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("occurredDateTime", m.GetOccurredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("propertyCausingError", m.GetPropertyCausingError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. Category of the provisioning error. Note: Currently, there is only one possible value. Possible value: PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
func (m *OnPremisesProvisioningError) SetCategory(value *string)() {
    m.category = value
}
// SetOccurredDateTime sets the occurredDateTime property value. The date and time at which the error occurred.
func (m *OnPremisesProvisioningError) SetOccurredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.occurredDateTime = value
}
// SetPropertyCausingError sets the propertyCausingError property value. Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
func (m *OnPremisesProvisioningError) SetPropertyCausingError(value *string)() {
    m.propertyCausingError = value
}
// SetValue sets the value property value. Value of the property causing the error.
func (m *OnPremisesProvisioningError) SetValue(value *string)() {
    m.value = value
}
// OnPremisesProvisioningErrorable 
type OnPremisesProvisioningErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*string)
    GetOccurredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPropertyCausingError()(*string)
    GetValue()(*string)
    SetCategory(value *string)()
    SetOccurredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPropertyCausingError(value *string)()
    SetValue(value *string)()
}
