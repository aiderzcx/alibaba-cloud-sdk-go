package cloudauth

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

// DescribeFaceGuardRisk invokes the cloudauth.DescribeFaceGuardRisk API synchronously
func (client *Client) DescribeFaceGuardRisk(request *DescribeFaceGuardRiskRequest) (response *DescribeFaceGuardRiskResponse, err error) {
	response = CreateDescribeFaceGuardRiskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFaceGuardRiskWithChan invokes the cloudauth.DescribeFaceGuardRisk API asynchronously
func (client *Client) DescribeFaceGuardRiskWithChan(request *DescribeFaceGuardRiskRequest) (<-chan *DescribeFaceGuardRiskResponse, <-chan error) {
	responseChan := make(chan *DescribeFaceGuardRiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFaceGuardRisk(request)
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

// DescribeFaceGuardRiskWithCallback invokes the cloudauth.DescribeFaceGuardRisk API asynchronously
func (client *Client) DescribeFaceGuardRiskWithCallback(request *DescribeFaceGuardRiskRequest, callback func(response *DescribeFaceGuardRiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFaceGuardRiskResponse
		var err error
		defer close(result)
		response, err = client.DescribeFaceGuardRisk(request)
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

// DescribeFaceGuardRiskRequest is the request struct for api DescribeFaceGuardRisk
type DescribeFaceGuardRiskRequest struct {
	*requests.RpcRequest
	ProductCode  string `position:"Query" name:"ProductCode"`
	DeviceToken  string `position:"Query" name:"DeviceToken"`
	OuterOrderNo string `position:"Query" name:"OuterOrderNo"`
	BizId        string `position:"Query" name:"BizId"`
}

// DescribeFaceGuardRiskResponse is the response struct for api DescribeFaceGuardRisk
type DescribeFaceGuardRiskResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateDescribeFaceGuardRiskRequest creates a request to invoke DescribeFaceGuardRisk API
func CreateDescribeFaceGuardRiskRequest() (request *DescribeFaceGuardRiskRequest) {
	request = &DescribeFaceGuardRiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeFaceGuardRisk", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFaceGuardRiskResponse creates a response to parse from DescribeFaceGuardRisk response
func CreateDescribeFaceGuardRiskResponse() (response *DescribeFaceGuardRiskResponse) {
	response = &DescribeFaceGuardRiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
