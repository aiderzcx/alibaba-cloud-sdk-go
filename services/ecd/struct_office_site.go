package ecd

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

// OfficeSite is a nested struct in ecd response
type OfficeSite struct {
	Status                   string        `json:"Status" xml:"Status"`
	CreationTime             string        `json:"CreationTime" xml:"CreationTime"`
	VpcId                    string        `json:"VpcId" xml:"VpcId"`
	VpcType                  string        `json:"VpcType" xml:"VpcType"`
	EnableAdminAccess        bool          `json:"EnableAdminAccess" xml:"EnableAdminAccess"`
	EnableCrossDesktopAccess bool          `json:"EnableCrossDesktopAccess" xml:"EnableCrossDesktopAccess"`
	DesktopVpcEndpoint       string        `json:"DesktopVpcEndpoint" xml:"DesktopVpcEndpoint"`
	DesktopAccessType        string        `json:"DesktopAccessType" xml:"DesktopAccessType"`
	DomainName               string        `json:"DomainName" xml:"DomainName"`
	SsoEnabled               bool          `json:"SsoEnabled" xml:"SsoEnabled"`
	CidrBlock                string        `json:"CidrBlock" xml:"CidrBlock"`
	Bandwidth                int           `json:"Bandwidth" xml:"Bandwidth"`
	TrustPassword            string        `json:"TrustPassword" xml:"TrustPassword"`
	Name                     string        `json:"Name" xml:"Name"`
	EnableInternetAccess     bool          `json:"EnableInternetAccess" xml:"EnableInternetAccess"`
	DomainPassword           string        `json:"DomainPassword" xml:"DomainPassword"`
	CustomSecurityGroupId    string        `json:"CustomSecurityGroupId" xml:"CustomSecurityGroupId"`
	OuName                   string        `json:"OuName" xml:"OuName"`
	DomainUserName           string        `json:"DomainUserName" xml:"DomainUserName"`
	SubDomainName            string        `json:"SubDomainName" xml:"SubDomainName"`
	OfficeSiteId             string        `json:"OfficeSiteId" xml:"OfficeSiteId"`
	CenId                    string        `json:"CenId" xml:"CenId"`
	CenAttachStatus          string        `json:"CenAttachStatus" xml:"CenAttachStatus"`
	MfaEnabled               bool          `json:"MfaEnabled" xml:"MfaEnabled"`
	NetworkPackageId         string        `json:"NetworkPackageId" xml:"NetworkPackageId"`
	DnsUserName              string        `json:"DnsUserName" xml:"DnsUserName"`
	OfficeSiteType           string        `json:"OfficeSiteType" xml:"OfficeSiteType"`
	NeedVerifyLoginRisk      bool          `json:"NeedVerifyLoginRisk" xml:"NeedVerifyLoginRisk"`
	DesktopCount             int64         `json:"DesktopCount" xml:"DesktopCount"`
	TotalEdsCount            int64         `json:"TotalEdsCount" xml:"TotalEdsCount"`
	TotalEdsCountForGroup    int64         `json:"TotalEdsCountForGroup" xml:"TotalEdsCountForGroup"`
	NeedVerifyZeroDevice     bool          `json:"NeedVerifyZeroDevice" xml:"NeedVerifyZeroDevice"`
	CloudBoxOfficeSite       bool          `json:"CloudBoxOfficeSite" xml:"CloudBoxOfficeSite"`
	SsoType                  string        `json:"SsoType" xml:"SsoType"`
	ProtocolType             string        `json:"ProtocolType" xml:"ProtocolType"`
	AdHostname               string        `json:"AdHostname" xml:"AdHostname"`
	RdsLicenseStatus         string        `json:"RdsLicenseStatus" xml:"RdsLicenseStatus"`
	RdsLicenseAddress        string        `json:"RdsLicenseAddress" xml:"RdsLicenseAddress"`
	RdsLicenseDomainName     string        `json:"RdsLicenseDomainName" xml:"RdsLicenseDomainName"`
	BackupDns                string        `json:"BackupDns" xml:"BackupDns"`
	BackupDCHostname         string        `json:"BackupDCHostname" xml:"BackupDCHostname"`
	EnableServiceRoute       bool          `json:"EnableServiceRoute" xml:"EnableServiceRoute"`
	SubnetMode               string        `json:"SubnetMode" xml:"SubnetMode"`
	SecurityProtection       string        `json:"SecurityProtection" xml:"SecurityProtection"`
	CustomAccessPoint        string        `json:"CustomAccessPoint" xml:"CustomAccessPoint"`
	VSwitchIds               []string      `json:"VSwitchIds" xml:"VSwitchIds"`
	FileSystemIds            []string      `json:"FileSystemIds" xml:"FileSystemIds"`
	SubDnsAddress            []string      `json:"SubDnsAddress" xml:"SubDnsAddress"`
	DnsAddress               []string      `json:"DnsAddress" xml:"DnsAddress"`
	ADConnectors             []ADConnector `json:"ADConnectors" xml:"ADConnectors"`
	Logs                     []Log         `json:"Logs" xml:"Logs"`
}
