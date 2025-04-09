export CLUSTER_ID=$(curl -X POST 10.73.116.40:8090/api/assisted-install/v2/clusters \
-H "Content-Type: application/json" \
-d "$(jq --null-input --slurpfile pull_secret /tmp/pull-secret ' { "name": "testcluster", "openshift_version": "4.19.0-0.test-2025-03-17-094626-ci-ln-4t067bk-latest", "high_availability_mode": "TNA", "control_plane_count": 2, "cpu_architecture": "x86_64", "base_dns_domain": "example.com", "api_vips": [{"ip": "192.168.122.20", "verification": "succeeded"}], "ingress_vips": [{"ip": "192.168.122.21", "verification": "succeeded"}], "pull_secret": $pull_secret[ 0 ] | tojson }')" | jq -r '.id')

curl -X PATCH 10.73.116.40:8090/api/assisted-install/v2/clusters/70777e5d-1359-48a2-97c8-af43c731f11a \
-H "Content-Type: application/json" \
-d '{"api_vips": [{"ip": "192.168.122.20", "verification": "succeeded"}], "ingress_vips": [{"ip": "192.168.122.21", "verification": "succeeded"}]}'

export INFRA_ENV_ID=$(curl -X POST 10.73.116.40:8090/api/assisted-install/v2/infra-envs \
-H "Content-Type: application/json" \
-d "$(jq --null-input --slurpfile pull_secret /tmp/pull-secret --arg cluster_id ${CLUSTER_ID} '{"name": "testcluster-infra-env", "image_type":"full-iso", "cluster_id": $cluster_id, "cpu_architecture": "x86_64", "pull_secret": $pull_secret[0] | tojson}')" | jq -r '.id')


curl -X POST localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts \
-H "Content-Type: application/json" \
-d '{"host_id": "2598e157-2b3d-41fe-b94c-1d5a6dae4454", "discovery_agent_version" : "quay.io/edge-infrastructure/assisted-installer-agent:latest}'

curl -X POST localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts \
-H "Content-Type: application/json" \
-d '{"host_id": "43c11d20-ce64-4d7e-9610-910da1a0c611", "discovery_agent_version" : "quay.io/edge-infrastructure/assisted-installer-agent:latest}'

curl -X POST localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts \
-H "Content-Type: application/json" \
-d '{"host_id": "29de31f7-9dd8-4e7c-bd15-13e09fec35b8", "discovery_agent_version" : "quay.io/edge-infrastructure/assisted-installer-agent:latest}'


curl -X PATCH localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts/2598e157-2b3d-41fe-b94c-1d5a6dae4454 \
-H "Content-Type: application/json" \
-d '{"host_role": "master", "host_name": "master-0"}'

curl -X PATCH localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts/43c11d20-ce64-4d7e-9610-910da1a0c611 \
-H "Content-Type: application/json" \
-d '{"host_role": "master", "host_name": "master-1"}'

curl -X PATCH 10.73.116.40:8090/api/assisted-install/v2/infra-envs/a25ac5db-60c7-4dc8-a137-1503739ac264/hosts/bf97fedf-06ec-4391-958d-6d413c5a35a9 \
-H "Content-Type: application/json" \
-d '{"host_role": "arbiter", "host_name": "arbiter-1"}'


curl -X POST localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts/2598e157-2b3d-41fe-b94c-1d5a6dae4454/instructions \
-H "Content-Type: application/json" \
-d '{"exit_code": 0, "step_id": "1fbf7d29-6416-4c44-b188-ce3977d3b350", "step_type": "inventory", "output": "{\"cpu\":{\"count\":4,\"flags\":null},\"disks\":[{\"bootable\":true,\"by_id\":\"wwn-0x1111111111111111111111\",\"drive_type\":\"HDD\",\"id\":\"wwn-0x1111111111111111111111\",\"installation_eligibility\":{\"not_eligible_reasons\":null},\"name\":\"sda1\",\"size_bytes\":128849018880}],\"gpus\":null,\"hostname\":\"0342b113-6d3b-4afa-8802-65f3dd24f7a2\",\"interfaces\":[{\"flags\":null,\"ipv4_addresses\":[\"192.168.122.1/24\"],\"ipv6_addresses\":null,\"name\":\"eth0\",\"speed_mbps\":40}],\"memory\":{\"physical_bytes\":17179869184,\"usable_bytes\":17179869184},\"routes\":null,\"system_vendor\":{\"manufacturer\":\"Red Hat\",\"product_name\":\"RHEL\",\"serial_number\":\"3534\"}}"}'

curl -X POST localhost:8090/api/assisted-install/v2/infra-envs/$INFRA_ENV_ID/hosts/43c11d20-ce64-4d7e-9610-910da1a0c611/instructions \
-H "Content-Type: application/json" \
-d '{"exit_code": 0, "step_id": "582f56c0-3420-4a95-b961-63c998cf75b5", "step_type": "inventory", "output": "{\"cpu\":{\"count\":4,\"flags\":null},\"disks\":[{\"bootable\":true,\"by_id\":\"wwn-0x1111111111111111111111\",\"drive_type\":\"HDD\",\"id\":\"wwn-0x1111111111111111111111\",\"installation_eligibility\":{\"not_eligible_reasons\":null},\"name\":\"sda1\",\"size_bytes\":128849018880}],\"gpus\":null,\"hostname\":\"0342b113-6d3b-4afa-8802-65f3dd24f7a2\",\"interfaces\":[{\"flags\":null,\"ipv4_addresses\":[\"192.168.122.2/24\"],\"ipv6_addresses\":null,\"name\":\"eth0\",\"speed_mbps\":40}],\"memory\":{\"physical_bytes\":17179869184,\"usable_bytes\":17179869184},\"routes\":null,\"system_vendor\":{\"manufacturer\":\"Red Hat\",\"product_name\":\"RHEL\",\"serial_number\":\"3534\"}}"}'

curl -X POST localhost:8090/api/assisted-install/v2/infra-envs/17227498-f9a3-4353-96f8-1b7087f5f6b9/hosts/5475e666-fa55-4afe-a20b-da5e990ae578/instructions \
-H "Content-Type: application/json" \
-d '{"exit_code": 0, "step_id": "637fed2f-83c9-414d-aabb-19faca2f74150", "step_type": "inventory", "output": "{\"cpu\":{\"architecture\":\"x86_64\",\"count\":16,\"flags\":null},\"disks\":[{\"by_id\":\"wwn-0x1111111111111111111111\",\"drive_type\":\"SSD\",\"id\":\"wwn-0x1111111111111111111111\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"loop0\",\"size_bytes\":128849018880},{\"by_id\":\"wwn-0x2222222222222222222222\",\"drive_type\":\"HDD\",\"id\":\"wwn-0x2222222222222222222222\",\"installation_eligibility\":{\"eligible\":true,\"not_eligible_reasons\":null},\"name\":\"sdb\",\"size_bytes\":128849018880}],\"gpus\":null,\"hostname\":\"h2\",\"interfaces\":[{\"flags\":null,\"ipv4_addresses\":[\"1.2.3.3/24\"],\"ipv6_addresses\":null,\"mac_address\":\"e6:53:3d:a7:77:b6\",\"type\":\"physical\"}],\"memory\":{\"physical_bytes\":34359738368,\"usable_bytes\":34359738368},\"routes\":[{\"destination\":\"0.0.0.0\",\"family\":2,\"gateway\":\"1.2.3.10\",\"interface\":\"eth0\",\"metric\":600}],\"system_vendor\":{\"manufacturer\":\"manu\",\"product_name\":\"prod\",\"serial_number\":\"3534\"},\"tpm_version\":\"2.0\"}"}'


