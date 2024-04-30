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

// DescribeDBInstances invokes the rds.DescribeDBInstances API synchronously
func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
	response = CreateDescribeDBInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancesWithChan invokes the rds.DescribeDBInstances API asynchronously
func (client *Client) DescribeDBInstancesWithChan(request *DescribeDBInstancesRequest) (<-chan *DescribeDBInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstances(request)
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

// DescribeDBInstancesWithCallback invokes the rds.DescribeDBInstances API asynchronously
func (client *Client) DescribeDBInstancesWithCallback(request *DescribeDBInstancesRequest, callback func(response *DescribeDBInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstances(request)
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

// DescribeDBInstancesRequest is the request struct for api DescribeDBInstances
type DescribeDBInstancesRequest struct {
	*requests.RpcRequest
	Tag4Value            string           `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string           `position:"Query" name:"Tag.2.key"`
	ConnectionString     string           `position:"Query" name:"ConnectionString"`
	NeedVpcName          requests.Boolean `position:"Query" name:"NeedVpcName"`
	Tag3Key              string           `position:"Query" name:"Tag.3.key"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	Tag1Value            string           `position:"Query" name:"Tag.1.value"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.key"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceType       string           `position:"Query" name:"DBInstanceType"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	Tags                 string           `position:"Query" name:"Tags"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	Tag4Key              string           `position:"Query" name:"Tag.4.key"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
	InstanceNetworkType  string           `position:"Query" name:"InstanceNetworkType"`
	ConnectionMode       string           `position:"Query" name:"ConnectionMode"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	InstanceLevel        requests.Integer `position:"Query" name:"InstanceLevel"`
	SearchKey            string           `position:"Query" name:"SearchKey"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Expired              string           `position:"Query" name:"Expired"`
	Engine               string           `position:"Query" name:"Engine"`
	NextToken            string           `position:"Query" name:"NextToken"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceStatus     string           `position:"Query" name:"DBInstanceStatus"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	DedicatedHostGroupId string           `position:"Query" name:"DedicatedHostGroupId"`
	Tag3Value            string           `position:"Query" name:"Tag.3.value"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	Tag5Value            string           `position:"Query" name:"Tag.5.value"`
	Tag1Key              string           `position:"Query" name:"Tag.1.key"`
	Filter               string           `position:"Query" name:"Filter"`
	VpcId                string           `position:"Query" name:"VpcId"`
	Tag2Value            string           `position:"Query" name:"Tag.2.value"`
	Category             string           `position:"Query" name:"Category"`
	PayType              string           `position:"Query" name:"PayType"`
}

// DescribeDBInstancesResponse is the response struct for api DescribeDBInstances
type DescribeDBInstancesResponse struct {
	*responses.BaseResponse
	RequestId        string                     `json:"RequestId" xml:"RequestId"`
	PageNumber       int                        `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                        `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                        `json:"TotalRecordCount" xml:"TotalRecordCount"`
	NextToken        string                     `json:"NextToken" xml:"NextToken"`
	Items            ItemsInDescribeDBInstances `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstancesRequest creates a request to invoke DescribeDBInstances API
func CreateDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstances", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstancesResponse creates a response to parse from DescribeDBInstances response
func CreateDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
	response = &DescribeDBInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
