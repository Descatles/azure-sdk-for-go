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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// VersionClient is the azure Service Fabric Resource Provider API Client
type VersionClient struct {
	BaseClient
}

// NewVersionClient creates an instance of the VersionClient client.
func NewVersionClient() VersionClient {
	return NewVersionClientWithBaseURI(DefaultBaseURI)
}

// NewVersionClientWithBaseURI creates an instance of the VersionClient client.
func NewVersionClientWithBaseURI(baseURI string) VersionClient {
	return VersionClient{NewWithBaseURI(baseURI)}
}

// Delete unprovisions an application type version resource.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationTypeName is the name of the application type name
// resource version is the application type version.
func (client VersionClient) Delete(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result VersionDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VersionClient) DeletePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
		"version":             autorest.Encode("path", version),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VersionClient) DeleteSender(req *http.Request) (future VersionDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VersionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns an application type version resource.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationTypeName is the name of the application type name
// resource version is the application type version.
func (client VersionClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, version string) (result VersionResource, err error) {
	req, err := client.GetPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VersionClient) GetPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
		"version":             autorest.Encode("path", version),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VersionClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VersionClient) GetResponder(resp *http.Response) (result VersionResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all versions for the specified application type.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationTypeName is the name of the application type name
// resource
func (client VersionClient) List(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string) (result VersionResourceList, err error) {
	req, err := client.ListPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VersionClient) ListPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VersionClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VersionClient) ListResponder(resp *http.Response) (result VersionResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put provisions an application type version resource.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationTypeName is the name of the application type name
// resource version is the application type version. parameters is the application type version resource.
func (client VersionClient) Put(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters VersionResource) (result VersionPutFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VersionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.VersionProperties.AppPackageURL", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("servicefabric.VersionClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationTypeName, version, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Put", nil, "Failure preparing request")
		return
	}

	result, err = client.PutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.VersionClient", "Put", result.Response(), "Failure sending request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client VersionClient) PutPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters VersionResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationTypeName": autorest.Encode("path", applicationTypeName),
		"clusterName":         autorest.Encode("path", clusterName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
		"version":             autorest.Encode("path", version),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client VersionClient) PutSender(req *http.Request) (future VersionPutFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client VersionClient) PutResponder(resp *http.Response) (result VersionResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
