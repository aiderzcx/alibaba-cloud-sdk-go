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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ApplyCoordinationForMonitoring invokes the ecd.ApplyCoordinationForMonitoring API synchronously
func (client *Client) ApplyCoordinationForMonitoring(request *ApplyCoordinationForMonitoringRequest) (response *ApplyCoordinationForMonitoringResponse, err error) {
	response = CreateApplyCoordinationForMonitoringResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyCoordinationForMonitoringWithChan invokes the ecd.ApplyCoordinationForMonitoring API asynchronously
func (client *Client) ApplyCoordinationForMonitoringWithChan(request *ApplyCoordinationForMonitoringRequest) (<-chan *ApplyCoordinationForMonitoringResponse, <-chan error) {
	responseChan := make(chan *ApplyCoordinationForMonitoringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyCoordinationForMonitoring(request)
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

// ApplyCoordinationForMonitoringWithCallback invokes the ecd.ApplyCoordinationForMonitoring API asynchronously
func (client *Client) ApplyCoordinationForMonitoringWithCallback(request *ApplyCoordinationForMonitoringRequest, callback func(response *ApplyCoordinationForMonitoringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyCoordinationForMonitoringResponse
		var err error
		defer close(result)
		response, err = client.ApplyCoordinationForMonitoring(request)
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

// ApplyCoordinationForMonitoringRequest is the request struct for api ApplyCoordinationForMonitoring
type ApplyCoordinationForMonitoringRequest struct {
	*requests.RpcRequest
	Uuid                 string                                              `position:"Query" name:"Uuid"`
	InitiatorType        string                                              `position:"Query" name:"InitiatorType"`
	CoordinatePolicyType string                                              `position:"Query" name:"CoordinatePolicyType"`
	ResourceCandidates   *[]ApplyCoordinationForMonitoringResourceCandidates `position:"Query" name:"ResourceCandidates"  type:"Repeated"`
	EndUserId            string                                              `position:"Query" name:"EndUserId"`
}

// ApplyCoordinationForMonitoringResourceCandidates is a repeated param struct in ApplyCoordinationForMonitoringRequest
type ApplyCoordinationForMonitoringResourceCandidates struct {
	ResourceId         string `name:"ResourceId"`
	ResourceProperties string `name:"ResourceProperties"`
	ResourceName       string `name:"ResourceName"`
	ResourceType       string `name:"ResourceType"`
	OwnerWyId          string `name:"OwnerWyId"`
	ResourceRegionId   string `name:"ResourceRegionId"`
	OwnerAliUid        string `name:"OwnerAliUid"`
	OwnerEndUserId     string `name:"OwnerEndUserId"`
}

// ApplyCoordinationForMonitoringResponse is the response struct for api ApplyCoordinationForMonitoring
type ApplyCoordinationForMonitoringResponse struct {
	*responses.BaseResponse
	RequestId            string                `json:"RequestId" xml:"RequestId"`
	CoordinateFlowModels []CoordinateFlowModel `json:"CoordinateFlowModels" xml:"CoordinateFlowModels"`
}

// CreateApplyCoordinationForMonitoringRequest creates a request to invoke ApplyCoordinationForMonitoring API
func CreateApplyCoordinationForMonitoringRequest() (request *ApplyCoordinationForMonitoringRequest) {
	request = &ApplyCoordinationForMonitoringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ApplyCoordinationForMonitoring", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateApplyCoordinationForMonitoringResponse creates a response to parse from ApplyCoordinationForMonitoring response
func CreateApplyCoordinationForMonitoringResponse() (response *ApplyCoordinationForMonitoringResponse) {
	response = &ApplyCoordinationForMonitoringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
