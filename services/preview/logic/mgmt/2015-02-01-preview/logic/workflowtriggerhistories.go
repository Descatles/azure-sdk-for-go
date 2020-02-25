package logic

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WorkflowTriggerHistoriesClient is the REST API for Azure Logic Apps.
type WorkflowTriggerHistoriesClient struct {
	BaseClient
}

// NewWorkflowTriggerHistoriesClient creates an instance of the WorkflowTriggerHistoriesClient client.
func NewWorkflowTriggerHistoriesClient(subscriptionID string) WorkflowTriggerHistoriesClient {
	return NewWorkflowTriggerHistoriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowTriggerHistoriesClientWithBaseURI creates an instance of the WorkflowTriggerHistoriesClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewWorkflowTriggerHistoriesClientWithBaseURI(baseURI string, subscriptionID string) WorkflowTriggerHistoriesClient {
	return WorkflowTriggerHistoriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow trigger history.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
// historyName - the workflow trigger history name.
func (client WorkflowTriggerHistoriesClient) Get(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string) (result WorkflowTriggerHistory, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggerHistoriesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, triggerName, historyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowTriggerHistoriesClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, historyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"historyName":       autorest.Encode("path", historyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggerHistoriesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowTriggerHistoriesClient) GetResponder(resp *http.Response) (result WorkflowTriggerHistory, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow trigger histories.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// triggerName - the workflow trigger name.
// top - the number of items to be included in the result.
// filter - the filter to apply on the operation.
func (client WorkflowTriggerHistoriesClient) List(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, top *int32, filter string) (result WorkflowTriggerHistoryListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggerHistoriesClient.List")
		defer func() {
			sc := -1
			if result.wthlr.Response.Response != nil {
				sc = result.wthlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workflowName, triggerName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wthlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "List", resp, "Failure sending request")
		return
	}

	result.wthlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowTriggerHistoriesClient) ListPreparer(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"triggerName":       autorest.Encode("path", triggerName),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowTriggerHistoriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowTriggerHistoriesClient) ListResponder(resp *http.Response) (result WorkflowTriggerHistoryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkflowTriggerHistoriesClient) listNextResults(ctx context.Context, lastResults WorkflowTriggerHistoryListResult) (result WorkflowTriggerHistoryListResult, err error) {
	req, err := lastResults.workflowTriggerHistoryListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowTriggerHistoriesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowTriggerHistoriesClient) ListComplete(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, top *int32, filter string) (result WorkflowTriggerHistoryListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowTriggerHistoriesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workflowName, triggerName, top, filter)
	return
}
