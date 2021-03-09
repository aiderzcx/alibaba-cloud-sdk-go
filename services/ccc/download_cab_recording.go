package ccc

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

// DownloadCabRecording invokes the ccc.DownloadCabRecording API synchronously
func (client *Client) DownloadCabRecording(request *DownloadCabRecordingRequest) (response *DownloadCabRecordingResponse, err error) {
	response = CreateDownloadCabRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadCabRecordingWithChan invokes the ccc.DownloadCabRecording API asynchronously
func (client *Client) DownloadCabRecordingWithChan(request *DownloadCabRecordingRequest) (<-chan *DownloadCabRecordingResponse, <-chan error) {
	responseChan := make(chan *DownloadCabRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadCabRecording(request)
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

// DownloadCabRecordingWithCallback invokes the ccc.DownloadCabRecording API asynchronously
func (client *Client) DownloadCabRecordingWithCallback(request *DownloadCabRecordingRequest, callback func(response *DownloadCabRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadCabRecordingResponse
		var err error
		defer close(result)
		response, err = client.DownloadCabRecording(request)
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

// DownloadCabRecordingRequest is the request struct for api DownloadCabRecording
type DownloadCabRecordingRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	TaskId     string `position:"Query" name:"TaskId"`
}

// DownloadCabRecordingResponse is the response struct for api DownloadCabRecording
type DownloadCabRecordingResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Success            bool               `json:"Success" xml:"Success"`
	Code               string             `json:"Code" xml:"Code"`
	Message            string             `json:"Message" xml:"Message"`
	HttpStatusCode     int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	MediaDownloadParam MediaDownloadParam `json:"MediaDownloadParam" xml:"MediaDownloadParam"`
}

// CreateDownloadCabRecordingRequest creates a request to invoke DownloadCabRecording API
func CreateDownloadCabRecordingRequest() (request *DownloadCabRecordingRequest) {
	request = &DownloadCabRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "DownloadCabRecording", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDownloadCabRecordingResponse creates a response to parse from DownloadCabRecording response
func CreateDownloadCabRecordingResponse() (response *DownloadCabRecordingResponse) {
	response = &DownloadCabRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
