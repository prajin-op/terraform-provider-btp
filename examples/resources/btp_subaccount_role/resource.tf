resource "btp_subaccount_role" "xsuaa_auditor" {
  subaccount_id      = "6aa64c2f-38c1-49a9-b2e8-cf9fea769b7f"
  name               = "XSUAA Auditor"
  role_template_name = "xsuaa_auditor"
  app_id             = "xsuaa!t1"
}
