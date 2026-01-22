// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobApplicationParameters 
type SynchronizationJobApplicationParameters struct {
    // The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
    ruleId *string
    // The identifiers of one or more objects to which a synchronizationJob is to be applied.
    subjects []SynchronizationJobSubjectable
}
// NewSynchronizationJobApplicationParameters instantiates a new synchronizationJobApplicationParameters and sets the default values.
func NewSynchronizationJobApplicationParameters()(*SynchronizationJobApplicationParameters) {
    m := &SynchronizationJobApplicationParameters{
    }
    return m
}
// CreateSynchronizationJobApplicationParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobApplicationParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJobApplicationParameters(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobApplicationParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ruleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleId(val)
        }
        return nil
    }
    res["subjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobSubjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SynchronizationJobSubjectable)
                }
            }
            m.SetSubjects(res)
        }
        return nil
    }
    return res
}
// GetRuleId gets the ruleId property value. The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
func (m *SynchronizationJobApplicationParameters) GetRuleId()(*string) {
    return m.ruleId
}
// GetSubjects gets the subjects property value. The identifiers of one or more objects to which a synchronizationJob is to be applied.
func (m *SynchronizationJobApplicationParameters) GetSubjects()([]SynchronizationJobSubjectable) {
    return m.subjects
}
// Serialize serializes information the current object
func (m *SynchronizationJobApplicationParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ruleId", m.GetRuleId())
        if err != nil {
            return err
        }
    }
    if m.GetSubjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjects()))
        for i, v := range m.GetSubjects() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("subjects", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRuleId sets the ruleId property value. The identifier of the synchronizationRule to be applied. This rule ID is defined in the schema for a given synchronization job or template.
func (m *SynchronizationJobApplicationParameters) SetRuleId(value *string)() {
    m.ruleId = value
}
// SetSubjects sets the subjects property value. The identifiers of one or more objects to which a synchronizationJob is to be applied.
func (m *SynchronizationJobApplicationParameters) SetSubjects(value []SynchronizationJobSubjectable)() {
    m.subjects = value
}
// SynchronizationJobApplicationParametersable 
type SynchronizationJobApplicationParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRuleId()(*string)
    GetSubjects()([]SynchronizationJobSubjectable)
    SetRuleId(value *string)()
    SetSubjects(value []SynchronizationJobSubjectable)()
}
