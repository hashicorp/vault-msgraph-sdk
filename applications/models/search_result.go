// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchResult 
type SearchResult struct {
    // A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if the user interacts with this item to improve the quality of results.
    onClickTelemetryUrl *string
}
// NewSearchResult instantiates a new searchResult and sets the default values.
func NewSearchResult()(*SearchResult) {
    m := &SearchResult{
    }
    return m
}
// CreateSearchResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["onClickTelemetryUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnClickTelemetryUrl(val)
        }
        return nil
    }
    return res
}
// GetOnClickTelemetryUrl gets the onClickTelemetryUrl property value. A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if the user interacts with this item to improve the quality of results.
func (m *SearchResult) GetOnClickTelemetryUrl()(*string) {
    return m.onClickTelemetryUrl
}
// Serialize serializes information the current object
func (m *SearchResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("onClickTelemetryUrl", m.GetOnClickTelemetryUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnClickTelemetryUrl sets the onClickTelemetryUrl property value. A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if the user interacts with this item to improve the quality of results.
func (m *SearchResult) SetOnClickTelemetryUrl(value *string)() {
    m.onClickTelemetryUrl = value
}
// SearchResultable 
type SearchResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOnClickTelemetryUrl()(*string)
    SetOnClickTelemetryUrl(value *string)()
}
