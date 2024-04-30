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

// DescribeResourceDetails invokes the rds.DescribeResourceDetails API synchronously
func (client *Client) DescribeResourceDetails(request *DescribeResourceDetailsRequest) (response *DescribeResourceDetailsResponse, err error) {
	response = CreateDescribeResourceDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceDetailsWithChan invokes the rds.DescribeResourceDetails API asynchronously
func (client *Client) DescribeResourceDetailsWithChan(request *DescribeResourceDetailsRequest) (<-chan *DescribeResourceDetailsResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourceDetails(request)
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

// DescribeResourceDetailsWithCallback invokes the rds.DescribeResourceDetails API asynchronously
func (client *Client) DescribeResourceDetailsWithCallback(request *DescribeResourceDetailsRequest, callback func(response *DescribeResourceDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceDetailsResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourceDetails(request)
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

// DescribeResourceDetailsRequest is the request struct for api DescribeResourceDetails
type DescribeResourceDetailsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeResourceDetailsResponse is the response struct for api DescribeResourceDetails
type DescribeResourceDetailsResponse struct {
	*responses.BaseResponse
	RequestId              string                       `json:"RequestId" xml:"RequestId"`
	Region                 string                       `json:"Region" xml:"Region"`
	InstanceStorageType    string                       `json:"InstanceStorageType" xml:"InstanceStorageType"`
	DbInstanceStorage      int64                        `json:"DbInstanceStorage" xml:"DbInstanceStorage"`
	DiskUsed               int64                        `json:"DiskUsed" xml:"DiskUsed"`
	BackupSize             int64                        `json:"BackupSize" xml:"BackupSize"`
	VpcId                  string                       `json:"VpcId" xml:"VpcId"`
	VSwitchId              string                       `json:"VSwitchId" xml:"VSwitchId"`
	SecurityIPList         string                       `json:"SecurityIPList" xml:"SecurityIPList"`
	DbProxyInstanceName    string                       `json:"DbProxyInstanceName" xml:"DbProxyInstanceName"`
	ResourceGroupId        string                       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	BackupLogSize          int64                        `json:"BackupLogSize" xml:"BackupLogSize"`
	BackupDataSize         int64                        `json:"BackupDataSize" xml:"BackupDataSize"`
	RdsEcsSecurityGroupRel []RdsEcsSecurityGroupRelItem `json:"RdsEcsSecurityGroupRel" xml:"RdsEcsSecurityGroupRel"`
}

// CreateDescribeResourceDetailsRequest creates a request to invoke DescribeResourceDetails API
func CreateDescribeResourceDetailsRequest() (request *DescribeResourceDetailsRequest) {
	request = &DescribeResourceDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeResourceDetails", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeResourceDetailsResponse creates a response to parse from DescribeResourceDetails response
func CreateDescribeResourceDetailsResponse() (response *DescribeResourceDetailsResponse) {
	response = &DescribeResourceDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
