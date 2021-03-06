package apsarastack

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccApsaraStackRouterInterfacesDataSourceBasic(t *testing.T) {
	preCheck := func() {
		testAccPreCheck(t)
		testAccPreCheckWithAccountSiteType(t, DomesticSite)
	}
	rand := acctest.RandIntRange(1000, 9999)

	oppositeInterfaceIdConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}_fake"`,
		}),
	}

	statusConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"status":                `"Active"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"status":                `"Inactive"`,
		}),
	}

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"name_regex": `"${var.name}_initiating"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"name_regex": `"${var.name}_fake"`,
		}),
	}

	specificationConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"specification":         `"Large.2"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"specification":         `"Large.1"`,
		}),
	}

	routerIdConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"router_id": `"${apsarastack_vpc.default.0.router_id}"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"router_id": `"${apsarastack_vpc.default.0.router_id}_fake"`,
		}),
	}

	routerTypeConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"router_type":           `"VRouter"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"router_type":           `"VBR"`,
		}),
	}

	roleConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"role":                  `"InitiatingSide"`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"role":                  `"AcceptingSide"`,
		}),
	}

	idsConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"ids": `[ "${apsarastack_router_interface.initiating.id}" ]`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"ids": `[ "${apsarastack_router_interface.initiating.id}_fake" ]`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"role":                  `"InitiatingSide"`,
			"name_regex":            `"${var.name}_initiating"`,
			"specification":         `"Large.2"`,
			"router_id":             `"${apsarastack_vpc.default.0.router_id}"`,
			"router_type":           `"VRouter"`,
			"ids":                   `[ "${apsarastack_router_interface.initiating.id}" ]`,
		}),
		fakeConfig: testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand, map[string]string{
			"opposite_interface_id": `"${apsarastack_router_interface_connection.foo.opposite_interface_id}"`,
			"role":                  `"AcceptingSide"`,
			"name_regex":            `"${var.name}_initiating"`,
			"specification":         `"Large.2"`,
			"router_id":             `"${apsarastack_vpc.default.0.router_id}"`,
			"router_type":           `"VRouter"`,
			"ids":                   `[ "${apsarastack_router_interface.initiating.id}" ]`,
		}),
	}

	routerInterfacesCheckInfo.dataSourceTestCheckWithPreCheck(t, rand, preCheck, oppositeInterfaceIdConf, statusConf, nameRegexConf, specificationConf, routerIdConf,
		routerTypeConf, roleConf, idsConf, allConf)

}

func testAccCheckApsaraStackRouterInterfacesDataSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}

	config := fmt.Sprintf(`
provider "apsarastack" {
  region = "${var.region}"
}
variable "region" {
  default = "cn-qingdao-env66-d01"
}
variable "name" {
  default = "tf-testAccCheckApsaraStackRouterInterfacesDataSourceConfig%d"
}
variable cidr_block_list {
	type = "list"
	default = [ "172.16.0.0/12", "192.168.0.0/16" ]
}

resource "apsarastack_vpc" "default" {
  count = 2
  name = "${var.name}"
  cidr_block = "${element(var.cidr_block_list,count.index)}"
}
resource "apsarastack_router_interface" "initiating" {
  opposite_region = "${var.region}"
  router_type = "VRouter"
  router_id = "${apsarastack_vpc.default.0.router_id}"
  role = "InitiatingSide"
  specification = "Large.2"
  name = "${var.name}_initiating"
  description = "${var.name}_decription"

}
resource "apsarastack_router_interface" "opposite" {
  provider = "apsarastack"
  opposite_region = "${var.region}"
  router_type = "VRouter"
  router_id = "${apsarastack_vpc.default.1.router_id}"
  role = "AcceptingSide"
  specification = "Large.1"
  name = "${var.name}_opposite"
  description = "${var.name}_decription"

}

resource "apsarastack_router_interface_connection" "foo" {
  interface_id = "${apsarastack_router_interface.initiating.id}"
  opposite_interface_id = "${apsarastack_router_interface.opposite.id}"
  depends_on = ["apsarastack_router_interface_connection.bar"]
  opposite_interface_owner_id = "1262302482727553"
  opposite_router_id = apsarastack_vpc.default.0.router_id
  opposite_router_type = "VRouter"
}

resource "apsarastack_router_interface_connection" "bar" {
  provider = "apsarastack"
  interface_id = "${apsarastack_router_interface.opposite.id}"
  opposite_interface_id = "${apsarastack_router_interface.initiating.id}"
  opposite_interface_owner_id =  "1262302482727553"
  opposite_router_id = apsarastack_vpc.default.1.router_id
  opposite_router_type = "VRouter"
}
data "apsarastack_router_interfaces" "default" {
  %s
}`, rand, strings.Join(pairs, "\n  "))
	return config
}

var existRouterInterfacesMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"ids.#":                                    "1",
		"names.#":                                  "1",
		"interfaces.#":                             "1",
		"interfaces.0.id":                          CHECKSET,
		"interfaces.0.status":                      "Active",
		"interfaces.0.name":                        fmt.Sprintf("tf-testAccCheckApsaraStackRouterInterfacesDataSourceConfig%d_initiating", rand),
		"interfaces.0.description":                 fmt.Sprintf("tf-testAccCheckApsaraStackRouterInterfacesDataSourceConfig%d_decription", rand),
		"interfaces.0.role":                        "InitiatingSide",
		"interfaces.0.specification":               "Large.2",
		"interfaces.0.router_id":                   CHECKSET,
		"interfaces.0.router_type":                 "VRouter",
		"interfaces.0.vpc_id":                      CHECKSET,
		"interfaces.0.access_point_id":             "",
		"interfaces.0.creation_time":               CHECKSET,
		"interfaces.0.opposite_region_id":          CHECKSET,
		"interfaces.0.opposite_interface_id":       CHECKSET,
		"interfaces.0.opposite_router_id":          CHECKSET,
		"interfaces.0.opposite_router_type":        "VRouter",
		"interfaces.0.opposite_interface_owner_id": CHECKSET,
		"interfaces.0.health_check_source_ip":      "",
		"interfaces.0.health_check_target_ip":      "",
	}
}

var fakeRouterInterfacesMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"ids.#":        "0",
		"names.#":      "0",
		"interfaces.#": "0",
	}
}

var routerInterfacesCheckInfo = dataSourceAttr{
	resourceId:   "data.apsarastack_router_interfaces.default",
	existMapFunc: existRouterInterfacesMapFunc,
	fakeMapFunc:  fakeRouterInterfacesMapFunc,
}
