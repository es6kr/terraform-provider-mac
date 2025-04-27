//go:build brew

package provider_test

import (
	"regexp"
	"testing"

	"github.com/es6kr/terraform-provider-mac/internal/xerrors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceBrew(t *testing.T) {
	t.Parallel()

	t.Run("data.mac_brew", func(t *testing.T) {
		t.Parallel()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccDataSourceBrew,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.mac_brew.test", "name", "sl"),
						resource.TestMatchResourceAttr("data.mac_brew.test", "path", regexp.MustCompile(`[\w\./]+bin/sl$`)),
					),
				},
			},
		})
	})

	t.Run("data.mac_brew error", func(t *testing.T) {
		t.Parallel()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      testAccDataSourceBrewError,
					ExpectError: regexp.MustCompile(xerrors.ErrNotInstalled.Error()),
				},
			},
		})
	})
}

const testAccDataSourceBrew = `
data "mac_brew" "test" {
  name = "sl"
}
`

const testAccDataSourceBrewError = `
data "mac_brew" "test" {
  name = "ls"
}
`
