// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ObjectIdentity 
type ObjectIdentity struct {
    // Specifies the issuer of the identity, for example facebook.com.For local accounts (where signInType isn't federated), this property is the local B2C tenant default domain name, for example contoso.onmicrosoft.com.For guests from other Microsoft Entra organization, this is the domain of the federated organization, for example contoso.com.Supports $filter. 512 character limit.
    issuer *string
    // Specifies the unique identifier assigned to the user by the issuer. The combination of issuer and issuerAssignedId must be unique within the organization. Represents the sign-in name for the user, when signInType is set to emailAddress or userName (also known as local accounts).When signInType is set to: emailAddress, (or a custom string that starts with emailAddress like emailAddress1) issuerAssignedId must be a valid email addressuserName, issuerAssignedId must begin with alphabetical character or number, and can only contain alphanumeric characters and the following symbols: - or Supports $filter. 64 character limit.
    issuerAssignedId *string
    // Specifies the user sign-in types in your directory, such as emailAddress, userName, federated, or userPrincipalName. federated represents a unique identifier for a user from an issuer, that can be in any format chosen by the issuer. Setting or updating a userPrincipalName identity will update the value of the userPrincipalName property on the user object. The validations performed on the userPrincipalName property on the user object, for example, verified domains and acceptable characters, will be performed when setting or updating a userPrincipalName identity. Other validation is enforced on issuerAssignedId when the sign-in type is set to emailAddress or userName. This property can also be set to any custom string.
    signInType *string
}
// NewObjectIdentity instantiates a new objectIdentity and sets the default values.
func NewObjectIdentity()(*ObjectIdentity) {
    m := &ObjectIdentity{
    }
    return m
}
// CreateObjectIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerAssignedId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerAssignedId(val)
        }
        return nil
    }
    res["signInType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInType(val)
        }
        return nil
    }
    return res
}
// GetIssuer gets the issuer property value. Specifies the issuer of the identity, for example facebook.com.For local accounts (where signInType isn't federated), this property is the local B2C tenant default domain name, for example contoso.onmicrosoft.com.For guests from other Microsoft Entra organization, this is the domain of the federated organization, for example contoso.com.Supports $filter. 512 character limit.
func (m *ObjectIdentity) GetIssuer()(*string) {
    return m.issuer
}
// GetIssuerAssignedId gets the issuerAssignedId property value. Specifies the unique identifier assigned to the user by the issuer. The combination of issuer and issuerAssignedId must be unique within the organization. Represents the sign-in name for the user, when signInType is set to emailAddress or userName (also known as local accounts).When signInType is set to: emailAddress, (or a custom string that starts with emailAddress like emailAddress1) issuerAssignedId must be a valid email addressuserName, issuerAssignedId must begin with alphabetical character or number, and can only contain alphanumeric characters and the following symbols: - or Supports $filter. 64 character limit.
func (m *ObjectIdentity) GetIssuerAssignedId()(*string) {
    return m.issuerAssignedId
}
// GetSignInType gets the signInType property value. Specifies the user sign-in types in your directory, such as emailAddress, userName, federated, or userPrincipalName. federated represents a unique identifier for a user from an issuer, that can be in any format chosen by the issuer. Setting or updating a userPrincipalName identity will update the value of the userPrincipalName property on the user object. The validations performed on the userPrincipalName property on the user object, for example, verified domains and acceptable characters, will be performed when setting or updating a userPrincipalName identity. Other validation is enforced on issuerAssignedId when the sign-in type is set to emailAddress or userName. This property can also be set to any custom string.
func (m *ObjectIdentity) GetSignInType()(*string) {
    return m.signInType
}
// Serialize serializes information the current object
func (m *ObjectIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuerAssignedId", m.GetIssuerAssignedId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signInType", m.GetSignInType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIssuer sets the issuer property value. Specifies the issuer of the identity, for example facebook.com.For local accounts (where signInType isn't federated), this property is the local B2C tenant default domain name, for example contoso.onmicrosoft.com.For guests from other Microsoft Entra organization, this is the domain of the federated organization, for example contoso.com.Supports $filter. 512 character limit.
func (m *ObjectIdentity) SetIssuer(value *string)() {
    m.issuer = value
}
// SetIssuerAssignedId sets the issuerAssignedId property value. Specifies the unique identifier assigned to the user by the issuer. The combination of issuer and issuerAssignedId must be unique within the organization. Represents the sign-in name for the user, when signInType is set to emailAddress or userName (also known as local accounts).When signInType is set to: emailAddress, (or a custom string that starts with emailAddress like emailAddress1) issuerAssignedId must be a valid email addressuserName, issuerAssignedId must begin with alphabetical character or number, and can only contain alphanumeric characters and the following symbols: - or Supports $filter. 64 character limit.
func (m *ObjectIdentity) SetIssuerAssignedId(value *string)() {
    m.issuerAssignedId = value
}
// SetSignInType sets the signInType property value. Specifies the user sign-in types in your directory, such as emailAddress, userName, federated, or userPrincipalName. federated represents a unique identifier for a user from an issuer, that can be in any format chosen by the issuer. Setting or updating a userPrincipalName identity will update the value of the userPrincipalName property on the user object. The validations performed on the userPrincipalName property on the user object, for example, verified domains and acceptable characters, will be performed when setting or updating a userPrincipalName identity. Other validation is enforced on issuerAssignedId when the sign-in type is set to emailAddress or userName. This property can also be set to any custom string.
func (m *ObjectIdentity) SetSignInType(value *string)() {
    m.signInType = value
}
// ObjectIdentityable 
type ObjectIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssuer()(*string)
    GetIssuerAssignedId()(*string)
    GetSignInType()(*string)
    SetIssuer(value *string)()
    SetIssuerAssignedId(value *string)()
    SetSignInType(value *string)()
}
