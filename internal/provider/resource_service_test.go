//go:build service

package provider_test

import (
	"context"
	"os"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/es6kr/terraform-provider-mac/internal/launchctl"
	"github.com/es6kr/terraform-provider-mac/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceServiceBasic(t *testing.T) {
	t.Parallel()

	t.Run("resource.mac_service", func(t *testing.T) {
		testServiceTargetFromServiceID(t)
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckServiceDestroy,
			Steps: []resource.TestStep{
				{
					Config: readTFFile("./testdata/resources/service/resources_service_nebula.tf"),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckServiceExists("homebrew.mxcl.nebula"),
					),
				},
			},
		})
	})
}

func testAccCheckServiceDestroy(s *terraform.State) error {
	for _, resource := range s.RootModule().Resources {
		if resource.Type != "mac_service" {
			continue
		}

		id := resource.Primary.ID
		plist := resource.Primary.Attributes["plist"]

		if _, err := os.Stat(plist); !errors.Is(err, os.ErrNotExist) {
			bootoutErr := launchctl.Bootout(context.Background(), provider.ServiceTargetFromServiceID(id))
			return errors.CombineErrors(err, bootoutErr)
		}
	}

	return nil
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

func testServiceTargetFromServiceID(t *testing.T) {
	tests := []struct {
		id             string
		expectedTarget string
	}{
		{"service:gui/501/homebrew.mxcl.nebula", "gui/501/homebrew.mxcl.nebula"},
		{"service:system/custom.label", "system/custom.label"},
		{"service:service/789/myapp", "service/789/myapp"},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			actual := provider.ServiceTargetFromServiceID(test.id)
			if actual != test.expectedTarget {
				t.Errorf("serviceTargetFromServiceID(%s) = %s; want %s", test.id, actual, test.expectedTarget)
			}
		})
	}
}
