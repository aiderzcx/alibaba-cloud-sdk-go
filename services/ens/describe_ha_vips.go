package ens

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

// DescribeHaVips invokes the ens.DescribeHaVips API synchronously
func (client *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
	response = CreateDescribeHaVipsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHaVipsWithChan invokes the ens.DescribeHaVips API asynchronously
func (client *Client) DescribeHaVipsWithChan(request *DescribeHaVipsRequest) (<-chan *DescribeHaVipsResponse, <-chan error) {
	responseChan := make(chan *DescribeHaVipsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHaVips(request)
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

// DescribeHaVipsWithCallback invokes the ens.DescribeHaVips API asynchronously
func (client *Client) DescribeHaVipsWithCallback(request *DescribeHaVipsRequest, callback func(response *DescribeHaVipsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHaVipsResponse
		var err error
		defer close(result)
		response, err = client.DescribeHaVips(request)
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

// DescribeHaVipsRequest is the request struct for api DescribeHaVips
type DescribeHaVipsRequest struct {
	*requests.RpcRequest
	EnsRegionId  string `position:"Query" name:"EnsRegionId"`
	HaVipId      string `position:"Query" name:"HaVipId"`
	HaVipAddress string `position:"Query" name:"HaVipAddress"`
	VSwitchId    string `position:"Query" name:"VSwitchId"`
	Name         string `position:"Query" name:"Name"`
	NetworkId    string `position:"Query" name:"NetworkId"`
	Status       string `position:"Query" name:"Status"`
	PageNumber   string `position:"Query" name:"PageNumber"`
	PageSize     string `position:"Query" name:"PageSize"`
}

// DescribeHaVipsResponse is the response struct for api DescribeHaVips
type DescribeHaVipsResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	TotalCount string  `json:"TotalCount" xml:"TotalCount"`
	PageSize   string  `json:"PageSize" xml:"PageSize"`
	PageNumber string  `json:"PageNumber" xml:"PageNumber"`
	HaVips     []HaVip `json:"HaVips" xml:"HaVips"`
}

// CreateDescribeHaVipsRequest creates a request to invoke DescribeHaVips API
func CreateDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
	request = &DescribeHaVipsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeHaVips", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeHaVipsResponse creates a response to parse from DescribeHaVips response
func CreateDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
	response = &DescribeHaVipsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
