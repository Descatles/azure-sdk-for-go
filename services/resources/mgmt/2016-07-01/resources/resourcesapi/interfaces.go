package resourcesapi

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
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-07-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

// DeploymentsClientAPI contains the set of methods on the DeploymentsClient type.
type DeploymentsClientAPI interface {
	CalculateTemplateHash(ctx context.Context, templateParameter interface{}) (result resources.TemplateHashResult, err error)
	Cancel(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error)
	CheckExistence(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, parameters resources.Deployment) (result resources.DeploymentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, deploymentName string) (result resources.DeploymentsDeleteFuture, err error)
	ExportTemplate(ctx context.Context, resourceGroupName string, deploymentName string) (result resources.DeploymentExportResult, err error)
	Get(ctx context.Context, resourceGroupName string, deploymentName string) (result resources.DeploymentExtended, err error)
	List(ctx context.Context, resourceGroupName string, filter string, top *int32) (result resources.DeploymentListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result resources.DeploymentListResultIterator, err error)
	Validate(ctx context.Context, resourceGroupName string, deploymentName string, parameters resources.Deployment) (result resources.DeploymentValidateResult, err error)
}

var _ DeploymentsClientAPI = (*resources.DeploymentsClient)(nil)

// ProvidersClientAPI contains the set of methods on the ProvidersClient type.
type ProvidersClientAPI interface {
	Get(ctx context.Context, resourceProviderNamespace string, expand string) (result resources.Provider, err error)
	List(ctx context.Context, top *int32, expand string) (result resources.ProviderListResultPage, err error)
	ListComplete(ctx context.Context, top *int32, expand string) (result resources.ProviderListResultIterator, err error)
	Register(ctx context.Context, resourceProviderNamespace string) (result resources.Provider, err error)
	Unregister(ctx context.Context, resourceProviderNamespace string) (result resources.Provider, err error)
}

var _ ProvidersClientAPI = (*resources.ProvidersClient)(nil)

// GroupsClientAPI contains the set of methods on the GroupsClient type.
type GroupsClientAPI interface {
	CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters resources.Group) (result resources.Group, err error)
	Delete(ctx context.Context, resourceGroupName string) (result resources.GroupsDeleteFuture, err error)
	ExportTemplate(ctx context.Context, resourceGroupName string, parameters resources.ExportTemplateRequest) (result resources.GroupExportResult, err error)
	Get(ctx context.Context, resourceGroupName string) (result resources.Group, err error)
	List(ctx context.Context, filter string, top *int32) (result resources.GroupListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32) (result resources.GroupListResultIterator, err error)
	ListResources(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
	ListResourcesComplete(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result resources.ListResultIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, parameters resources.Group) (result resources.Group, err error)
}

var _ GroupsClientAPI = (*resources.GroupsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckExistence(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, parameters resources.GenericResource) (result resources.GenericResource, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string) (result resources.GenericResource, err error)
	List(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
	ListComplete(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultIterator, err error)
	MoveResources(ctx context.Context, sourceResourceGroupName string, parameters resources.MoveInfo) (result resources.MoveResourcesFuture, err error)
	Update(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, parameters resources.GenericResource) (result resources.UpdateFuture, err error)
}

var _ ClientAPI = (*resources.Client)(nil)

// TagsClientAPI contains the set of methods on the TagsClient type.
type TagsClientAPI interface {
	CreateOrUpdate(ctx context.Context, tagName string) (result resources.TagDetails, err error)
	CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string) (result resources.TagValue, err error)
	Delete(ctx context.Context, tagName string) (result autorest.Response, err error)
	DeleteValue(ctx context.Context, tagName string, tagValue string) (result autorest.Response, err error)
	List(ctx context.Context) (result resources.TagsListResultPage, err error)
	ListComplete(ctx context.Context) (result resources.TagsListResultIterator, err error)
}

var _ TagsClientAPI = (*resources.TagsClient)(nil)

// DeploymentOperationsClientAPI contains the set of methods on the DeploymentOperationsClient type.
type DeploymentOperationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, deploymentName string, operationID string) (result resources.DeploymentOperation, err error)
	List(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result resources.DeploymentOperationsListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result resources.DeploymentOperationsListResultIterator, err error)
}

var _ DeploymentOperationsClientAPI = (*resources.DeploymentOperationsClient)(nil)
