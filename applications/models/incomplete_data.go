package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IncompleteData 
type IncompleteData struct {
    // The service does not have source data before the specified time.
    missingDataBeforeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Some data was not recorded due to excessive activity.
    wasThrottled *bool
}
// NewIncompleteData instantiates a new incompleteData and sets the default values.
func NewIncompleteData()(*IncompleteData) {
    m := &IncompleteData{
    }
    return m
}
// CreateIncompleteDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIncompleteDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIncompleteData(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IncompleteData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["missingDataBeforeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMissingDataBeforeDateTime(val)
        }
        return nil
    }
    res["wasThrottled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWasThrottled(val)
        }
        return nil
    }
    return res
}
// GetMissingDataBeforeDateTime gets the missingDataBeforeDateTime property value. The service does not have source data before the specified time.
func (m *IncompleteData) GetMissingDataBeforeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.missingDataBeforeDateTime
}
// GetWasThrottled gets the wasThrottled property value. Some data was not recorded due to excessive activity.
func (m *IncompleteData) GetWasThrottled()(*bool) {
    return m.wasThrottled
}
// Serialize serializes information the current object
func (m *IncompleteData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("missingDataBeforeDateTime", m.GetMissingDataBeforeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wasThrottled", m.GetWasThrottled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMissingDataBeforeDateTime sets the missingDataBeforeDateTime property value. The service does not have source data before the specified time.
func (m *IncompleteData) SetMissingDataBeforeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.missingDataBeforeDateTime = value
}
// SetWasThrottled sets the wasThrottled property value. Some data was not recorded due to excessive activity.
func (m *IncompleteData) SetWasThrottled(value *bool)() {
    m.wasThrottled = value
}
// IncompleteDataable 
type IncompleteDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMissingDataBeforeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetWasThrottled()(*bool)
    SetMissingDataBeforeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetWasThrottled(value *bool)()
}
