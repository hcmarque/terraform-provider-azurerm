package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccAzureRMHDInsightStormCluster_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.RequiresImportErrorStep(testAccAzureRMHDInsightStormCluster_requiresImport),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
			{
				Config: testAccAzureRMHDInsightStormCluster_updated(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep(
				"roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_sshKeys(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_sshKeys(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("storage_account",
				"roles.0.head_node.0.ssh_keys",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.ssh_keys",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.ssh_keys",
				"roles.0.zookeeper_node.0.vm_size"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_virtualNetwork(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_virtualNetwork(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_complete(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_tls(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_tls(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_allMetastores(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_allMetastores(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account",
				"metastores.0.hive.0.password",
				"metastores.0.oozie.0.password",
				"metastores.0.ambari.0.password"),
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_hiveMetastore(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_hiveMetastore(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
		},
	})
}

func TestAccAzureRMHDInsightStormCluster_updateMetastore(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_hdinsight_storm_cluster", "test")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMHDInsightClusterDestroy(data.ResourceType),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureRMHDInsightStormCluster_hiveMetastore(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account",
				"metastores.0.hive.0.password",
				"metastores.0.oozie.0.password",
				"metastores.0.ambari.0.password"),
			{
				Config: testAccAzureRMHDInsightStormCluster_allMetastores(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMHDInsightClusterExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "https_endpoint"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "ssh_endpoint"),
				),
			},
			data.ImportStep("roles.0.head_node.0.password",
				"roles.0.head_node.0.vm_size",
				"roles.0.worker_node.0.password",
				"roles.0.worker_node.0.vm_size",
				"roles.0.zookeeper_node.0.password",
				"roles.0.zookeeper_node.0.vm_size",
				"storage_account",
				"metastores.0.hive.0.password",
				"metastores.0.oozie.0.password",
				"metastores.0.ambari.0.password"),
		},
	})
}

func testAccAzureRMHDInsightStormCluster_basic(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    storm = "1.1"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_requiresImport(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_basic(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_storm_cluster" "import" {
  name                = azurerm_hdinsight_storm_cluster.test.name
  resource_group_name = azurerm_hdinsight_storm_cluster.test.resource_group_name
  location            = azurerm_hdinsight_storm_cluster.test.location
  cluster_version     = azurerm_hdinsight_storm_cluster.test.cluster_version
  tier                = azurerm_hdinsight_storm_cluster.test.tier
  dynamic "component_version" {
    for_each = azurerm_hdinsight_storm_cluster.test.component_version
    content {
      storm = component_version.value.storm
    }
  }
  dynamic "gateway" {
    for_each = azurerm_hdinsight_storm_cluster.test.gateway
    content {
      enabled  = gateway.value.enabled
      password = gateway.value.password
      username = gateway.value.username
    }
  }
  dynamic "storage_account" {
    for_each = azurerm_hdinsight_storm_cluster.test.storage_account
    content {
      is_default           = storage_account.value.is_default
      storage_account_key  = storage_account.value.storage_account_key
      storage_container_id = storage_account.value.storage_container_id
    }
  }
  dynamic "roles" {
    for_each = azurerm_hdinsight_storm_cluster.test.roles
    content {
      dynamic "head_node" {
        for_each = lookup(roles.value, "head_node", [])
        content {
          password           = lookup(head_node.value, "password", null)
          ssh_keys           = lookup(head_node.value, "ssh_keys", null)
          subnet_id          = lookup(head_node.value, "subnet_id", null)
          username           = head_node.value.username
          virtual_network_id = lookup(head_node.value, "virtual_network_id", null)
          vm_size            = head_node.value.vm_size
        }
      }

      dynamic "worker_node" {
        for_each = lookup(roles.value, "worker_node", [])
        content {
          password              = lookup(worker_node.value, "password", null)
          ssh_keys              = lookup(worker_node.value, "ssh_keys", null)
          subnet_id             = lookup(worker_node.value, "subnet_id", null)
          target_instance_count = worker_node.value.target_instance_count
          username              = worker_node.value.username
          virtual_network_id    = lookup(worker_node.value, "virtual_network_id", null)
          vm_size               = worker_node.value.vm_size
        }
      }

      dynamic "zookeeper_node" {
        for_each = lookup(roles.value, "zookeeper_node", [])
        content {
          password           = lookup(zookeeper_node.value, "password", null)
          ssh_keys           = lookup(zookeeper_node.value, "ssh_keys", null)
          subnet_id          = lookup(zookeeper_node.value, "subnet_id", null)
          username           = zookeeper_node.value.username
          virtual_network_id = lookup(zookeeper_node.value, "virtual_network_id", null)
          vm_size            = zookeeper_node.value.vm_size
        }
      }
    }
  }
}
`, template)
}

func testAccAzureRMHDInsightStormCluster_sshKeys(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s

variable "ssh_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCqaZoyiz1qbdOQ8xEf6uEu1cCwYowo5FHtsBhqLoDnnp7KUTEBN+L2NxRIfQ781rxV6Iq5jSav6b2Q8z5KiseOlvKA/RF2wqU0UPYqQviQhLmW6THTpmrv/YkUCuzxDpsH7DUDhZcwySLKVVe0Qm3+5N2Ta6UYH3lsDf9R9wTP2K/+vAnflKebuypNlmocIvakFWoZda18FOmsOoIVXQ8HWFNCuw9ZCunMSN62QGamCe3dL5cXlkgHYv7ekJE15IA9aOJcM7e90oeTqo+7HTcWfdu0qQqPWY5ujyMw/llas8tsXY85LFqRnr3gJ02bAscjc477+X+j/gkpFoN1QEmt terraform@demo.tld"
}

resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    storm = "1.1"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      ssh_keys = [var.ssh_key]
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      ssh_keys              = [var.ssh_key]
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      ssh_keys = [var.ssh_key]
    }
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_updated(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    storm = "1.1"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 5
    }

    zookeeper_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }

  tags = {
    Hello = "World"
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_virtualNetwork(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    storm = "1.1"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size            = "Standard_A4_V2"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
      subnet_id             = azurerm_subnet.test.id
      virtual_network_id    = azurerm_virtual_network.test.id
    }

    zookeeper_node {
      vm_size            = "Standard_A4_V2"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }
  }
}
`, template, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_complete(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_virtual_network" "test" {
  name                = "acctestvirtnet%d"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
}

resource "azurerm_subnet" "test" {
  name                 = "acctestsubnet%d"
  resource_group_name  = azurerm_resource_group.test.name
  virtual_network_name = azurerm_virtual_network.test.name
  address_prefix       = "10.0.2.0/24"
}

resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"

  component_version {
    storm = "1.1"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size            = "Standard_A4_V2"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
      subnet_id             = azurerm_subnet.test.id
      virtual_network_id    = azurerm_virtual_network.test.id
    }

    zookeeper_node {
      vm_size            = "Standard_A4_V2"
      username           = "acctestusrvm"
      password           = "AccTestvdSC4daf986!"
      subnet_id          = azurerm_subnet.test.id
      virtual_network_id = azurerm_virtual_network.test.id
    }
  }

  tags = {
    Hello = "World"
  }
}
`, template, data.RandomInteger, data.RandomInteger, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsa%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "test" {
  name                  = "acctest"
  storage_account_name  = azurerm_storage_account.test.name
  container_access_type = "private"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}

func testAccAzureRMHDInsightStormCluster_tls(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s

resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"
  tls_min_version     = "1.2"

  component_version {
    storm = "1.1"
  }

  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }

  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }

  roles {
    head_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }

    worker_node {
      vm_size               = "Standard_A4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 3
    }

    zookeeper_node {
      vm_size  = "Standard_A4_V2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
}
`, template, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_allMetastores(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s
resource "azurerm_sql_server" "test" {
  name                         = "acctestsql-%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  administrator_login          = "sql_admin"
  administrator_login_password = "TerrAform123!"
  version                      = "12.0"
}
resource "azurerm_sql_database" "hive" {
  name                             = "hive"
  resource_group_name              = azurerm_resource_group.test.name
  location                         = azurerm_resource_group.test.location
  server_name                      = azurerm_sql_server.test.name
  collation                        = "SQL_Latin1_General_CP1_CI_AS"
  create_mode                      = "Default"
  requested_service_objective_name = "GP_Gen5_2"
}
resource "azurerm_sql_database" "oozie" {
  name                             = "oozie"
  resource_group_name              = azurerm_resource_group.test.name
  location                         = azurerm_resource_group.test.location
  server_name                      = azurerm_sql_server.test.name
  collation                        = "SQL_Latin1_General_CP1_CI_AS"
  create_mode                      = "Default"
  requested_service_objective_name = "GP_Gen5_2"
}
resource "azurerm_sql_database" "ambari" {
  name                             = "ambari"
  resource_group_name              = azurerm_resource_group.test.name
  location                         = azurerm_resource_group.test.location
  server_name                      = azurerm_sql_server.test.name
  collation                        = "SQL_Latin1_General_CP1_CI_AS"
  create_mode                      = "Default"
  requested_service_objective_name = "GP_Gen5_2"
}
resource "azurerm_sql_firewall_rule" "AzureServices" {
  name                = "allow-azure-services"
  resource_group_name = azurerm_resource_group.test.name
  server_name         = azurerm_sql_server.test.name
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "0.0.0.0"
}
resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"
  component_version {
    storm = "1.1"
  }
  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }
  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }
  roles {
    head_node {
      vm_size  = "Standard_D3_v2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
    worker_node {
      vm_size               = "Standard_D4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 2
    }
    zookeeper_node {
      vm_size  = "Standard_D3_v2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
  metastores {
    hive {
      server        = azurerm_sql_server.test.fully_qualified_domain_name
      database_name = azurerm_sql_database.hive.name
      username      = azurerm_sql_server.test.administrator_login
      password      = azurerm_sql_server.test.administrator_login_password
    }
    oozie {
      server        = azurerm_sql_server.test.fully_qualified_domain_name
      database_name = azurerm_sql_database.oozie.name
      username      = azurerm_sql_server.test.administrator_login
      password      = azurerm_sql_server.test.administrator_login_password
    }
    ambari {
      server        = azurerm_sql_server.test.fully_qualified_domain_name
      database_name = azurerm_sql_database.ambari.name
      username      = azurerm_sql_server.test.administrator_login
      password      = azurerm_sql_server.test.administrator_login_password
    }
  }
}
`, template, data.RandomInteger, data.RandomInteger)
}

func testAccAzureRMHDInsightStormCluster_hiveMetastore(data acceptance.TestData) string {
	template := testAccAzureRMHDInsightStormCluster_template(data)
	return fmt.Sprintf(`
%s
resource "azurerm_sql_server" "test" {
  name                         = "acctestsql-%d"
  resource_group_name          = azurerm_resource_group.test.name
  location                     = azurerm_resource_group.test.location
  administrator_login          = "sql_admin"
  administrator_login_password = "TerrAform123!"
  version                      = "12.0"
}
resource "azurerm_sql_database" "hive" {
  name                             = "hive"
  resource_group_name              = azurerm_resource_group.test.name
  location                         = azurerm_resource_group.test.location
  server_name                      = azurerm_sql_server.test.name
  collation                        = "SQL_Latin1_General_CP1_CI_AS"
  create_mode                      = "Default"
  requested_service_objective_name = "GP_Gen5_2"
}
resource "azurerm_sql_firewall_rule" "AzureServices" {
  name                = "allow-azure-services"
  resource_group_name = azurerm_resource_group.test.name
  server_name         = azurerm_sql_server.test.name
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "0.0.0.0"
}
resource "azurerm_hdinsight_storm_cluster" "test" {
  name                = "acctesthdi-%d"
  resource_group_name = azurerm_resource_group.test.name
  location            = azurerm_resource_group.test.location
  cluster_version     = "3.6"
  tier                = "Standard"
  component_version {
    storm = "1.1"
  }
  gateway {
    enabled  = true
    username = "acctestusrgw"
    password = "TerrAform123!"
  }
  storage_account {
    storage_container_id = azurerm_storage_container.test.id
    storage_account_key  = azurerm_storage_account.test.primary_access_key
    is_default           = true
  }
  roles {
    head_node {
      vm_size  = "Standard_D3_v2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
    worker_node {
      vm_size               = "Standard_D4_V2"
      username              = "acctestusrvm"
      password              = "AccTestvdSC4daf986!"
      target_instance_count = 2
    }
    zookeeper_node {
      vm_size  = "Standard_D3_v2"
      username = "acctestusrvm"
      password = "AccTestvdSC4daf986!"
    }
  }
  metastores {
    hive {
      server        = azurerm_sql_server.test.fully_qualified_domain_name
      database_name = azurerm_sql_database.hive.name
      username      = azurerm_sql_server.test.administrator_login
      password      = azurerm_sql_server.test.administrator_login_password
    }
  }
}
`, template, data.RandomInteger, data.RandomInteger)
}
