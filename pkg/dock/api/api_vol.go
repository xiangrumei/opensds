// Copyright (c) 2016 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

/*
This module implements the entry into operations of storageDock module.

*/

package api

import (
	"log"

	dock "github.com/opensds/opensds/pkg/dock/volume"
	pb "github.com/opensds/opensds/pkg/grpc/opensds"
)

func CreateVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.CreateVolume(vr.GetResourceType(),
		vr.GetName(),
		vr.GetVolumeType(),
		vr.GetSize())
	if err != nil {
		log.Println("Error occured in dock module when create volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func GetVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.GetVolume(vr.GetResourceType(), vr.GetId())
	if err != nil {
		log.Println("Error occured in dock module when get volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func ListVolumes(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.GetAllVolumes(vr.GetResourceType(),
		vr.GetAllowDetails())
	if err != nil {
		log.Println("Error occured in dock module when list volumes:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func DeleteVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.DeleteVolume(vr.GetResourceType(),
		vr.GetId())
	if err != nil {
		log.Println("Error occured in dock module when delete volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func AttachVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.AttachVolume(vr.GetResourceType(),
		vr.GetId())
	if err != nil {
		log.Println("Error occured in dock module when attach volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func DetachVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.DetachVolume(vr.GetResourceType(),
		vr.GetDevice())
	if err != nil {
		log.Println("Error occured in dock module when detach volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func MountVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.MountVolume(vr.GetMountDir(),
		vr.GetDevice(),
		vr.GetFsType())
	if err != nil {
		log.Println("Error occured in dock module when mount volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func UnmountVolume(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.UnmountVolume(vr.GetMountDir())
	if err != nil {
		log.Println("Error occured in dock module when unmount volume:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func CreateVolumeSnapshot(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.CreateSnapshot(vr.GetResourceType(),
		vr.GetSnapshotName(),
		vr.GetId(),
		vr.GetDescription(),
		vr.GetForceSnapshoted())
	if err != nil {
		log.Println("Error occured in dock module when create snapshot:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func GetVolumeSnapshot(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.GetSnapshot(vr.GetResourceType(), vr.GetSnapshotId())
	if err != nil {
		log.Println("Error occured in dock module when get snapshot:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func ListVolumeSnapshots(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.GetAllSnapshots(vr.GetResourceType())
	if err != nil {
		log.Println("Error occured in dock module when list snapshots:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}

func DeleteVolumeSnapshot(vr *pb.VolumeRequest) (*pb.Response, error) {
	result, err := dock.DeleteSnapshot(vr.GetResourceType(),
		vr.GetSnapshotId())
	if err != nil {
		log.Println("Error occured in dock module when delete snapshot:", err)
		return &pb.Response{}, err
	}

	return &pb.Response{
		Status:  "Success",
		Message: result,
	}, nil
}
