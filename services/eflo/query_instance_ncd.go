package eflo

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

// QueryInstanceNcd invokes the eflo.QueryInstanceNcd API synchronously
func (client *Client) QueryInstanceNcd(request *QueryInstanceNcdRequest) (response *QueryInstanceNcdResponse, err error) {
	response = CreateQueryInstanceNcdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInstanceNcdWithChan invokes the eflo.QueryInstanceNcd API asynchronously
func (client *Client) QueryInstanceNcdWithChan(request *QueryInstanceNcdRequest) (<-chan *QueryInstanceNcdResponse, <-chan error) {
	responseChan := make(chan *QueryInstanceNcdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInstanceNcd(request)
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

// QueryInstanceNcdWithCallback invokes the eflo.QueryInstanceNcd API asynchronously
func (client *Client) QueryInstanceNcdWithCallback(request *QueryInstanceNcdRequest, callback func(response *QueryInstanceNcdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInstanceNcdResponse
		var err error
		defer close(result)
		response, err = client.QueryInstanceNcd(request)
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

// QueryInstanceNcdRequest is the request struct for api QueryInstanceNcd
type QueryInstanceNcdRequest struct {
	*requests.RpcRequest
	InstanceId2  string `position:"Body" name:"InstanceId2"`
	InstanceId1  string `position:"Body" name:"InstanceId1"`
	InstanceType string `position:"Body" name:"InstanceType"`
}

// QueryInstanceNcdResponse is the response struct for api QueryInstanceNcd
type QueryInstanceNcdResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateQueryInstanceNcdRequest creates a request to invoke QueryInstanceNcd API
func CreateQueryInstanceNcdRequest() (request *QueryInstanceNcdRequest) {
	request = &QueryInstanceNcdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "QueryInstanceNcd", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryInstanceNcdResponse creates a response to parse from QueryInstanceNcd response
func CreateQueryInstanceNcdResponse() (response *QueryInstanceNcdResponse) {
	response = &QueryInstanceNcdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
