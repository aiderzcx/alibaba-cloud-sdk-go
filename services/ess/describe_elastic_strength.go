package ess

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

// DescribeElasticStrength invokes the ess.DescribeElasticStrength API synchronously
func (client *Client) DescribeElasticStrength(request *DescribeElasticStrengthRequest) (response *DescribeElasticStrengthResponse, err error) {
	response = CreateDescribeElasticStrengthResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeElasticStrengthWithChan invokes the ess.DescribeElasticStrength API asynchronously
func (client *Client) DescribeElasticStrengthWithChan(request *DescribeElasticStrengthRequest) (<-chan *DescribeElasticStrengthResponse, <-chan error) {
	responseChan := make(chan *DescribeElasticStrengthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeElasticStrength(request)
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

// DescribeElasticStrengthWithCallback invokes the ess.DescribeElasticStrength API asynchronously
func (client *Client) DescribeElasticStrengthWithCallback(request *DescribeElasticStrengthRequest, callback func(response *DescribeElasticStrengthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeElasticStrengthResponse
		var err error
		defer close(result)
		response, err = client.DescribeElasticStrength(request)
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

// DescribeElasticStrengthRequest is the request struct for api DescribeElasticStrength
type DescribeElasticStrengthRequest struct {
	*requests.RpcRequest
	ImageId              string           `position:"Query" name:"ImageId"`
	DataDiskCategories   *[]string        `position:"Query" name:"DataDiskCategories"  type:"Repeated"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	VSwitchIds           *[]string        `position:"Query" name:"VSwitchIds"  type:"Repeated"`
	InstanceTypes        *[]string        `position:"Query" name:"InstanceTypes"  type:"Repeated"`
	ImageName            string           `position:"Query" name:"ImageName"`
	ScalingGroupIds      *[]string        `position:"Query" name:"ScalingGroupIds"  type:"Repeated"`
	Ipv6AddressCount     requests.Integer `position:"Query" name:"Ipv6AddressCount"`
	SystemDiskCategories *[]string        `position:"Query" name:"SystemDiskCategories"  type:"Repeated"`
	SpotStrategy         string           `position:"Query" name:"SpotStrategy"`
	PriorityStrategy     string           `position:"Query" name:"PriorityStrategy"`
	ImageFamily          string           `position:"Query" name:"ImageFamily"`
}

// DescribeElasticStrengthResponse is the response struct for api DescribeElasticStrength
type DescribeElasticStrengthResponse struct {
	*responses.BaseResponse
	RequestId             string                 `json:"RequestId" xml:"RequestId"`
	TotalStrength         string                 `json:"TotalStrength" xml:"TotalStrength"`
	ResourcePools         []ResourcePool         `json:"ResourcePools" xml:"ResourcePools"`
	ElasticStrengthModels []ElasticStrengthModel `json:"ElasticStrengthModels" xml:"ElasticStrengthModels"`
}

// CreateDescribeElasticStrengthRequest creates a request to invoke DescribeElasticStrength API
func CreateDescribeElasticStrengthRequest() (request *DescribeElasticStrengthRequest) {
	request = &DescribeElasticStrengthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeElasticStrength", "ess", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeElasticStrengthResponse creates a response to parse from DescribeElasticStrength response
func CreateDescribeElasticStrengthResponse() (response *DescribeElasticStrengthResponse) {
	response = &DescribeElasticStrengthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
