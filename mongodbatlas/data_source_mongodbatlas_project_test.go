package mongodbatlas

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccMongodbatlasDataSource_Project(t *testing.T) {
	projectName := "testAcc"
	dataSourceName := "data.mongodbatlas_project.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMongodbatlasDataSourceProject(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "org_id", testGetOrgID()),
					resource.TestCheckResourceAttrSet(dataSourceName, "created"),
					resource.TestCheckResourceAttr(dataSourceName, "cluster_count", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "name", projectName),
				),
			},
		},
	})
}

func testAccMongodbatlasDataSourceProject(projectName string) string {
	return fmt.Sprintf(`resource "mongodbatlas_project" "test" {
  org_id = "%s"
  name = "%s"
}
data "mongodbatlas_project" "test" {
  depends_on = ["mongodbatlas_project.test"]
  name = "%s"
}`, testGetOrgID(), projectName, projectName)
}
