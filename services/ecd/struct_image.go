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

// Image is a nested struct in ecd response
type Image struct {
	CreationTime            string   `json:"CreationTime" xml:"CreationTime"`
	Status                  string   `json:"Status" xml:"Status"`
	Progress                string   `json:"Progress" xml:"Progress"`
	DataDiskSize            int      `json:"DataDiskSize" xml:"DataDiskSize"`
	ImageType               string   `json:"ImageType" xml:"ImageType"`
	Description             string   `json:"Description" xml:"Description"`
	Size                    int      `json:"Size" xml:"Size"`
	OsType                  string   `json:"OsType" xml:"OsType"`
	ProtocolType            string   `json:"ProtocolType" xml:"ProtocolType"`
	Name                    string   `json:"Name" xml:"Name"`
	ImageId                 string   `json:"ImageId" xml:"ImageId"`
	GpuCategory             bool     `json:"GpuCategory" xml:"GpuCategory"`
	GpuDriverVersion        string   `json:"GpuDriverVersion" xml:"GpuDriverVersion"`
	AppVersion              string   `json:"AppVersion" xml:"AppVersion"`
	VolumeEncryptionEnabled bool     `json:"VolumeEncryptionEnabled" xml:"VolumeEncryptionEnabled"`
	VolumeEncryptionKey     string   `json:"VolumeEncryptionKey" xml:"VolumeEncryptionKey"`
	SharedCount             int      `json:"SharedCount" xml:"SharedCount"`
	SessionType             string   `json:"SessionType" xml:"SessionType"`
	UpdateTime              string   `json:"UpdateTime" xml:"UpdateTime"`
	Platform                string   `json:"Platform" xml:"Platform"`
	SupportedLanguages      []string `json:"SupportedLanguages" xml:"SupportedLanguages"`
}
