/*
 * Authorization
 *
 * Provides functions to administrate the Authorization and Trust Management service (XSUAA) of SAP BTP, Cloud Foundry environment. You can manage service instances of the Authorization and Trust Management service. You can also manage roles, role templates, and role collections of your subaccount.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package xsuaa_authz

type UserReference struct {
	Id         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	ZoneId     string `json:"zoneId,omitempty"`
	Email      string `json:"email,omitempty"`
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	Origin     string `json:"origin,omitempty"`

	// FIXME additional fields not mentioned in the swagger file
	Verified                   bool     `json:"verified,omitempty"`
	LegacyVerificationBehavior bool     `json:"legacyVerificationBehavior,omitempty"`
	PasswordChangeRequired     bool     `json:"passwordChangeRequired,omitempty"`
	Version                    int32    `json"version,omitempty"`
	Active                     bool     `json:"active,omitempty"`
	RoleCollections            []string `json:"roleCollections,omitempty"`
}
