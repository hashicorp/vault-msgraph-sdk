// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Hashes 
type Hashes struct {
    // The CRC32 value of the file in little endian (if available). Read-only.
    crc32Hash *string
    // A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
    quickXorHash *string
    // SHA1 hash for the contents of the file (if available). Read-only.
    sha1Hash *string
    // SHA256 hash for the contents of the file (if available). Read-only.
    sha256Hash *string
}
// NewHashes instantiates a new hashes and sets the default values.
func NewHashes()(*Hashes) {
    m := &Hashes{
    }
    return m
}
// CreateHashesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHashesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHashes(), nil
}
// GetCrc32Hash gets the crc32Hash property value. The CRC32 value of the file in little endian (if available). Read-only.
func (m *Hashes) GetCrc32Hash()(*string) {
    return m.crc32Hash
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Hashes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["crc32Hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrc32Hash(val)
        }
        return nil
    }
    res["quickXorHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickXorHash(val)
        }
        return nil
    }
    res["sha1Hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha1Hash(val)
        }
        return nil
    }
    res["sha256Hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256Hash(val)
        }
        return nil
    }
    return res
}
// GetQuickXorHash gets the quickXorHash property value. A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
func (m *Hashes) GetQuickXorHash()(*string) {
    return m.quickXorHash
}
// GetSha1Hash gets the sha1Hash property value. SHA1 hash for the contents of the file (if available). Read-only.
func (m *Hashes) GetSha1Hash()(*string) {
    return m.sha1Hash
}
// GetSha256Hash gets the sha256Hash property value. SHA256 hash for the contents of the file (if available). Read-only.
func (m *Hashes) GetSha256Hash()(*string) {
    return m.sha256Hash
}
// Serialize serializes information the current object
func (m *Hashes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("crc32Hash", m.GetCrc32Hash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("quickXorHash", m.GetQuickXorHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha1Hash", m.GetSha1Hash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha256Hash", m.GetSha256Hash())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCrc32Hash sets the crc32Hash property value. The CRC32 value of the file in little endian (if available). Read-only.
func (m *Hashes) SetCrc32Hash(value *string)() {
    m.crc32Hash = value
}
// SetQuickXorHash sets the quickXorHash property value. A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
func (m *Hashes) SetQuickXorHash(value *string)() {
    m.quickXorHash = value
}
// SetSha1Hash sets the sha1Hash property value. SHA1 hash for the contents of the file (if available). Read-only.
func (m *Hashes) SetSha1Hash(value *string)() {
    m.sha1Hash = value
}
// SetSha256Hash sets the sha256Hash property value. SHA256 hash for the contents of the file (if available). Read-only.
func (m *Hashes) SetSha256Hash(value *string)() {
    m.sha256Hash = value
}
// Hashesable 
type Hashesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCrc32Hash()(*string)
    GetQuickXorHash()(*string)
    GetSha1Hash()(*string)
    GetSha256Hash()(*string)
    SetCrc32Hash(value *string)()
    SetQuickXorHash(value *string)()
    SetSha1Hash(value *string)()
    SetSha256Hash(value *string)()
}
