package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc "github.com/hashicorp/vault-msgraph-sdk/applications/models"
    i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887 "github.com/hashicorp/vault-msgraph-sdk/applications/models/odataerrors"
)

// ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\memberOf\microsoft.graph.directoryRole
type ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderGetQueryParameters get the items of type microsoft.graph.directoryRole in the microsoft.graph.directoryObject collection
type ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderGetQueryParameters struct {
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
// NewItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderInternal instantiates a new MicrosoftGraphDirectoryRoleRequestBuilder and sets the default values.
func NewItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) {
    m := &ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/memberOf/microsoft.graph.directoryRole{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder instantiates a new MicrosoftGraphDirectoryRoleRequestBuilder and sets the default values.
func NewItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) Count()(*ItemMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilder) {
    return NewItemMemberOfMicrosoftGraphDirectoryRoleCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.directoryRole in the microsoft.graph.directoryObject collection
func (m *ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderGetQueryParameters])(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.DirectoryRoleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
        "5XX": i343a7f1f2f3bfa96ea6adc2d6a090ea976eacea28fbba19ae939f63dd5132887.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i39f8001d2b9b38b5cf71c15d33bad53f1b997f356b85cb9fd9fee820940d88bc.DirectoryRoleCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.directoryRole in the microsoft.graph.directoryObject collection
func (m *ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) WithUrl(rawUrl string)(*ItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder) {
    return NewItemMemberOfMicrosoftGraphDirectoryRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
