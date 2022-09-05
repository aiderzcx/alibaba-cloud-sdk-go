package cdn

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

// GetDomainDetailModel is a nested struct in cdn response
type GetDomainDetailModel struct {
	SourcePort              int                              `json:"SourcePort" xml:"SourcePort"`
	HttpsCname              string                           `json:"HttpsCname" xml:"HttpsCname"`
	SourceType              string                           `json:"SourceType" xml:"SourceType"`
	ServerCertificateStatus string                           `json:"ServerCertificateStatus" xml:"ServerCertificateStatus"`
	GmtModified             string                           `json:"GmtModified" xml:"GmtModified"`
	DomainName              string                           `json:"DomainName" xml:"DomainName"`
	GmtCreated              string                           `json:"GmtCreated" xml:"GmtCreated"`
	Description             string                           `json:"Description" xml:"Description"`
	Region                  string                           `json:"Region" xml:"Region"`
	ResourceGroupId         string                           `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Scope                   string                           `json:"Scope" xml:"Scope"`
	DomainStatus            string                           `json:"DomainStatus" xml:"DomainStatus"`
	Cname                   string                           `json:"Cname" xml:"Cname"`
	CdnType                 string                           `json:"CdnType" xml:"CdnType"`
	Sources                 SourcesInDescribeCdnDomainDetail `json:"Sources" xml:"Sources"`
	SourceModels            SourceModels                     `json:"SourceModels" xml:"SourceModels"`
}
