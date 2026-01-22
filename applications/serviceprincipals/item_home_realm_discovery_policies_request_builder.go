// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemHomeRealmDiscoveryPoliciesRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\homeRealmDiscoveryPolicies
type ItemHomeRealmDiscoveryPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters list the homeRealmDiscoveryPolicy objects that are assigned to a servicePrincipal.
type ItemHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters struct {
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
// ByHomeRealmDiscoveryPolicyId gets an item from the github.com/hashicorp/vault-msgraph-sdk/applications.servicePrincipals.item.homeRealmDiscoveryPolicies.item collection
func (m *ItemHomeRealmDiscoveryPoliciesRequestBuilder) ByHomeRealmDiscoveryPolicyId(homeRealmDiscoveryPolicyId string)(*ItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if homeRealmDiscoveryPolicyId != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = homeRealmDiscoveryPolicyId
    }
    return NewItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemHomeRealmDiscoveryPoliciesRequestBuilderInternal instantiates a new HomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    m := &ItemHomeRealmDiscoveryPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/homeRealmDiscoveryPolicies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemHomeRealmDiscoveryPoliciesRequestBuilder instantiates a new HomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewItemHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHomeRealmDiscoveryPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ItemHomeRealmDiscoveryPoliciesRequestBuilder) Count()(*ItemHomeRealmDiscoveryPoliciesCountRequestBuilder) {
    return NewItemHomeRealmDiscoveryPoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list the homeRealmDiscoveryPolicy objects that are assigned to a servicePrincipal.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-list-homerealmdiscoverypolicies?view=graph-rest-1.0
func (m *ItemHomeRealmDiscoveryPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.HomeRealmDiscoveryPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateHomeRealmDiscoveryPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.HomeRealmDiscoveryPolicyCollectionResponseable), nil
}
// Ref the Ref property
func (m *ItemHomeRealmDiscoveryPoliciesRequestBuilder) Ref()(*ItemHomeRealmDiscoveryPoliciesRefRequestBuilder) {
    return NewItemHomeRealmDiscoveryPoliciesRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation list the homeRealmDiscoveryPolicy objects that are assigned to a servicePrincipal.
func (m *ItemHomeRealmDiscoveryPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemHomeRealmDiscoveryPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewItemHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
