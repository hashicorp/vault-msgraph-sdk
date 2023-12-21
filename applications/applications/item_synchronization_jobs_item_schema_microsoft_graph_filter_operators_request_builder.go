package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder builds and executes requests for operations under \applications\{application-id}\synchronization\jobs\{synchronizationJob-id}\schema\microsoft.graph.filterOperators()
type ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderGetQueryParameters invoke function filterOperators
type ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderGetQueryParameters struct {
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
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderInternal instantiates a new MicrosoftGraphFilterOperatorsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/microsoft.graph.filterOperators(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder instantiates a new MicrosoftGraphFilterOperatorsRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function filterOperators
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderGetQueryParameters])(ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsFilterOperatorsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsFilterOperatorsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsFilterOperatorsGetResponseable), nil
}
// ToGetRequestInformation invoke function filterOperators
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder) {
    return NewItemSynchronizationJobsItemSchemaMicrosoftGraphFilterOperatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
