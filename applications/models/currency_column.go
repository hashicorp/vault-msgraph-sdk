// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CurrencyColumn 
type CurrencyColumn struct {
    // Specifies the locale from which to infer the currency symbol.
    locale *string
}
// NewCurrencyColumn instantiates a new currencyColumn and sets the default values.
func NewCurrencyColumn()(*CurrencyColumn) {
    m := &CurrencyColumn{
    }
    return m
}
// CreateCurrencyColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCurrencyColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCurrencyColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CurrencyColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    return res
}
// GetLocale gets the locale property value. Specifies the locale from which to infer the currency symbol.
func (m *CurrencyColumn) GetLocale()(*string) {
    return m.locale
}
// Serialize serializes information the current object
func (m *CurrencyColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLocale sets the locale property value. Specifies the locale from which to infer the currency symbol.
func (m *CurrencyColumn) SetLocale(value *string)() {
    m.locale = value
}
// CurrencyColumnable 
type CurrencyColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocale()(*string)
    SetLocale(value *string)()
}
