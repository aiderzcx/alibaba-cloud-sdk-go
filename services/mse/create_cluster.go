package mse

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

// CreateCluster invokes the mse.CreateCluster API synchronously
func (client *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	response = CreateCreateClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateClusterWithChan invokes the mse.CreateCluster API asynchronously
func (client *Client) CreateClusterWithChan(request *CreateClusterRequest) (<-chan *CreateClusterResponse, <-chan error) {
	responseChan := make(chan *CreateClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCluster(request)
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

// CreateClusterWithCallback invokes the mse.CreateCluster API asynchronously
func (client *Client) CreateClusterWithCallback(request *CreateClusterRequest, callback func(response *CreateClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateCluster(request)
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

// CreateClusterRequest is the request struct for api CreateCluster
type CreateClusterRequest struct {
	*requests.RpcRequest
	ClusterSpecification    string              `position:"Query" name:"ClusterSpecification"`
	MseSessionId            string              `position:"Query" name:"MseSessionId"`
	PubSlbSpecification     string              `position:"Query" name:"PubSlbSpecification"`
	PrivateSlbSpecification string              `position:"Query" name:"PrivateSlbSpecification"`
	ResourceGroupId         string              `position:"Query" name:"ResourceGroupId"`
	InstanceCount           requests.Integer    `position:"Query" name:"InstanceCount"`
	RequestPars             string              `position:"Query" name:"RequestPars"`
	Tag                     *[]CreateClusterTag `position:"Query" name:"Tag"  type:"Repeated"`
	ConnectionType          string              `position:"Query" name:"ConnectionType"`
	ClusterVersion          string              `position:"Query" name:"ClusterVersion"`
	DiskCapacity            requests.Integer    `position:"Query" name:"DiskCapacity"`
	DiskType                string              `position:"Query" name:"DiskType"`
	VSwitchId               string              `position:"Query" name:"VSwitchId"`
	ClusterType             string              `position:"Query" name:"ClusterType"`
	InstanceName            string              `position:"Query" name:"InstanceName"`
	PubNetworkFlow          string              `position:"Query" name:"PubNetworkFlow"`
	VpcId                   string              `position:"Query" name:"VpcId"`
	NetType                 string              `position:"Query" name:"NetType"`
	MseVersion              string              `position:"Query" name:"MseVersion"`
	AcceptLanguage          string              `position:"Query" name:"AcceptLanguage"`
	Region                  string              `position:"Query" name:"Region"`
}

// CreateClusterTag is a repeated param struct in CreateClusterRequest
type CreateClusterTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateClusterResponse is the response struct for api CreateCluster
type CreateClusterResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Message    string `json:"Message" xml:"Message"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	OrderId    string `json:"OrderId" xml:"OrderId"`
	Success    bool   `json:"Success" xml:"Success"`
}

// CreateCreateClusterRequest creates a request to invoke CreateCluster API
func CreateCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateCluster", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateClusterResponse creates a response to parse from CreateCluster response
func CreateCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
