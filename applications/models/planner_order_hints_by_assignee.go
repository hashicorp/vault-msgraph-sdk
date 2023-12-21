package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerOrderHintsByAssignee 
type PlannerOrderHintsByAssignee struct {
}
// NewPlannerOrderHintsByAssignee instantiates a new plannerOrderHintsByAssignee and sets the default values.
func NewPlannerOrderHintsByAssignee()(*PlannerOrderHintsByAssignee) {
    m := &PlannerOrderHintsByAssignee{
    }
    return m
}
// CreatePlannerOrderHintsByAssigneeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerOrderHintsByAssigneeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerOrderHintsByAssignee(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerOrderHintsByAssignee) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *PlannerOrderHintsByAssignee) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// PlannerOrderHintsByAssigneeable 
type PlannerOrderHintsByAssigneeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
