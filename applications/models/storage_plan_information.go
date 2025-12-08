// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StoragePlanInformation 
type StoragePlanInformation struct {
    // Indicates whether there are higher storage quota plans available. Read-only.
    upgradeAvailable *bool
}
// NewStoragePlanInformation instantiates a new storagePlanInformation and sets the default values.
func NewStoragePlanInformation()(*StoragePlanInformation) {
    m := &StoragePlanInformation{
    }
    return m
}
// CreateStoragePlanInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStoragePlanInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStoragePlanInformation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StoragePlanInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["upgradeAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAvailable(val)
        }
        return nil
    }
    return res
}
// GetUpgradeAvailable gets the upgradeAvailable property value. Indicates whether there are higher storage quota plans available. Read-only.
func (m *StoragePlanInformation) GetUpgradeAvailable()(*bool) {
    return m.upgradeAvailable
}
// Serialize serializes information the current object
func (m *StoragePlanInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("upgradeAvailable", m.GetUpgradeAvailable())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpgradeAvailable sets the upgradeAvailable property value. Indicates whether there are higher storage quota plans available. Read-only.
func (m *StoragePlanInformation) SetUpgradeAvailable(value *bool)() {
    m.upgradeAvailable = value
}
// StoragePlanInformationable 
type StoragePlanInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUpgradeAvailable()(*bool)
    SetUpgradeAvailable(value *bool)()
}
