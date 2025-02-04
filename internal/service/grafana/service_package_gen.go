// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package grafana

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_grafana_workspace": DataSourceWorkspace,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_grafana_license_association":          ResourceLicenseAssociation,
		"aws_grafana_role_association":             ResourceRoleAssociation,
		"aws_grafana_workspace":                    ResourceWorkspace,
		"aws_grafana_workspace_api_key":            ResourceWorkspaceAPIKey,
		"aws_grafana_workspace_saml_configuration": ResourceWorkspaceSAMLConfiguration,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Grafana
}

var ServicePackage = &servicePackage{}
