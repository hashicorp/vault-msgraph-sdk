package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessagePolicyViolation 
type ChatMessagePolicyViolation struct {
    // The dlpAction property
    dlpAction *ChatMessagePolicyViolationDlpActionTypes
    // Justification text provided by the sender of the message when overriding a policy violation.
    justificationText *string
    // The policyTip property
    policyTip ChatMessagePolicyViolationPolicyTipable
    // The userAction property
    userAction *ChatMessagePolicyViolationUserActionTypes
    // The verdictDetails property
    verdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes
}
// NewChatMessagePolicyViolation instantiates a new chatMessagePolicyViolation and sets the default values.
func NewChatMessagePolicyViolation()(*ChatMessagePolicyViolation) {
    m := &ChatMessagePolicyViolation{
    }
    return m
}
// CreateChatMessagePolicyViolationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessagePolicyViolationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessagePolicyViolation(), nil
}
// GetDlpAction gets the dlpAction property value. The dlpAction property
func (m *ChatMessagePolicyViolation) GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes) {
    return m.dlpAction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessagePolicyViolation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dlpAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationDlpActionTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDlpAction(val.(*ChatMessagePolicyViolationDlpActionTypes))
        }
        return nil
    }
    res["justificationText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationText(val)
        }
        return nil
    }
    res["policyTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTip(val.(ChatMessagePolicyViolationPolicyTipable))
        }
        return nil
    }
    res["userAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationUserActionTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAction(val.(*ChatMessagePolicyViolationUserActionTypes))
        }
        return nil
    }
    res["verdictDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationVerdictDetailsTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerdictDetails(val.(*ChatMessagePolicyViolationVerdictDetailsTypes))
        }
        return nil
    }
    return res
}
// GetJustificationText gets the justificationText property value. Justification text provided by the sender of the message when overriding a policy violation.
func (m *ChatMessagePolicyViolation) GetJustificationText()(*string) {
    return m.justificationText
}
// GetPolicyTip gets the policyTip property value. The policyTip property
func (m *ChatMessagePolicyViolation) GetPolicyTip()(ChatMessagePolicyViolationPolicyTipable) {
    return m.policyTip
}
// GetUserAction gets the userAction property value. The userAction property
func (m *ChatMessagePolicyViolation) GetUserAction()(*ChatMessagePolicyViolationUserActionTypes) {
    return m.userAction
}
// GetVerdictDetails gets the verdictDetails property value. The verdictDetails property
func (m *ChatMessagePolicyViolation) GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes) {
    return m.verdictDetails
}
// Serialize serializes information the current object
func (m *ChatMessagePolicyViolation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDlpAction() != nil {
        cast := (*m.GetDlpAction()).String()
        err := writer.WriteStringValue("dlpAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationText", m.GetJustificationText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("policyTip", m.GetPolicyTip())
        if err != nil {
            return err
        }
    }
    if m.GetUserAction() != nil {
        cast := (*m.GetUserAction()).String()
        err := writer.WriteStringValue("userAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVerdictDetails() != nil {
        cast := (*m.GetVerdictDetails()).String()
        err := writer.WriteStringValue("verdictDetails", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDlpAction sets the dlpAction property value. The dlpAction property
func (m *ChatMessagePolicyViolation) SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)() {
    m.dlpAction = value
}
// SetJustificationText sets the justificationText property value. Justification text provided by the sender of the message when overriding a policy violation.
func (m *ChatMessagePolicyViolation) SetJustificationText(value *string)() {
    m.justificationText = value
}
// SetPolicyTip sets the policyTip property value. The policyTip property
func (m *ChatMessagePolicyViolation) SetPolicyTip(value ChatMessagePolicyViolationPolicyTipable)() {
    m.policyTip = value
}
// SetUserAction sets the userAction property value. The userAction property
func (m *ChatMessagePolicyViolation) SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)() {
    m.userAction = value
}
// SetVerdictDetails sets the verdictDetails property value. The verdictDetails property
func (m *ChatMessagePolicyViolation) SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)() {
    m.verdictDetails = value
}
// ChatMessagePolicyViolationable 
type ChatMessagePolicyViolationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes)
    GetJustificationText()(*string)
    GetPolicyTip()(ChatMessagePolicyViolationPolicyTipable)
    GetUserAction()(*ChatMessagePolicyViolationUserActionTypes)
    GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes)
    SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)()
    SetJustificationText(value *string)()
    SetPolicyTip(value ChatMessagePolicyViolationPolicyTipable)()
    SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)()
    SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)()
}
