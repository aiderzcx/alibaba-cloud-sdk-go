package rds

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

// DeleteGadInstance invokes the rds.DeleteGadInstance API synchronously
func (client *Client) DeleteGadInstance(request *DeleteGadInstanceRequest) (response *DeleteGadInstanceResponse, err error) {
	response = CreateDeleteGadInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGadInstanceWithChan invokes the rds.DeleteGadInstance API asynchronously
func (client *Client) DeleteGadInstanceWithChan(request *DeleteGadInstanceRequest) (<-chan *DeleteGadInstanceResponse, <-chan error) {
	responseChan := make(chan *DeleteGadInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGadInstance(request)
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

// DeleteGadInstanceWithCallback invokes the rds.DeleteGadInstance API asynchronously
func (client *Client) DeleteGadInstanceWithCallback(request *DeleteGadInstanceRequest, callback func(response *DeleteGadInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGadInstanceResponse
		var err error
		defer close(result)
		response, err = client.DeleteGadInstance(request)
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

// DeleteGadInstanceRequest is the request struct for api DeleteGadInstance
type DeleteGadInstanceRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	GadInstanceName string `position:"Query" name:"GadInstanceName"`
}

// DeleteGadInstanceResponse is the response struct for api DeleteGadInstance
type DeleteGadInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteGadInstanceRequest creates a request to invoke DeleteGadInstance API
func CreateDeleteGadInstanceRequest() (request *DeleteGadInstanceRequest) {
	request = &DeleteGadInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteGadInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGadInstanceResponse creates a response to parse from DeleteGadInstance response
func CreateDeleteGadInstanceResponse() (response *DeleteGadInstanceResponse) {
	response = &DeleteGadInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
