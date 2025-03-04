---
page_title: "btp_globalaccount_user Data Source - terraform-provider-btp"
subcategory: ""
description: |-
  Shows registered users in a global account. Users belong to one of the identity providers (IdPs) of the global account.
---

# btp_globalaccount_user (Data Source)

Shows registered users in a global account. Users belong to one of the identity providers (IdPs) of the global account.

## Example Usage

```terraform
# look up user details which belongs to the default identity provider on global account level
data "btp_globalaccount_user" "someone" {
  user_name = "john.doe@mycompany.com"
}

# look up user details which belongs to a custom identity provider on global account level
data "btp_globalaccount_user" "someone_else" {
  user_name = "jane.doe@mycompany.com"
  origin    = "my-custom-idp"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `user_name` (String) The username of the user.

### Optional

- `origin` (String) The identity provider that hosts the user. Only needed for custom identity provider.

### Read-Only

- `active` (Boolean) Shows if the account is still in use.
- `email` (String) The e-mail address of the user.
- `family_name` (String) The last name of the user.
- `given_name` (String) The given name of the user.
- `id` (String) The ID of the user.
- `role_collections` (Set of String) The set of role collections, which are assigned to the user.
- `verified` (Boolean) The verification status of the user.