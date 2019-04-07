package mongodbatlas

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccMongodbatlasProject_basic(t *testing.T) {
	projectName := "testAcc"
	resourceName := "mongodbatlas_project.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMongodbatlasProject(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", testGetOrgID()),
					resource.TestCheckResourceAttrSet(resourceName, "created"),
					resource.TestCheckResourceAttr(resourceName, "cluster_count", "0"),
					resource.TestCheckResourceAttr(resourceName, "name", projectName),
				),
			},
		},
	})
}

func testAccMongodbatlasProject(projectName string) string {
	return fmt.Sprintf(`resource "mongodbatlas_project" "test" {
  org_id = "%s"
  name = "%s"
}`, testGetOrgID(), projectName)
}
