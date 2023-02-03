package mse

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

// Gateways is a nested struct in mse response
type Gateways struct {
	Id              int64             `json:"Id" xml:"Id"`
	Name            string            `json:"Name" xml:"Name"`
	GatewayUniqueId string            `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	GatewayType     string            `json:"GatewayType" xml:"GatewayType"`
	Region          string            `json:"Region" xml:"Region"`
	PrimaryUser     string            `json:"PrimaryUser" xml:"PrimaryUser"`
	Status          int               `json:"Status" xml:"Status"`
	AhasOn          bool              `json:"AhasOn" xml:"AhasOn"`
	ArmsOn          bool              `json:"ArmsOn" xml:"ArmsOn"`
	Spec            string            `json:"Spec" xml:"Spec"`
	Replica         int               `json:"Replica" xml:"Replica"`
	GmtCreate       string            `json:"GmtCreate" xml:"GmtCreate"`
	GmtModified     string            `json:"GmtModified" xml:"GmtModified"`
	StatusDesc      string            `json:"StatusDesc" xml:"StatusDesc"`
	Upgrade         bool              `json:"Upgrade" xml:"Upgrade"`
	MustUpgrade     bool              `json:"MustUpgrade" xml:"MustUpgrade"`
	SupportWasm     bool              `json:"SupportWasm" xml:"SupportWasm"`
	CurrentVersion  string            `json:"CurrentVersion" xml:"CurrentVersion"`
	AppVersion      string            `json:"AppVersion" xml:"AppVersion"`
	LatestVersion   string            `json:"LatestVersion" xml:"LatestVersion"`
	Vswitch2        string            `json:"Vswitch2" xml:"Vswitch2"`
	InstanceId      string            `json:"InstanceId" xml:"InstanceId"`
	ChargeType      string            `json:"ChargeType" xml:"ChargeType"`
	EndDate         string            `json:"EndDate" xml:"EndDate"`
	Tag             string            `json:"Tag" xml:"Tag"`
	GatewayVersion  string            `json:"GatewayVersion" xml:"GatewayVersion"`
	RollBack        bool              `json:"RollBack" xml:"RollBack"`
	MseTag          string            `json:"MseTag" xml:"MseTag"`
	ResourceGroupId string            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InitConfig      InitConfig        `json:"InitConfig" xml:"InitConfig"`
	Slb             []SlbItem         `json:"Slb" xml:"Slb"`
	InternetSlb     []InternetSlbItem `json:"InternetSlb" xml:"InternetSlb"`
}
