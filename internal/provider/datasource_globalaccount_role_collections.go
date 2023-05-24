package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/SAP/terraform-provider-btp/internal/btpcli"
)

func newGlobalaccountRoleCollectionsDataSource() datasource.DataSource {
	return &globalaccountRoleCollectionsDataSource{}
}

type globalaccountRoleCollectionsRoleReferences struct {
	/* OUTPUT */
	RoleTemplateName  types.String `tfsdk:"role_template_name"`
	RoleTemplateAppId types.String `tfsdk:"role_template_app_id"`
	Description       types.String `tfsdk:"description"`
	Name              types.String `tfsdk:"name"`
}

type globalaccountRoleCollectionsValueConfig struct {
	/* OUTPUT */
	Name           types.String                                 `tfsdk:"name"`
	IsReadOnly     types.Bool                                   `tfsdk:"read_only"`
	Description    types.String                                 `tfsdk:"description"`
	RoleReferences []globalaccountRoleCollectionsRoleReferences `tfsdk:"role_references"`
}

type globalaccountRoleCollectionsDataSourceConfig struct {
	/* OUTPUT */
	Values []globalaccountRoleCollectionsValueConfig `tfsdk:"values"`
}

type globalaccountRoleCollectionsDataSource struct {
	cli *btpcli.ClientFacade
}

func (ds *globalaccountRoleCollectionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_globalaccount_role_collections", req.ProviderTypeName)
}

func (ds *globalaccountRoleCollectionsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	ds.cli = req.ProviderData.(*btpcli.ClientFacade)
}

func (ds *globalaccountRoleCollectionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `List all role collections.`,
		Attributes: map[string]schema.Attribute{
			"values": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the role collection.",
							Computed:            true,
						},
						"read_only": schema.BoolAttribute{
							MarkdownDescription: "Whether the role collection is readonly.",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "The description of the role collection.",
							Computed:            true,
						},
						"role_references": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"role_template_name": schema.StringAttribute{
										MarkdownDescription: "The name of the referenced role template.",
										Computed:            true,
									},
									"role_template_app_id": schema.StringAttribute{
										MarkdownDescription: "The name of the referenced template app id",
										Computed:            true,
									},
									"description": schema.StringAttribute{
										MarkdownDescription: "The description of the referenced role",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "The name of the referenced role.",
										Computed:            true,
									},
								},
							},
							Computed: true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}

func (ds *globalaccountRoleCollectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data globalaccountRoleCollectionsDataSourceConfig

	diags := req.Config.Get(ctx, &data)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	cliRes, _, err := ds.cli.Security.RoleCollection.ListByGlobalAccount(ctx)
	if err != nil {
		resp.Diagnostics.AddError("API Error Reading Resource Role Collections (Global Account)", fmt.Sprintf("%s", err))
		return
	}

	data.Values = []globalaccountRoleCollectionsValueConfig{}

	for _, rolecollection := range cliRes {
		val := globalaccountRoleCollectionsValueConfig{
			Name:           types.StringValue(rolecollection.Name),
			Description:    types.StringValue(rolecollection.Description),
			IsReadOnly:     types.BoolValue(rolecollection.IsReadOnly),
			RoleReferences: []globalaccountRoleCollectionsRoleReferences{},
		}

		for _, ref := range rolecollection.RoleReferences {
			val.RoleReferences = append(val.RoleReferences, globalaccountRoleCollectionsRoleReferences{
				RoleTemplateName:  types.StringValue(ref.RoleTemplateName),
				RoleTemplateAppId: types.StringValue(ref.RoleTemplateAppId),
				Description:       types.StringValue(ref.Description),
				Name:              types.StringValue(ref.Name),
			})
		}
		// TODO map additional attributes
		data.Values = append(data.Values, val)
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
