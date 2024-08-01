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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// Personalizedtxt2imgQueryInferenceJobInfo invokes the aicontent.Personalizedtxt2imgQueryInferenceJobInfo API synchronously
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfo(request *Personalizedtxt2imgQueryInferenceJobInfoRequest) (response *Personalizedtxt2imgQueryInferenceJobInfoResponse, err error) {
	response = CreatePersonalizedtxt2imgQueryInferenceJobInfoResponse()
	err = client.DoAction(request, response)
	return
}

// Personalizedtxt2imgQueryInferenceJobInfoWithChan invokes the aicontent.Personalizedtxt2imgQueryInferenceJobInfo API asynchronously
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfoWithChan(request *Personalizedtxt2imgQueryInferenceJobInfoRequest) (<-chan *Personalizedtxt2imgQueryInferenceJobInfoResponse, <-chan error) {
	responseChan := make(chan *Personalizedtxt2imgQueryInferenceJobInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Personalizedtxt2imgQueryInferenceJobInfo(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// Personalizedtxt2imgQueryInferenceJobInfoWithCallback invokes the aicontent.Personalizedtxt2imgQueryInferenceJobInfo API asynchronously
func (client *Client) Personalizedtxt2imgQueryInferenceJobInfoWithCallback(request *Personalizedtxt2imgQueryInferenceJobInfoRequest, callback func(response *Personalizedtxt2imgQueryInferenceJobInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Personalizedtxt2imgQueryInferenceJobInfoResponse
		var err error
		defer close(result)
		response, err = client.Personalizedtxt2imgQueryInferenceJobInfo(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// Personalizedtxt2imgQueryInferenceJobInfoRequest is the request struct for api Personalizedtxt2imgQueryInferenceJobInfo
type Personalizedtxt2imgQueryInferenceJobInfoRequest struct {
	*requests.RoaRequest
	InferenceJobId string `position:"Query" name:"inferenceJobId"`
}

// Personalizedtxt2imgQueryInferenceJobInfoResponse is the response struct for api Personalizedtxt2imgQueryInferenceJobInfo
type Personalizedtxt2imgQueryInferenceJobInfoResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"requestId" xml:"requestId"`
	Success        bool   `json:"success" xml:"success"`
	ErrCode        string `json:"errCode" xml:"errCode"`
	ErrMessage     string `json:"errMessage" xml:"errMessage"`
	HttpStatusCode int    `json:"httpStatusCode" xml:"httpStatusCode"`
	Data           Data   `json:"data" xml:"data"`
}

// CreatePersonalizedtxt2imgQueryInferenceJobInfoRequest creates a request to invoke Personalizedtxt2imgQueryInferenceJobInfo API
func CreatePersonalizedtxt2imgQueryInferenceJobInfoRequest() (request *Personalizedtxt2imgQueryInferenceJobInfoRequest) {
	request = &Personalizedtxt2imgQueryInferenceJobInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AiContent", "20240611", "Personalizedtxt2imgQueryInferenceJobInfo", "/api/v1/personalizedtxt2img/queryInferenceJobInfo", "", "")
	request.Method = requests.GET
	return
}

// CreatePersonalizedtxt2imgQueryInferenceJobInfoResponse creates a response to parse from Personalizedtxt2imgQueryInferenceJobInfo response
func CreatePersonalizedtxt2imgQueryInferenceJobInfoResponse() (response *Personalizedtxt2imgQueryInferenceJobInfoResponse) {
	response = &Personalizedtxt2imgQueryInferenceJobInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
