package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Certification 
type Certification struct {
    // URL that shows certification details for the application.
    certificationDetailsUrl *string
    // The timestamp when the current certification for the application expires.
    certificationExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates whether the application is certified by Microsoft.
    isCertifiedByMicrosoft *bool
    // Indicates whether the application has been self-attested by the application developer or the publisher.
    isPublisherAttested *bool
    // The timestamp when the certification for the application was most recently added or updated.
    lastCertificationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCertification instantiates a new certification and sets the default values.
func NewCertification()(*Certification) {
    m := &Certification{
    }
    return m
}
// CreateCertificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertification(), nil
}
// GetCertificationDetailsUrl gets the certificationDetailsUrl property value. URL that shows certification details for the application.
func (m *Certification) GetCertificationDetailsUrl()(*string) {
    return m.certificationDetailsUrl
}
// GetCertificationExpirationDateTime gets the certificationExpirationDateTime property value. The timestamp when the current certification for the application expires.
func (m *Certification) GetCertificationExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificationExpirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Certification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificationDetailsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationDetailsUrl(val)
        }
        return nil
    }
    res["certificationExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationExpirationDateTime(val)
        }
        return nil
    }
    res["isCertifiedByMicrosoft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCertifiedByMicrosoft(val)
        }
        return nil
    }
    res["isPublisherAttested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPublisherAttested(val)
        }
        return nil
    }
    res["lastCertificationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCertificationDateTime(val)
        }
        return nil
    }
    return res
}
// GetIsCertifiedByMicrosoft gets the isCertifiedByMicrosoft property value. Indicates whether the application is certified by Microsoft.
func (m *Certification) GetIsCertifiedByMicrosoft()(*bool) {
    return m.isCertifiedByMicrosoft
}
// GetIsPublisherAttested gets the isPublisherAttested property value. Indicates whether the application has been self-attested by the application developer or the publisher.
func (m *Certification) GetIsPublisherAttested()(*bool) {
    return m.isPublisherAttested
}
// GetLastCertificationDateTime gets the lastCertificationDateTime property value. The timestamp when the certification for the application was most recently added or updated.
func (m *Certification) GetLastCertificationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCertificationDateTime
}
// Serialize serializes information the current object
func (m *Certification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("certificationExpirationDateTime", m.GetCertificationExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPublisherAttested", m.GetIsPublisherAttested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastCertificationDateTime", m.GetLastCertificationDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificationDetailsUrl sets the certificationDetailsUrl property value. URL that shows certification details for the application.
func (m *Certification) SetCertificationDetailsUrl(value *string)() {
    m.certificationDetailsUrl = value
}
// SetCertificationExpirationDateTime sets the certificationExpirationDateTime property value. The timestamp when the current certification for the application expires.
func (m *Certification) SetCertificationExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificationExpirationDateTime = value
}
// SetIsCertifiedByMicrosoft sets the isCertifiedByMicrosoft property value. Indicates whether the application is certified by Microsoft.
func (m *Certification) SetIsCertifiedByMicrosoft(value *bool)() {
    m.isCertifiedByMicrosoft = value
}
// SetIsPublisherAttested sets the isPublisherAttested property value. Indicates whether the application has been self-attested by the application developer or the publisher.
func (m *Certification) SetIsPublisherAttested(value *bool)() {
    m.isPublisherAttested = value
}
// SetLastCertificationDateTime sets the lastCertificationDateTime property value. The timestamp when the certification for the application was most recently added or updated.
func (m *Certification) SetLastCertificationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCertificationDateTime = value
}
// Certificationable 
type Certificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificationDetailsUrl()(*string)
    GetCertificationExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsCertifiedByMicrosoft()(*bool)
    GetIsPublisherAttested()(*bool)
    GetLastCertificationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCertificationDetailsUrl(value *string)()
    SetCertificationExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsCertifiedByMicrosoft(value *bool)()
    SetIsPublisherAttested(value *bool)()
    SetLastCertificationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
