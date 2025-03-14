package polardb

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

// LicenseOrder is a nested struct in polardb response
type LicenseOrder struct {
	ActivatedCodeCount         int    `json:"ActivatedCodeCount" xml:"ActivatedCodeCount"`
	ActivationCodeQuota        int    `json:"ActivationCodeQuota" xml:"ActivationCodeQuota"`
	AliyunOrderId              string `json:"AliyunOrderId" xml:"AliyunOrderId"`
	AllowEmptySystemIdentifier bool   `json:"AllowEmptySystemIdentifier" xml:"AllowEmptySystemIdentifier"`
	Engine                     string `json:"Engine" xml:"Engine"`
	GmtCreated                 string `json:"GmtCreated" xml:"GmtCreated"`
	GmtModified                string `json:"GmtModified" xml:"GmtModified"`
	IsVirtualOrder             bool   `json:"IsVirtualOrder" xml:"IsVirtualOrder"`
	IsVirtualOrderFrozen       bool   `json:"IsVirtualOrderFrozen" xml:"IsVirtualOrderFrozen"`
	PackageType                string `json:"PackageType" xml:"PackageType"`
	PackageValidity            string `json:"PackageValidity" xml:"PackageValidity"`
	PurchaseChannel            string `json:"PurchaseChannel" xml:"PurchaseChannel"`
	VirtualAliyunOrderId       string `json:"VirtualAliyunOrderId" xml:"VirtualAliyunOrderId"`
}
