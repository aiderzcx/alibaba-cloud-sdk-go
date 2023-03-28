package pts

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

// RemoveOpenJMeterScene invokes the pts.RemoveOpenJMeterScene API synchronously
func (client *Client) RemoveOpenJMeterScene(request *RemoveOpenJMeterSceneRequest) (response *RemoveOpenJMeterSceneResponse, err error) {
	response = CreateRemoveOpenJMeterSceneResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveOpenJMeterSceneWithChan invokes the pts.RemoveOpenJMeterScene API asynchronously
func (client *Client) RemoveOpenJMeterSceneWithChan(request *RemoveOpenJMeterSceneRequest) (<-chan *RemoveOpenJMeterSceneResponse, <-chan error) {
	responseChan := make(chan *RemoveOpenJMeterSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveOpenJMeterScene(request)
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

// RemoveOpenJMeterSceneWithCallback invokes the pts.RemoveOpenJMeterScene API asynchronously
func (client *Client) RemoveOpenJMeterSceneWithCallback(request *RemoveOpenJMeterSceneRequest, callback func(response *RemoveOpenJMeterSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveOpenJMeterSceneResponse
		var err error
		defer close(result)
		response, err = client.RemoveOpenJMeterScene(request)
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

// RemoveOpenJMeterSceneRequest is the request struct for api RemoveOpenJMeterScene
type RemoveOpenJMeterSceneRequest struct {
	*requests.RpcRequest
	SceneId string `position:"Query" name:"SceneId"`
}

// RemoveOpenJMeterSceneResponse is the response struct for api RemoveOpenJMeterScene
type RemoveOpenJMeterSceneResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateRemoveOpenJMeterSceneRequest creates a request to invoke RemoveOpenJMeterScene API
func CreateRemoveOpenJMeterSceneRequest() (request *RemoveOpenJMeterSceneRequest) {
	request = &RemoveOpenJMeterSceneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "RemoveOpenJMeterScene", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveOpenJMeterSceneResponse creates a response to parse from RemoveOpenJMeterScene response
func CreateRemoveOpenJMeterSceneResponse() (response *RemoveOpenJMeterSceneResponse) {
	response = &RemoveOpenJMeterSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
