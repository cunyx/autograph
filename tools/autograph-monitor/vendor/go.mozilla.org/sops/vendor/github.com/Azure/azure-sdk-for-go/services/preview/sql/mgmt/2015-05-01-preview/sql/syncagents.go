package sql

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

// SyncAgentsClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type SyncAgentsClient struct {
	BaseClient
}

// NewSyncAgentsClient creates an instance of the SyncAgentsClient client.
func NewSyncAgentsClient(subscriptionID string) SyncAgentsClient {
	return NewSyncAgentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSyncAgentsClientWithBaseURI creates an instance of the SyncAgentsClient client.
func NewSyncAgentsClientWithBaseURI(baseURI string, subscriptionID string) SyncAgentsClient {
	return SyncAgentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a sync agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server on which the sync agent is hosted.
// syncAgentName - the name of the sync agent.
// parameters - the requested sync agent resource state.
func (client SyncAgentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent) (result SyncAgentsCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, syncAgentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SyncAgentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters SyncAgent) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) CreateOrUpdateSender(req *http.Request) (future SyncAgentsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) CreateOrUpdateResponder(resp *http.Response) (result SyncAgent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a sync agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server on which the sync agent is hosted.
// syncAgentName - the name of the sync agent.
func (client SyncAgentsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result SyncAgentsDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SyncAgentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) DeleteSender(req *http.Request) (future SyncAgentsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	err = autorest.Respond(resp, azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateKey generates a sync agent key.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server on which the sync agent is hosted.
// syncAgentName - the name of the sync agent.
func (client SyncAgentsClient) GenerateKey(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result SyncAgentKeyProperties, err error) {
	req, err := client.GenerateKeyPreparer(ctx, resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "GenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "GenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "GenerateKey", resp, "Failure responding to request")
	}

	return
}

// GenerateKeyPreparer prepares the GenerateKey request.
func (client SyncAgentsClient) GenerateKeyPreparer(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/generateKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateKeySender sends the GenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) GenerateKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GenerateKeyResponder handles the response to the GenerateKey request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) GenerateKeyResponder(resp *http.Response) (result SyncAgentKeyProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a sync agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server on which the sync agent is hosted.
// syncAgentName - the name of the sync agent.
func (client SyncAgentsClient) Get(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result SyncAgent, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SyncAgentsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) GetResponder(resp *http.Response) (result SyncAgent, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer lists sync agents in a server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server on which the sync agent is hosted.
func (client SyncAgentsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result SyncAgentListResultPage, err error) {
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.salr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.salr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client SyncAgentsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) ListByServerResponder(resp *http.Response) (result SyncAgentListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client SyncAgentsClient) listByServerNextResults(lastResults SyncAgentListResult) (result SyncAgentListResult, err error) {
	req, err := lastResults.syncAgentListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client SyncAgentsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result SyncAgentListResultIterator, err error) {
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}

// ListLinkedDatabases lists databases linked to a sync agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server on which the sync agent is hosted.
// syncAgentName - the name of the sync agent.
func (client SyncAgentsClient) ListLinkedDatabases(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result SyncAgentLinkedDatabaseListResultPage, err error) {
	result.fn = client.listLinkedDatabasesNextResults
	req, err := client.ListLinkedDatabasesPreparer(ctx, resourceGroupName, serverName, syncAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListLinkedDatabasesSender(req)
	if err != nil {
		result.saldlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", resp, "Failure sending request")
		return
	}

	result.saldlr, err = client.ListLinkedDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "ListLinkedDatabases", resp, "Failure responding to request")
	}

	return
}

// ListLinkedDatabasesPreparer prepares the ListLinkedDatabases request.
func (client SyncAgentsClient) ListLinkedDatabasesPreparer(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncAgentName":     autorest.Encode("path", syncAgentName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/syncAgents/{syncAgentName}/linkedDatabases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListLinkedDatabasesSender sends the ListLinkedDatabases request. The method will close the
// http.Response Body if it receives an error.
func (client SyncAgentsClient) ListLinkedDatabasesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListLinkedDatabasesResponder handles the response to the ListLinkedDatabases request. The method always
// closes the http.Response Body.
func (client SyncAgentsClient) ListLinkedDatabasesResponder(resp *http.Response) (result SyncAgentLinkedDatabaseListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listLinkedDatabasesNextResults retrieves the next set of results, if any.
func (client SyncAgentsClient) listLinkedDatabasesNextResults(lastResults SyncAgentLinkedDatabaseListResult) (result SyncAgentLinkedDatabaseListResult, err error) {
	req, err := lastResults.syncAgentLinkedDatabaseListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "listLinkedDatabasesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListLinkedDatabasesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "listLinkedDatabasesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListLinkedDatabasesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncAgentsClient", "listLinkedDatabasesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListLinkedDatabasesComplete enumerates all values, automatically crossing page boundaries as required.
func (client SyncAgentsClient) ListLinkedDatabasesComplete(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result SyncAgentLinkedDatabaseListResultIterator, err error) {
	result.page, err = client.ListLinkedDatabases(ctx, resourceGroupName, serverName, syncAgentName)
	return
}
