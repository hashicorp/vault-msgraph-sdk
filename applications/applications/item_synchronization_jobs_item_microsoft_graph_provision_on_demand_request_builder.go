// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder builds and executes requests for operations under \applications\{application-id}\synchronization\jobs\{synchronizationJob-id}\microsoft.graph.provisionOnDemand
type ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilderInternal instantiates a new MicrosoftGraphProvisionOnDemandRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder) {
    m := &ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/microsoft.graph.provisionOnDemand", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder instantiates a new MicrosoftGraphProvisionOnDemandRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilderInternal(urlParams, requestAdapter)
}
// Post select a user and provision the account on-demand. The rate limit for this API is 5 requests per 10 seconds. 
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronizationjob-provisionondemand?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.StringKeyStringValuePairable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateStringKeyStringValuePairFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.StringKeyStringValuePairable), nil
}
// ToPostRequestInformation select a user and provision the account on-demand. The rate limit for this API is 5 requests per 10 seconds. 
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder) {
    return NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
