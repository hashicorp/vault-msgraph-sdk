package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParentalControlSettings 
type ParentalControlSettings struct {
    // Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
    countriesBlockedForMinors []string
    // Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
    legalAgeGroupRule *string
}
// NewParentalControlSettings instantiates a new parentalControlSettings and sets the default values.
func NewParentalControlSettings()(*ParentalControlSettings) {
    m := &ParentalControlSettings{
    }
    return m
}
// CreateParentalControlSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParentalControlSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParentalControlSettings(), nil
}
// GetCountriesBlockedForMinors gets the countriesBlockedForMinors property value. Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
func (m *ParentalControlSettings) GetCountriesBlockedForMinors()([]string) {
    return m.countriesBlockedForMinors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParentalControlSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["countriesBlockedForMinors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCountriesBlockedForMinors(res)
        }
        return nil
    }
    res["legalAgeGroupRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegalAgeGroupRule(val)
        }
        return nil
    }
    return res
}
// GetLegalAgeGroupRule gets the legalAgeGroupRule property value. Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
func (m *ParentalControlSettings) GetLegalAgeGroupRule()(*string) {
    return m.legalAgeGroupRule
}
// Serialize serializes information the current object
func (m *ParentalControlSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCountriesBlockedForMinors() != nil {
        err := writer.WriteCollectionOfStringValues("countriesBlockedForMinors", m.GetCountriesBlockedForMinors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("legalAgeGroupRule", m.GetLegalAgeGroupRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCountriesBlockedForMinors sets the countriesBlockedForMinors property value. Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list.
func (m *ParentalControlSettings) SetCountriesBlockedForMinors(value []string)() {
    m.countriesBlockedForMinors = value
}
// SetLegalAgeGroupRule sets the legalAgeGroupRule property value. Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app.
func (m *ParentalControlSettings) SetLegalAgeGroupRule(value *string)() {
    m.legalAgeGroupRule = value
}
// ParentalControlSettingsable 
type ParentalControlSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountriesBlockedForMinors()([]string)
    GetLegalAgeGroupRule()(*string)
    SetCountriesBlockedForMinors(value []string)()
    SetLegalAgeGroupRule(value *string)()
}
