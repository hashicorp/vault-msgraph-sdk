// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\jobs\{synchronizationJob-id}\schema\directories\{directoryDefinition-id}\microsoft.graph.discover
type ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderInternal instantiates a new MicrosoftGraphDiscoverRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/directories/{directoryDefinition%2Did}/microsoft.graph.discover", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder instantiates a new MicrosoftGraphDiscoverRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderInternal(urlParams, requestAdapter)
}
// Post discover the latest schema definition for provisioning to an application. 
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-directorydefinition-discover?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.DirectoryDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateDirectoryDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.DirectoryDefinitionable), nil
}
// ToPostRequestInformation discover the latest schema definition for provisioning to an application. 
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) {
    return NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
