package csas

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

// Strategy is a nested struct in csas response
type Strategy struct {
	UpdateTime      string       `json:"UpdateTime" xml:"UpdateTime"`
	PolicyId        string       `json:"PolicyId" xml:"PolicyId"`
	AllowReport     bool         `json:"AllowReport" xml:"AllowReport"`
	CreateTime      string       `json:"CreateTime" xml:"CreateTime"`
	IsAntiUninstall bool         `json:"IsAntiUninstall" xml:"IsAntiUninstall"`
	ReportProcessId string       `json:"ReportProcessId" xml:"ReportProcessId"`
	IsBoot          bool         `json:"IsBoot" xml:"IsBoot"`
	UserGroupIds    []string     `json:"UserGroupIds" xml:"UserGroupIds"`
	WhitelistUsers  []string     `json:"WhitelistUsers" xml:"WhitelistUsers"`
	BlockContent    BlockContent `json:"BlockContent" xml:"BlockContent"`
}
