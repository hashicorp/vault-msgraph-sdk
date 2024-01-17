// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApprovalStatusColumn 
type ContentApprovalStatusColumn struct {
}
// NewContentApprovalStatusColumn instantiates a new contentApprovalStatusColumn and sets the default values.
func NewContentApprovalStatusColumn()(*ContentApprovalStatusColumn) {
    m := &ContentApprovalStatusColumn{
    }
    return m
}
// CreateContentApprovalStatusColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentApprovalStatusColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentApprovalStatusColumn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentApprovalStatusColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ContentApprovalStatusColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// ContentApprovalStatusColumnable 
type ContentApprovalStatusColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
