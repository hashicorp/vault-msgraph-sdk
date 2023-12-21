package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicationFacet 
type PublicationFacet struct {
    // The checkedOutBy property
    checkedOutBy IdentitySetable
    // The state of publication for this document. Either published or checkout. Read-only.
    level *string
    // The unique identifier for the version that is visible to the current caller. Read-only.
    versionId *string
}
// NewPublicationFacet instantiates a new publicationFacet and sets the default values.
func NewPublicationFacet()(*PublicationFacet) {
    m := &PublicationFacet{
    }
    return m
}
// CreatePublicationFacetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicationFacetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicationFacet(), nil
}
// GetCheckedOutBy gets the checkedOutBy property value. The checkedOutBy property
func (m *PublicationFacet) GetCheckedOutBy()(IdentitySetable) {
    return m.checkedOutBy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicationFacet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checkedOutBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckedOutBy(val.(IdentitySetable))
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["versionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionId(val)
        }
        return nil
    }
    return res
}
// GetLevel gets the level property value. The state of publication for this document. Either published or checkout. Read-only.
func (m *PublicationFacet) GetLevel()(*string) {
    return m.level
}
// GetVersionId gets the versionId property value. The unique identifier for the version that is visible to the current caller. Read-only.
func (m *PublicationFacet) GetVersionId()(*string) {
    return m.versionId
}
// Serialize serializes information the current object
func (m *PublicationFacet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("checkedOutBy", m.GetCheckedOutBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionId", m.GetVersionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheckedOutBy sets the checkedOutBy property value. The checkedOutBy property
func (m *PublicationFacet) SetCheckedOutBy(value IdentitySetable)() {
    m.checkedOutBy = value
}
// SetLevel sets the level property value. The state of publication for this document. Either published or checkout. Read-only.
func (m *PublicationFacet) SetLevel(value *string)() {
    m.level = value
}
// SetVersionId sets the versionId property value. The unique identifier for the version that is visible to the current caller. Read-only.
func (m *PublicationFacet) SetVersionId(value *string)() {
    m.versionId = value
}
// PublicationFacetable 
type PublicationFacetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckedOutBy()(IdentitySetable)
    GetLevel()(*string)
    GetVersionId()(*string)
    SetCheckedOutBy(value IdentitySetable)()
    SetLevel(value *string)()
    SetVersionId(value *string)()
}
