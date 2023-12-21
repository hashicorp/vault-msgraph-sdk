package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder builds and executes requests for operations under \applications\{application-id}\synchronization\jobs\{synchronizationJob-id}\schema\microsoft.graph.functions()
type ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderGetQueryParameters invoke function functions
type ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
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
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderInternal instantiates a new MicrosoftGraphFunctionsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/microsoft.graph.functions(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder instantiates a new MicrosoftGraphFunctionsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function functions
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderGetQueryParameters])(ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsFunctionsGetResponseable), nil
}
// ToGetRequestInformation invoke function functions
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder) {
    return NewItemSynchronizationJobsItemSchemaMicrosoftGraphFunctionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
