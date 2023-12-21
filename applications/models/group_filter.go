package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupFilter 
type GroupFilter struct {
    // The includedGroups property
    includedGroups []string
}
// NewGroupFilter instantiates a new groupFilter and sets the default values.
func NewGroupFilter()(*GroupFilter) {
    m := &GroupFilter{
    }
    return m
}
// CreateGroupFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetIncludedGroups(res)
        }
        return nil
    }
    return res
}
// GetIncludedGroups gets the includedGroups property value. The includedGroups property
func (m *GroupFilter) GetIncludedGroups()([]string) {
    return m.includedGroups
}
// Serialize serializes information the current object
func (m *GroupFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludedGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includedGroups", m.GetIncludedGroups())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIncludedGroups sets the includedGroups property value. The includedGroups property
func (m *GroupFilter) SetIncludedGroups(value []string)() {
    m.includedGroups = value
}
// GroupFilterable 
type GroupFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedGroups()([]string)
    SetIncludedGroups(value []string)()
}
