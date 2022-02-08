package iot

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

// CreateEdgeDriver invokes the iot.CreateEdgeDriver API synchronously
func (client *Client) CreateEdgeDriver(request *CreateEdgeDriverRequest) (response *CreateEdgeDriverResponse, err error) {
	response = CreateCreateEdgeDriverResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEdgeDriverWithChan invokes the iot.CreateEdgeDriver API asynchronously
func (client *Client) CreateEdgeDriverWithChan(request *CreateEdgeDriverRequest) (<-chan *CreateEdgeDriverResponse, <-chan error) {
	responseChan := make(chan *CreateEdgeDriverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEdgeDriver(request)
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

// CreateEdgeDriverWithCallback invokes the iot.CreateEdgeDriver API asynchronously
func (client *Client) CreateEdgeDriverWithCallback(request *CreateEdgeDriverRequest, callback func(response *CreateEdgeDriverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEdgeDriverResponse
		var err error
		defer close(result)
		response, err = client.CreateEdgeDriver(request)
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

// CreateEdgeDriverRequest is the request struct for api CreateEdgeDriver
type CreateEdgeDriverRequest struct {
	*requests.RpcRequest
	DriverProtocol    string           `position:"Query" name:"DriverProtocol"`
	DriverName        string           `position:"Query" name:"DriverName"`
	IsBuiltIn         requests.Boolean `position:"Query" name:"IsBuiltIn"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	Runtime           string           `position:"Query" name:"Runtime"`
	UseOfficialConfig requests.Integer `position:"Query" name:"UseOfficialConfig"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	CpuArch           string           `position:"Query" name:"CpuArch"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
}

// CreateEdgeDriverResponse is the response struct for api CreateEdgeDriver
type CreateEdgeDriverResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	DriverId     string `json:"DriverId" xml:"DriverId"`
}

// CreateCreateEdgeDriverRequest creates a request to invoke CreateEdgeDriver API
func CreateCreateEdgeDriverRequest() (request *CreateEdgeDriverRequest) {
	request = &CreateEdgeDriverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateEdgeDriver", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateEdgeDriverResponse creates a response to parse from CreateEdgeDriver response
func CreateCreateEdgeDriverResponse() (response *CreateEdgeDriverResponse) {
	response = &CreateEdgeDriverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
