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

// QueryGovernanceKubernetesCluster invokes the mse.QueryGovernanceKubernetesCluster API synchronously
func (client *Client) QueryGovernanceKubernetesCluster(request *QueryGovernanceKubernetesClusterRequest) (response *QueryGovernanceKubernetesClusterResponse, err error) {
	response = CreateQueryGovernanceKubernetesClusterResponse()
	err = client.DoAction(request, response)
	return
}

// QueryGovernanceKubernetesClusterWithChan invokes the mse.QueryGovernanceKubernetesCluster API asynchronously
func (client *Client) QueryGovernanceKubernetesClusterWithChan(request *QueryGovernanceKubernetesClusterRequest) (<-chan *QueryGovernanceKubernetesClusterResponse, <-chan error) {
	responseChan := make(chan *QueryGovernanceKubernetesClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryGovernanceKubernetesCluster(request)
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

// QueryGovernanceKubernetesClusterWithCallback invokes the mse.QueryGovernanceKubernetesCluster API asynchronously
func (client *Client) QueryGovernanceKubernetesClusterWithCallback(request *QueryGovernanceKubernetesClusterRequest, callback func(response *QueryGovernanceKubernetesClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryGovernanceKubernetesClusterResponse
		var err error
		defer close(result)
		response, err = client.QueryGovernanceKubernetesCluster(request)
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

// QueryGovernanceKubernetesClusterRequest is the request struct for api QueryGovernanceKubernetesCluster
type QueryGovernanceKubernetesClusterRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	ClusterName    string           `position:"Query" name:"ClusterName"`
	ClusterId      string           `position:"Query" name:"ClusterId"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// QueryGovernanceKubernetesClusterResponse is the response struct for api QueryGovernanceKubernetesCluster
type QueryGovernanceKubernetesClusterResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                                 `json:"Message" xml:"Message"`
	RequestId      string                                 `json:"RequestId" xml:"RequestId"`
	Code           int                                    `json:"Code" xml:"Code"`
	Success        bool                                   `json:"Success" xml:"Success"`
	Data           DataInQueryGovernanceKubernetesCluster `json:"Data" xml:"Data"`
}

// CreateQueryGovernanceKubernetesClusterRequest creates a request to invoke QueryGovernanceKubernetesCluster API
func CreateQueryGovernanceKubernetesClusterRequest() (request *QueryGovernanceKubernetesClusterRequest) {
	request = &QueryGovernanceKubernetesClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryGovernanceKubernetesCluster", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryGovernanceKubernetesClusterResponse creates a response to parse from QueryGovernanceKubernetesCluster response
func CreateQueryGovernanceKubernetesClusterResponse() (response *QueryGovernanceKubernetesClusterResponse) {
	response = &QueryGovernanceKubernetesClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
