package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVICertificateManagementProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVICertificateManagementProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVICertificateManagementProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICertificateManagementProfileExists("avi_certificatemanagementprofile.testcertificatemanagementprofile"),
					resource.TestCheckResourceAttr(
						"avi_certificatemanagementprofile.testcertificatemanagementprofile", "name", "cert-test")),
			},
			{
				Config: testAccUpdatedAVICertificateManagementProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVICertificateManagementProfileExists("avi_certificatemanagementprofile.testcertificatemanagementprofile"),
					resource.TestCheckResourceAttr(
						"avi_certificatemanagementprofile.testcertificatemanagementprofile", "name", "cert-abc")),
			},
		},
	})

}

func testAccCheckAVICertificateManagementProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Certificate Management Profile ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVICertificateManagementProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_certificatemanagementprofile" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Certificate Management Profile still exists")
		}
	}
	return nil
}

const testAccAVICertificateManagementProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_certificatemanagementprofile" "testcertificatemanagementprofile" {
	name = "cert-test"
	script_path= "test script path"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVICertificateManagementProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_certificatemanagementprofile" "testcertificatemanagementprofile" {
	name = "cert-abc"
	script_path= "test script path"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
