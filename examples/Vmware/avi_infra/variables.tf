
// Avi Provider variables
variable "avi_username" {
  type = string
  default = "admin"
}

variable "avi_password" {
  type = string
  default = ""
}

variable "avi_controller" {
  type = string
  default = ""
}

variable "project_name" {
  type = string
  default = "grastogi"
}

variable "vcenter_url" {
  type    = string
  default = ""
}

variable "vsphere_datacenter" {
  type    = string
  default = "wdc-02-vc14"
}

variable "resource_pool" {
  type    = string
  default = ""
}

variable "content_library" {
  type    = string
  default = ""
}

variable "vm_datastore" {
  type    = string
  default = ""
}

variable "vm_name" {
  type    = string
  default = ""
}

variable "vm_network" {
  type    = string
  default = "vxw-dvs-34-virtualwire-71-sid-2140070-wdc-02-vc14-avi-dev064"
}

variable "vm_mgmt_network" {
  type = string
  default = "vxw-dvs-34-virtualwire-3-sid-2140002-wdc-02-vc14-avi-mgmt"
}

variable "vm_subnet" {
  type = string
  default = "100.64.72.0"
}

variable "vm_cluster_name" {
  type    = string
  default = "/api/vimgrclusterruntime?name=wdc-02-vc14c01"
}

variable "vm_folder" {
  type    = string
  default = ""
}

variable "vm_template" {
  type    = string
  default = ""
}

variable "vsphere_user" {
  type = string
  default = ""
}

variable "vsphere_password" {
  type = string
  default = ""
}

variable "vsphere_server" {
  type = string
  default = ""
}