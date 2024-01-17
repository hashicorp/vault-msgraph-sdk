// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContainerFilter 
type ContainerFilter struct {
    // The includedContainers property
    includedContainers []string
}
// NewContainerFilter instantiates a new containerFilter and sets the default values.
func NewContainerFilter()(*ContainerFilter) {
    m := &ContainerFilter{
    }
    return m
}
// CreateContainerFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContainerFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContainerFilter(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContainerFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["includedContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIncludedContainers(res)
        }
        return nil
    }
    return res
}
// GetIncludedContainers gets the includedContainers property value. The includedContainers property
func (m *ContainerFilter) GetIncludedContainers()([]string) {
    return m.includedContainers
}
// Serialize serializes information the current object
func (m *ContainerFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludedContainers() != nil {
        err := writer.WriteCollectionOfStringValues("includedContainers", m.GetIncludedContainers())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIncludedContainers sets the includedContainers property value. The includedContainers property
func (m *ContainerFilter) SetIncludedContainers(value []string)() {
    m.includedContainers = value
}
// ContainerFilterable 
type ContainerFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedContainers()([]string)
    SetIncludedContainers(value []string)()
}
