package assisted_service

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/openshift/assisted-service/models"
)

func defaultInventory() string {
	inventory := models.Inventory{
		Interfaces: []*models.Interface{
			{
				Name: "eth0",
				IPV4Addresses: []string{
					"1.2.3.4/24",
				},
				SpeedMbps: 20,
			},
			{
				Name: "eth1",
				IPV4Addresses: []string{
					"1.2.5.4/24",
				},
				SpeedMbps: 40,
			},
		},

		// CPU, Disks, and Memory were added here to prevent the case that assisted-service crashes in case the monitor starts
		// working in the middle of the test and this inventory is in the database.
		CPU: &models.CPU{
			Count: 4,
		},
		Disks: []*models.Disk{
			{
				ID:        "wwn-0x1111111111111111111111",
				ByID:      "wwn-0x1111111111111111111111",
				DriveType: "HDD",
				Name:      "sda1",
				SizeBytes: int64(120) * (int64(1) << 30),
				Bootable:  true,
			},
		},
		Hostname: uuid.New().String(),
		Memory: &models.Memory{
			PhysicalBytes: int64(16) * (int64(1) << 30),
			UsableBytes:   int64(16) * (int64(1) << 30),
		},
		SystemVendor: &models.SystemVendor{Manufacturer: "Red Hat", ProductName: "RHEL", SerialNumber: "3534"},
	}
	b, _ := json.Marshal(&inventory)
	return string(b)
}

func main() {
	fmt.Println(defaultInventory())
}
