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

// VehicleInsureQuery invokes the cloudauth.VehicleInsureQuery API synchronously
func (client *Client) VehicleInsureQuery(request *VehicleInsureQueryRequest) (response *VehicleInsureQueryResponse, err error) {
	response = CreateVehicleInsureQueryResponse()
	err = client.DoAction(request, response)
	return
}

// VehicleInsureQueryWithChan invokes the cloudauth.VehicleInsureQuery API asynchronously
func (client *Client) VehicleInsureQueryWithChan(request *VehicleInsureQueryRequest) (<-chan *VehicleInsureQueryResponse, <-chan error) {
	responseChan := make(chan *VehicleInsureQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VehicleInsureQuery(request)
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

// VehicleInsureQueryWithCallback invokes the cloudauth.VehicleInsureQuery API asynchronously
func (client *Client) VehicleInsureQueryWithCallback(request *VehicleInsureQueryRequest, callback func(response *VehicleInsureQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VehicleInsureQueryResponse
		var err error
		defer close(result)
		response, err = client.VehicleInsureQuery(request)
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

// VehicleInsureQueryRequest is the request struct for api VehicleInsureQuery
type VehicleInsureQueryRequest struct {
	*requests.RpcRequest
	VehicleType string `position:"Query" name:"VehicleType"`
	ParamType   string `position:"Query" name:"ParamType"`
	VehicleNum  string `position:"Query" name:"VehicleNum"`
	Vin         string `position:"Query" name:"Vin"`
}

// VehicleInsureQueryResponse is the response struct for api VehicleInsureQuery
type VehicleInsureQueryResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateVehicleInsureQueryRequest creates a request to invoke VehicleInsureQuery API
func CreateVehicleInsureQueryRequest() (request *VehicleInsureQueryRequest) {
	request = &VehicleInsureQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "VehicleInsureQuery", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVehicleInsureQueryResponse creates a response to parse from VehicleInsureQuery response
func CreateVehicleInsureQueryResponse() (response *VehicleInsureQueryResponse) {
	response = &VehicleInsureQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
