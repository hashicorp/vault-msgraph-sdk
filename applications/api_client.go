// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MIT

package applications

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i00582015c38e43f5aeb10599c7983ce4615ab83fc9d40ad567423cf75d3754d6 "github.com/hashicorp/vault-msgraph-sdk/applications/users"
    i1003f028e186a911c00edcd68f52da9f5615dcc2f3738600b952acf07f5d7be6 "github.com/hashicorp/vault-msgraph-sdk/applications/applications"
    i153e248007b547ca7451cce6e7f8fdec3f563406944c4940aa5b04e710486dd6 "github.com/hashicorp/vault-msgraph-sdk/applications/applicationtemplates"
    i1f07a7dd921a18fcc1d60220c1a8f0156b262f654e877e04cb974ffda4f1787a "github.com/hashicorp/vault-msgraph-sdk/applications/applicationswithappid"
    i36fadfe2c5b17d018fbcdbaffaee9e0c41e72c75927d494e370cec45c8807e7b "github.com/hashicorp/vault-msgraph-sdk/applications/serviceprincipals"
    i3e87c1ecf33f3dee7e1ff01cba6fd29d752ec8fa1e7d30399148552864b0676b "github.com/hashicorp/vault-msgraph-sdk/applications/groups"
    i4454e7993473786929b4859d03bce64bf2bea814172983d098f065330c205e91 "github.com/hashicorp/vault-msgraph-sdk/applications/serviceprincipalswithappid"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Applications the applications property
func (m *ApiClient) Applications()(*i1003f028e186a911c00edcd68f52da9f5615dcc2f3738600b952acf07f5d7be6.ApplicationsRequestBuilder) {
    return i1003f028e186a911c00edcd68f52da9f5615dcc2f3738600b952acf07f5d7be6.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplicationsWithAppId builds and executes requests for operations under \applications(appId='{appId}')
func (m *ApiClient) ApplicationsWithAppId(appId *string)(*i1f07a7dd921a18fcc1d60220c1a8f0156b262f654e877e04cb974ffda4f1787a.ApplicationsWithAppIdRequestBuilder) {
    return i1f07a7dd921a18fcc1d60220c1a8f0156b262f654e877e04cb974ffda4f1787a.NewApplicationsWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// ApplicationTemplates the applicationTemplates property
func (m *ApiClient) ApplicationTemplates()(*i153e248007b547ca7451cce6e7f8fdec3f563406944c4940aa5b04e710486dd6.ApplicationTemplatesRequestBuilder) {
    return i153e248007b547ca7451cce6e7f8fdec3f563406944c4940aa5b04e710486dd6.NewApplicationTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Groups the groups property
func (m *ApiClient) Groups()(*i3e87c1ecf33f3dee7e1ff01cba6fd29d752ec8fa1e7d30399148552864b0676b.GroupsRequestBuilder) {
    return i3e87c1ecf33f3dee7e1ff01cba6fd29d752ec8fa1e7d30399148552864b0676b.NewGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePrincipals the servicePrincipals property
func (m *ApiClient) ServicePrincipals()(*i36fadfe2c5b17d018fbcdbaffaee9e0c41e72c75927d494e370cec45c8807e7b.ServicePrincipalsRequestBuilder) {
    return i36fadfe2c5b17d018fbcdbaffaee9e0c41e72c75927d494e370cec45c8807e7b.NewServicePrincipalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePrincipalsWithAppId builds and executes requests for operations under \servicePrincipals(appId='{appId}')
func (m *ApiClient) ServicePrincipalsWithAppId(appId *string)(*i4454e7993473786929b4859d03bce64bf2bea814172983d098f065330c205e91.ServicePrincipalsWithAppIdRequestBuilder) {
    return i4454e7993473786929b4859d03bce64bf2bea814172983d098f065330c205e91.NewServicePrincipalsWithAppIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, appId)
}
// Users the users property
func (m *ApiClient) Users()(*i00582015c38e43f5aeb10599c7983ce4615ab83fc9d40ad567423cf75d3754d6.UsersRequestBuilder) {
    return i00582015c38e43f5aeb10599c7983ce4615ab83fc9d40ad567423cf75d3754d6.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
