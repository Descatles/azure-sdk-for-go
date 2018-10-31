package cognitiveservices

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AccountsClient is the cognitive Services Management Client
type AccountsClient struct {
	BaseClient
}

// NewAccountsClient creates an instance of the AccountsClient client.
func NewAccountsClient(subscriptionID string) AccountsClient {
	return NewAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccountsClientWithBaseURI creates an instance of the AccountsClient client.
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return AccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for
// developer to access intelligent APIs. It's also the resource type for billing.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
// parameters - the parameters to provide for the created account.
func (client AccountsClient) Create(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters) (result Account, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Sku", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AccountsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AccountsClient) CreateResponder(resp *http.Response) (result Account, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Cognitive Services account from the resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
func (client AccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AccountsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AccountsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetProperties returns a Cognitive Services account specified by the parameters.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
func (client AccountsClient) GetProperties(ctx context.Context, resourceGroupName string, accountName string) (result Account, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.GetProperties")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "GetProperties", err.Error())
	}

	req, err := client.GetPropertiesPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetProperties", resp, "Failure sending request")
		return
	}

	result, err = client.GetPropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetProperties", resp, "Failure responding to request")
	}

	return
}

// GetPropertiesPreparer prepares the GetProperties request.
func (client AccountsClient) GetPropertiesPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPropertiesSender sends the GetProperties request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) GetPropertiesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetPropertiesResponder handles the response to the GetProperties request. The method always
// closes the http.Response Body.
func (client AccountsClient) GetPropertiesResponder(resp *http.Response) (result Account, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUsages get usages for the requested Cognitive Services account
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
// filter - an OData filter expression that describes a subset of usages to return. The supported parameter is
// name.value (name of the metric, can have an or of multiple names).
func (client AccountsClient) GetUsages(ctx context.Context, resourceGroupName string, accountName string, filter string) (result UsagesResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.GetUsages")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "GetUsages", err.Error())
	}

	req, err := client.GetUsagesPreparer(ctx, resourceGroupName, accountName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetUsages", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUsagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetUsages", resp, "Failure sending request")
		return
	}

	result, err = client.GetUsagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "GetUsages", resp, "Failure responding to request")
	}

	return
}

// GetUsagesPreparer prepares the GetUsages request.
func (client AccountsClient) GetUsagesPreparer(ctx context.Context, resourceGroupName string, accountName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUsagesSender sends the GetUsages request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) GetUsagesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetUsagesResponder handles the response to the GetUsages request. The method always
// closes the http.Response Body.
func (client AccountsClient) GetUsagesResponder(resp *http.Response) (result UsagesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all the resources of a particular type belonging to a subscription.
func (client AccountsClient) List(ctx context.Context) (result AccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.List")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "List", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AccountsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListResponder(resp *http.Response) (result AccountListResult, err error) {
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
func (client AccountsClient) listNextResults(ctx context.Context, lastResults AccountListResult) (result AccountListResult, err error) {
	req, err := lastResults.accountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountsClient) ListComplete(ctx context.Context) (result AccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup returns all the resources of a particular type belonging to a resource group
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
func (client AccountsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result AccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AccountsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListByResourceGroupResponder(resp *http.Response) (result AccountListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AccountsClient) listByResourceGroupNextResults(ctx context.Context, lastResults AccountListResult) (result AccountListResult, err error) {
	req, err := lastResults.accountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result AccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListKeys lists the account keys for the specified Cognitive Services account.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
func (client AccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result AccountKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.ListKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "ListKeys", err.Error())
	}

	req, err := client.ListKeysPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client AccountsClient) ListKeysPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListKeysResponder(resp *http.Response) (result AccountKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSkus list available SKUs for the requested Cognitive Services account
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
func (client AccountsClient) ListSkus(ctx context.Context, resourceGroupName string, accountName string) (result AccountEnumerateSkusResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.ListSkus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "ListSkus", err.Error())
	}

	req, err := client.ListSkusPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListSkus", resp, "Failure sending request")
		return
	}

	result, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "ListSkus", resp, "Failure responding to request")
	}

	return
}

// ListSkusPreparer prepares the ListSkus request.
func (client AccountsClient) ListSkusPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSkusSender sends the ListSkus request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) ListSkusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client AccountsClient) ListSkusResponder(resp *http.Response) (result AccountEnumerateSkusResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateKey regenerates the specified account key for the specified Cognitive Services account.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
// parameters - regenerate key parameters.
func (client AccountsClient) RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, parameters RegenerateKeyParameters) (result AccountKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.RegenerateKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "RegenerateKey", err.Error())
	}

	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "RegenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "RegenerateKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client AccountsClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, accountName string, parameters RegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/regenerateKey", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client AccountsClient) RegenerateKeyResponder(resp *http.Response) (result AccountKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a Cognitive Services account
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription.
// accountName - the name of Cognitive Services account.
// parameters - the parameters to provide for the created account.
func (client AccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, parameters AccountUpdateParameters) (result Account, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.AccountsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.AccountsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AccountsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters AccountUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AccountsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AccountsClient) UpdateResponder(resp *http.Response) (result Account, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
