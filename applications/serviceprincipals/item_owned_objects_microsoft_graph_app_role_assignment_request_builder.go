// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\ownedObjects\microsoft.graph.appRoleAssignment
type ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters get the items of type microsoft.graph.appRoleAssignment in the microsoft.graph.directoryObject collection
type ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters struct {
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
// NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderInternal instantiates a new MicrosoftGraphAppRoleAssignmentRequestBuilder and sets the default values.
func NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    m := &ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects/microsoft.graph.appRoleAssignment{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder instantiates a new MicrosoftGraphAppRoleAssignmentRequestBuilder and sets the default values.
func NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) Count()(*ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentCountRequestBuilder) {
    return NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.appRoleAssignment in the microsoft.graph.directoryObject collection
func (m *ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AppRoleAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateAppRoleAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.AppRoleAssignmentCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.appRoleAssignment in the microsoft.graph.directoryObject collection
func (m *ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) WithUrl(rawUrl string)(*ItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder) {
    return NewItemOwnedObjectsMicrosoftGraphAppRoleAssignmentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
