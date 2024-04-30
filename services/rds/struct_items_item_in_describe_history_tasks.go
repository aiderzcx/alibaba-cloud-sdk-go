package rds

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

// ItemsItemInDescribeHistoryTasks is a nested struct in rds response
type ItemsItemInDescribeHistoryTasks struct {
	Status          string  `json:"Status" xml:"Status"`
	TaskId          string  `json:"TaskId" xml:"TaskId"`
	CurrentStepName string  `json:"CurrentStepName" xml:"CurrentStepName"`
	StartTime       string  `json:"StartTime" xml:"StartTime"`
	EndTime         string  `json:"EndTime" xml:"EndTime"`
	TaskType        string  `json:"TaskType" xml:"TaskType"`
	RemainTime      int     `json:"RemainTime" xml:"RemainTime"`
	Progress        float64 `json:"Progress" xml:"Progress"`
	RegionId        string  `json:"RegionId" xml:"RegionId"`
	InstanceType    string  `json:"InstanceType" xml:"InstanceType"`
	InstanceId      string  `json:"InstanceId" xml:"InstanceId"`
	InstanceName    string  `json:"InstanceName" xml:"InstanceName"`
	DbType          string  `json:"DbType" xml:"DbType"`
	Product         string  `json:"Product" xml:"Product"`
	TaskDetail      string  `json:"TaskDetail" xml:"TaskDetail"`
	ReasonCode      string  `json:"ReasonCode" xml:"ReasonCode"`
	ActionInfo      string  `json:"ActionInfo" xml:"ActionInfo"`
	Uid             string  `json:"Uid" xml:"Uid"`
	CallerSource    string  `json:"CallerSource" xml:"CallerSource"`
	CallerUid       string  `json:"CallerUid" xml:"CallerUid"`
}
