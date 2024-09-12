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

// DescribeDesktopOversoldUserGroup invokes the ecd.DescribeDesktopOversoldUserGroup API synchronously
func (client *Client) DescribeDesktopOversoldUserGroup(request *DescribeDesktopOversoldUserGroupRequest) (response *DescribeDesktopOversoldUserGroupResponse, err error) {
	response = CreateDescribeDesktopOversoldUserGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDesktopOversoldUserGroupWithChan invokes the ecd.DescribeDesktopOversoldUserGroup API asynchronously
func (client *Client) DescribeDesktopOversoldUserGroupWithChan(request *DescribeDesktopOversoldUserGroupRequest) (<-chan *DescribeDesktopOversoldUserGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeDesktopOversoldUserGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDesktopOversoldUserGroup(request)
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

// DescribeDesktopOversoldUserGroupWithCallback invokes the ecd.DescribeDesktopOversoldUserGroup API asynchronously
func (client *Client) DescribeDesktopOversoldUserGroupWithCallback(request *DescribeDesktopOversoldUserGroupRequest, callback func(response *DescribeDesktopOversoldUserGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDesktopOversoldUserGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeDesktopOversoldUserGroup(request)
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

// DescribeDesktopOversoldUserGroupRequest is the request struct for api DescribeDesktopOversoldUserGroup
type DescribeDesktopOversoldUserGroupRequest struct {
	*requests.RpcRequest
	OversoldGroupId string           `position:"Query" name:"OversoldGroupId"`
	UserGroupIds    *[]string        `position:"Query" name:"UserGroupIds"  type:"Repeated"`
	NextToken       string           `position:"Query" name:"NextToken"`
	MaxResults      requests.Integer `position:"Query" name:"MaxResults"`
}

// DescribeDesktopOversoldUserGroupResponse is the response struct for api DescribeDesktopOversoldUserGroup
type DescribeDesktopOversoldUserGroupResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Count     int        `json:"Count" xml:"Count"`
	NextToken string     `json:"NextToken" xml:"NextToken"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeDesktopOversoldUserGroupRequest creates a request to invoke DescribeDesktopOversoldUserGroup API
func CreateDescribeDesktopOversoldUserGroupRequest() (request *DescribeDesktopOversoldUserGroupRequest) {
	request = &DescribeDesktopOversoldUserGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeDesktopOversoldUserGroup", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDesktopOversoldUserGroupResponse creates a response to parse from DescribeDesktopOversoldUserGroup response
func CreateDescribeDesktopOversoldUserGroupResponse() (response *DescribeDesktopOversoldUserGroupResponse) {
	response = &DescribeDesktopOversoldUserGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
