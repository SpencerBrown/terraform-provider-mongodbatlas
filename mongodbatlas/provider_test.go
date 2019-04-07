package mongodbatlas

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"mongodbatlas": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

// checks that environment variables are set for acceptance test
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MONGODB_ATLAS_USERNAME"); v == "" {
		t.Fatal("MONGODB_ATLAS_USERNAME must be set for acceptance tests")
		if v := os.Getenv("MONGODB_ATLAS_API_KEY"); v == "" {
			t.Fatal("MONGODB_ATLAS_API_KEY must be set for acceptance tests")
			if v := os.Getenv("MONGODB_ATLAS_ORG_ID"); v == "" {
				t.Fatal("MONGODB_ATLAS_ORG_ID must be set for acceptance tests")
			}
		}
	}
}

// Returns the org ID for testing
func testGetOrgID() string {
	return os.Getenv("MONGODB_ATLAS_ORG_ID")
}
