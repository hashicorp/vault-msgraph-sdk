package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingParticipantInfo 
type MeetingParticipantInfo struct {
    // The identity property
    identity IdentitySetable
    // The role property
    role *OnlineMeetingRole
    // User principal name of the participant.
    upn *string
}
// NewMeetingParticipantInfo instantiates a new meetingParticipantInfo and sets the default values.
func NewMeetingParticipantInfo()(*MeetingParticipantInfo) {
    m := &MeetingParticipantInfo{
    }
    return m
}
// CreateMeetingParticipantInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingParticipantInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingParticipantInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingParticipantInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(IdentitySetable))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*OnlineMeetingRole))
        }
        return nil
    }
    res["upn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity property
func (m *MeetingParticipantInfo) GetIdentity()(IdentitySetable) {
    return m.identity
}
// GetRole gets the role property value. The role property
func (m *MeetingParticipantInfo) GetRole()(*OnlineMeetingRole) {
    return m.role
}
// GetUpn gets the upn property value. User principal name of the participant.
func (m *MeetingParticipantInfo) GetUpn()(*string) {
    return m.upn
}
// Serialize serializes information the current object
func (m *MeetingParticipantInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err := writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("upn", m.GetUpn())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentity sets the identity property value. The identity property
func (m *MeetingParticipantInfo) SetIdentity(value IdentitySetable)() {
    m.identity = value
}
// SetRole sets the role property value. The role property
func (m *MeetingParticipantInfo) SetRole(value *OnlineMeetingRole)() {
    m.role = value
}
// SetUpn sets the upn property value. User principal name of the participant.
func (m *MeetingParticipantInfo) SetUpn(value *string)() {
    m.upn = value
}
// MeetingParticipantInfoable 
type MeetingParticipantInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentity()(IdentitySetable)
    GetRole()(*OnlineMeetingRole)
    GetUpn()(*string)
    SetIdentity(value IdentitySetable)()
    SetRole(value *OnlineMeetingRole)()
    SetUpn(value *string)()
}
