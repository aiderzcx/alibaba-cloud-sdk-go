package ga

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

// DescribeListener invokes the ga.DescribeListener API synchronously
func (client *Client) DescribeListener(request *DescribeListenerRequest) (response *DescribeListenerResponse, err error) {
	response = CreateDescribeListenerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeListenerWithChan invokes the ga.DescribeListener API asynchronously
func (client *Client) DescribeListenerWithChan(request *DescribeListenerRequest) (<-chan *DescribeListenerResponse, <-chan error) {
	responseChan := make(chan *DescribeListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeListener(request)
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

// DescribeListenerWithCallback invokes the ga.DescribeListener API asynchronously
func (client *Client) DescribeListenerWithCallback(request *DescribeListenerRequest, callback func(response *DescribeListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeListenerResponse
		var err error
		defer close(result)
		response, err = client.DescribeListener(request)
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

// DescribeListenerRequest is the request struct for api DescribeListener
type DescribeListenerRequest struct {
	*requests.RpcRequest
	ListenerId string `position:"Query" name:"ListenerId"`
}

// DescribeListenerResponse is the response struct for api DescribeListener
type DescribeListenerResponse struct {
	*responses.BaseResponse
	ClientAffinity string           `json:"ClientAffinity" xml:"ClientAffinity"`
	CreateTime     string           `json:"CreateTime" xml:"CreateTime"`
	Description    string           `json:"Description" xml:"Description"`
	ListenerId     string           `json:"ListenerId" xml:"ListenerId"`
	Name           string           `json:"Name" xml:"Name"`
	Protocol       string           `json:"Protocol" xml:"Protocol"`
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	State          string           `json:"State" xml:"State"`
	PortRanges     []PortRangesItem `json:"PortRanges" xml:"PortRanges"`
	Certificates   []Certificate    `json:"Certificates" xml:"Certificates"`
	BackendPorts   []BackendPort    `json:"BackendPorts" xml:"BackendPorts"`
}

// CreateDescribeListenerRequest creates a request to invoke DescribeListener API
func CreateDescribeListenerRequest() (request *DescribeListenerRequest) {
	request = &DescribeListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "DescribeListener", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeListenerResponse creates a response to parse from DescribeListener response
func CreateDescribeListenerResponse() (response *DescribeListenerResponse) {
	response = &DescribeListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
