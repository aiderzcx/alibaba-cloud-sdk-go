package aicontent

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

// DataItem is a nested struct in aicontent response
type DataItem struct {
	InferenceImageCount  int                    `json:"inferenceImageCount" xml:"inferenceImageCount"`
	FreeConcurrencyCount int                    `json:"FreeConcurrencyCount" xml:"FreeConcurrencyCount"`
	ModelId              string                 `json:"modelId" xml:"modelId"`
	ServiceCode          string                 `json:"ServiceCode" xml:"ServiceCode"`
	JobTrainProgress     string                 `json:"jobTrainProgress" xml:"jobTrainProgress"`
	Name                 string                 `json:"name" xml:"name"`
	FreeCount            int                    `json:"FreeCount" xml:"FreeCount"`
	ObjectType           string                 `json:"objectType" xml:"objectType"`
	ServiceName          string                 `json:"ServiceName" xml:"ServiceName"`
	Id                   string                 `json:"id" xml:"id"`
	CreateTime           string                 `json:"createTime" xml:"createTime"`
	JobStatus            string                 `json:"jobStatus" xml:"jobStatus"`
	ImageUrl             []string               `json:"imageUrl" xml:"imageUrl"`
	InferenceJobList     []InferenceJobListItem `json:"inferenceJobList" xml:"inferenceJobList"`
}
