/*
 * Provisioning Service
 *
 * The Provisioning service provides REST APIs that are responsible for the provisioning and deprovisioning of environment instances and tenants in the corresponding region.  Provisioning is performed after validation by the Entitlements service. Use the APIs in this service to manage and create environment instances, such as a Cloud Foundry org, in a subaccount and to retrieve the plans and quota assignments for a subaccount.  NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package provisioning

type AvailableEnvironmentResponseObject struct {
	// The availability level of the environment broker.
	AvailabilityLevel string `json:"availabilityLevel,omitempty"`
	// The create schema of the environment broker.
	CreateSchema string `json:"createSchema,omitempty"`
	// Description of the service plan for the available environment.
	Description string `json:"description,omitempty"`
	// The type of environment that is available (for example: cloudfoundry).
	EnvironmentType string `json:"environmentType,omitempty"`
	// The landscape label of the environment broker.
	LandscapeLabel string `json:"landscapeLabel,omitempty"`
	// Name of the service plan for the available environment.
	PlanName string `json:"planName,omitempty"`
	// Specifies if the consumer can change the plan of an existing instance of the environment.
	PlanUpdatable bool `json:"planUpdatable,omitempty"`
	// The short description of the service.
	ServiceDescription string `json:"serviceDescription,omitempty"`
	// The display name of the service.
	ServiceDisplayName string `json:"serviceDisplayName,omitempty"`
	// The URL of the documentation link for the service.
	ServiceDocumentationUrl string `json:"serviceDocumentationUrl,omitempty"`
	// The URL of the image for the service.
	ServiceImageUrl string `json:"serviceImageUrl,omitempty"`
	// The long description of the service.
	ServiceLongDescription string `json:"serviceLongDescription,omitempty"`
	// Name of the service offered in the catalog of the corresponding environment broker (for example, cloudfoundry).
	ServiceName string `json:"serviceName,omitempty"`
	// The URL of the support link for the service.
	ServiceSupportUrl string `json:"serviceSupportUrl,omitempty"`
	// Technical key of the corresponding environment broker.
	TechnicalKey string `json:"technicalKey,omitempty"`
	// The update schema of the environment broker.
	UpdateSchema string `json:"updateSchema,omitempty"`
}
