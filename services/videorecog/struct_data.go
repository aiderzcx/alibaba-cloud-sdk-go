package videorecog

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

// Data is a nested struct in videorecog response
type Data struct {
	JobId                 string                      `json:"JobId" xml:"JobId"`
	Result                string                      `json:"Result" xml:"Result"`
	ErrorMessage          string                      `json:"ErrorMessage" xml:"ErrorMessage"`
	JsonUrl               string                      `json:"JsonUrl" xml:"JsonUrl"`
	TagInfo               map[string]interface{}      `json:"TagInfo" xml:"TagInfo"`
	PdfUrl                string                      `json:"PdfUrl" xml:"PdfUrl"`
	Status                string                      `json:"Status" xml:"Status"`
	ErrorCode             string                      `json:"ErrorCode" xml:"ErrorCode"`
	ShotFrameIds          []int                       `json:"ShotFrameIds" xml:"ShotFrameIds"`
	VideoQualityInfo      VideoQualityInfo            `json:"VideoQualityInfo" xml:"VideoQualityInfo"`
	VideoInfo             VideoInfo                   `json:"VideoInfo" xml:"VideoInfo"`
	SplitVideoPartResults []SplitVideoPartResultsItem `json:"SplitVideoPartResults" xml:"SplitVideoPartResults"`
	Elements              []ElementsItem              `json:"Elements" xml:"Elements"`
	Outputs               []Output                    `json:"Outputs" xml:"Outputs"`
}
