package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
)

// ItemMicrosoftGraphAddKeyAddKeyPostRequestBody 
type ItemMicrosoftGraphAddKeyAddKeyPostRequestBody struct {
    // The keyCredential property
    keyCredential i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.KeyCredentialable
    // The passwordCredential property
    passwordCredential i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable
    // The proof property
    proof *string
}
// NewItemMicrosoftGraphAddKeyAddKeyPostRequestBody instantiates a new ItemMicrosoftGraphAddKeyAddKeyPostRequestBody and sets the default values.
func NewItemMicrosoftGraphAddKeyAddKeyPostRequestBody()(*ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) {
    m := &ItemMicrosoftGraphAddKeyAddKeyPostRequestBody{
    }
    return m
}
// CreateItemMicrosoftGraphAddKeyAddKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphAddKeyAddKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphAddKeyAddKeyPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredential(val.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.KeyCredentialable))
        }
        return nil
    }
    res["passwordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable))
        }
        return nil
    }
    res["proof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
// GetKeyCredential gets the keyCredential property value. The keyCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetKeyCredential()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.KeyCredentialable) {
    return m.keyCredential
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetPasswordCredential()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable) {
    return m.passwordCredential
}
// GetProof gets the proof property value. The proof property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("keyCredential", m.GetKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKeyCredential sets the keyCredential property value. The keyCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetKeyCredential(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.KeyCredentialable)() {
    m.keyCredential = value
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetPasswordCredential(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable)() {
    m.passwordCredential = value
}
// SetProof sets the proof property value. The proof property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
// ItemMicrosoftGraphAddKeyAddKeyPostRequestBodyable 
type ItemMicrosoftGraphAddKeyAddKeyPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeyCredential()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.KeyCredentialable)
    GetPasswordCredential()(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable)
    GetProof()(*string)
    SetKeyCredential(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.KeyCredentialable)()
    SetPasswordCredential(value i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.PasswordCredentialable)()
    SetProof(value *string)()
}
