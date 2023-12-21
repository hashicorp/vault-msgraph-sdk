package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanContainer 
type PlannerPlanContainer struct {
    // The identifier of the resource that contains the plan. Optional.
    containerId *string
    // The type property
    typeEscaped *PlannerContainerType
    // The full canonical URL of the container. Optional.
    url *string
}
// NewPlannerPlanContainer instantiates a new plannerPlanContainer and sets the default values.
func NewPlannerPlanContainer()(*PlannerPlanContainer) {
    m := &PlannerPlanContainer{
    }
    return m
}
// CreatePlannerPlanContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanContainer(), nil
}
// GetContainerId gets the containerId property value. The identifier of the resource that contains the plan. Optional.
func (m *PlannerPlanContainer) GetContainerId()(*string) {
    return m.containerId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["containerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerContainerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*PlannerContainerType))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The type property
func (m *PlannerPlanContainer) GetTypeEscaped()(*PlannerContainerType) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The full canonical URL of the container. Optional.
func (m *PlannerPlanContainer) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PlannerPlanContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerId", m.GetContainerId())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContainerId sets the containerId property value. The identifier of the resource that contains the plan. Optional.
func (m *PlannerPlanContainer) SetContainerId(value *string)() {
    m.containerId = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *PlannerPlanContainer) SetTypeEscaped(value *PlannerContainerType)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The full canonical URL of the container. Optional.
func (m *PlannerPlanContainer) SetUrl(value *string)() {
    m.url = value
}
// PlannerPlanContainerable 
type PlannerPlanContainerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainerId()(*string)
    GetTypeEscaped()(*PlannerContainerType)
    GetUrl()(*string)
    SetContainerId(value *string)()
    SetTypeEscaped(value *PlannerContainerType)()
    SetUrl(value *string)()
}
