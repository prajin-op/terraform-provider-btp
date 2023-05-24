/*
 * SaaS Provisioning Service
 *
 * The SAP SaaS Provisioning service provides REST APIs that are responsible for the registration and provisioning of multitenant applications and services.   Use the APIs in this service to perform various operations related to your multitenant applications and services. For example, to get application registration details, subscribe a tenant to your application, unsubscribe a tenant from your application, retrieve all your application subscriptions, update subscription dependencies, and to get subscription job information.  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package saas_manager_service

type ApplicationSubscriptionsResponseObject struct {
	// Specifies the ability to use the service plan of the subscribed application. The actual amount has no bearing on the maximum consumption limit of the application.
	Amount int64 `json:"amount,omitempty"`
	// The unique registration name of the deployed multitenant application, as defined by the app developer.
	AppName string `json:"appName,omitempty"`
	// The date and time the subscription was last modified. Dates and times are in UTC format.
	ChangedOn string `json:"changedOn,omitempty"`
	// A subscription code for the application.
	Code string `json:"code,omitempty"`
	// Tenant ID of the global account or subaccount of the consumer that has subscribed to the multitenant application.
	ConsumerTenantId string `json:"consumerTenantId,omitempty"`
	// The date and time the subscription was created. Dates and times are in UTC format.
	CreatedOn string `json:"createdOn,omitempty"`
	// Any reuse services used or required by a subscribed application and its services.
	Dependencies []DependenciesResponseObject `json:"dependencies,omitempty"`
	// Error description for the following statuses: SUBSCRIBE_FAILED, UNSUBSCRIBE_FAILED, UPDATE_FAILED.
	Error_ string `json:"error,omitempty"`
	// ID of the associated global account.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	// Whether the consumer tenant is active. This field is returned only if one of the following query parameters was used during the API call: tenantId, subaccountId
	IsConsumerTenantActive bool `json:"isConsumerTenantActive,omitempty"`
	// The license type of the associated global account.
	LicenseType string `json:"licenseType,omitempty"`
	// The ID of the multitenant application that is registered to the SAP SaaS Provisioning registry.
	ServiceInstanceId string `json:"serviceInstanceId,omitempty"`
	// State of the subscriptions. Possible states: IN_PROCESS, SUBSCRIBED, SUBSCRIBE_FAILED, UNSUBSCRIBE_FAILED, UPDATE_FAILED.
	State string `json:"state,omitempty"`
	// ID of the associated subaccount.
	SubaccountId string `json:"subaccountId,omitempty"`
	// Consumer Subdomain
	Subdomain string `json:"subdomain,omitempty"`
	// Application URL
	Url string `json:"url,omitempty"`
}
