/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add_networks_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmAddNetworksEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"all_seupgrade_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAllSeUpgradeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"anomaly_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAnomalyEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_agent_bd_vrf_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentBridgeDomainVrfChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_agent_generic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentGenericEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_agent_vs_network_error": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicAgentVsNetworkErrorSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"avg_uptime_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAvgUptimeChangeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"aws_asg_notif_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSASGNotifDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"aws_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"azure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bind_vs_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmBindVsSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bm_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBMSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bootup_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSeBootupFailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_cluster_vip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudClusterVipSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_dns_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudDnsUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_health_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudHealthSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudGenericSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_ip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudIpChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_parkintf_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVipParkingIntfSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_se_vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSeVmChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_sync_services_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSyncServicesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_tenant_del_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudTenantsDeletedSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_vip_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVipUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_vnic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudVnicChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_config_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterConfigFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_leader_failover_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterLeaderFailoverEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_add_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeAddEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_db_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeDbFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_remove_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeRemoveEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_shutdown_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeShutdownEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_node_started_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeStartedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_service_critical_failure_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceCriticalFailureEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_service_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_service_restored_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterServiceRestoredEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cluster_warm_reboot_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterWarmRebootEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cntlr_host_list_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraCntlrHostUnreachableListSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_action_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigActionDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_create_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigCreateDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigDeleteDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_password_change_request_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserPasswordChangeRequestSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_update_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUpdateDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_authrz_rule_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserAuthrzByRuleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_login_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_logout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserLogoutSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_user_not_authrz_rule_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigUserNotAuthrzByRuleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"container_cloud_setup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"container_cloud_sevice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContainerCloudServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cs_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudStackSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"delete_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmDeleteSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"disable_se_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDisableSeMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"disc_summary": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraDiscSummaryDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_sync_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSVsSyncInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"docker_ucp_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerUCPSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dos_attack_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosAttackEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gcp_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"glb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gs_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"host_unavail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostUnavailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthScoreDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSeIpFailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"license_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLicenseDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"license_expiry_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLicenseExpiryDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"marathon_service_port_conflict_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMarathonServicePortConflictSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metric_threshold_up_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricThresoldUpDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metrics_db_disk_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDbDiskEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mgmt_nw_change_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraMgmtNwChangeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"modify_networks_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmModifyNetworksEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network_subnet_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSubnetInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nw_subnet_clash_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSubnetClashSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nw_summarized_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSummarizedInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackClusterSetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_ip_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackIpChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_lbaudit_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackLbProvAuditCheckSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_lbplugin_op_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackLbPluginOpSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_se_vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackSeVmChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_sync_services_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackSyncServicesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_vnic_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenStackVnicChangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_deployment_failure_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentFailureInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_deployment_success_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolDeploymentSuccessInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_server_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraPoolServerDeleteDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rebalance_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rebalance_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceScaleinEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rebalance_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRebalanceScaleoutEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reboot_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmRebootSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"scheduler_action_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSchedulerActionDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMgrEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_dupip_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeDupipEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_gateway_heartbeat_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGatewayHeartbeatFailedDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_gateway_heartbeat_success_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGatewayHeartbeatSuccessDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_geo_db_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeGeoDbDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hb_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHBEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_gs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGSDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_gsgroup_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventGslbPoolDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_pool_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventPoolDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_hm_vs_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeHmEventVsDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ip_added_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpAddedEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ip_removed_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpRemovedEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ipfailure_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeIpfailureEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_persistence_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePersistenceEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_pool_lb_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePoolLbEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_thresh_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeThreshEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_version_check_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVersionCheckFailedEventSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_vnic_down_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicDownEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_vnic_tx_queue_stall_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVnicTxQueueStallEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"semigrate_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscale_failed_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleFailedInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscalein_complete_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleInCompleteInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscalein_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleInInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscaleout_complete_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleOutCompleteInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_autoscaleout_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScaleOutInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_disrupted_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeVsDisruptedEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_event_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeScaleinEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"seupgrade_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeScaleoutEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"spawn_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmSpawnSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_expire_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLExpireDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_export_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLExportDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_renew_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRenewDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_renew_failed_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRenewFailedDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switchover_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switchover_fail_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverFailEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"system_upgrade_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSystemUpgradeDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"unbind_vs_se_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRmUnbindVsSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vca_infra_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVCASetupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_connectivity_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterConnectivityStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterBadCredentialsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_disc_failure": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterDiscoveryFailureSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_obj_delete_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVcenterObjDeleteDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_dns_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDNSRegisterInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVinfraVmDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_awaitingse_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsAwaitingSeEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_error_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsErrorEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_fsm_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsFsmEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_initialplacement_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsInitialPlacementEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_migrate_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMigrateEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_pool_nw_fltr_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsPoolNwFilterEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_scalein_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleInEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_scaleout_details": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsScaleOutEventDetailsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
