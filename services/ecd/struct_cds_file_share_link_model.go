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

// CdsFileShareLinkModel is a nested struct in ecd response
type CdsFileShareLinkModel struct {
	Description       string `json:"Description" xml:"Description"`
	SaveLimit         int64  `json:"SaveLimit" xml:"SaveLimit"`
	DownloadCount     int64  `json:"DownloadCount" xml:"DownloadCount"`
	Creator           string `json:"Creator" xml:"Creator"`
	SaveCount         int64  `json:"SaveCount" xml:"SaveCount"`
	DisableSave       bool   `json:"DisableSave" xml:"DisableSave"`
	SharePwd          string `json:"SharePwd" xml:"SharePwd"`
	DisablePreview    bool   `json:"DisablePreview" xml:"DisablePreview"`
	AccessCount       int64  `json:"AccessCount" xml:"AccessCount"`
	ShareId           string `json:"ShareId" xml:"ShareId"`
	DownloadLimit     int64  `json:"DownloadLimit" xml:"DownloadLimit"`
	Status            string `json:"Status" xml:"Status"`
	DriveId           string `json:"DriveId" xml:"DriveId"`
	ModifiyTime       string `json:"ModifiyTime" xml:"ModifiyTime"`
	DisableDownload   bool   `json:"DisableDownload" xml:"DisableDownload"`
	Expiration        string `json:"Expiration" xml:"Expiration"`
	Expired           bool   `json:"Expired" xml:"Expired"`
	PreviewCount      int64  `json:"PreviewCount" xml:"PreviewCount"`
	PreviewLimit      int64  `json:"PreviewLimit" xml:"PreviewLimit"`
	VideoPreviewCount int64  `json:"VideoPreviewCount" xml:"VideoPreviewCount"`
	ReportCount       int64  `json:"ReportCount" xml:"ReportCount"`
	FileIds           string `json:"FileIds" xml:"FileIds"`
	ShareName         string `json:"ShareName" xml:"ShareName"`
	CreateTime        string `json:"CreateTime" xml:"CreateTime"`
	ShareLink         string `json:"ShareLink" xml:"ShareLink"`
}
