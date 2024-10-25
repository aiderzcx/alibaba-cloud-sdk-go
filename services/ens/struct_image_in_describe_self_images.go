package ens

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ImageInDescribeSelfImages is a nested struct in ens response
type ImageInDescribeSelfImages struct {
	Architecture       string                                 `json:"Architecture" xml:"Architecture"`
	CreationTime       string                                 `json:"CreationTime" xml:"CreationTime"`
	ImageId            string                                 `json:"ImageId" xml:"ImageId"`
	ImageName          string                                 `json:"ImageName" xml:"ImageName"`
	ImageOwnerAlias    string                                 `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
	ImageSize          string                                 `json:"ImageSize" xml:"ImageSize"`
	InstanceId         string                                 `json:"InstanceId" xml:"InstanceId"`
	OsVersion          string                                 `json:"OsVersion" xml:"OsVersion"`
	Platform           string                                 `json:"Platform" xml:"Platform"`
	Status             string                                 `json:"Status" xml:"Status"`
	ComputeType        string                                 `json:"ComputeType" xml:"ComputeType"`
	SnapshotId         string                                 `json:"SnapshotId" xml:"SnapshotId"`
	ImageStorageSize   string                                 `json:"ImageStorageSize" xml:"ImageStorageSize"`
	DiskDeviceMappings DiskDeviceMappingsInDescribeSelfImages `json:"DiskDeviceMappings" xml:"DiskDeviceMappings"`
}
