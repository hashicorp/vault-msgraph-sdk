package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerChecklistItems 
type PlannerChecklistItems struct {
}
// NewPlannerChecklistItems instantiates a new plannerChecklistItems and sets the default values.
func NewPlannerChecklistItems()(*PlannerChecklistItems) {
    m := &PlannerChecklistItems{
    }
    return m
}
// CreatePlannerChecklistItemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerChecklistItemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerChecklistItems(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerChecklistItems) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *PlannerChecklistItems) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// PlannerChecklistItemsable 
type PlannerChecklistItemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
