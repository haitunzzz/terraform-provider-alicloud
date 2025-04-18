package alicloud

import (
	"fmt"
	"testing"

	"log"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("alicloud_ram_account_alias", &resource.Sweeper{
		Name: "alicloud_ram_account_alias",
		F:    testSweepAccountAliases,
	})
}

func testSweepAccountAliases(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapError(err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}

	request := ram.CreateGetAccountAliasRequest()
	raw, err := client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
		return ramClient.GetAccountAlias(request)
	})
	if err != nil {
		return WrapError(err)
	}
	sweeped := false
	resp, _ := raw.(*ram.GetAccountAliasResponse)
	name := resp.AccountAlias
	skip := true
	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			skip = false
			break
		}
	}
	if skip {
		log.Printf("[INFO] Skipping Ram account alias: %s", name)
		return nil
	}
	sweeped = true
	log.Printf("[INFO] Deleting Ram account alias: %s", name)

	_, err = client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
		request := ram.CreateClearAccountAliasRequest()
		return ramClient.ClearAccountAlias(request)
	})
	if err != nil {
		log.Printf("[ERROR] Failed to delete Ram account alias (%s): %s", name, err)
	}

	if sweeped {
		time.Sleep(5 * time.Second)
	}
	return nil
}

func TestAccAliCloudRAMAccountAlias_basic(t *testing.T) {
	randInt := acctest.RandIntRange(1000, 9999)
	var v *ram.GetAccountAliasResponse
	resourceId := "alicloud_ram_account_alias.default"
	ra := resourceAttrInit(resourceId, nil)
	serviceFunc := func() interface{} {
		return &RamService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.RamNoSkipRegions)
		},

		// module name
		IDRefreshName: resourceId,

		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRamAccountAliasDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRamAccountAliasConfig(randInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": fmt.Sprintf("tf-testacc%s%d", defaultRegionToTest, randInt),
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRamAccountAliasConfig(rand int) string {
	return fmt.Sprintf(`
	resource "alicloud_ram_account_alias" "default" {
	  account_alias = "tf-testacc%s%d"
	}`, defaultRegionToTest, rand)
}

func testAccCheckRamAccountAliasDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_ram_account_alias" {
			continue
		}

		// Try to find the alias
		client := testAccProvider.Meta().(*connectivity.AliyunClient)

		request := ram.CreateGetAccountAliasRequest()
		_, err := client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
			return ramClient.GetAccountAlias(request)
		})

		if err != nil && !IsExpectedErrors(err, []string{"EntityNotExist"}) {
			return WrapError(err)
		}
	}
	return nil
}

// Test Ram AccountAlias. >>> Resource test cases, automatically generated.
// Case  AccountAlias测试 9000
func TestAccAliCloudRamAccountAlias_basic9000(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ram_account_alias.default"
	ra := resourceAttrInit(resourceId, AlicloudRamAccountAliasMap9000)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RamServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRamAccountAlias")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccram%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRamAccountAliasBasicDependence9000)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"account_alias": "1511928242963727forcreate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_alias": "1511928242963727forupdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_alias": "1511928242963727forcreate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_alias": "1511928242963727forupdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_alias": "1511928242963727forcreate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_alias": "1511928242963727forupdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_alias": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudRamAccountAliasMap9000 = map[string]string{}

func AlicloudRamAccountAliasBasicDependence9000(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Ram AccountAlias. <<< Resource test cases, automatically generated.
