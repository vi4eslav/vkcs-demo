package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

const testCreateAPP = `
provider "myapp" {
	my_option = "app1"
}

resource "myapp_hello" "myhello" {
	content_from_conf = "cfg"
}
`

const testUpdateAPP = testCreateAPP

func testAccCheckAgent(n string, content *struct{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		if !strings.Contains(rs.Primary.Attributes["content_from_conf"], "cfg") {
			return fmt.Errorf("can't find '%s' in '%s'", "cfg", rs.Primary.Attributes["content_from_conf"])
		}

		return nil
	}
}

// acceptance tests
func TestAccCreateAgent(t *testing.T) {
	var content struct{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testAccCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testCreateAPP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAgent("myapp_hello.myhello", &content),
				),
			},
			{
				Config: testUpdateAPP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAgent("myapp_hello.myhello", &content),
				),
			},
		},
	})

	//	pre-check

	//  apply to create
	//  apply to check not diff
	//  ...
	//  ...

	//  destroy

}

func testAccPreCheck(t *testing.T) {
}

func testAccCheckDestroy(s *terraform.State) error {
	return nil
}

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]func() (*schema.Provider, error){
		"myapp": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}
