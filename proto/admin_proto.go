// Copyright 2018 The Container File System Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package proto

const (
	// Admin APIs
	AdminGetCluster                = "/admin/getCluster"
	AdminGetDataPartition          = "/dataPartition/get"
	AdminLoadDataPartition         = "/dataPartition/load"
	AdminCreateDataPartition       = "/dataPartition/create"
	AdminDecommissionDataPartition = "/dataPartition/decommission"
	AdminDeleteVol                 = "/vol/delete"
	AdminUpdateVol                 = "/vol/update"
	AdminCreateVol                 = "/admin/createVol"
	AdminClusterFreeze             = "/cluster/freeze"
	AdminGetIP                     = "/admin/getIp"
	AdminCreateMP                  = "/metaPartition/create"
	AdminSetMetaNodeThreshold      = "/threshold/set"

	// Client APIs
	ClientDataPartitions = "/client/partitions"
	ClientVol            = "/client/vol"
	ClientMetaPartition  = "/client/metaPartition"
	ClientVolStat        = "/client/volStat"

	//raft node APIs
	AddRaftNode    = "/raftNode/add"
	RemoveRaftNode = "/raftNode/remove"

	// Node APIs
	AddDataNode                    = "/dataNode/add"
	DecommissionDataNode           = "/dataNode/decommission"
	DecommissionDisk               = "/disk/decommission"
	GetDataNode                    = "/dataNode/get"
	AddMetaNode                    = "/metaNode/add"
	DecommissionMetaNode           = "/metaNode/decommission"
	GetMetaNode                    = "/metaNode/get"
	AdminLoadMetaPartition         = "/metaPartition/load"
	AdminDecommissionMetaPartition = "/metaPartition/decommission"

	// Operation response
	GetMetaNodeTaskResponse = "/metaNode/response" // Method: 'POST', ContentType: 'application/json'
	GetDataNodeTaskResponse = "/dataNode/response" // Method: 'POST', ContentType: 'application/json'

	GetTopologyView = "/topo/get"
)

// RegisterMetaNodeResp defines the response to register a meta node.
type RegisterMetaNodeResp struct {
	ID uint64
}

// ClusterInfo defines the cluster infomation.
type ClusterInfo struct {
	Cluster string
	Ip      string
}

// CreateDataPartitionRequest defines the request to create a data partition.
type CreateDataPartitionRequest struct {
	PartitionType string
	PartitionId   uint64
	PartitionSize int
	VolumeId      string
	IsRandomWrite bool
	Members       []Peer
}

// CreateDataPartitionResponse defines the response to the request of creating a data partition.
type CreateDataPartitionResponse struct {
	PartitionId uint64
	Status      uint8
	Result      string
}

// DeleteDataPartitionRequest defines the request to delete a data partition.
type DeleteDataPartitionRequest struct {
	DataPartitionType string
	PartitionId       uint64
	PartitionSize     int
}

// DeleteDataPartitionResponse defines the response to the request of deleting a data partition.
type DeleteDataPartitionResponse struct {
	Status      uint8
	Result      string
	PartitionId uint64
}

// DataPartitionDecommissionRequest defines the request of decommissioning a data partition.
type DataPartitionDecommissionRequest struct {
	PartitionId uint64
	RemovePeer  Peer
	AddPeer     Peer
}

// DataPartitionDecommissionResponse defines the response to the request of decommissioning a data partition.
type DataPartitionDecommissionResponse struct {
	Status      uint8
	Result      string
	PartitionId uint64
}

// LoadDataPartitionRequest defines the request of loading a data partition.
type LoadDataPartitionRequest struct {
	PartitionId uint64
}

// LoadDataPartitionResponse defines the response to the request of loading a data partition.
type LoadDataPartitionResponse struct {
	PartitionId       uint64
	Used              uint64
	PartitionSnapshot []*File
	Status            uint8
	PartitionStatus   int
	Result            string
}

// File defines the file struct.
type File struct {
	Name     string
	Crc      uint32
	Size     uint32
	Modified int64
}

// LoadMetaPartitionMetricRequest defines the request of loading the meta partition metrics.
type LoadMetaPartitionMetricRequest struct {
	PartitionID uint64
	Start       uint64
	End         uint64
}

// LoadMetaPartitionMetricResponse defines the response to the request of loading the meta partition metrics.
type LoadMetaPartitionMetricResponse struct {
	Start    uint64
	End      uint64
	MaxInode uint64
	Status   uint8
	Result   string
}

// HeartBeatRequest define the heartbeat request.
type HeartBeatRequest struct {
	CurrTime   int64
	MasterAddr string
}

// PartitionReport defines the partition report.
type PartitionReport struct {
	PartitionID     uint64
	PartitionStatus int
	Total           uint64
	Used            uint64
	DiskPath        string
	IsLeader        bool
	ExtentCount     int
	NeedCompare     bool
}

// DataNodeHeartbeatResponse defines the response to the data node heartbeat.
type DataNodeHeartbeatResponse struct {
	Total               uint64
	Used                uint64
	Available           uint64
	TotalPartitionSize  uint64 // volCnt * volsize
	RemainingCapacity   uint64 // remaining capacity to create partition
	CreatedPartitionCnt uint32
	MaxCapacity         uint64 // maximum capacity to create partition
	RackName            string
	PartitionReports    []*PartitionReport
	Status              uint8
	Result              string
}

// MetaPartitionReport defines the meta partition report.
type MetaPartitionReport struct {
	PartitionID uint64
	Start       uint64
	End         uint64
	Status      int
	MaxInodeID  uint64
	IsLeader    bool
}

// MetaNodeHeartbeatResponse defines the response to the meta node heartbeat request.
type MetaNodeHeartbeatResponse struct {
	RackName             string
	Total                uint64
	Used                 uint64
	MetaPartitionReports []*MetaPartitionReport
	Status               uint8
	Result               string
}

// DeleteFileRequest defines the request to delete a file.
type DeleteFileRequest struct {
	VolId uint64
	Name  string
}

// DeleteFileResponse defines the response to the request of deleting a file.
type DeleteFileResponse struct {
	Status uint8
	Result string
	VolId  uint64
	Name   string
}

// DeleteMetaPartitionRequest defines the request of deleting a meta partition.
type DeleteMetaPartitionRequest struct {
	PartitionID uint64
}

// DeleteMetaPartitionResponse defines the response to the request of deleting a meta partition.
type DeleteMetaPartitionResponse struct {
	PartitionID uint64
	Status      uint8
	Result      string
}

// UpdateMetaPartitionRequest defines the request to update a meta partition.
type UpdateMetaPartitionRequest struct {
	PartitionID uint64
	VolName     string
	Start       uint64
	End         uint64
}

// UpdateMetaPartitionResponse defines the response to the request of updating the meta partition.
type UpdateMetaPartitionResponse struct {
	PartitionID uint64
	VolName     string
	End         uint64
	Status      uint8
	Result      string
}

// MetaPartitionDecommissionRequest defines the request of decommissioning a meta partition.
type MetaPartitionDecommissionRequest struct {
	PartitionID uint64
	VolName     string
	RemovePeer  Peer
	AddPeer     Peer
}

// MetaPartitionDecommissionResponse defines the response to the request of decommissioning a meta partition.
type MetaPartitionDecommissionResponse struct {
	PartitionID uint64
	VolName     string
	Status      uint8
	Result      string
}

// MetaPartitionLoadRequest defines the request to load meta partition.
type MetaPartitionLoadRequest struct {
	PartitionID uint64
}

// MetaPartitionLoadResponse defines the response to the request of loading meta partition.
type MetaPartitionLoadResponse struct {
	PartitionID uint64
	DoCompare   bool
	ApplyID     uint64
	InodeSign   uint32
	DentrySign  uint32
	Addr        string
}
