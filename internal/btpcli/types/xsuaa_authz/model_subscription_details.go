/*
 * Authorization
 *
 * Provides functions to administrate the Authorization and Trust Management service (XSUAA) of SAP BTP, Cloud Foundry environment. You can manage service instances of the Authorization and Trust Management service. You can also manage roles, role templates, and role collections of your subaccount.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package xsuaa_authz

type SubscriptionDetails struct {
	// The identifier for the subscription to the application.
	Id string `json:"id,omitempty"`
	// The application ID is the xsappname plus the identifier, which consists of an exclamation mark (!), an identifier for the plan underwhich the application is deployed, and an index number.
	Appid string `json:"appid,omitempty"`
	// Tenant ID of your subaccount. The zoneId parameter is the same ID and can be used interchangably with tenant ID.
	TenantId string `json:"tenantId,omitempty"`
}
