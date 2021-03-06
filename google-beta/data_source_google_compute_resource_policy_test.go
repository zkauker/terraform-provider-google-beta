package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceComputeResourcePolicy(t *testing.T) {
	t.Parallel()

	randomSuffix := randString(t, 10)

	rsName := "foo_" + randomSuffix
	rsFullName := fmt.Sprintf("google_compute_resource_policy.%s", rsName)
	dsName := "my_policy_" + randomSuffix
	dsFullName := fmt.Sprintf("data.google_compute_resource_policy.%s", dsName)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDataSourceComputeResourcePolicyDestroy(t, rsFullName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceComputeResourcePolicyConfig(rsName, dsName, randomSuffix),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceComputeResourcePolicyCheck(t, dsFullName, rsFullName),
				),
			},
		},
	})
}

func testAccDataSourceComputeResourcePolicyCheck(t *testing.T, dataSourceName string, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", dataSourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName)
		}

		dsAttr := ds.Primary.Attributes
		rsAttr := rs.Primary.Attributes

		policyAttrsToTest := []string{
			"name",
		}

		for _, attrToCheck := range policyAttrsToTest {
			if dsAttr[attrToCheck] != rsAttr[attrToCheck] {
				return fmt.Errorf(
					"%s is %s; want %s",
					attrToCheck,
					dsAttr[attrToCheck],
					rsAttr[attrToCheck],
				)
			}
		}

		if !compareSelfLinkOrResourceName("", dsAttr["self_link"], rsAttr["self_link"], nil) && dsAttr["self_link"] != rsAttr["self_link"] {
			return fmt.Errorf("self link does not match: %s vs %s", dsAttr["self_link"], rsAttr["self_link"])
		}

		return nil
	}
}

func testAccCheckDataSourceComputeResourcePolicyDestroy(t *testing.T, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_resource_policy" {
				continue
			}

			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			policyAttrs := rs.Primary.Attributes

			_, err := config.NewComputeClient(config.userAgent).ResourcePolicies.Get(
				config.Project, policyAttrs["region"], policyAttrs["name"]).Do()
			if err == nil {
				return fmt.Errorf("Resource Policy still exists")
			}
		}

		return nil
	}
}

func testAccDataSourceComputeResourcePolicyConfig(rsName, dsName, randomSuffix string) string {
	return fmt.Sprintf(`
resource "google_compute_resource_policy" "%s" {
  name   = "policy-%s"
  region = "us-central1"
  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time    = "04:00"
      }
    }
  }
}

data "google_compute_resource_policy" "%s" {
  name     = google_compute_resource_policy.%s.name
  region   = google_compute_resource_policy.%s.region
}
`, rsName, randomSuffix, dsName, rsName, rsName)
}
