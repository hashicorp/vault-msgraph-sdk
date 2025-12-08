// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPrint 
type UserPrint struct {
    // The recentPrinterShares property
    recentPrinterShares []PrinterShareable
}
// NewUserPrint instantiates a new userPrint and sets the default values.
func NewUserPrint()(*UserPrint) {
    m := &UserPrint{
    }
    return m
}
// CreateUserPrintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserPrintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPrint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserPrint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recentPrinterShares"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrinterShareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterShareable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrinterShareable)
                }
            }
            m.SetRecentPrinterShares(res)
        }
        return nil
    }
    return res
}
// GetRecentPrinterShares gets the recentPrinterShares property value. The recentPrinterShares property
func (m *UserPrint) GetRecentPrinterShares()([]PrinterShareable) {
    return m.recentPrinterShares
}
// Serialize serializes information the current object
func (m *UserPrint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRecentPrinterShares() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecentPrinterShares()))
        for i, v := range m.GetRecentPrinterShares() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("recentPrinterShares", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecentPrinterShares sets the recentPrinterShares property value. The recentPrinterShares property
func (m *UserPrint) SetRecentPrinterShares(value []PrinterShareable)() {
    m.recentPrinterShares = value
}
// UserPrintable 
type UserPrintable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRecentPrinterShares()([]PrinterShareable)
    SetRecentPrinterShares(value []PrinterShareable)()
}
