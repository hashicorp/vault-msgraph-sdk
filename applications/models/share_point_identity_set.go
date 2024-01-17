// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharePointIdentitySet 
type SharePointIdentitySet struct {
    IdentitySet
    // The group property
    group Identityable
    // The siteGroup property
    siteGroup SharePointIdentityable
    // The siteUser property
    siteUser SharePointIdentityable
}
// NewSharePointIdentitySet instantiates a new sharePointIdentitySet and sets the default values.
func NewSharePointIdentitySet()(*SharePointIdentitySet) {
    m := &SharePointIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// CreateSharePointIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharePointIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Identityable))
        }
        return nil
    }
    res["siteGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteGroup(val.(SharePointIdentityable))
        }
        return nil
    }
    res["siteUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteUser(val.(SharePointIdentityable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *SharePointIdentitySet) GetGroup()(Identityable) {
    return m.group
}
// GetSiteGroup gets the siteGroup property value. The siteGroup property
func (m *SharePointIdentitySet) GetSiteGroup()(SharePointIdentityable) {
    return m.siteGroup
}
// GetSiteUser gets the siteUser property value. The siteUser property
func (m *SharePointIdentitySet) GetSiteUser()(SharePointIdentityable) {
    return m.siteUser
}
// Serialize serializes information the current object
func (m *SharePointIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteGroup", m.GetSiteGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteUser", m.GetSiteUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group property
func (m *SharePointIdentitySet) SetGroup(value Identityable)() {
    m.group = value
}
// SetSiteGroup sets the siteGroup property value. The siteGroup property
func (m *SharePointIdentitySet) SetSiteGroup(value SharePointIdentityable)() {
    m.siteGroup = value
}
// SetSiteUser sets the siteUser property value. The siteUser property
func (m *SharePointIdentitySet) SetSiteUser(value SharePointIdentityable)() {
    m.siteUser = value
}
// SharePointIdentitySetable 
type SharePointIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(Identityable)
    GetSiteGroup()(SharePointIdentityable)
    GetSiteUser()(SharePointIdentityable)
    SetGroup(value Identityable)()
    SetSiteGroup(value SharePointIdentityable)()
    SetSiteUser(value SharePointIdentityable)()
}
