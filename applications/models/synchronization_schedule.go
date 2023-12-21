package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationSchedule 
type SynchronizationSchedule struct {
    // Date and time when this job expires. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    expiration *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The interval between synchronization iterations. The value is represented in ISO 8601 format for durations. For example, PT1M represents a period of one month.
    interval *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The state property
    state *SynchronizationScheduleState
}
// NewSynchronizationSchedule instantiates a new synchronizationSchedule and sets the default values.
func NewSynchronizationSchedule()(*SynchronizationSchedule) {
    m := &SynchronizationSchedule{
    }
    return m
}
// CreateSynchronizationScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationSchedule(), nil
}
// GetExpiration gets the expiration property value. Date and time when this job expires. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationSchedule) GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val)
        }
        return nil
    }
    res["interval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterval(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationScheduleState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SynchronizationScheduleState))
        }
        return nil
    }
    return res
}
// GetInterval gets the interval property value. The interval between synchronization iterations. The value is represented in ISO 8601 format for durations. For example, PT1M represents a period of one month.
func (m *SynchronizationSchedule) GetInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.interval
}
// GetState gets the state property value. The state property
func (m *SynchronizationSchedule) GetState()(*SynchronizationScheduleState) {
    return m.state
}
// Serialize serializes information the current object
func (m *SynchronizationSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expiration", m.GetExpiration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("interval", m.GetInterval())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpiration sets the expiration property value. Date and time when this job expires. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationSchedule) SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiration = value
}
// SetInterval sets the interval property value. The interval between synchronization iterations. The value is represented in ISO 8601 format for durations. For example, PT1M represents a period of one month.
func (m *SynchronizationSchedule) SetInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.interval = value
}
// SetState sets the state property value. The state property
func (m *SynchronizationSchedule) SetState(value *SynchronizationScheduleState)() {
    m.state = value
}
// SynchronizationScheduleable 
type SynchronizationScheduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiration()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetState()(*SynchronizationScheduleState)
    SetExpiration(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetState(value *SynchronizationScheduleState)()
}
