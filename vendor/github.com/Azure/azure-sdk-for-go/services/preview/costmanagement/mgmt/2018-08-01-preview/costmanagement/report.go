package costmanagement

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

// ReportClient is the client for the Report methods of the Costmanagement service.
type ReportClient struct {
	BaseClient
}

// NewReportClient creates an instance of the ReportClient client.
func NewReportClient(subscriptionID string) ReportClient {
	return NewReportClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReportClientWithBaseURI creates an instance of the ReportClient client.
func NewReportClientWithBaseURI(baseURI string, subscriptionID string) ReportClient {
	return ReportClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update a report. Update operation requires latest eTag to be set in the
// request mandatorily. You may obtain the latest eTag by performing a get operation. Create operation does not require
// eTag.
// Parameters:
// reportName - report Name.
// parameters - parameters supplied to the CreateOrUpdate Report operation.
func (client ReportClient) CreateOrUpdate(ctx context.Context, reportName string, parameters Report) (result Report, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ReportProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ReportProperties.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ReportProperties.Schedule.RecurrencePeriod", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ReportProperties.Schedule.RecurrencePeriod.From", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.ReportClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, reportName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ReportClient) CreateOrUpdatePreparer(ctx context.Context, reportName string, parameters Report) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":     autorest.Encode("path", reportName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reports/{reportName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ReportClient) CreateOrUpdateResponder(resp *http.Response) (result Report, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByResourceGroupName the operation to create or update a report. Update operation requires latest eTag
// to be set in the request mandatorily. You may obtain the latest eTag by performing a get operation. Create operation
// does not require eTag.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportName - report Name.
// parameters - parameters supplied to the CreateOrUpdate Report operation.
func (client ReportClient) CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string, parameters Report) (result Report, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ReportProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ReportProperties.Schedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ReportProperties.Schedule.RecurrencePeriod", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ReportProperties.Schedule.RecurrencePeriod.From", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("costmanagement.ReportClient", "CreateOrUpdateByResourceGroupName", err.Error())
	}

	req, err := client.CreateOrUpdateByResourceGroupNamePreparer(ctx, resourceGroupName, reportName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "CreateOrUpdateByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "CreateOrUpdateByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "CreateOrUpdateByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateByResourceGroupNamePreparer prepares the CreateOrUpdateByResourceGroupName request.
func (client ReportClient) CreateOrUpdateByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportName string, parameters Report) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":        autorest.Encode("path", reportName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reports/{reportName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateByResourceGroupNameSender sends the CreateOrUpdateByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) CreateOrUpdateByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateByResourceGroupNameResponder handles the response to the CreateOrUpdateByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportClient) CreateOrUpdateByResourceGroupNameResponder(resp *http.Response) (result Report, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a report.
// Parameters:
// reportName - report Name.
func (client ReportClient) Delete(ctx context.Context, reportName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReportClient) DeletePreparer(ctx context.Context, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":     autorest.Encode("path", reportName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reports/{reportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReportClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByResourceGroupName the operation to delete a report.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportName - report Name.
func (client ReportClient) DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result autorest.Response, err error) {
	req, err := client.DeleteByResourceGroupNamePreparer(ctx, resourceGroupName, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "DeleteByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByResourceGroupNameSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "DeleteByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "DeleteByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// DeleteByResourceGroupNamePreparer prepares the DeleteByResourceGroupName request.
func (client ReportClient) DeleteByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":        autorest.Encode("path", reportName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reports/{reportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByResourceGroupNameSender sends the DeleteByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) DeleteByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteByResourceGroupNameResponder handles the response to the DeleteByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportClient) DeleteByResourceGroupNameResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Execute the operation to execute a report.
// Parameters:
// reportName - report Name.
func (client ReportClient) Execute(ctx context.Context, reportName string) (result autorest.Response, err error) {
	req, err := client.ExecutePreparer(ctx, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Execute", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Execute", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Execute", resp, "Failure responding to request")
	}

	return
}

// ExecutePreparer prepares the Execute request.
func (client ReportClient) ExecutePreparer(ctx context.Context, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":     autorest.Encode("path", reportName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reports/{reportName}/run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteSender sends the Execute request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) ExecuteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ExecuteResponder handles the response to the Execute request. The method always
// closes the http.Response Body.
func (client ReportClient) ExecuteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ExecuteByResourceGroupName the operation to execute a report.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportName - report Name.
func (client ReportClient) ExecuteByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result autorest.Response, err error) {
	req, err := client.ExecuteByResourceGroupNamePreparer(ctx, resourceGroupName, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "ExecuteByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecuteByResourceGroupNameSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "ExecuteByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.ExecuteByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "ExecuteByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// ExecuteByResourceGroupNamePreparer prepares the ExecuteByResourceGroupName request.
func (client ReportClient) ExecuteByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":        autorest.Encode("path", reportName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reports/{reportName}/run", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecuteByResourceGroupNameSender sends the ExecuteByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) ExecuteByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ExecuteByResourceGroupNameResponder handles the response to the ExecuteByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportClient) ExecuteByResourceGroupNameResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the report for a subscription by report name.
// Parameters:
// reportName - report Name.
func (client ReportClient) Get(ctx context.Context, reportName string) (result Report, err error) {
	req, err := client.GetPreparer(ctx, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReportClient) GetPreparer(ctx context.Context, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":     autorest.Encode("path", reportName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reports/{reportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReportClient) GetResponder(resp *http.Response) (result Report, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByResourceGroupName gets the report for a resource group under a subscription by report name.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportName - report Name.
func (client ReportClient) GetByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result Report, err error) {
	req, err := client.GetByResourceGroupNamePreparer(ctx, resourceGroupName, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// GetByResourceGroupNamePreparer prepares the GetByResourceGroupName request.
func (client ReportClient) GetByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":        autorest.Encode("path", reportName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reports/{reportName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByResourceGroupNameSender sends the GetByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) GetByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByResourceGroupNameResponder handles the response to the GetByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportClient) GetByResourceGroupNameResponder(resp *http.Response) (result Report, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetExecutionHistory gets the execution history of a report for a subscription by report name.
// Parameters:
// reportName - report Name.
func (client ReportClient) GetExecutionHistory(ctx context.Context, reportName string) (result ReportExecutionListResult, err error) {
	req, err := client.GetExecutionHistoryPreparer(ctx, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetExecutionHistory", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExecutionHistorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetExecutionHistory", resp, "Failure sending request")
		return
	}

	result, err = client.GetExecutionHistoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetExecutionHistory", resp, "Failure responding to request")
	}

	return
}

// GetExecutionHistoryPreparer prepares the GetExecutionHistory request.
func (client ReportClient) GetExecutionHistoryPreparer(ctx context.Context, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":     autorest.Encode("path", reportName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reports/{reportName}/runHistory", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetExecutionHistorySender sends the GetExecutionHistory request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) GetExecutionHistorySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetExecutionHistoryResponder handles the response to the GetExecutionHistory request. The method always
// closes the http.Response Body.
func (client ReportClient) GetExecutionHistoryResponder(resp *http.Response) (result ReportExecutionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetExecutionHistoryByResourceGroupName gets the execution history of a report for a resource group by report name.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// reportName - report Name.
func (client ReportClient) GetExecutionHistoryByResourceGroupName(ctx context.Context, resourceGroupName string, reportName string) (result ReportExecutionListResult, err error) {
	req, err := client.GetExecutionHistoryByResourceGroupNamePreparer(ctx, resourceGroupName, reportName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetExecutionHistoryByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExecutionHistoryByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetExecutionHistoryByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.GetExecutionHistoryByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "GetExecutionHistoryByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// GetExecutionHistoryByResourceGroupNamePreparer prepares the GetExecutionHistoryByResourceGroupName request.
func (client ReportClient) GetExecutionHistoryByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, reportName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"reportName":        autorest.Encode("path", reportName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reports/{reportName}/runHistory", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetExecutionHistoryByResourceGroupNameSender sends the GetExecutionHistoryByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) GetExecutionHistoryByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetExecutionHistoryByResourceGroupNameResponder handles the response to the GetExecutionHistoryByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportClient) GetExecutionHistoryByResourceGroupNameResponder(resp *http.Response) (result ReportExecutionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all reports for a subscription.
func (client ReportClient) List(ctx context.Context) (result ReportListResult, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReportClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReportClient) ListResponder(resp *http.Response) (result ReportListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupName lists all reports for a resource group under a subscription.
// Parameters:
// resourceGroupName - azure Resource Group Name.
func (client ReportClient) ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result ReportListResult, err error) {
	req, err := client.ListByResourceGroupNamePreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "ListByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "ListByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costmanagement.ReportClient", "ListByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupNamePreparer prepares the ListByResourceGroupName request.
func (client ReportClient) ListByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CostManagement/reports", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupNameSender sends the ListByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client ReportClient) ListByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupNameResponder handles the response to the ListByResourceGroupName request. The method always
// closes the http.Response Body.
func (client ReportClient) ListByResourceGroupNameResponder(resp *http.Response) (result ReportListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
