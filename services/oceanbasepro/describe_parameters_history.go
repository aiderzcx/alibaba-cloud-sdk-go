package oceanbasepro

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

// DescribeParametersHistory invokes the oceanbasepro.DescribeParametersHistory API synchronously
func (client *Client) DescribeParametersHistory(request *DescribeParametersHistoryRequest) (response *DescribeParametersHistoryResponse, err error) {
	response = CreateDescribeParametersHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeParametersHistoryWithChan invokes the oceanbasepro.DescribeParametersHistory API asynchronously
func (client *Client) DescribeParametersHistoryWithChan(request *DescribeParametersHistoryRequest) (<-chan *DescribeParametersHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeParametersHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeParametersHistory(request)
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

// DescribeParametersHistoryWithCallback invokes the oceanbasepro.DescribeParametersHistory API asynchronously
func (client *Client) DescribeParametersHistoryWithCallback(request *DescribeParametersHistoryRequest, callback func(response *DescribeParametersHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeParametersHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeParametersHistory(request)
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

// DescribeParametersHistoryRequest is the request struct for api DescribeParametersHistory
type DescribeParametersHistoryRequest struct {
	*requests.RpcRequest
	StartTime      string           `position:"Body" name:"StartTime"`
	PageNumber     requests.Integer `position:"Body" name:"PageNumber"`
	DimensionValue string           `position:"Body" name:"DimensionValue"`
	PageSize       requests.Integer `position:"Body" name:"PageSize"`
	Dimension      string           `position:"Body" name:"Dimension"`
	EndTime        string           `position:"Body" name:"EndTime"`
	InstanceId     string           `position:"Body" name:"InstanceId"`
}

// DescribeParametersHistoryResponse is the response struct for api DescribeParametersHistory
type DescribeParametersHistoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Respond   []Data `json:"Respond" xml:"Respond"`
}

// CreateDescribeParametersHistoryRequest creates a request to invoke DescribeParametersHistory API
func CreateDescribeParametersHistoryRequest() (request *DescribeParametersHistoryRequest) {
	request = &DescribeParametersHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeParametersHistory", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeParametersHistoryResponse creates a response to parse from DescribeParametersHistory response
func CreateDescribeParametersHistoryResponse() (response *DescribeParametersHistoryResponse) {
	response = &DescribeParametersHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
