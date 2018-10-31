package storsimple

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

// BackupPoliciesClient is the client for the BackupPolicies methods of the Storsimple service.
type BackupPoliciesClient struct {
	BaseClient
}

// NewBackupPoliciesClient creates an instance of the BackupPoliciesClient client.
func NewBackupPoliciesClient(subscriptionID string) BackupPoliciesClient {
	return NewBackupPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupPoliciesClientWithBaseURI creates an instance of the BackupPoliciesClient client.
func NewBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BackupPoliciesClient {
	return BackupPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// BackupNow backup the backup policy now.
// Parameters:
// deviceName - the device name
// backupPolicyName - the backup policy name.
// backupType - the backup Type. This can be cloudSnapshot or localSnapshot.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupPoliciesClient) BackupNow(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string) (result BackupPoliciesBackupNowFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupPoliciesClient.BackupNow")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupPoliciesClient", "BackupNow", err.Error())
	}

	req, err := client.BackupNowPreparer(ctx, deviceName, backupPolicyName, backupType, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "BackupNow", nil, "Failure preparing request")
		return
	}

	result, err = client.BackupNowSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "BackupNow", result.Response(), "Failure sending request")
		return
	}

	return
}

// BackupNowPreparer prepares the BackupNow request.
func (client BackupPoliciesClient) BackupNowPreparer(ctx context.Context, deviceName string, backupPolicyName string, backupType string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":  backupPolicyName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"backupType":  backupType,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/backup", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BackupNowSender sends the BackupNow request. The method will close the
// http.Response Body if it receives an error.
func (client BackupPoliciesClient) BackupNowSender(req *http.Request) (future BackupPoliciesBackupNowFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// BackupNowResponder handles the response to the BackupNow request. The method always
// closes the http.Response Body.
func (client BackupPoliciesClient) BackupNowResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate creates or updates the backup policy.
// Parameters:
// deviceName - the device name
// backupPolicyName - the name of the backup policy to be created/updated.
// parameters - the backup policy.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupPoliciesClient) CreateOrUpdate(ctx context.Context, deviceName string, backupPolicyName string, parameters BackupPolicy, resourceGroupName string, managerName string) (result BackupPoliciesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BackupPolicyProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.BackupPolicyProperties.VolumeIds", Name: validation.Null, Rule: true, Chain: nil}}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, backupPolicyName, parameters, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, backupPolicyName string, parameters BackupPolicy, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":  backupPolicyName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupPoliciesClient) CreateOrUpdateSender(req *http.Request) (future BackupPoliciesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result BackupPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the backup policy.
// Parameters:
// deviceName - the device name
// backupPolicyName - the name of the backup policy.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupPoliciesClient) Delete(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string) (result BackupPoliciesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupPoliciesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupPoliciesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, deviceName, backupPolicyName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BackupPoliciesClient) DeletePreparer(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":  backupPolicyName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BackupPoliciesClient) DeleteSender(req *http.Request) (future BackupPoliciesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BackupPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified backup policy name.
// Parameters:
// deviceName - the device name
// backupPolicyName - the name of backup policy to be fetched.
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupPoliciesClient) Get(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string) (result BackupPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupPoliciesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, deviceName, backupPolicyName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupPoliciesClient) GetPreparer(ctx context.Context, deviceName string, backupPolicyName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupPolicyName":  backupPolicyName,
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupPoliciesClient) GetResponder(resp *http.Response) (result BackupPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDevice gets all the backup policies in a device.
// Parameters:
// deviceName - the device name
// resourceGroupName - the resource group name
// managerName - the manager name
func (client BackupPoliciesClient) ListByDevice(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (result BackupPolicyList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupPoliciesClient.ListByDevice")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storsimple.BackupPoliciesClient", "ListByDevice", err.Error())
	}

	req, err := client.ListByDevicePreparer(ctx, deviceName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "ListByDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "ListByDevice", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.BackupPoliciesClient", "ListByDevice", resp, "Failure responding to request")
	}

	return
}

// ListByDevicePreparer prepares the ListByDevice request.
func (client BackupPoliciesClient) ListByDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        deviceName,
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDeviceSender sends the ListByDevice request. The method will close the
// http.Response Body if it receives an error.
func (client BackupPoliciesClient) ListByDeviceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByDeviceResponder handles the response to the ListByDevice request. The method always
// closes the http.Response Body.
func (client BackupPoliciesClient) ListByDeviceResponder(resp *http.Response) (result BackupPolicyList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
