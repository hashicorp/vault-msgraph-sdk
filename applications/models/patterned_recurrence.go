package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PatternedRecurrence 
type PatternedRecurrence struct {
    // The pattern property
    pattern RecurrencePatternable
    // The range property
    rangeEscaped RecurrenceRangeable
}
// NewPatternedRecurrence instantiates a new patternedRecurrence and sets the default values.
func NewPatternedRecurrence()(*PatternedRecurrence) {
    m := &PatternedRecurrence{
    }
    return m
}
// CreatePatternedRecurrenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePatternedRecurrenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPatternedRecurrence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PatternedRecurrence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["pattern"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecurrencePatternFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val.(RecurrencePatternable))
        }
        return nil
    }
    res["range"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecurrenceRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRangeEscaped(val.(RecurrenceRangeable))
        }
        return nil
    }
    return res
}
// GetPattern gets the pattern property value. The pattern property
func (m *PatternedRecurrence) GetPattern()(RecurrencePatternable) {
    return m.pattern
}
// GetRangeEscaped gets the range property value. The range property
func (m *PatternedRecurrence) GetRangeEscaped()(RecurrenceRangeable) {
    return m.rangeEscaped
}
// Serialize serializes information the current object
func (m *PatternedRecurrence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pattern", m.GetPattern())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("range", m.GetRangeEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPattern sets the pattern property value. The pattern property
func (m *PatternedRecurrence) SetPattern(value RecurrencePatternable)() {
    m.pattern = value
}
// SetRangeEscaped sets the range property value. The range property
func (m *PatternedRecurrence) SetRangeEscaped(value RecurrenceRangeable)() {
    m.rangeEscaped = value
}
// PatternedRecurrenceable 
type PatternedRecurrenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPattern()(RecurrencePatternable)
    GetRangeEscaped()(RecurrenceRangeable)
    SetPattern(value RecurrencePatternable)()
    SetRangeEscaped(value RecurrenceRangeable)()
}
