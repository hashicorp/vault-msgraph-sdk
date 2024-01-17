// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// Relation 
type Relation struct {
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Entity
    // The fromTerm property
    fromTerm Termable
    // The relationship property
    relationship *RelationType
    // The set property
    set Setable
    // The toTerm property
    toTerm Termable
}
// NewRelation instantiates a new relation and sets the default values.
func NewRelation()(*Relation) {
    m := &Relation{
        Entity: *i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.NewEntity(),
    }
    return m
}
// CreateRelationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRelationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Relation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fromTerm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFromTerm(val.(Termable))
        }
        return nil
    }
    res["relationship"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRelationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationship(val.(*RelationType))
        }
        return nil
    }
    res["set"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSet(val.(Setable))
        }
        return nil
    }
    res["toTerm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToTerm(val.(Termable))
        }
        return nil
    }
    return res
}
// GetFromTerm gets the fromTerm property value. The fromTerm property
func (m *Relation) GetFromTerm()(Termable) {
    return m.fromTerm
}
// GetRelationship gets the relationship property value. The relationship property
func (m *Relation) GetRelationship()(*RelationType) {
    return m.relationship
}
// GetSet gets the set property value. The set property
func (m *Relation) GetSet()(Setable) {
    return m.set
}
// GetToTerm gets the toTerm property value. The toTerm property
func (m *Relation) GetToTerm()(Termable) {
    return m.toTerm
}
// Serialize serializes information the current object
func (m *Relation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fromTerm", m.GetFromTerm())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := (*m.GetRelationship()).String()
        err = writer.WriteStringValue("relationship", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("set", m.GetSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("toTerm", m.GetToTerm())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFromTerm sets the fromTerm property value. The fromTerm property
func (m *Relation) SetFromTerm(value Termable)() {
    m.fromTerm = value
}
// SetRelationship sets the relationship property value. The relationship property
func (m *Relation) SetRelationship(value *RelationType)() {
    m.relationship = value
}
// SetSet sets the set property value. The set property
func (m *Relation) SetSet(value Setable)() {
    m.set = value
}
// SetToTerm sets the toTerm property value. The toTerm property
func (m *Relation) SetToTerm(value Termable)() {
    m.toTerm = value
}
// Relationable 
type Relationable interface {
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFromTerm()(Termable)
    GetRelationship()(*RelationType)
    GetSet()(Setable)
    GetToTerm()(Termable)
    SetFromTerm(value Termable)()
    SetRelationship(value *RelationType)()
    SetSet(value Setable)()
    SetToTerm(value Termable)()
}
