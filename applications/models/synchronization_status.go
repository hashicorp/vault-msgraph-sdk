// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationStatus 
type SynchronizationStatus struct {
    // The code property
    code *SynchronizationStatusCode
    // Number of consecutive times this job failed.
    countSuccessiveCompleteFailures *int64
    // true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
    escrowsPruned *bool
    // The lastExecution property
    lastExecution SynchronizationTaskExecutionable
    // The lastSuccessfulExecution property
    lastSuccessfulExecution SynchronizationTaskExecutionable
    // The lastSuccessfulExecutionWithExports property
    lastSuccessfulExecutionWithExports SynchronizationTaskExecutionable
    // Details of the progress of a job toward completion.
    progress []SynchronizationProgressable
    // The quarantine property
    quarantine SynchronizationQuarantineable
    // The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateFirstAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateLastAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Count of synchronized objects, listed by object type.
    synchronizedEntryCountByType []StringKeyLongValuePairable
    // In the event of an error, the URL with the troubleshooting steps for the issue.
    troubleshootingUrl *string
}
// NewSynchronizationStatus instantiates a new synchronizationStatus and sets the default values.
func NewSynchronizationStatus()(*SynchronizationStatus) {
    m := &SynchronizationStatus{
    }
    return m
}
// CreateSynchronizationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationStatus(), nil
}
// GetCode gets the code property value. The code property
func (m *SynchronizationStatus) GetCode()(*SynchronizationStatusCode) {
    return m.code
}
// GetCountSuccessiveCompleteFailures gets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) GetCountSuccessiveCompleteFailures()(*int64) {
    return m.countSuccessiveCompleteFailures
}
// GetEscrowsPruned gets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) GetEscrowsPruned()(*bool) {
    return m.escrowsPruned
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationStatusCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val.(*SynchronizationStatusCode))
        }
        return nil
    }
    res["countSuccessiveCompleteFailures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountSuccessiveCompleteFailures(val)
        }
        return nil
    }
    res["escrowsPruned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscrowsPruned(val)
        }
        return nil
    }
    res["lastExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastExecution(val.(SynchronizationTaskExecutionable))
        }
        return nil
    }
    res["lastSuccessfulExecution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulExecution(val.(SynchronizationTaskExecutionable))
        }
        return nil
    }
    res["lastSuccessfulExecutionWithExports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulExecutionWithExports(val.(SynchronizationTaskExecutionable))
        }
        return nil
    }
    res["progress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationProgressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationProgressable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SynchronizationProgressable)
                }
            }
            m.SetProgress(res)
        }
        return nil
    }
    res["quarantine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationQuarantineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantine(val.(SynchronizationQuarantineable))
        }
        return nil
    }
    res["steadyStateFirstAchievedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteadyStateFirstAchievedTime(val)
        }
        return nil
    }
    res["steadyStateLastAchievedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteadyStateLastAchievedTime(val)
        }
        return nil
    }
    res["synchronizedEntryCountByType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStringKeyLongValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyLongValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StringKeyLongValuePairable)
                }
            }
            m.SetSynchronizedEntryCountByType(res)
        }
        return nil
    }
    res["troubleshootingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingUrl(val)
        }
        return nil
    }
    return res
}
// GetLastExecution gets the lastExecution property value. The lastExecution property
func (m *SynchronizationStatus) GetLastExecution()(SynchronizationTaskExecutionable) {
    return m.lastExecution
}
// GetLastSuccessfulExecution gets the lastSuccessfulExecution property value. The lastSuccessfulExecution property
func (m *SynchronizationStatus) GetLastSuccessfulExecution()(SynchronizationTaskExecutionable) {
    return m.lastSuccessfulExecution
}
// GetLastSuccessfulExecutionWithExports gets the lastSuccessfulExecutionWithExports property value. The lastSuccessfulExecutionWithExports property
func (m *SynchronizationStatus) GetLastSuccessfulExecutionWithExports()(SynchronizationTaskExecutionable) {
    return m.lastSuccessfulExecutionWithExports
}
// GetProgress gets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) GetProgress()([]SynchronizationProgressable) {
    return m.progress
}
// GetQuarantine gets the quarantine property value. The quarantine property
func (m *SynchronizationStatus) GetQuarantine()(SynchronizationQuarantineable) {
    return m.quarantine
}
// GetSteadyStateFirstAchievedTime gets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.steadyStateFirstAchievedTime
}
// GetSteadyStateLastAchievedTime gets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.steadyStateLastAchievedTime
}
// GetSynchronizedEntryCountByType gets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) GetSynchronizedEntryCountByType()([]StringKeyLongValuePairable) {
    return m.synchronizedEntryCountByType
}
// GetTroubleshootingUrl gets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) GetTroubleshootingUrl()(*string) {
    return m.troubleshootingUrl
}
// Serialize serializes information the current object
func (m *SynchronizationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCode() != nil {
        cast := (*m.GetCode()).String()
        err := writer.WriteStringValue("code", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countSuccessiveCompleteFailures", m.GetCountSuccessiveCompleteFailures())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("escrowsPruned", m.GetEscrowsPruned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastExecution", m.GetLastExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastSuccessfulExecution", m.GetLastSuccessfulExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastSuccessfulExecutionWithExports", m.GetLastSuccessfulExecutionWithExports())
        if err != nil {
            return err
        }
    }
    if m.GetProgress() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProgress()))
        for i, v := range m.GetProgress() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("progress", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("quarantine", m.GetQuarantine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("steadyStateFirstAchievedTime", m.GetSteadyStateFirstAchievedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("steadyStateLastAchievedTime", m.GetSteadyStateLastAchievedTime())
        if err != nil {
            return err
        }
    }
    if m.GetSynchronizedEntryCountByType() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSynchronizedEntryCountByType()))
        for i, v := range m.GetSynchronizedEntryCountByType() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("synchronizedEntryCountByType", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("troubleshootingUrl", m.GetTroubleshootingUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCode sets the code property value. The code property
func (m *SynchronizationStatus) SetCode(value *SynchronizationStatusCode)() {
    m.code = value
}
// SetCountSuccessiveCompleteFailures sets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) SetCountSuccessiveCompleteFailures(value *int64)() {
    m.countSuccessiveCompleteFailures = value
}
// SetEscrowsPruned sets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) SetEscrowsPruned(value *bool)() {
    m.escrowsPruned = value
}
// SetLastExecution sets the lastExecution property value. The lastExecution property
func (m *SynchronizationStatus) SetLastExecution(value SynchronizationTaskExecutionable)() {
    m.lastExecution = value
}
// SetLastSuccessfulExecution sets the lastSuccessfulExecution property value. The lastSuccessfulExecution property
func (m *SynchronizationStatus) SetLastSuccessfulExecution(value SynchronizationTaskExecutionable)() {
    m.lastSuccessfulExecution = value
}
// SetLastSuccessfulExecutionWithExports sets the lastSuccessfulExecutionWithExports property value. The lastSuccessfulExecutionWithExports property
func (m *SynchronizationStatus) SetLastSuccessfulExecutionWithExports(value SynchronizationTaskExecutionable)() {
    m.lastSuccessfulExecutionWithExports = value
}
// SetProgress sets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) SetProgress(value []SynchronizationProgressable)() {
    m.progress = value
}
// SetQuarantine sets the quarantine property value. The quarantine property
func (m *SynchronizationStatus) SetQuarantine(value SynchronizationQuarantineable)() {
    m.quarantine = value
}
// SetSteadyStateFirstAchievedTime sets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.steadyStateFirstAchievedTime = value
}
// SetSteadyStateLastAchievedTime sets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.steadyStateLastAchievedTime = value
}
// SetSynchronizedEntryCountByType sets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) SetSynchronizedEntryCountByType(value []StringKeyLongValuePairable)() {
    m.synchronizedEntryCountByType = value
}
// SetTroubleshootingUrl sets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) SetTroubleshootingUrl(value *string)() {
    m.troubleshootingUrl = value
}
// SynchronizationStatusable 
type SynchronizationStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*SynchronizationStatusCode)
    GetCountSuccessiveCompleteFailures()(*int64)
    GetEscrowsPruned()(*bool)
    GetLastExecution()(SynchronizationTaskExecutionable)
    GetLastSuccessfulExecution()(SynchronizationTaskExecutionable)
    GetLastSuccessfulExecutionWithExports()(SynchronizationTaskExecutionable)
    GetProgress()([]SynchronizationProgressable)
    GetQuarantine()(SynchronizationQuarantineable)
    GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSynchronizedEntryCountByType()([]StringKeyLongValuePairable)
    GetTroubleshootingUrl()(*string)
    SetCode(value *SynchronizationStatusCode)()
    SetCountSuccessiveCompleteFailures(value *int64)()
    SetEscrowsPruned(value *bool)()
    SetLastExecution(value SynchronizationTaskExecutionable)()
    SetLastSuccessfulExecution(value SynchronizationTaskExecutionable)()
    SetLastSuccessfulExecutionWithExports(value SynchronizationTaskExecutionable)()
    SetProgress(value []SynchronizationProgressable)()
    SetQuarantine(value SynchronizationQuarantineable)()
    SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSynchronizedEntryCountByType(value []StringKeyLongValuePairable)()
    SetTroubleshootingUrl(value *string)()
}
