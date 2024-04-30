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

// PreCheckCreateOrderForDeleteDBNodes invokes the rds.PreCheckCreateOrderForDeleteDBNodes API synchronously
func (client *Client) PreCheckCreateOrderForDeleteDBNodes(request *PreCheckCreateOrderForDeleteDBNodesRequest) (response *PreCheckCreateOrderForDeleteDBNodesResponse, err error) {
	response = CreatePreCheckCreateOrderForDeleteDBNodesResponse()
	err = client.DoAction(request, response)
	return
}

// PreCheckCreateOrderForDeleteDBNodesWithChan invokes the rds.PreCheckCreateOrderForDeleteDBNodes API asynchronously
func (client *Client) PreCheckCreateOrderForDeleteDBNodesWithChan(request *PreCheckCreateOrderForDeleteDBNodesRequest) (<-chan *PreCheckCreateOrderForDeleteDBNodesResponse, <-chan error) {
	responseChan := make(chan *PreCheckCreateOrderForDeleteDBNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreCheckCreateOrderForDeleteDBNodes(request)
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

// PreCheckCreateOrderForDeleteDBNodesWithCallback invokes the rds.PreCheckCreateOrderForDeleteDBNodes API asynchronously
func (client *Client) PreCheckCreateOrderForDeleteDBNodesWithCallback(request *PreCheckCreateOrderForDeleteDBNodesRequest, callback func(response *PreCheckCreateOrderForDeleteDBNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreCheckCreateOrderForDeleteDBNodesResponse
		var err error
		defer close(result)
		response, err = client.PreCheckCreateOrderForDeleteDBNodes(request)
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

// PreCheckCreateOrderForDeleteDBNodesRequest is the request struct for api PreCheckCreateOrderForDeleteDBNodes
type PreCheckCreateOrderForDeleteDBNodesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NodeType             string           `position:"Query" name:"NodeType"`
	DBNodeId             string           `position:"Query" name:"DBNodeId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Resource             string           `position:"Query" name:"Resource"`
	CommodityCode        string           `position:"Query" name:"CommodityCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PromotionCode        string           `position:"Query" name:"PromotionCode"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// PreCheckCreateOrderForDeleteDBNodesResponse is the response struct for api PreCheckCreateOrderForDeleteDBNodes
type PreCheckCreateOrderForDeleteDBNodesResponse struct {
	*responses.BaseResponse
	PreCheckResult bool     `json:"PreCheckResult" xml:"PreCheckResult"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Failures       Failures `json:"Failures" xml:"Failures"`
}

// CreatePreCheckCreateOrderForDeleteDBNodesRequest creates a request to invoke PreCheckCreateOrderForDeleteDBNodes API
func CreatePreCheckCreateOrderForDeleteDBNodesRequest() (request *PreCheckCreateOrderForDeleteDBNodesRequest) {
	request = &PreCheckCreateOrderForDeleteDBNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "PreCheckCreateOrderForDeleteDBNodes", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePreCheckCreateOrderForDeleteDBNodesResponse creates a response to parse from PreCheckCreateOrderForDeleteDBNodes response
func CreatePreCheckCreateOrderForDeleteDBNodesResponse() (response *PreCheckCreateOrderForDeleteDBNodesResponse) {
	response = &PreCheckCreateOrderForDeleteDBNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
