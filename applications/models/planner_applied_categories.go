// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerAppliedCategories 
type PlannerAppliedCategories struct {
}
// NewPlannerAppliedCategories instantiates a new plannerAppliedCategories and sets the default values.
func NewPlannerAppliedCategories()(*PlannerAppliedCategories) {
    m := &PlannerAppliedCategories{
    }
    return m
}
// CreatePlannerAppliedCategoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerAppliedCategoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerAppliedCategories(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerAppliedCategories) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *PlannerAppliedCategories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// PlannerAppliedCategoriesable 
type PlannerAppliedCategoriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
