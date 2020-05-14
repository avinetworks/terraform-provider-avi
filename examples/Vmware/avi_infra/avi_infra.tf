provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_tenant     = "admin"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_network" "mgmt_network" {
  name = var.vm_mgmt_network
}


resource "avi_cloud" "vmware_cloud" {
  name         = "Vcenter"
  vtype        = "CLOUD_VCENTER"
  dhcp_enabled = true
  vcenter_configuration {
    datacenter = var.vsphere_datacenter
    management_network= data.avi_network.mgmt_network.id
    privilege= "WRITE_ACCESS"
    vcenter_url= var.vcenter_url
    password= var.vsphere_password
    username= var.vsphere_user
  }

  license_tier = "ENTERPRISE_18"
  license_type = "LIC_CORES"
  tenant_ref   = data.avi_tenant.default_tenant.id
}

resource "avi_serviceenginegroup" "vcenter_se_group" {
  name                         = "Default-Group"
  cloud_ref                    = avi_cloud.vmware_cloud.id
  vcenter_folder = "Gaurav Rastogi/se"
  vcenter_clusters {
    cluster_refs = [var.vm_cluster_name]
    include = true
  }
  max_se= 4
  buffer_se= 1
  se_name_prefix= var.project_name
  tenant_ref= data.avi_tenant.default_tenant.id
}

resource "avi_network" "mgmt_network" {
  name = var.vm_network
  exclude_discovered_subnets = true
  //configured_subnets {
  //  prefix = {
  //    ip_addr = {
  //      type = "V4"
  //      addr = var.vm_subnet
  //    }
  //    mask = 24
  //  }
    //vimgrnw_ref = "/api/vimgrnwruntime?name=vxw-dvs-34-virtualwire-71-sid-2140070-wdc-02-vc14-avi-dev064"
  //}
}
