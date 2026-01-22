// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemEndpointsRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\endpoints
type ItemEndpointsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEndpointsRequestBuilderGetQueryParameters get endpoints from servicePrincipals
type ItemEndpointsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ByEndpointId gets an item from the github.com/hashicorp/vault-msgraph-sdk/applications.servicePrincipals.item.endpoints.item collection
func (m *ItemEndpointsRequestBuilder) ByEndpointId(endpointId string)(*ItemEndpointsEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if endpointId != "" {
        urlTplParams["endpoint%2Did"] = endpointId
    }
    return NewItemEndpointsEndpointItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEndpointsRequestBuilderInternal instantiates a new EndpointsRequestBuilder and sets the default values.
func NewItemEndpointsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEndpointsRequestBuilder) {
    m := &ItemEndpointsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/endpoints{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemEndpointsRequestBuilder instantiates a new EndpointsRequestBuilder and sets the default values.
func NewItemEndpointsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEndpointsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEndpointsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ItemEndpointsRequestBuilder) Count()(*ItemEndpointsCountRequestBuilder) {
    return NewItemEndpointsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get endpoints from servicePrincipals
func (m *ItemEndpointsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemEndpointsRequestBuilderGetQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.EndpointCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateEndpointCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.EndpointCollectionResponseable), nil
}
// Post create new navigation property to endpoints for servicePrincipals
func (m *ItemEndpointsRequestBuilder) Post(ctx context.Context, body i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Endpointable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Endpointable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Endpointable), nil
}
// ToGetRequestInformation get endpoints from servicePrincipals
func (m *ItemEndpointsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemEndpointsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to endpoints for servicePrincipals
func (m *ItemEndpointsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.Endpointable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemEndpointsRequestBuilder) WithUrl(rawUrl string)(*ItemEndpointsRequestBuilder) {
    return NewItemEndpointsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
