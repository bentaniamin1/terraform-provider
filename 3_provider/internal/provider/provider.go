package provider

import (
	"context"
	"net/http"

	tfdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	tffunction "github.com/hashicorp/terraform-plugin-framework/function"
	tfprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	tfschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

type UserProviderModel struct {
	Endpoint tftypes.String `tfsdk:"endpoint"`
}

type UserProvider struct {
	endpoint string
	client   *http.Client
}

var _ tfprovider.Provider = &UserProvider{}
var _ tfprovider.ProviderWithFunctions = &UserProvider{}

func New() func() tfprovider.Provider {
	return func() tfprovider.Provider {
		return &UserProvider{}
	}
}

func (p *UserProvider) Metadata(ctx context.Context, req tfprovider.MetadataRequest, resp *tfprovider.MetadataResponse) {
	resp.TypeName = "myuserprovider" // matches in your .tf file `resource "myuserprovider_user" "john_doe" {`
}

func (p *UserProvider) Schema(ctx context.Context, req tfprovider.SchemaRequest, resp *tfprovider.SchemaResponse) {
	resp.Schema = tfschema.Schema{
		Attributes: map[string]tfschema.Attribute{
			"endpoint": tfschema.StringAttribute{
				MarkdownDescription: "Endpoint of the API, e.g. - http://localhost:6251/",
				Required:            true,
			},
		},
	}
}

func (p *UserProvider) Configure(ctx context.Context, req tfprovider.ConfigureRequest, resp *tfprovider.ConfigureResponse) {
	var data UserProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	p.endpoint = data.Endpoint.ValueString()
	p.client = http.DefaultClient
	resp.DataSourceData = p
	resp.ResourceData = p
}

func (p *UserProvider) Resources(ctx context.Context) []func() tfresource.Resource {
	return []func() tfresource.Resource{
		NewUserResource,
	}
}

func (p *UserProvider) DataSources(ctx context.Context) []func() tfdatasource.DataSource {
	return []func() tfdatasource.DataSource{}
}

func (p *UserProvider) Functions(ctx context.Context) []func() tffunction.Function {
	return []func() tffunction.Function{}
}
