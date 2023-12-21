package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResponseStatus 
type ResponseStatus struct {
    // The response property
    response *ResponseType
    // The date and time when the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewResponseStatus instantiates a new responseStatus and sets the default values.
func NewResponseStatus()(*ResponseStatus) {
    m := &ResponseStatus{
    }
    return m
}
// CreateResponseStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResponseStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResponseStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResponseStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["response"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResponseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponse(val.(*ResponseType))
        }
        return nil
    }
    res["time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTime(val)
        }
        return nil
    }
    return res
}
// GetResponse gets the response property value. The response property
func (m *ResponseStatus) GetResponse()(*ResponseType) {
    return m.response
}
// GetTime gets the time property value. The date and time when the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ResponseStatus) GetTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.time
}
// Serialize serializes information the current object
func (m *ResponseStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetResponse() != nil {
        cast := (*m.GetResponse()).String()
        err := writer.WriteStringValue("response", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("time", m.GetTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResponse sets the response property value. The response property
func (m *ResponseStatus) SetResponse(value *ResponseType)() {
    m.response = value
}
// SetTime sets the time property value. The date and time when the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ResponseStatus) SetTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.time = value
}
// ResponseStatusable 
type ResponseStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResponse()(*ResponseType)
    GetTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetResponse(value *ResponseType)()
    SetTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
