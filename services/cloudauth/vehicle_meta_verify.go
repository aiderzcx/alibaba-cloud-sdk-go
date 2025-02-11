package cloudauth

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

// VehicleMetaVerify invokes the cloudauth.VehicleMetaVerify API synchronously
func (client *Client) VehicleMetaVerify(request *VehicleMetaVerifyRequest) (response *VehicleMetaVerifyResponse, err error) {
	response = CreateVehicleMetaVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// VehicleMetaVerifyWithChan invokes the cloudauth.VehicleMetaVerify API asynchronously
func (client *Client) VehicleMetaVerifyWithChan(request *VehicleMetaVerifyRequest) (<-chan *VehicleMetaVerifyResponse, <-chan error) {
	responseChan := make(chan *VehicleMetaVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VehicleMetaVerify(request)
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

// VehicleMetaVerifyWithCallback invokes the cloudauth.VehicleMetaVerify API asynchronously
func (client *Client) VehicleMetaVerifyWithCallback(request *VehicleMetaVerifyRequest, callback func(response *VehicleMetaVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VehicleMetaVerifyResponse
		var err error
		defer close(result)
		response, err = client.VehicleMetaVerify(request)
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

// VehicleMetaVerifyRequest is the request struct for api VehicleMetaVerify
type VehicleMetaVerifyRequest struct {
	*requests.RpcRequest
	VehicleType    string `position:"Query" name:"VehicleType"`
	ParamType      string `position:"Query" name:"ParamType"`
	VehicleNum     string `position:"Query" name:"VehicleNum"`
	IdentifyNum    string `position:"Query" name:"IdentifyNum"`
	VerifyMetaType string `position:"Query" name:"VerifyMetaType"`
	UserName       string `position:"Query" name:"UserName"`
}

// VehicleMetaVerifyResponse is the response struct for api VehicleMetaVerify
type VehicleMetaVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateVehicleMetaVerifyRequest creates a request to invoke VehicleMetaVerify API
func CreateVehicleMetaVerifyRequest() (request *VehicleMetaVerifyRequest) {
	request = &VehicleMetaVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "VehicleMetaVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVehicleMetaVerifyResponse creates a response to parse from VehicleMetaVerify response
func CreateVehicleMetaVerifyResponse() (response *VehicleMetaVerifyResponse) {
	response = &VehicleMetaVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
