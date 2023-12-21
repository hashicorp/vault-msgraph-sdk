package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmployeeOrgData 
type EmployeeOrgData struct {
    // The cost center associated with the user. Returned only on $select. Supports $filter.
    costCenter *string
    // The name of the division in which the user works. Returned only on $select. Supports $filter.
    division *string
}
// NewEmployeeOrgData instantiates a new employeeOrgData and sets the default values.
func NewEmployeeOrgData()(*EmployeeOrgData) {
    m := &EmployeeOrgData{
    }
    return m
}
// CreateEmployeeOrgDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmployeeOrgDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmployeeOrgData(), nil
}
// GetCostCenter gets the costCenter property value. The cost center associated with the user. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) GetCostCenter()(*string) {
    return m.costCenter
}
// GetDivision gets the division property value. The name of the division in which the user works. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) GetDivision()(*string) {
    return m.division
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmployeeOrgData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["costCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCostCenter(val)
        }
        return nil
    }
    res["division"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDivision(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EmployeeOrgData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("costCenter", m.GetCostCenter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("division", m.GetDivision())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCostCenter sets the costCenter property value. The cost center associated with the user. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) SetCostCenter(value *string)() {
    m.costCenter = value
}
// SetDivision sets the division property value. The name of the division in which the user works. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) SetDivision(value *string)() {
    m.division = value
}
// EmployeeOrgDataable 
type EmployeeOrgDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCostCenter()(*string)
    GetDivision()(*string)
    SetCostCenter(value *string)()
    SetDivision(value *string)()
}