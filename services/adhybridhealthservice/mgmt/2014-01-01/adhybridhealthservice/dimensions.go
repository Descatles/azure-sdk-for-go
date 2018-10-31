package adhybridhealthservice

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

// DimensionsClient is the REST APIs for Azure Active Drectory Connect Health
type DimensionsClient struct {
	BaseClient
}

// NewDimensionsClient creates an instance of the DimensionsClient client.
func NewDimensionsClient() DimensionsClient {
	return NewDimensionsClientWithBaseURI(DefaultBaseURI)
}

// NewDimensionsClientWithBaseURI creates an instance of the DimensionsClient client.
func NewDimensionsClientWithBaseURI(baseURI string) DimensionsClient {
	return DimensionsClient{NewWithBaseURI(baseURI)}
}

// ListAddsDimensions gets the dimensions for a given dimension type in a server.
// Parameters:
// serviceName - the name of the service.
// dimension - the dimension type.
func (client DimensionsClient) ListAddsDimensions(ctx context.Context, serviceName string, dimension string) (result DimensionsPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListAddsDimensions")
		defer func() {
			sc := -1
			if result.d.Response.Response != nil {
				sc = result.d.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAddsDimensionsNextResults
	req, err := client.ListAddsDimensionsPreparer(ctx, serviceName, dimension)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.DimensionsClient", "ListAddsDimensions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAddsDimensionsSender(req)
	if err != nil {
		result.d.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.DimensionsClient", "ListAddsDimensions", resp, "Failure sending request")
		return
	}

	result.d, err = client.ListAddsDimensionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.DimensionsClient", "ListAddsDimensions", resp, "Failure responding to request")
	}

	return
}

// ListAddsDimensionsPreparer prepares the ListAddsDimensions request.
func (client DimensionsClient) ListAddsDimensionsPreparer(ctx context.Context, serviceName string, dimension string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dimension":   autorest.Encode("path", dimension),
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/dimensions/{dimension}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAddsDimensionsSender sends the ListAddsDimensions request. The method will close the
// http.Response Body if it receives an error.
func (client DimensionsClient) ListAddsDimensionsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListAddsDimensionsResponder handles the response to the ListAddsDimensions request. The method always
// closes the http.Response Body.
func (client DimensionsClient) ListAddsDimensionsResponder(resp *http.Response) (result Dimensions, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAddsDimensionsNextResults retrieves the next set of results, if any.
func (client DimensionsClient) listAddsDimensionsNextResults(ctx context.Context, lastResults Dimensions) (result Dimensions, err error) {
	req, err := lastResults.dimensionsPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.DimensionsClient", "listAddsDimensionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAddsDimensionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.DimensionsClient", "listAddsDimensionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAddsDimensionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.DimensionsClient", "listAddsDimensionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAddsDimensionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client DimensionsClient) ListAddsDimensionsComplete(ctx context.Context, serviceName string, dimension string) (result DimensionsIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DimensionsClient.ListAddsDimensions")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAddsDimensions(ctx, serviceName, dimension)
	return
}
