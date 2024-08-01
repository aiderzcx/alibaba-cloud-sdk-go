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

// Personalizedtxt2imgAddInferenceJob invokes the aicontent.Personalizedtxt2imgAddInferenceJob API synchronously
func (client *Client) Personalizedtxt2imgAddInferenceJob(request *Personalizedtxt2imgAddInferenceJobRequest) (response *Personalizedtxt2imgAddInferenceJobResponse, err error) {
	response = CreatePersonalizedtxt2imgAddInferenceJobResponse()
	err = client.DoAction(request, response)
	return
}

// Personalizedtxt2imgAddInferenceJobWithChan invokes the aicontent.Personalizedtxt2imgAddInferenceJob API asynchronously
func (client *Client) Personalizedtxt2imgAddInferenceJobWithChan(request *Personalizedtxt2imgAddInferenceJobRequest) (<-chan *Personalizedtxt2imgAddInferenceJobResponse, <-chan error) {
	responseChan := make(chan *Personalizedtxt2imgAddInferenceJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Personalizedtxt2imgAddInferenceJob(request)
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

// Personalizedtxt2imgAddInferenceJobWithCallback invokes the aicontent.Personalizedtxt2imgAddInferenceJob API asynchronously
func (client *Client) Personalizedtxt2imgAddInferenceJobWithCallback(request *Personalizedtxt2imgAddInferenceJobRequest, callback func(response *Personalizedtxt2imgAddInferenceJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Personalizedtxt2imgAddInferenceJobResponse
		var err error
		defer close(result)
		response, err = client.Personalizedtxt2imgAddInferenceJob(request)
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

// Personalizedtxt2imgAddInferenceJobRequest is the request struct for api Personalizedtxt2imgAddInferenceJob
type Personalizedtxt2imgAddInferenceJobRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// Personalizedtxt2imgAddInferenceJobResponse is the response struct for api Personalizedtxt2imgAddInferenceJob
type Personalizedtxt2imgAddInferenceJobResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"requestId" xml:"requestId"`
	Success        bool   `json:"success" xml:"success"`
	ErrCode        string `json:"errCode" xml:"errCode"`
	ErrMessage     string `json:"errMessage" xml:"errMessage"`
	HttpStatusCode int    `json:"httpStatusCode" xml:"httpStatusCode"`
	Data           Data   `json:"data" xml:"data"`
}

// CreatePersonalizedtxt2imgAddInferenceJobRequest creates a request to invoke Personalizedtxt2imgAddInferenceJob API
func CreatePersonalizedtxt2imgAddInferenceJobRequest() (request *Personalizedtxt2imgAddInferenceJobRequest) {
	request = &Personalizedtxt2imgAddInferenceJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AiContent", "20240611", "Personalizedtxt2imgAddInferenceJob", "/api/v1/personalizedtxt2img/addInferenceJob", "", "")
	request.Method = requests.POST
	return
}

// CreatePersonalizedtxt2imgAddInferenceJobResponse creates a response to parse from Personalizedtxt2imgAddInferenceJob response
func CreatePersonalizedtxt2imgAddInferenceJobResponse() (response *Personalizedtxt2imgAddInferenceJobResponse) {
	response = &Personalizedtxt2imgAddInferenceJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
