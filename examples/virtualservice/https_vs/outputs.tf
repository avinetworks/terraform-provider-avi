output "virtual_service" {
  value = avi_virtualservice.https_vs
}
output "pool" {
  value = avi_pool.lb_pool
}
output "pool_group" {
  value = avi_poolgroup.poolgroup1
}