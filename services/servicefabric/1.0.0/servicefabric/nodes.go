package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NodesClient is the client for the Nodes methods of the Servicefabric service.
type NodesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewNodesClient creates an instance of the NodesClient client.
func NewNodesClient(timeout *int32) NodesClient {
	return NewNodesClientWithBaseURI(DefaultBaseURI, timeout)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewNodesClientWithBaseURI creates an instance of the NodesClient client.
func NewNodesClientWithBaseURI(baseURI string, timeout *int32) NodesClient {
	return NodesClient{NewWithBaseURI(baseURI, timeout)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Disable disable nodes
// Parameters:
// nodeName - the name of the node
// disableNode - the node
func (client NodesClient) Disable(ctx context.Context, nodeName string, disableNode DisableNode) (result String, err error) {
	req, err := client.DisablePreparer(ctx, nodeName, disableNode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Disable", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Disable", resp, "Failure sending request")
		return
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// DisablePreparer prepares the Disable request.
func (client NodesClient) DisablePreparer(ctx context.Context, nodeName string, disableNode DisableNode) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName": autorest.Encode("path", nodeName),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Nodes/{nodeName}/$/Deactivate", pathParameters),
		autorest.WithJSON(disableNode),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client NodesClient) DisableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client NodesClient) DisableResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Enable enable nodes
// Parameters:
// nodeName - the name of the node
func (client NodesClient) Enable(ctx context.Context, nodeName string) (result String, err error) {
	req, err := client.EnablePreparer(ctx, nodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// EnablePreparer prepares the Enable request.
func (client NodesClient) EnablePreparer(ctx context.Context, nodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName": autorest.Encode("path", nodeName),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Nodes/{nodeName}/$/Activate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client NodesClient) EnableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client NodesClient) EnableResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Get get nodes
// Parameters:
// nodeName - the name of the node
func (client NodesClient) Get(ctx context.Context, nodeName string) (result Node, err error) {
	req, err := client.GetPreparer(ctx, nodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// GetPreparer prepares the Get request.
func (client NodesClient) GetPreparer(ctx context.Context, nodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName": autorest.Encode("path", nodeName),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Nodes/{nodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NodesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NodesClient) GetResponder(resp *http.Response) (result Node, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// List list nodes
// Parameters:
// continuationToken - the token of the continuation
func (client NodesClient) List(ctx context.Context, continuationToken string) (result NodeList, err error) {
	req, err := client.ListPreparer(ctx, continuationToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.NodesClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// ListPreparer prepares the List request.
func (client NodesClient) ListPreparer(ctx context.Context, continuationToken string) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}
	if len(continuationToken) > 0 {
		queryParameters["continuation-token"] = autorest.Encode("query", continuationToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/Nodes"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NodesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NodesClient) ListResponder(resp *http.Response) (result NodeList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
