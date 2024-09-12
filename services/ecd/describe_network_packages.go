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

// DescribeNetworkPackages invokes the ecd.DescribeNetworkPackages API synchronously
func (client *Client) DescribeNetworkPackages(request *DescribeNetworkPackagesRequest) (response *DescribeNetworkPackagesResponse, err error) {
	response = CreateDescribeNetworkPackagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkPackagesWithChan invokes the ecd.DescribeNetworkPackages API asynchronously
func (client *Client) DescribeNetworkPackagesWithChan(request *DescribeNetworkPackagesRequest) (<-chan *DescribeNetworkPackagesResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkPackagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkPackages(request)
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

// DescribeNetworkPackagesWithCallback invokes the ecd.DescribeNetworkPackages API asynchronously
func (client *Client) DescribeNetworkPackagesWithCallback(request *DescribeNetworkPackagesRequest, callback func(response *DescribeNetworkPackagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkPackagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkPackages(request)
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

// DescribeNetworkPackagesRequest is the request struct for api DescribeNetworkPackages
type DescribeNetworkPackagesRequest struct {
	*requests.RpcRequest
	NetworkPackageId   *[]string        `position:"Query" name:"NetworkPackageId"  type:"Repeated"`
	NextToken          string           `position:"Query" name:"NextToken"`
	InternetChargeType string           `position:"Query" name:"InternetChargeType"`
	MaxResults         requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeNetworkPackagesResponse is the response struct for api DescribeNetworkPackages
type DescribeNetworkPackagesResponse struct {
	*responses.BaseResponse
	NextToken       string           `json:"NextToken" xml:"NextToken"`
	RequestId       string           `json:"RequestId" xml:"RequestId"`
	NetworkPackages []NetworkPackage `json:"NetworkPackages" xml:"NetworkPackages"`
}

// CreateDescribeNetworkPackagesRequest creates a request to invoke DescribeNetworkPackages API
func CreateDescribeNetworkPackagesRequest() (request *DescribeNetworkPackagesRequest) {
	request = &DescribeNetworkPackagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeNetworkPackages", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkPackagesResponse creates a response to parse from DescribeNetworkPackages response
func CreateDescribeNetworkPackagesResponse() (response *DescribeNetworkPackagesResponse) {
	response = &DescribeNetworkPackagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
