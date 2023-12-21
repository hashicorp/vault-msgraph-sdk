package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalLockConfiguration 
type ServicePrincipalLockConfiguration struct {
    // Enables locking all sensitive properties. The sensitive properties are keyCredentials, passwordCredentials, and tokenEncryptionKeyId.
    allProperties *bool
    // Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Sign.
    credentialsWithUsageSign *bool
    // Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Verify. This locks OAuth service principals.
    credentialsWithUsageVerify *bool
    // Enables or disables service principal lock configuration. To allow the sensitive properties to be updated, update this property to false to disable the lock on the service principal.
    isEnabled *bool
    // Locks the tokenEncryptionKeyId property for modification on the service principal.
    tokenEncryptionKeyId *bool
}
// NewServicePrincipalLockConfiguration instantiates a new servicePrincipalLockConfiguration and sets the default values.
func NewServicePrincipalLockConfiguration()(*ServicePrincipalLockConfiguration) {
    m := &ServicePrincipalLockConfiguration{
    }
    return m
}
// CreateServicePrincipalLockConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalLockConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalLockConfiguration(), nil
}
// GetAllProperties gets the allProperties property value. Enables locking all sensitive properties. The sensitive properties are keyCredentials, passwordCredentials, and tokenEncryptionKeyId.
func (m *ServicePrincipalLockConfiguration) GetAllProperties()(*bool) {
    return m.allProperties
}
// GetCredentialsWithUsageSign gets the credentialsWithUsageSign property value. Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Sign.
func (m *ServicePrincipalLockConfiguration) GetCredentialsWithUsageSign()(*bool) {
    return m.credentialsWithUsageSign
}
// GetCredentialsWithUsageVerify gets the credentialsWithUsageVerify property value. Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Verify. This locks OAuth service principals.
func (m *ServicePrincipalLockConfiguration) GetCredentialsWithUsageVerify()(*bool) {
    return m.credentialsWithUsageVerify
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalLockConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllProperties(val)
        }
        return nil
    }
    res["credentialsWithUsageSign"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialsWithUsageSign(val)
        }
        return nil
    }
    res["credentialsWithUsageVerify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentialsWithUsageVerify(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["tokenEncryptionKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenEncryptionKeyId(val)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Enables or disables service principal lock configuration. To allow the sensitive properties to be updated, update this property to false to disable the lock on the service principal.
func (m *ServicePrincipalLockConfiguration) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Locks the tokenEncryptionKeyId property for modification on the service principal.
func (m *ServicePrincipalLockConfiguration) GetTokenEncryptionKeyId()(*bool) {
    return m.tokenEncryptionKeyId
}
// Serialize serializes information the current object
func (m *ServicePrincipalLockConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allProperties", m.GetAllProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("credentialsWithUsageSign", m.GetCredentialsWithUsageSign())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("credentialsWithUsageVerify", m.GetCredentialsWithUsageVerify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllProperties sets the allProperties property value. Enables locking all sensitive properties. The sensitive properties are keyCredentials, passwordCredentials, and tokenEncryptionKeyId.
func (m *ServicePrincipalLockConfiguration) SetAllProperties(value *bool)() {
    m.allProperties = value
}
// SetCredentialsWithUsageSign sets the credentialsWithUsageSign property value. Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Sign.
func (m *ServicePrincipalLockConfiguration) SetCredentialsWithUsageSign(value *bool)() {
    m.credentialsWithUsageSign = value
}
// SetCredentialsWithUsageVerify sets the credentialsWithUsageVerify property value. Locks the keyCredentials and passwordCredentials properties for modification where credential usage type is Verify. This locks OAuth service principals.
func (m *ServicePrincipalLockConfiguration) SetCredentialsWithUsageVerify(value *bool)() {
    m.credentialsWithUsageVerify = value
}
// SetIsEnabled sets the isEnabled property value. Enables or disables service principal lock configuration. To allow the sensitive properties to be updated, update this property to false to disable the lock on the service principal.
func (m *ServicePrincipalLockConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Locks the tokenEncryptionKeyId property for modification on the service principal.
func (m *ServicePrincipalLockConfiguration) SetTokenEncryptionKeyId(value *bool)() {
    m.tokenEncryptionKeyId = value
}
// ServicePrincipalLockConfigurationable 
type ServicePrincipalLockConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllProperties()(*bool)
    GetCredentialsWithUsageSign()(*bool)
    GetCredentialsWithUsageVerify()(*bool)
    GetIsEnabled()(*bool)
    GetTokenEncryptionKeyId()(*bool)
    SetAllProperties(value *bool)()
    SetCredentialsWithUsageSign(value *bool)()
    SetCredentialsWithUsageVerify(value *bool)()
    SetIsEnabled(value *bool)()
    SetTokenEncryptionKeyId(value *bool)()
}
