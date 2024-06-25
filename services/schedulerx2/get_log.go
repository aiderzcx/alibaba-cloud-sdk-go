package schedulerx2

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

// GetLog invokes the schedulerx2.GetLog API synchronously
func (client *Client) GetLog(request *GetLogRequest) (response *GetLogResponse, err error) {
	response = CreateGetLogResponse()
	err = client.DoAction(request, response)
	return
}

// GetLogWithChan invokes the schedulerx2.GetLog API asynchronously
func (client *Client) GetLogWithChan(request *GetLogRequest) (<-chan *GetLogResponse, <-chan error) {
	responseChan := make(chan *GetLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLog(request)
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

// GetLogWithCallback invokes the schedulerx2.GetLog API asynchronously
func (client *Client) GetLogWithCallback(request *GetLogRequest, callback func(response *GetLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLogResponse
		var err error
		defer close(result)
		response, err = client.GetLog(request)
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

// GetLogRequest is the request struct for api GetLog
type GetLogRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	Line            requests.Integer `position:"Query" name:"Line"`
	StartTimestamp  requests.Integer `position:"Query" name:"StartTimestamp"`
	EndTimestamp    requests.Integer `position:"Query" name:"EndTimestamp"`
	JobId           string           `position:"Query" name:"JobId"`
	Keyword         string           `position:"Query" name:"Keyword"`
	Offset          requests.Integer `position:"Query" name:"Offset"`
	GroupId         string           `position:"Query" name:"GroupId"`
	Reverse         requests.Boolean `position:"Query" name:"Reverse"`
	Namespace       string           `position:"Query" name:"Namespace"`
	JobInstanceId   string           `position:"Query" name:"JobInstanceId"`
}

// GetLogResponse is the response struct for api GetLog
type GetLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetLogRequest creates a request to invoke GetLog API
func CreateGetLogRequest() (request *GetLogRequest) {
	request = &GetLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "GetLog", "schedulerx2", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetLogResponse creates a response to parse from GetLog response
func CreateGetLogResponse() (response *GetLogResponse) {
	response = &GetLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
