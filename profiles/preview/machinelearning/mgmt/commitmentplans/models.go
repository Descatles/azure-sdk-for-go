// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package commitmentplans

import original "github.com/Azure/azure-sdk-for-go/services/machinelearning/mgmt/2016-05-01-preview/commitmentplans"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type CommitmentAssociationsClient = original.CommitmentAssociationsClient

func NewCommitmentAssociationsClient(subscriptionID string) CommitmentAssociationsClient {
	return original.NewCommitmentAssociationsClient(subscriptionID)
}
func NewCommitmentAssociationsClientWithBaseURI(baseURI string, subscriptionID string) CommitmentAssociationsClient {
	return original.NewCommitmentAssociationsClientWithBaseURI(baseURI, subscriptionID)
}

type Client = original.Client

func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}

type ResourceSkuRestrictionsReasonCode = original.ResourceSkuRestrictionsReasonCode

const (
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = original.NotAvailableForSubscription
	QuotaID                     ResourceSkuRestrictionsReasonCode = original.QuotaID
)

type ResourceSkuRestrictionsType = original.ResourceSkuRestrictionsType

const (
	Location ResourceSkuRestrictionsType = original.Location
	Zone     ResourceSkuRestrictionsType = original.Zone
)

type SkuCapacityScaleType = original.SkuCapacityScaleType

const (
	Automatic SkuCapacityScaleType = original.Automatic
	Manual    SkuCapacityScaleType = original.Manual
	None      SkuCapacityScaleType = original.None
)

type CatalogSku = original.CatalogSku
type CommitmentAssociation = original.CommitmentAssociation
type CommitmentAssociationListResult = original.CommitmentAssociationListResult
type CommitmentAssociationListResultIterator = original.CommitmentAssociationListResultIterator
type CommitmentAssociationListResultPage = original.CommitmentAssociationListResultPage
type CommitmentAssociationProperties = original.CommitmentAssociationProperties
type CommitmentPlan = original.CommitmentPlan
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type MoveCommitmentAssociationRequest = original.MoveCommitmentAssociationRequest
type PatchPayload = original.PatchPayload
type PlanQuantity = original.PlanQuantity
type PlanUsageHistory = original.PlanUsageHistory
type PlanUsageHistoryListResult = original.PlanUsageHistoryListResult
type PlanUsageHistoryListResultIterator = original.PlanUsageHistoryListResultIterator
type PlanUsageHistoryListResultPage = original.PlanUsageHistoryListResultPage
type Properties = original.Properties
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type SkuCapability = original.SkuCapability
type SkuCapacity = original.SkuCapacity
type SkuCost = original.SkuCost
type SkuListResult = original.SkuListResult
type SkuRestrictions = original.SkuRestrictions
type SkusClient = original.SkusClient

func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}

type UsageHistoryClient = original.UsageHistoryClient

func NewUsageHistoryClient(subscriptionID string) UsageHistoryClient {
	return original.NewUsageHistoryClient(subscriptionID)
}
func NewUsageHistoryClientWithBaseURI(baseURI string, subscriptionID string) UsageHistoryClient {
	return original.NewUsageHistoryClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
