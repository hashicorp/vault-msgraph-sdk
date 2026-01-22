// Copyright IBM Corp. 2023, 2025
// SPDX-License-Identifier: MIT

package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder builds and executes requests for operations under \applications\microsoft.graph.getAvailableExtensionProperties
type MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) {
    m := &MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/microsoft.graph.getAvailableExtensionProperties", pathParameters),
    }
    return m
}
// NewMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder instantiates a new MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body MicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(MicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostResponseable), nil
}
// ToPostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
func (m *MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftGraphGetAvailableExtensionPropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) WithUrl(rawUrl string)(*MicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder) {
    return NewMicrosoftGraphGetAvailableExtensionPropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
