{
  "api_vips": [
    {
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "ip": "1.2.3.8"
    }
  ],
  "base_dns_domain": "example.com",
  "cluster_networks": [
    {
      "cidr": "10.128.0.0/14",
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "host_prefix": 23
    }
  ],
  "connectivity_majority_groups": "{\"majority_groups\":{\"1.2.3.0/24\":[\"4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\",\"5475e666-fa55-4afe-a20b-da5e990ae578\",\"78eb4750-4f8c-439a-b6f7-e859faea3506\"],\"IPv4\":[\"4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\",\"5475e666-fa55-4afe-a20b-da5e990ae578\",\"78eb4750-4f8c-439a-b6f7-e859faea3506\"],\"IPv6\":[]},\"l3_connected_addresses\":{\"4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\":[\"1.2.3.1\"],\"5475e666-fa55-4afe-a20b-da5e990ae578\":[\"1.2.3.3\"],\"78eb4750-4f8c-439a-b6f7-e859faea3506\":[\"1.2.3.2\"]}}",
  "control_plane_count": 2,
  "controller_logs_collected_at": "0001-01-01T00:00:00.000Z",
  "controller_logs_started_at": "0001-01-01T00:00:00.000Z",
  "cpu_architecture": "x86_64",
  "created_at": "2025-03-06T14:39:17.476789Z",
  "deleted_at": null,
  "disk_encryption": {
    "enable_on": "none",
    "mode": "tpmv2"
  },
  "email_domain": "Unknown",
  "enabled_host_count": 3,
  "feature_usage": "{\"Hyperthreading\":{\"data\":{\"hyperthreading_enabled\":\"all\"},\"id\":\"HYPERTHREADING\",\"name\":\"Hyperthreading\"},\"OVN network type\":{\"id\":\"OVN_NETWORK_TYPE\",\"name\":\"OVN network type\"}}",
  "high_availability_mode": "TNA",
  "host_networks": [
    {
      "cidr": "1.2.3.0/24",
      "host_ids": [
        "78eb4750-4f8c-439a-b6f7-e859faea3506",
        "4a78e1c4-caa2-48f6-a020-7d74dd7a3a46",
        "5475e666-fa55-4afe-a20b-da5e990ae578"
      ]
    }
  ],
  "hosts": [
    {
      "checked_in_at": "2025-03-06T14:39:24.992Z",
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "connectivity": "{\"remote_hosts\":[{\"host_id\":\"4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.1\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.1\",\"successful\":true}],\"mtu_report\":null},{\"host_id\":\"78eb4750-4f8c-439a-b6f7-e859faea3506\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.2\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.2\",\"successful\":true}],\"mtu_report\":null},{\"host_id\":\"5475e666-fa55-4afe-a20b-da5e990ae578\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.3\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.3\",\"successful\":true}],\"mtu_report\":null}]}",
      "created_at": "2025-03-06T14:39:24.971309Z",
      "deleted_at": null,
      "discovery_agent_version": "quay.io/edge-infrastructure/assisted-installer-agent:latest",
      "disks_info": "{\"wwn-0x2222222222222222222222\":{\"disk_speed\":{\"speed_ms\":10,\"tested\":true},\"path\":\"wwn-0x2222222222222222222222\"}}",
      "domain_name_resolutions": "{\"resolutions\":[{\"cnames\":null,\"domain_name\":\"api.test-cluster.example.com\",\"ipv4_addresses\":[\"1.2.3.4/24\"],\"ipv6_addresses\":[\"1001:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"api-int.test-cluster.example.com\",\"ipv4_addresses\":[\"4.5.6.7/24\"],\"ipv6_addresses\":[\"1002:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"console-openshift-console.apps.test-cluster.example.com\",\"ipv4_addresses\":[\"7.8.9.10/24\"],\"ipv6_addresses\":[\"1003:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"quay.io\",\"ipv4_addresses\":[\"7.8.9.11/24\"],\"ipv6_addresses\":[\"1003:db8::11/120\"]},{\"cnames\":null,\"domain_name\":\"validateNoWildcardDNS.test-cluster.example.com\",\"ipv4_addresses\":[],\"ipv6_addresses\":[]},{\"cnames\":null,\"domain_name\":\"validateNoWildcardDNS.test-cluster.example.com.\",\"ipv4_addresses\":[],\"ipv6_addresses\":[]}]}",
      "href": "/api/assisted-install/v2/infra-envs/17227498-f9a3-4353-96f8-1b7087f5f6b9/hosts/78eb4750-4f8c-439a-b6f7-e859faea3506",
      "id": "78eb4750-4f8c-439a-b6f7-e859faea3506",
      "images_status": "{\"image\":{\"download_rate\":33.3,\"name\":\"image\",\"result\":\"success\",\"size_bytes\":333000000,\"time\":10}}",
      "infra_env_id": "17227498-f9a3-4353-96f8-1b7087f5f6b9",
      "installation_disk_id": "wwn-0x2222222222222222222222",
      "installation_disk_path": "/dev/sdb",
      "inventory": "{\"cpu\":{\"architecture\":\"x86_64\",\"count\":16,\"flags\":null},\"disks\":[{\"by_id\":\"wwn-0x1111111111111111111111\",\"drive_type\":\"SSD\",\"id\":\"wwn-0x1111111111111111111111\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"loop0\",\"size_bytes\":128849018880},{\"by_id\":\"wwn-0x2222222222222222222222\",\"drive_type\":\"HDD\",\"id\":\"wwn-0x2222222222222222222222\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"sdb\",\"size_bytes\":128849018880}],\"gpus\":null,\"hostname\":\"h1\",\"interfaces\":[{\"flags\":null,\"ipv4_addresses\":[\"1.2.3.2/24\"],\"ipv6_addresses\":null,\"mac_address\":\"e6:53:3d:a7:77:b4\",\"type\":\"physical\"}],\"memory\":{\"physical_bytes\":34359738368,\"usable_bytes\":34359738368},\"routes\":[{\"destination\":\"0.0.0.0\",\"family\":2,\"gateway\":\"1.2.3.10\",\"interface\":\"eth0\",\"metric\":600}],\"system_vendor\":{\"manufacturer\":\"manu\",\"product_name\":\"prod\",\"serial_number\":\"3534\"},\"tpm_version\":\"2.0\"}",
      "kind": "Host",
      "logs_collected_at": "0001-01-01T00:00:00.000Z",
      "logs_started_at": "0001-01-01T00:00:00.000Z",
      "ntp_sources": "[{\"source_name\":\"clock.dummy.test\",\"source_state\":\"synced\"}]",
      "progress": {
        "stage_started_at": "0001-01-01T00:00:00.000Z",
        "stage_updated_at": "0001-01-01T00:00:00.000Z"
      },
      "progress_stages": [
        "Starting installation",
        "Installing",
        "Writing image to disk",
        "Rebooting",
        "Configuring",
        "Joined",
        "Done"
      ],
      "registered_at": "2025-03-06T14:39:24.970Z",
      "requested_hostname": "h1",
      "role": "master",
      "stage_started_at": "0001-01-01T00:00:00.000Z",
      "stage_updated_at": "0001-01-01T00:00:00.000Z",
      "status": "disconnected",
      "status_info": "Host has stopped communicating with the installation service",
      "status_updated_at": "2025-03-06T14:42:31.921Z",
      "suggested_role": "master",
      "timestamp": 1741271964,
      "updated_at": "2025-03-06T14:42:31.922292Z",
      "user_name": "admin",
      "validations_info": "{\"hardware\":[{\"id\":\"has-inventory\",\"status\":\"success\",\"message\":\"Valid inventory exists for the host\"},{\"id\":\"has-min-cpu-cores\",\"status\":\"success\",\"message\":\"Sufficient CPU cores\"},{\"id\":\"has-min-memory\",\"status\":\"success\",\"message\":\"Sufficient minimum RAM\"},{\"id\":\"has-min-valid-disks\",\"status\":\"success\",\"message\":\"Sufficient disk capacity\"},{\"id\":\"has-cpu-cores-for-role\",\"status\":\"success\",\"message\":\"Sufficient CPU cores for role master\"},{\"id\":\"has-memory-for-role\",\"status\":\"success\",\"message\":\"Sufficient RAM for role master\"},{\"id\":\"hostname-unique\",\"status\":\"success\",\"message\":\"Hostname h1 is unique in cluster\"},{\"id\":\"hostname-valid\",\"status\":\"success\",\"message\":\"Hostname h1 is allowed\"},{\"id\":\"sufficient-installation-disk-speed\",\"status\":\"success\",\"message\":\"Speed of installation disk is sufficient\"},{\"id\":\"compatible-with-cluster-platform\",\"status\":\"success\",\"message\":\"Host is compatible with cluster platform baremetal\"},{\"id\":\"compatible-agent\",\"status\":\"success\",\"message\":\"Host agent is compatible with the service\"},{\"id\":\"no-skip-installation-disk\",\"status\":\"success\",\"message\":\"No request to skip formatting of the installation disk\"},{\"id\":\"no-skip-missing-disk\",\"status\":\"success\",\"message\":\"All disks that have skipped formatting are present in the host inventory\"}],\"network\":[{\"id\":\"connected\",\"status\":\"failure\",\"message\":\"Host is disconnected\"},{\"id\":\"media-connected\",\"status\":\"success\",\"message\":\"Media device is connected\"},{\"id\":\"machine-cidr-defined\",\"status\":\"success\",\"message\":\"Machine Network CIDR is defined\"},{\"id\":\"belongs-to-machine-cidr\",\"status\":\"success\",\"message\":\"Host belongs to all machine network CIDRs\"},{\"id\":\"belongs-to-majority-group\",\"status\":\"success\",\"message\":\"Host has connectivity to the majority of hosts in the cluster\"},{\"id\":\"mtu-valid\",\"status\":\"success\",\"message\":\"MTU is ok\"},{\"id\":\"valid-platform-network-settings\",\"status\":\"success\",\"message\":\"Platform prod is allowed\"},{\"id\":\"ntp-synced\",\"status\":\"success\",\"message\":\"Host NTP is synced\"},{\"id\":\"time-synced-between-host-and-service\",\"status\":\"success\",\"message\":\"Host clock is synchronized with service\"},{\"id\":\"container-images-available\",\"status\":\"success\",\"message\":\"All required container images were either pulled successfully or no attempt was made to pull them\"},{\"id\":\"sufficient-network-latency-requirement-for-role\",\"status\":\"success\",\"message\":\"Network latency requirement has been satisfied.\"},{\"id\":\"sufficient-packet-loss-requirement-for-role\",\"status\":\"success\",\"message\":\"Packet loss requirement has been satisfied.\"},{\"id\":\"has-default-route\",\"status\":\"success\",\"message\":\"Host has been configured with at least one default route.\"},{\"id\":\"api-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the api.test-cluster.example.com domain was successful or not required\"},{\"id\":\"api-int-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the api-int.test-cluster.example.com domain was successful or not required\"},{\"id\":\"apps-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the *.apps.test-cluster.example.com domain was successful or not required\"},{\"id\":\"dns-wildcard-not-configured\",\"status\":\"success\",\"message\":\"DNS wildcard check was successful\"},{\"id\":\"non-overlapping-subnets\",\"status\":\"success\",\"message\":\"Host subnets are not overlapping\"},{\"id\":\"no-ip-collisions-in-network\",\"status\":\"success\",\"message\":\"No IP collisions were detected by host 78eb4750-4f8c-439a-b6f7-e859faea3506\"}],\"operators\":[{\"id\":\"authorino-requirements-satisfied\",\"status\":\"success\",\"message\":\"authorino is disabled\"},{\"id\":\"cnv-requirements-satisfied\",\"status\":\"success\",\"message\":\"cnv is disabled\"},{\"id\":\"lso-requirements-satisfied\",\"status\":\"success\",\"message\":\"lso is disabled\"},{\"id\":\"lvm-requirements-satisfied\",\"status\":\"success\",\"message\":\"lvm is disabled\"},{\"id\":\"mce-requirements-satisfied\",\"status\":\"success\",\"message\":\"mce is disabled\"},{\"id\":\"mtv-requirements-satisfied\",\"status\":\"success\",\"message\":\"mtv is disabled\"},{\"id\":\"nmstate-requirements-satisfied\",\"status\":\"success\",\"message\":\"nmstate is disabled\"},{\"id\":\"node-feature-discovery-requirements-satisfied\",\"status\":\"success\",\"message\":\"node-feature-discovery is disabled\"},{\"id\":\"nvidia-gpu-requirements-satisfied\",\"status\":\"success\",\"message\":\"nvidia-gpu is disabled\"},{\"id\":\"odf-requirements-satisfied\",\"status\":\"success\",\"message\":\"odf is disabled\"},{\"id\":\"openshift-ai-requirements-satisfied\",\"status\":\"success\",\"message\":\"openshift-ai is disabled\"},{\"id\":\"osc-requirements-satisfied\",\"status\":\"success\",\"message\":\"osc is disabled\"},{\"id\":\"pipelines-requirements-satisfied\",\"status\":\"success\",\"message\":\"pipelines is disabled\"},{\"id\":\"serverless-requirements-satisfied\",\"status\":\"success\",\"message\":\"serverless is disabled\"},{\"id\":\"servicemesh-requirements-satisfied\",\"status\":\"success\",\"message\":\"servicemesh is disabled\"}]}"
    },
    {
      "bootstrap": true,
      "checked_in_at": "2025-03-06T14:39:24.808Z",
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "connectivity": "{\"remote_hosts\":[{\"host_id\":\"4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.1\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.1\",\"successful\":true}],\"mtu_report\":null},{\"host_id\":\"78eb4750-4f8c-439a-b6f7-e859faea3506\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.2\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.2\",\"successful\":true}],\"mtu_report\":null},{\"host_id\":\"5475e666-fa55-4afe-a20b-da5e990ae578\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.3\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.3\",\"successful\":true}],\"mtu_report\":null}]}",
      "created_at": "2025-03-06T14:39:24.786585Z",
      "deleted_at": null,
      "discovery_agent_version": "quay.io/edge-infrastructure/assisted-installer-agent:latest",
      "disks_info": "{\"wwn-0x2222222222222222222222\":{\"disk_speed\":{\"speed_ms\":10,\"tested\":true},\"path\":\"wwn-0x2222222222222222222222\"}}",
      "domain_name_resolutions": "{\"resolutions\":[{\"cnames\":null,\"domain_name\":\"api.test-cluster.example.com\",\"ipv4_addresses\":[\"1.2.3.4/24\"],\"ipv6_addresses\":[\"1001:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"api-int.test-cluster.example.com\",\"ipv4_addresses\":[\"4.5.6.7/24\"],\"ipv6_addresses\":[\"1002:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"console-openshift-console.apps.test-cluster.example.com\",\"ipv4_addresses\":[\"7.8.9.10/24\"],\"ipv6_addresses\":[\"1003:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"quay.io\",\"ipv4_addresses\":[\"7.8.9.11/24\"],\"ipv6_addresses\":[\"1003:db8::11/120\"]},{\"cnames\":null,\"domain_name\":\"validateNoWildcardDNS.test-cluster.example.com\",\"ipv4_addresses\":[],\"ipv6_addresses\":[]},{\"cnames\":null,\"domain_name\":\"validateNoWildcardDNS.test-cluster.example.com.\",\"ipv4_addresses\":[],\"ipv6_addresses\":[]}]}",
      "href": "/api/assisted-install/v2/infra-envs/17227498-f9a3-4353-96f8-1b7087f5f6b9/hosts/4a78e1c4-caa2-48f6-a020-7d74dd7a3a46",
      "id": "4a78e1c4-caa2-48f6-a020-7d74dd7a3a46",
      "images_status": "{\"image\":{\"download_rate\":33.3,\"name\":\"image\",\"result\":\"success\",\"size_bytes\":333000000,\"time\":10}}",
      "infra_env_id": "17227498-f9a3-4353-96f8-1b7087f5f6b9",
      "installation_disk_id": "wwn-0x2222222222222222222222",
      "installation_disk_path": "/dev/sdb",
      "inventory": "{\"cpu\":{\"architecture\":\"x86_64\",\"count\":16,\"flags\":null},\"disks\":[{\"by_id\":\"wwn-0x1111111111111111111111\",\"drive_type\":\"SSD\",\"id\":\"wwn-0x1111111111111111111111\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"loop0\",\"size_bytes\":128849018880},{\"by_id\":\"wwn-0x2222222222222222222222\",\"drive_type\":\"HDD\",\"id\":\"wwn-0x2222222222222222222222\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"sdb\",\"size_bytes\":128849018880}],\"gpus\":null,\"hostname\":\"h0\",\"interfaces\":[{\"flags\":null,\"ipv4_addresses\":[\"1.2.3.1/24\"],\"ipv6_addresses\":null,\"mac_address\":\"e6:53:3d:a7:77:b4\",\"type\":\"physical\"}],\"memory\":{\"physical_bytes\":34359738368,\"usable_bytes\":34359738368},\"routes\":[{\"destination\":\"0.0.0.0\",\"family\":2,\"gateway\":\"1.2.3.10\",\"interface\":\"eth0\",\"metric\":600}],\"system_vendor\":{\"manufacturer\":\"manu\",\"product_name\":\"prod\",\"serial_number\":\"3534\"},\"tpm_version\":\"2.0\"}",
      "kind": "Host",
      "logs_collected_at": "0001-01-01T00:00:00.000Z",
      "logs_started_at": "0001-01-01T00:00:00.000Z",
      "ntp_sources": "[{\"source_name\":\"clock.dummy.test\",\"source_state\":\"synced\"}]",
      "progress": {
        "stage_started_at": "0001-01-01T00:00:00.000Z",
        "stage_updated_at": "0001-01-01T00:00:00.000Z"
      },
      "progress_stages": [
        "Starting installation",
        "Installing",
        "Writing image to disk",
        "Waiting for control plane",
        "Waiting for bootkube",
        "Waiting for controller",
        "Rebooting",
        "Configuring",
        "Joined",
        "Done"
      ],
      "registered_at": "2025-03-06T14:39:24.784Z",
      "requested_hostname": "h0",
      "role": "master",
      "stage_started_at": "0001-01-01T00:00:00.000Z",
      "stage_updated_at": "0001-01-01T00:00:00.000Z",
      "status": "known",
      "status_info": "Host is ready to be installed",
      "status_updated_at": "2025-03-06T14:40:39.913Z",
      "suggested_role": "master",
      "timestamp": 1741271964,
      "updated_at": "2025-03-06T14:40:39.914101Z",
      "user_name": "admin",
      "validations_info": "{\"hardware\":[{\"id\":\"has-inventory\",\"status\":\"success\",\"message\":\"Valid inventory exists for the host\"},{\"id\":\"has-min-cpu-cores\",\"status\":\"success\",\"message\":\"Sufficient CPU cores\"},{\"id\":\"has-min-memory\",\"status\":\"success\",\"message\":\"Sufficient minimum RAM\"},{\"id\":\"has-min-valid-disks\",\"status\":\"success\",\"message\":\"Sufficient disk capacity\"},{\"id\":\"has-cpu-cores-for-role\",\"status\":\"success\",\"message\":\"Sufficient CPU cores for role master\"},{\"id\":\"has-memory-for-role\",\"status\":\"success\",\"message\":\"Sufficient RAM for role master\"},{\"id\":\"hostname-unique\",\"status\":\"success\",\"message\":\"Hostname h0 is unique in cluster\"},{\"id\":\"hostname-valid\",\"status\":\"success\",\"message\":\"Hostname h0 is allowed\"},{\"id\":\"sufficient-installation-disk-speed\",\"status\":\"success\",\"message\":\"Speed of installation disk is sufficient\"},{\"id\":\"compatible-with-cluster-platform\",\"status\":\"success\",\"message\":\"Host is compatible with cluster platform baremetal\"},{\"id\":\"compatible-agent\",\"status\":\"success\",\"message\":\"Host agent is compatible with the service\"},{\"id\":\"no-skip-installation-disk\",\"status\":\"success\",\"message\":\"No request to skip formatting of the installation disk\"},{\"id\":\"no-skip-missing-disk\",\"status\":\"success\",\"message\":\"All disks that have skipped formatting are present in the host inventory\"}],\"network\":[{\"id\":\"connected\",\"status\":\"success\",\"message\":\"Host is connected\"},{\"id\":\"media-connected\",\"status\":\"success\",\"message\":\"Media device is connected\"},{\"id\":\"machine-cidr-defined\",\"status\":\"success\",\"message\":\"Machine Network CIDR is defined\"},{\"id\":\"belongs-to-machine-cidr\",\"status\":\"success\",\"message\":\"Host belongs to all machine network CIDRs\"},{\"id\":\"belongs-to-majority-group\",\"status\":\"success\",\"message\":\"Host has connectivity to the majority of hosts in the cluster\"},{\"id\":\"mtu-valid\",\"status\":\"success\",\"message\":\"MTU is ok\"},{\"id\":\"valid-platform-network-settings\",\"status\":\"success\",\"message\":\"Platform prod is allowed\"},{\"id\":\"ntp-synced\",\"status\":\"success\",\"message\":\"Host NTP is synced\"},{\"id\":\"time-synced-between-host-and-service\",\"status\":\"success\",\"message\":\"Host clock is synchronized with service\"},{\"id\":\"container-images-available\",\"status\":\"success\",\"message\":\"All required container images were either pulled successfully or no attempt was made to pull them\"},{\"id\":\"sufficient-network-latency-requirement-for-role\",\"status\":\"success\",\"message\":\"Network latency requirement has been satisfied.\"},{\"id\":\"sufficient-packet-loss-requirement-for-role\",\"status\":\"success\",\"message\":\"Packet loss requirement has been satisfied.\"},{\"id\":\"has-default-route\",\"status\":\"success\",\"message\":\"Host has been configured with at least one default route.\"},{\"id\":\"api-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the api.test-cluster.example.com domain was successful or not required\"},{\"id\":\"api-int-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the api-int.test-cluster.example.com domain was successful or not required\"},{\"id\":\"apps-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the *.apps.test-cluster.example.com domain was successful or not required\"},{\"id\":\"dns-wildcard-not-configured\",\"status\":\"success\",\"message\":\"DNS wildcard check was successful\"},{\"id\":\"non-overlapping-subnets\",\"status\":\"success\",\"message\":\"Host subnets are not overlapping\"},{\"id\":\"no-ip-collisions-in-network\",\"status\":\"success\",\"message\":\"No IP collisions were detected by host 4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\"}],\"operators\":[{\"id\":\"authorino-requirements-satisfied\",\"status\":\"success\",\"message\":\"authorino is disabled\"},{\"id\":\"cnv-requirements-satisfied\",\"status\":\"success\",\"message\":\"cnv is disabled\"},{\"id\":\"lso-requirements-satisfied\",\"status\":\"success\",\"message\":\"lso is disabled\"},{\"id\":\"lvm-requirements-satisfied\",\"status\":\"success\",\"message\":\"lvm is disabled\"},{\"id\":\"mce-requirements-satisfied\",\"status\":\"success\",\"message\":\"mce is disabled\"},{\"id\":\"mtv-requirements-satisfied\",\"status\":\"success\",\"message\":\"mtv is disabled\"},{\"id\":\"nmstate-requirements-satisfied\",\"status\":\"success\",\"message\":\"nmstate is disabled\"},{\"id\":\"node-feature-discovery-requirements-satisfied\",\"status\":\"success\",\"message\":\"node-feature-discovery is disabled\"},{\"id\":\"nvidia-gpu-requirements-satisfied\",\"status\":\"success\",\"message\":\"nvidia-gpu is disabled\"},{\"id\":\"odf-requirements-satisfied\",\"status\":\"success\",\"message\":\"odf is disabled\"},{\"id\":\"openshift-ai-requirements-satisfied\",\"status\":\"success\",\"message\":\"openshift-ai is disabled\"},{\"id\":\"osc-requirements-satisfied\",\"status\":\"success\",\"message\":\"osc is disabled\"},{\"id\":\"pipelines-requirements-satisfied\",\"status\":\"success\",\"message\":\"pipelines is disabled\"},{\"id\":\"serverless-requirements-satisfied\",\"status\":\"success\",\"message\":\"serverless is disabled\"},{\"id\":\"servicemesh-requirements-satisfied\",\"status\":\"success\",\"message\":\"servicemesh is disabled\"}]}"
    },
    {
      "checked_in_at": "2025-03-06T14:39:25.200Z",
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "connectivity": "{\"remote_hosts\":[{\"host_id\":\"4a78e1c4-caa2-48f6-a020-7d74dd7a3a46\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.1\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.1\",\"successful\":true}],\"mtu_report\":null},{\"host_id\":\"78eb4750-4f8c-439a-b6f7-e859faea3506\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.2\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.2\",\"successful\":true}],\"mtu_report\":null},{\"host_id\":\"5475e666-fa55-4afe-a20b-da5e990ae578\",\"l2_connectivity\":[{\"remote_ip_address\":\"1.2.3.3\",\"successful\":true}],\"l3_connectivity\":[{\"remote_ip_address\":\"1.2.3.3\",\"successful\":true}],\"mtu_report\":null}]}",
      "created_at": "2025-03-06T14:39:25.178573Z",
      "deleted_at": null,
      "discovery_agent_version": "quay.io/edge-infrastructure/assisted-installer-agent:latest",
      "disks_info": "{\"wwn-0x2222222222222222222222\":{\"disk_speed\":{\"speed_ms\":10,\"tested\":true},\"path\":\"wwn-0x2222222222222222222222\"}}",
      "domain_name_resolutions": "{\"resolutions\":[{\"cnames\":null,\"domain_name\":\"api.test-cluster.example.com\",\"ipv4_addresses\":[\"1.2.3.4/24\"],\"ipv6_addresses\":[\"1001:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"api-int.test-cluster.example.com\",\"ipv4_addresses\":[\"4.5.6.7/24\"],\"ipv6_addresses\":[\"1002:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"console-openshift-console.apps.test-cluster.example.com\",\"ipv4_addresses\":[\"7.8.9.10/24\"],\"ipv6_addresses\":[\"1003:db8::10/120\"]},{\"cnames\":null,\"domain_name\":\"quay.io\",\"ipv4_addresses\":[\"7.8.9.11/24\"],\"ipv6_addresses\":[\"1003:db8::11/120\"]},{\"cnames\":null,\"domain_name\":\"validateNoWildcardDNS.test-cluster.example.com\",\"ipv4_addresses\":[],\"ipv6_addresses\":[]},{\"cnames\":null,\"domain_name\":\"validateNoWildcardDNS.test-cluster.example.com.\",\"ipv4_addresses\":[],\"ipv6_addresses\":[]}]}",
      "href": "/api/assisted-install/v2/infra-envs/17227498-f9a3-4353-96f8-1b7087f5f6b9/hosts/5475e666-fa55-4afe-a20b-da5e990ae578",
      "id": "5475e666-fa55-4afe-a20b-da5e990ae578",
      "images_status": "{\"image\":{\"download_rate\":33.3,\"name\":\"image\",\"result\":\"success\",\"size_bytes\":333000000,\"time\":10}}",
      "infra_env_id": "17227498-f9a3-4353-96f8-1b7087f5f6b9",
      "installation_disk_id": "wwn-0x2222222222222222222222",
      "installation_disk_path": "/dev/sdb",
      "inventory": "{\"cpu\":{\"architecture\":\"x86_64\",\"count\":16,\"flags\":null},\"disks\":[{\"by_id\":\"wwn-0x1111111111111111111111\",\"drive_type\":\"SSD\",\"id\":\"wwn-0x1111111111111111111111\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"loop0\",\"size_bytes\":128849018880},{\"by_id\":\"wwn-0x2222222222222222222222\",\"drive_type\":\"HDD\",\"id\":\"wwn-0x2222222222222222222222\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"sdb\",\"size_bytes\":128849018880}],\"gpus\":null,\"hostname\":\"h2\",\"interfaces\":[{\"flags\":null,\"ipv4_addresses\":[\"1.2.3.3/24\"],\"ipv6_addresses\":null,\"mac_address\":\"e6:53:3d:a7:77:b4\",\"type\":\"physical\"}],\"memory\":{\"physical_bytes\":34359738368,\"usable_bytes\":34359738368},\"routes\":[{\"destination\":\"0.0.0.0\",\"family\":2,\"gateway\":\"1.2.3.10\",\"interface\":\"eth0\",\"metric\":600}],\"system_vendor\":{\"manufacturer\":\"manu\",\"product_name\":\"prod\",\"serial_number\":\"3534\"},\"tpm_version\":\"2.0\"}",
      "kind": "Host",
      "logs_collected_at": "0001-01-01T00:00:00.000Z",
      "logs_started_at": "0001-01-01T00:00:00.000Z",
      "ntp_sources": "[{\"source_name\":\"clock.dummy.test\",\"source_state\":\"synced\"}]",
      "progress": {
        "stage_started_at": "0001-01-01T00:00:00.000Z",
        "stage_updated_at": "0001-01-01T00:00:00.000Z"
      },
      "progress_stages": [],
      "registered_at": "2025-03-06T14:39:25.177Z",
      "requested_hostname": "h2",
      "role": "arbiter",
      "stage_started_at": "0001-01-01T00:00:00.000Z",
      "stage_updated_at": "0001-01-01T00:00:00.000Z",
      "status": "disconnected",
      "status_info": "Host has stopped communicating with the installation service",
      "status_updated_at": "2025-03-06T14:42:31.999Z",
      "suggested_role": "arbiter",
      "timestamp": 1741271965,
      "updated_at": "2025-03-06T14:42:31.999708Z",
      "user_name": "admin",
      "validations_info": "{\"hardware\":[{\"id\":\"has-inventory\",\"status\":\"success\",\"message\":\"Valid inventory exists for the host\"},{\"id\":\"has-min-cpu-cores\",\"status\":\"success\",\"message\":\"Sufficient CPU cores\"},{\"id\":\"has-min-memory\",\"status\":\"success\",\"message\":\"Sufficient minimum RAM\"},{\"id\":\"has-min-valid-disks\",\"status\":\"success\",\"message\":\"Sufficient disk capacity\"},{\"id\":\"has-cpu-cores-for-role\",\"status\":\"success\",\"message\":\"Sufficient CPU cores for role arbiter\"},{\"id\":\"has-memory-for-role\",\"status\":\"success\",\"message\":\"Sufficient RAM for role arbiter\"},{\"id\":\"hostname-unique\",\"status\":\"success\",\"message\":\"Hostname h2 is unique in cluster\"},{\"id\":\"hostname-valid\",\"status\":\"success\",\"message\":\"Hostname h2 is allowed\"},{\"id\":\"sufficient-installation-disk-speed\",\"status\":\"success\",\"message\":\"Speed of installation disk is sufficient\"},{\"id\":\"compatible-with-cluster-platform\",\"status\":\"success\",\"message\":\"Host is compatible with cluster platform baremetal\"},{\"id\":\"compatible-agent\",\"status\":\"success\",\"message\":\"Host agent is compatible with the service\"},{\"id\":\"no-skip-installation-disk\",\"status\":\"success\",\"message\":\"No request to skip formatting of the installation disk\"},{\"id\":\"no-skip-missing-disk\",\"status\":\"success\",\"message\":\"All disks that have skipped formatting are present in the host inventory\"}],\"network\":[{\"id\":\"connected\",\"status\":\"failure\",\"message\":\"Host is disconnected\"},{\"id\":\"media-connected\",\"status\":\"success\",\"message\":\"Media device is connected\"},{\"id\":\"machine-cidr-defined\",\"status\":\"success\",\"message\":\"Machine Network CIDR is defined\"},{\"id\":\"belongs-to-machine-cidr\",\"status\":\"success\",\"message\":\"Host belongs to all machine network CIDRs\"},{\"id\":\"belongs-to-majority-group\",\"status\":\"success\",\"message\":\"Host has connectivity to the majority of hosts in the cluster\"},{\"id\":\"mtu-valid\",\"status\":\"success\",\"message\":\"MTU is ok\"},{\"id\":\"valid-platform-network-settings\",\"status\":\"success\",\"message\":\"Platform prod is allowed\"},{\"id\":\"ntp-synced\",\"status\":\"success\",\"message\":\"Host NTP is synced\"},{\"id\":\"time-synced-between-host-and-service\",\"status\":\"success\",\"message\":\"Host clock is synchronized with service\"},{\"id\":\"container-images-available\",\"status\":\"success\",\"message\":\"All required container images were either pulled successfully or no attempt was made to pull them\"},{\"id\":\"sufficient-network-latency-requirement-for-role\",\"status\":\"success\",\"message\":\"Network latency requirement has been satisfied.\"},{\"id\":\"sufficient-packet-loss-requirement-for-role\",\"status\":\"success\",\"message\":\"Packet loss requirement has been satisfied.\"},{\"id\":\"has-default-route\",\"status\":\"success\",\"message\":\"Host has been configured with at least one default route.\"},{\"id\":\"api-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the api.test-cluster.example.com domain was successful or not required\"},{\"id\":\"api-int-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the api-int.test-cluster.example.com domain was successful or not required\"},{\"id\":\"apps-domain-name-resolved-correctly\",\"status\":\"success\",\"message\":\"Domain name resolution for the *.apps.test-cluster.example.com domain was successful or not required\"},{\"id\":\"dns-wildcard-not-configured\",\"status\":\"success\",\"message\":\"DNS wildcard check was successful\"},{\"id\":\"non-overlapping-subnets\",\"status\":\"success\",\"message\":\"Host subnets are not overlapping\"},{\"id\":\"no-ip-collisions-in-network\",\"status\":\"success\",\"message\":\"No IP collisions were detected by host 5475e666-fa55-4afe-a20b-da5e990ae578\"}],\"operators\":[{\"id\":\"authorino-requirements-satisfied\",\"status\":\"success\",\"message\":\"authorino is disabled\"},{\"id\":\"cnv-requirements-satisfied\",\"status\":\"success\",\"message\":\"cnv is disabled\"},{\"id\":\"lso-requirements-satisfied\",\"status\":\"success\",\"message\":\"lso is disabled\"},{\"id\":\"lvm-requirements-satisfied\",\"status\":\"success\",\"message\":\"lvm is disabled\"},{\"id\":\"mce-requirements-satisfied\",\"status\":\"success\",\"message\":\"mce is disabled\"},{\"id\":\"mtv-requirements-satisfied\",\"status\":\"success\",\"message\":\"mtv is disabled\"},{\"id\":\"nmstate-requirements-satisfied\",\"status\":\"success\",\"message\":\"nmstate is disabled\"},{\"id\":\"node-feature-discovery-requirements-satisfied\",\"status\":\"success\",\"message\":\"node-feature-discovery is disabled\"},{\"id\":\"nvidia-gpu-requirements-satisfied\",\"status\":\"success\",\"message\":\"nvidia-gpu is disabled\"},{\"id\":\"odf-requirements-satisfied\",\"status\":\"success\",\"message\":\"odf is disabled\"},{\"id\":\"openshift-ai-requirements-satisfied\",\"status\":\"success\",\"message\":\"openshift-ai is disabled\"},{\"id\":\"osc-requirements-satisfied\",\"status\":\"success\",\"message\":\"osc is disabled\"},{\"id\":\"pipelines-requirements-satisfied\",\"status\":\"success\",\"message\":\"pipelines is disabled\"},{\"id\":\"serverless-requirements-satisfied\",\"status\":\"success\",\"message\":\"serverless is disabled\"},{\"id\":\"servicemesh-requirements-satisfied\",\"status\":\"success\",\"message\":\"servicemesh is disabled\"}]}"
    }
  ],
  "href": "/api/assisted-install/v2/clusters/4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
  "hyperthreading": "all",
  "id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
  "ignition_endpoint": {},
  "image_info": {
    "created_at": "0001-01-01T00:00:00Z",
    "expires_at": "0001-01-01T00:00:00.000Z"
  },
  "ingress_vips": [
    {
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "ip": "1.2.3.9"
    }
  ],
  "install_completed_at": "0001-01-01T00:00:00.000Z",
  "install_started_at": "2025-03-06T14:39:27.825Z",
  "ip_collisions": "{}",
  "kind": "Cluster",
  "last-installation-preparation": {
    "reason": "failed generating install config for cluster 4a6c4e2a-ce9a-430d-910a-4c94e4c41e36: error running openshift-install manifests,  level=warning msg=Found override for release image (quay.io/rh-ee-gravid/ocp-release:latest). Release Image Architecture is unknown\nlevel=error msg=failed to fetch Master Machines: failed to load asset \"Install Config\": failed to create install config: invalid \"install-config.yaml\" file: [platform.baremetal.hosts[1].BootMACAddress: Duplicate value: \"e6:53:3d:a7:77:b4\", platform.baremetal.hosts[2].BootMACAddress: Duplicate value: \"e6:53:3d:a7:77:b4\"]\n: exit status 3",
    "status": "failed"
  },
  "load_balancer": {
    "type": "cluster-managed"
  },
  "machine_networks": [
    {
      "cidr": "1.2.3.0/24",
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36"
    }
  ],
  "monitored_operators": [
    {
      "bundles": null,
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36",
      "name": "console",
      "operator_type": "builtin",
      "status_updated_at": "0001-01-01T00:00:00.000Z",
      "timeout_seconds": 3600
    }
  ],
  "name": "test-cluster",
  "network_type": "OVNKubernetes",
  "ocp_release_image": "quay.io/rh-ee-gravid/ocp-release:latest",
  "openshift_version": "4.19.0-0.nightly-2025-03-06-121124",
  "org_soft_timeouts_enabled": true,
  "platform": {
    "external": {},
    "type": "baremetal"
  },
  "progress": {
    "finalizing_stage_started_at": "0001-01-01T00:00:00.000Z"
  },
  "pull_secret_set": true,
  "ready_host_count": 1,
  "schedulable_masters": false,
  "schedulable_masters_forced_true": true,
  "service_networks": [
    {
      "cidr": "172.30.0.0/16",
      "cluster_id": "4a6c4e2a-ce9a-430d-910a-4c94e4c41e36"
    }
  ],
  "ssh_public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC50TuHS7aYci+U+5PLe/aW/I6maBi9PBDucLje6C6gtArfjy7udWA1DCSIQd+DkHhi57/s+PmvEjzfAfzqo+L+/8/O2l2seR1pPhHDxMR/rSyo/6rZP6KIL8HwFqXHHpDUM4tLXdgwKAe1LxBevLt/yNl8kOiHJESUSl+2QSf8z4SIbo/frDD8OwOvtfKBEG4WCb8zEsEuIPNF/Vo/UxPtS9pPTecEsWKDHR67yFjjamoyLvAzMAJotYgyMoxm8PTyCgEzHk3s3S4iO956d6KVOEJVXnTVhAxrtLuubjskd7N4hVN7h2s4Z584wYLKYhrIBL0EViihOMzY4mH3YE4KZusfIx6oMcggKX9b3NHm0la7cj2zg0r6zjUn6ZCP4gXM99e5q4auc0OEfoSfQwofGi3WmxkG3tEozCB8Zz0wGbi2CzR8zlcF+BNV5I2LESlLzjPY5B4dvv5zjxsYoz94p3rUhKnnPM2zTx1kkilDK5C5fC1k9l/I/r5Qk4ebLQU= oscohen@localhost.localdomain",
  "status": "insufficient",
  "status_info": "Cluster is not ready for install",
  "status_updated_at": "2025-03-06T14:42:37.922Z",
  "total_host_count": 3,
  "updated_at": "2025-03-06T14:42:37.923122Z",
  "user_managed_networking": false,
  "user_name": "admin",
  "validations_info": "{\"configuration\":[{\"id\":\"platform-requirements-satisfied\",\"status\":\"success\",\"message\":\"Platform requirements satisfied\"},{\"id\":\"pull-secret-set\",\"status\":\"success\",\"message\":\"The pull secret is set.\"}],\"hosts-data\":[{\"id\":\"all-hosts-are-ready-to-install\",\"status\":\"failure\",\"message\":\"The cluster has hosts that are not ready to install.\"},{\"id\":\"sufficient-masters-count\",\"status\":\"success\",\"message\":\"The cluster has the exact amount of dedicated control plane nodes.\"}],\"network\":[{\"id\":\"api-vips-defined\",\"status\":\"success\",\"message\":\"API virtual IPs are defined.\"},{\"id\":\"api-vips-valid\",\"status\":\"success\",\"message\":\"api vips 1.2.3.8 belongs to the Machine CIDR and is not in use.\"},{\"id\":\"cluster-cidr-defined\",\"status\":\"success\",\"message\":\"The Cluster Network CIDR is defined.\"},{\"id\":\"dns-domain-defined\",\"status\":\"success\",\"message\":\"The base domain is defined.\"},{\"id\":\"ingress-vips-defined\",\"status\":\"success\",\"message\":\"Ingress virtual IPs are defined.\"},{\"id\":\"ingress-vips-valid\",\"status\":\"success\",\"message\":\"ingress vips 1.2.3.9 belongs to the Machine CIDR and is not in use.\"},{\"id\":\"machine-cidr-defined\",\"status\":\"success\",\"message\":\"The Machine Network CIDR is defined.\"},{\"id\":\"machine-cidr-equals-to-calculated-cidr\",\"status\":\"success\",\"message\":\"The Cluster Machine CIDR is equivalent to the calculated CIDR.\"},{\"id\":\"network-prefix-valid\",\"status\":\"success\",\"message\":\"The Cluster Network prefix is valid.\"},{\"id\":\"network-type-valid\",\"status\":\"success\",\"message\":\"The cluster has a valid network type\"},{\"id\":\"networks-same-address-families\",\"status\":\"success\",\"message\":\"Same address families for all networks.\"},{\"id\":\"no-cidrs-overlapping\",\"status\":\"success\",\"message\":\"No CIDRS are overlapping.\"},{\"id\":\"ntp-server-configured\",\"status\":\"success\",\"message\":\"No ntp problems found\"},{\"id\":\"service-cidr-defined\",\"status\":\"success\",\"message\":\"The Service Network CIDR is defined.\"}],\"operators\":[{\"id\":\"authorino-requirements-satisfied\",\"status\":\"success\",\"message\":\"authorino is disabled\"},{\"id\":\"cnv-requirements-satisfied\",\"status\":\"success\",\"message\":\"cnv is disabled\"},{\"id\":\"lso-requirements-satisfied\",\"status\":\"success\",\"message\":\"lso is disabled\"},{\"id\":\"lvm-requirements-satisfied\",\"status\":\"success\",\"message\":\"lvm is disabled\"},{\"id\":\"mce-requirements-satisfied\",\"status\":\"success\",\"message\":\"mce is disabled\"},{\"id\":\"mtv-requirements-satisfied\",\"status\":\"success\",\"message\":\"mtv is disabled\"},{\"id\":\"nmstate-requirements-satisfied\",\"status\":\"success\",\"message\":\"nmstate is disabled\"},{\"id\":\"node-feature-discovery-requirements-satisfied\",\"status\":\"success\",\"message\":\"node-feature-discovery is disabled\"},{\"id\":\"nvidia-gpu-requirements-satisfied\",\"status\":\"success\",\"message\":\"nvidia-gpu is disabled\"},{\"id\":\"odf-requirements-satisfied\",\"status\":\"success\",\"message\":\"odf is disabled\"},{\"id\":\"openshift-ai-requirements-satisfied\",\"status\":\"success\",\"message\":\"openshift-ai is disabled\"},{\"id\":\"osc-requirements-satisfied\",\"status\":\"success\",\"message\":\"osc is disabled\"},{\"id\":\"pipelines-requirements-satisfied\",\"status\":\"success\",\"message\":\"pipelines is disabled\"},{\"id\":\"serverless-requirements-satisfied\",\"status\":\"success\",\"message\":\"serverless is disabled\"},{\"id\":\"servicemesh-requirements-satisfied\",\"status\":\"success\",\"message\":\"servicemesh is disabled\"}]}",
  "vip_dhcp_allocation": false
}
