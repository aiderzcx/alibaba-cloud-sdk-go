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

// DBClusterNode is a nested struct in rds response
type DBClusterNode struct {
	NodeRegionId string `json:"NodeRegionId" xml:"NodeRegionId"`
	NodeZoneId   string `json:"NodeZoneId" xml:"NodeZoneId"`
	NodeId       string `json:"NodeId" xml:"NodeId"`
	NodeRole     string `json:"NodeRole" xml:"NodeRole"`
	ClassCode    string `json:"ClassCode" xml:"ClassCode"`
	ClassType    string `json:"ClassType" xml:"ClassType"`
	Cpu          string `json:"Cpu" xml:"Cpu"`
	Memory       string `json:"Memory" xml:"Memory"`
	Status       string `json:"Status" xml:"Status"`
}
