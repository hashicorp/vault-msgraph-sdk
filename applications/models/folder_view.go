// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FolderView 
type FolderView struct {
    // The method by which the folder should be sorted.
    sortBy *string
    // If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
    sortOrder *string
    // The type of view that should be used to represent the folder.
    viewType *string
}
// NewFolderView instantiates a new folderView and sets the default values.
func NewFolderView()(*FolderView) {
    m := &FolderView{
    }
    return m
}
// CreateFolderViewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFolderViewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFolderView(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FolderView) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sortBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortBy(val)
        }
        return nil
    }
    res["sortOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortOrder(val)
        }
        return nil
    }
    res["viewType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewType(val)
        }
        return nil
    }
    return res
}
// GetSortBy gets the sortBy property value. The method by which the folder should be sorted.
func (m *FolderView) GetSortBy()(*string) {
    return m.sortBy
}
// GetSortOrder gets the sortOrder property value. If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
func (m *FolderView) GetSortOrder()(*string) {
    return m.sortOrder
}
// GetViewType gets the viewType property value. The type of view that should be used to represent the folder.
func (m *FolderView) GetViewType()(*string) {
    return m.viewType
}
// Serialize serializes information the current object
func (m *FolderView) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sortBy", m.GetSortBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sortOrder", m.GetSortOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("viewType", m.GetViewType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSortBy sets the sortBy property value. The method by which the folder should be sorted.
func (m *FolderView) SetSortBy(value *string)() {
    m.sortBy = value
}
// SetSortOrder sets the sortOrder property value. If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
func (m *FolderView) SetSortOrder(value *string)() {
    m.sortOrder = value
}
// SetViewType sets the viewType property value. The type of view that should be used to represent the folder.
func (m *FolderView) SetViewType(value *string)() {
    m.viewType = value
}
// FolderViewable 
type FolderViewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSortBy()(*string)
    GetSortOrder()(*string)
    GetViewType()(*string)
    SetSortBy(value *string)()
    SetSortOrder(value *string)()
    SetViewType(value *string)()
}
