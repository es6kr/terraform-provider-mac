//go:build service

package provider_test

import (
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceServiceBasic(t *testing.T) {
	t.Parallel()

	t.Run("resource.mac_service", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckServiceDestroy,
			Steps: []resource.TestStep{
				{
					Config: readTFFile("./testdata/resources/service/resources_service_nebula.tf"),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckServiceExists("com.slackhq.nebula"),
					),
				},
			},
		})
	})
}

func testAccCheckServiceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Check if the service is bootstrapped and running
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return errors.New("Service resource not found")
		}

		// Check the `id` field
		if rs.Primary.ID == "" {
			return errors.New("Service resource ID is empty")
		}

		// Here you can add specific validation like checking service state, etc.
		// Currently, we are simply asserting the existence of the service in the state.

		return nil
	}
}
