package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Folder 
type Folder struct {
    // Number of children contained immediately within this container.
    childCount *int32
    // The view property
    view FolderViewable
}
// NewFolder instantiates a new folder and sets the default values.
func NewFolder()(*Folder) {
    m := &Folder{
    }
    return m
}
// CreateFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFolder(), nil
}
// GetChildCount gets the childCount property value. Number of children contained immediately within this container.
func (m *Folder) GetChildCount()(*int32) {
    return m.childCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Folder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["childCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChildCount(val)
        }
        return nil
    }
    res["view"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFolderViewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetView(val.(FolderViewable))
        }
        return nil
    }
    return res
}
// GetView gets the view property value. The view property
func (m *Folder) GetView()(FolderViewable) {
    return m.view
}
// Serialize serializes information the current object
func (m *Folder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("childCount", m.GetChildCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("view", m.GetView())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildCount sets the childCount property value. Number of children contained immediately within this container.
func (m *Folder) SetChildCount(value *int32)() {
    m.childCount = value
}
// SetView sets the view property value. The view property
func (m *Folder) SetView(value FolderViewable)() {
    m.view = value
}
// Folderable 
type Folderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildCount()(*int32)
    GetView()(FolderViewable)
    SetChildCount(value *int32)()
    SetView(value FolderViewable)()
}
