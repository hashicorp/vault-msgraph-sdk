package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PendingOperations 
type PendingOperations struct {
    // The pendingContentUpdate property
    pendingContentUpdate PendingContentUpdateable
}
// NewPendingOperations instantiates a new pendingOperations and sets the default values.
func NewPendingOperations()(*PendingOperations) {
    m := &PendingOperations{
    }
    return m
}
// CreatePendingOperationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePendingOperationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPendingOperations(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PendingOperations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pendingContentUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingContentUpdateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingContentUpdate(val.(PendingContentUpdateable))
        }
        return nil
    }
    return res
}
// GetPendingContentUpdate gets the pendingContentUpdate property value. The pendingContentUpdate property
func (m *PendingOperations) GetPendingContentUpdate()(PendingContentUpdateable) {
    return m.pendingContentUpdate
}
// Serialize serializes information the current object
func (m *PendingOperations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pendingContentUpdate", m.GetPendingContentUpdate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPendingContentUpdate sets the pendingContentUpdate property value. The pendingContentUpdate property
func (m *PendingOperations) SetPendingContentUpdate(value PendingContentUpdateable)() {
    m.pendingContentUpdate = value
}
// PendingOperationsable 
type PendingOperationsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPendingContentUpdate()(PendingContentUpdateable)
    SetPendingContentUpdate(value PendingContentUpdateable)()
}
