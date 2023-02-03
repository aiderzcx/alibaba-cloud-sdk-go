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

// UpdateConfig invokes the mse.UpdateConfig API synchronously
func (client *Client) UpdateConfig(request *UpdateConfigRequest) (response *UpdateConfigResponse, err error) {
	response = CreateUpdateConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConfigWithChan invokes the mse.UpdateConfig API asynchronously
func (client *Client) UpdateConfigWithChan(request *UpdateConfigRequest) (<-chan *UpdateConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConfig(request)
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

// UpdateConfigWithCallback invokes the mse.UpdateConfig API asynchronously
func (client *Client) UpdateConfigWithCallback(request *UpdateConfigRequest, callback func(response *UpdateConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateConfig(request)
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

// UpdateConfigRequest is the request struct for api UpdateConfig
type UpdateConfigRequest struct {
	*requests.RpcRequest
	OpenSuperAcl             string           `position:"Body" name:"OpenSuperAcl"`
	MseSessionId             string           `position:"Query" name:"MseSessionId"`
	ConfigAuthEnabled        requests.Boolean `position:"Query" name:"ConfigAuthEnabled"`
	PassWord                 string           `position:"Query" name:"PassWord"`
	SnapshotCount            string           `position:"Query" name:"SnapshotCount"`
	MinSessionTimeout        string           `position:"Query" name:"MinSessionTimeout"`
	MaxClientCnxns           string           `position:"Query" name:"MaxClientCnxns"`
	RequestPars              string           `position:"Query" name:"RequestPars"`
	JuteMaxbuffer            string           `position:"Query" name:"JuteMaxbuffer"`
	NamingAuthEnabled        requests.Boolean `position:"Query" name:"NamingAuthEnabled"`
	ConfigType               string           `position:"Query" name:"ConfigType"`
	ExtendedTypesEnable      string           `position:"Query" name:"ExtendedTypesEnable"`
	AutopurgeSnapRetainCount string           `position:"Query" name:"AutopurgeSnapRetainCount"`
	MaxSessionTimeout        string           `position:"Query" name:"MaxSessionTimeout"`
	ConfigSecretEnabled      requests.Boolean `position:"Query" name:"ConfigSecretEnabled"`
	MCPEnabled               requests.Boolean `position:"Query" name:"MCPEnabled"`
	TickTime                 string           `position:"Query" name:"TickTime"`
	ClusterId                string           `position:"Query" name:"ClusterId"`
	SyncLimit                string           `position:"Query" name:"SyncLimit"`
	InstanceId               string           `position:"Query" name:"InstanceId"`
	AutopurgePurgeInterval   string           `position:"Query" name:"AutopurgePurgeInterval"`
	AcceptLanguage           string           `position:"Query" name:"AcceptLanguage"`
	InitLimit                string           `position:"Query" name:"InitLimit"`
	UserName                 string           `position:"Query" name:"UserName"`
}

// UpdateConfigResponse is the response struct for api UpdateConfig
type UpdateConfigResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           int    `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateUpdateConfigRequest creates a request to invoke UpdateConfig API
func CreateUpdateConfigRequest() (request *UpdateConfigRequest) {
	request = &UpdateConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateConfig", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateConfigResponse creates a response to parse from UpdateConfig response
func CreateUpdateConfigResponse() (response *UpdateConfigResponse) {
	response = &UpdateConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
