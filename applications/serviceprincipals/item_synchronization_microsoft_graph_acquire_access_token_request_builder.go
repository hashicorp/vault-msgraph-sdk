// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\microsoft.graph.acquireAccessToken
type ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilderInternal instantiates a new MicrosoftGraphAcquireAccessTokenRequestBuilder and sets the default values.
func NewItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder) {
    m := &ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/microsoft.graph.acquireAccessToken", pathParameters),
    }
    return m
}
// NewItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder instantiates a new MicrosoftGraphAcquireAccessTokenRequestBuilder and sets the default values.
func NewItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// Post acquire an OAuth access token to authorize the Microsoft Entra provisioning service to provision users into an application.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-synchronization-acquireaccesstoken?view=graph-rest-1.0
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder) Post(ctx context.Context, body ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation acquire an OAuth access token to authorize the Microsoft Entra provisioning service to provision users into an application.
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationMicrosoftGraphAcquireAccessTokenAcquireAccessTokenPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder) {
    return NewItemSynchronizationMicrosoftGraphAcquireAccessTokenRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
