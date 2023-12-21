package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\transitiveMemberOf\microsoft.graph.directoryRole\$count
type ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderGetQueryParameters get the number of the resource
type ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// NewItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) {
    m := &ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/transitiveMemberOf/microsoft.graph.directoryRole/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderGetQueryParameters])(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) WithUrl(rawUrl string)(*ItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) {
    return NewItemTransitiveMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
