package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBigqueryConnectionConnection_bigqueryConnectionBasic(t *testing.T) {
	// Uses random provider
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckBigqueryConnectionConnectionDestroyProducer(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionBasic(context),
			},
			{
				Config: testAccBigqueryConnectionConnection_bigqueryConnectionBasicUpdate(context),
			},
		},
	})
}

func testAccBigqueryConnectionConnection_bigqueryConnectionBasic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
    provider         = google-beta
    name             = "tf-test-pg-database-instance%{random_suffix}"
    database_version = "POSTGRES_11"
    region           = "us-central1"
    settings {
		tier = "db-f1-micro"
	}

    deletion_protection = false
}

resource "google_sql_database" "db" {
    provider = google-beta
    instance = google_sql_database_instance.instance.name
    name     = "db"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    provider = google-beta
    name = "username"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_bigquery_connection" "connection" {
    provider      = google-beta
    connection_id = "tf-test-my-connection%{random_suffix}"
    location      = "US"
    friendly_name = "👋"
    description   = "a riveting description"
    cloud_sql {
        instance_id = google_sql_database_instance.instance.connection_name
        database    = google_sql_database.db.name
        type        = "POSTGRES"
        credential {
            username = google_sql_user.user.name
            password = google_sql_user.user.password
        }
    }
}
`, context)
}

func testAccBigqueryConnectionConnection_bigqueryConnectionBasicUpdate(context map[string]interface{}) string {
	return Nprintf(`
resource "google_sql_database_instance" "instance" {
    provider         = google-beta
    name             = "tf-test-mysql-database-instance%{random_suffix}"
    database_version = "MYSQL_5_6"
    region           = "us-central1"
    settings {
		tier = "db-f1-micro"
	}

    deletion_protection = false
}

resource "google_sql_database" "db" {
    provider = google-beta
    instance = google_sql_database_instance.instance.name
    name     = "db2"
}

resource "random_password" "pwd" {
    length = 16
    special = false
}

resource "google_sql_user" "user" {
    provider = google-beta
    name = "username"
    instance = google_sql_database_instance.instance.name
    password = random_password.pwd.result
}

resource "google_bigquery_connection" "connection" {
    provider      = google-beta
    connection_id = "tf-test-my-connection%{random_suffix}"
    location      = "US"
    friendly_name = "👋👋"
    description   = "a very riveting description"
    cloud_sql {
        instance_id = google_sql_database_instance.instance.connection_name
        database    = google_sql_database.db.name
        type        = "MYSQL"
        credential {
            username = google_sql_user.user.name
            password = google_sql_user.user.password
        }
    }
}
`, context)
}
