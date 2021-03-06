// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotebooksInstanceIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstanceIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccNotebooksInstanceIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccNotebooksInstanceIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccNotebooksInstanceIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccNotebooksInstanceIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstanceIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccNotebooksInstanceIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccNotebooksInstanceIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  provider = google-beta
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

resource "google_notebooks_instance_iam_member" "foo" {
  provider = google-beta
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccNotebooksInstanceIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  provider = google-beta
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_notebooks_instance_iam_policy" "foo" {
  provider = google-beta
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccNotebooksInstanceIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  provider = google-beta
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_notebooks_instance_iam_policy" "foo" {
  provider = google-beta
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccNotebooksInstanceIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  provider = google-beta
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

resource "google_notebooks_instance_iam_binding" "foo" {
 
  provider = google-beta
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccNotebooksInstanceIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_notebooks_instance" "instance" {
  provider = google-beta
  name = "tf-test-notebooks-instance%{random_suffix}"
  location = "us-west1-a"
  machine_type = "e2-medium"
  vm_image {
    project      = "deeplearning-platform-release"
    image_family = "tf-latest-cpu"
  }
}

resource "google_notebooks_instance_iam_binding" "foo" {
  provider = google-beta
  project = google_notebooks_instance.instance.project
  location = google_notebooks_instance.instance.location
  instance_name = google_notebooks_instance.instance.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}
