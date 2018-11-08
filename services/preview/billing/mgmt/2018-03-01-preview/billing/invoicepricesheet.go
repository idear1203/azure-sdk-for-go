package billing

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

// InvoicePricesheetClient is the billing client provides access to billing resources for Azure subscriptions.
type InvoicePricesheetClient struct {
	BaseClient
}

// NewInvoicePricesheetClient creates an instance of the InvoicePricesheetClient client.
func NewInvoicePricesheetClient(subscriptionID string) InvoicePricesheetClient {
	return NewInvoicePricesheetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoicePricesheetClientWithBaseURI creates an instance of the InvoicePricesheetClient client.
func NewInvoicePricesheetClientWithBaseURI(baseURI string, subscriptionID string) InvoicePricesheetClient {
	return InvoicePricesheetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get pricesheet data for invoice id (invoiceName).
// Parameters:
// billingAccountID - azure Billing Account ID.
// invoiceName - the name of an invoice resource.
func (client InvoicePricesheetClient) Get(ctx context.Context, billingAccountID string, invoiceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicePricesheetClient.Get")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountID, invoiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicePricesheetClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "billing.InvoicePricesheetClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicePricesheetClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvoicePricesheetClient) GetPreparer(ctx context.Context, billingAccountID string, invoiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
		"invoiceName":      autorest.Encode("path", invoiceName),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoices/{invoiceName}/pricesheet/download", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicePricesheetClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvoicePricesheetClient) GetResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
