package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BitlockerRecoveryKeysCountRequestBuilder provides operations to count the resources in the collection.
type BitlockerRecoveryKeysCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BitlockerRecoveryKeysCountRequestBuilderGetQueryParameters get the number of the resource
type BitlockerRecoveryKeysCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// BitlockerRecoveryKeysCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRecoveryKeysCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRecoveryKeysCountRequestBuilderGetQueryParameters
}
// NewBitlockerRecoveryKeysCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysCountRequestBuilder) {
    m := &BitlockerRecoveryKeysCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/bitlocker/recoveryKeys/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewBitlockerRecoveryKeysCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRecoveryKeysCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *BitlockerRecoveryKeysCountRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *BitlockerRecoveryKeysCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *BitlockerRecoveryKeysCountRequestBuilder) WithUrl(rawUrl string)(*BitlockerRecoveryKeysCountRequestBuilder) {
    return NewBitlockerRecoveryKeysCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
