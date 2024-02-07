package models

import "time"

type Vm struct {
	// Unique identifier for the VM
	Uuid string `json:"uuid,omitempty"`
	// Name of the VM
	Name string `json:"name"`

	Description string `json:"description,omitempty"`
	// Number of CPU cores
	Cpucores int32 `json:"cpucores"`
	// Amount of memory in GB
	MemoryGb int32 `json:"memoryGb"`
	// List of disk sizes in GB
	DisksGb []int32 `json:"disksGb"`
	// VLAN ID
	Vlan int32 `json:"vlan,omitempty"`
	// Hypervisor name
	Hypervisor string `json:"hypervisor,omitempty"`

	Created time.Time `json:"created,omitempty"`

	Updated time.Time `json:"updated,omitempty"`
}
