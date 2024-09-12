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

// SetDesktopGroupScaleTimer invokes the ecd.SetDesktopGroupScaleTimer API synchronously
func (client *Client) SetDesktopGroupScaleTimer(request *SetDesktopGroupScaleTimerRequest) (response *SetDesktopGroupScaleTimerResponse, err error) {
	response = CreateSetDesktopGroupScaleTimerResponse()
	err = client.DoAction(request, response)
	return
}

// SetDesktopGroupScaleTimerWithChan invokes the ecd.SetDesktopGroupScaleTimer API asynchronously
func (client *Client) SetDesktopGroupScaleTimerWithChan(request *SetDesktopGroupScaleTimerRequest) (<-chan *SetDesktopGroupScaleTimerResponse, <-chan error) {
	responseChan := make(chan *SetDesktopGroupScaleTimerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDesktopGroupScaleTimer(request)
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

// SetDesktopGroupScaleTimerWithCallback invokes the ecd.SetDesktopGroupScaleTimer API asynchronously
func (client *Client) SetDesktopGroupScaleTimerWithCallback(request *SetDesktopGroupScaleTimerRequest, callback func(response *SetDesktopGroupScaleTimerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDesktopGroupScaleTimerResponse
		var err error
		defer close(result)
		response, err = client.SetDesktopGroupScaleTimer(request)
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

// SetDesktopGroupScaleTimerRequest is the request struct for api SetDesktopGroupScaleTimer
type SetDesktopGroupScaleTimerRequest struct {
	*requests.RpcRequest
	DesktopGroupId  string                                      `position:"Query" name:"DesktopGroupId"`
	ScaleTimerInfos *[]SetDesktopGroupScaleTimerScaleTimerInfos `position:"Query" name:"ScaleTimerInfos"  type:"Repeated"`
}

// SetDesktopGroupScaleTimerScaleTimerInfos is a repeated param struct in SetDesktopGroupScaleTimerRequest
type SetDesktopGroupScaleTimerScaleTimerInfos struct {
	KeepDuration   string `name:"KeepDuration"`
	MinResAmount   string `name:"MinResAmount"`
	Cron           string `name:"Cron"`
	BuyResAmount   string `name:"BuyResAmount"`
	MaxResAmount   string `name:"MaxResAmount"`
	Type           string `name:"Type"`
	LoadPolicy     string `name:"LoadPolicy"`
	RatioThreshold string `name:"RatioThreshold"`
}

// SetDesktopGroupScaleTimerResponse is the response struct for api SetDesktopGroupScaleTimer
type SetDesktopGroupScaleTimerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDesktopGroupScaleTimerRequest creates a request to invoke SetDesktopGroupScaleTimer API
func CreateSetDesktopGroupScaleTimerRequest() (request *SetDesktopGroupScaleTimerRequest) {
	request = &SetDesktopGroupScaleTimerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "SetDesktopGroupScaleTimer", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDesktopGroupScaleTimerResponse creates a response to parse from SetDesktopGroupScaleTimer response
func CreateSetDesktopGroupScaleTimerResponse() (response *SetDesktopGroupScaleTimerResponse) {
	response = &SetDesktopGroupScaleTimerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
