// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationLinkedObjects 
type SynchronizationLinkedObjects struct {
    // The manager property
    manager SynchronizationJobSubjectable
    // All group members that you would like to provision.
    members []SynchronizationJobSubjectable
    // The owners property
    owners []SynchronizationJobSubjectable
}
// NewSynchronizationLinkedObjects instantiates a new synchronizationLinkedObjects and sets the default values.
func NewSynchronizationLinkedObjects()(*SynchronizationLinkedObjects) {
    m := &SynchronizationLinkedObjects{
    }
    return m
}
// CreateSynchronizationLinkedObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationLinkedObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationLinkedObjects(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationLinkedObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(SynchronizationJobSubjectable))
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMembers(res)
        }
        return nil
    }
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetOwners(res)
        }
        return nil
    }
    return res
}
// GetManager gets the manager property value. The manager property
func (m *SynchronizationLinkedObjects) GetManager()(SynchronizationJobSubjectable) {
    return m.manager
}
// GetMembers gets the members property value. All group members that you would like to provision.
func (m *SynchronizationLinkedObjects) GetMembers()([]SynchronizationJobSubjectable) {
    return m.members
}
// GetOwners gets the owners property value. The owners property
func (m *SynchronizationLinkedObjects) GetOwners()([]SynchronizationJobSubjectable) {
    return m.owners
}
// Serialize serializes information the current object
func (m *SynchronizationLinkedObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetManager sets the manager property value. The manager property
func (m *SynchronizationLinkedObjects) SetManager(value SynchronizationJobSubjectable)() {
    m.manager = value
}
// SetMembers sets the members property value. All group members that you would like to provision.
func (m *SynchronizationLinkedObjects) SetMembers(value []SynchronizationJobSubjectable)() {
    m.members = value
}
// SetOwners sets the owners property value. The owners property
func (m *SynchronizationLinkedObjects) SetOwners(value []SynchronizationJobSubjectable)() {
    m.owners = value
}
// SynchronizationLinkedObjectsable 
type SynchronizationLinkedObjectsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManager()(SynchronizationJobSubjectable)
    GetMembers()([]SynchronizationJobSubjectable)
    GetOwners()([]SynchronizationJobSubjectable)
    SetManager(value SynchronizationJobSubjectable)()
    SetMembers(value []SynchronizationJobSubjectable)()
    SetOwners(value []SynchronizationJobSubjectable)()
}
