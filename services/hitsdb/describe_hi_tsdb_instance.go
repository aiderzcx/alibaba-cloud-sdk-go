package hitsdb

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

// DescribeHiTSDBInstance invokes the hitsdb.DescribeHiTSDBInstance API synchronously
func (client *Client) DescribeHiTSDBInstance(request *DescribeHiTSDBInstanceRequest) (response *DescribeHiTSDBInstanceResponse, err error) {
	response = CreateDescribeHiTSDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHiTSDBInstanceWithChan invokes the hitsdb.DescribeHiTSDBInstance API asynchronously
func (client *Client) DescribeHiTSDBInstanceWithChan(request *DescribeHiTSDBInstanceRequest) (<-chan *DescribeHiTSDBInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeHiTSDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHiTSDBInstance(request)
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

// DescribeHiTSDBInstanceWithCallback invokes the hitsdb.DescribeHiTSDBInstance API asynchronously
func (client *Client) DescribeHiTSDBInstanceWithCallback(request *DescribeHiTSDBInstanceRequest, callback func(response *DescribeHiTSDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHiTSDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeHiTSDBInstance(request)
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

// DescribeHiTSDBInstanceRequest is the request struct for api DescribeHiTSDBInstance
type DescribeHiTSDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeHiTSDBInstanceResponse is the response struct for api DescribeHiTSDBInstance
type DescribeHiTSDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId              string       `json:"RequestId" xml:"RequestId"`
	InstanceId             string       `json:"InstanceId" xml:"InstanceId"`
	InstanceAlias          string       `json:"InstanceAlias" xml:"InstanceAlias"`
	InstanceDescription    string       `json:"InstanceDescription" xml:"InstanceDescription"`
	RegionId               string       `json:"RegionId" xml:"RegionId"`
	ZoneId                 string       `json:"ZoneId" xml:"ZoneId"`
	InstanceStatus         string       `json:"InstanceStatus" xml:"InstanceStatus"`
	ChargeType             string       `json:"ChargeType" xml:"ChargeType"`
	NetworkType            string       `json:"NetworkType" xml:"NetworkType"`
	GmtCreated             string       `json:"GmtCreated" xml:"GmtCreated"`
	GmtExpire              string       `json:"GmtExpire" xml:"GmtExpire"`
	InstanceClass          string       `json:"InstanceClass" xml:"InstanceClass"`
	MaxTimelineLimit       string       `json:"MaxTimelineLimit" xml:"MaxTimelineLimit"`
	InstanceStorage        string       `json:"InstanceStorage" xml:"InstanceStorage"`
	InstanceTps            string       `json:"InstanceTps" xml:"InstanceTps"`
	ReverseVpcIp           string       `json:"ReverseVpcIp" xml:"ReverseVpcIp"`
	ReverseVpcPort         string       `json:"ReverseVpcPort" xml:"ReverseVpcPort"`
	VpcId                  string       `json:"VpcId" xml:"VpcId"`
	VswitchId              string       `json:"VswitchId" xml:"VswitchId"`
	ConnectionString       string       `json:"ConnectionString" xml:"ConnectionString"`
	PublicConnectionString string       `json:"PublicConnectionString" xml:"PublicConnectionString"`
	AutoRenew              bool         `json:"AutoRenew" xml:"AutoRenew"`
	EngineType             string       `json:"EngineType" xml:"EngineType"`
	CpuNumber              string       `json:"CpuNumber" xml:"CpuNumber"`
	MemSize                string       `json:"MemSize" xml:"MemSize"`
	Series                 int          `json:"Series" xml:"Series"`
	RDSStatus              string       `json:"RDSStatus" xml:"RDSStatus"`
	DiskCategory           string       `json:"DiskCategory" xml:"DiskCategory"`
	SecurityIpList         []SecurityIp `json:"SecurityIpList" xml:"SecurityIpList"`
}

// CreateDescribeHiTSDBInstanceRequest creates a request to invoke DescribeHiTSDBInstance API
func CreateDescribeHiTSDBInstanceRequest() (request *DescribeHiTSDBInstanceRequest) {
	request = &DescribeHiTSDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2017-06-01", "DescribeHiTSDBInstance", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHiTSDBInstanceResponse creates a response to parse from DescribeHiTSDBInstance response
func CreateDescribeHiTSDBInstanceResponse() (response *DescribeHiTSDBInstanceResponse) {
	response = &DescribeHiTSDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
