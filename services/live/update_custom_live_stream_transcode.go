package live

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

// UpdateCustomLiveStreamTranscode invokes the live.UpdateCustomLiveStreamTranscode API synchronously
func (client *Client) UpdateCustomLiveStreamTranscode(request *UpdateCustomLiveStreamTranscodeRequest) (response *UpdateCustomLiveStreamTranscodeResponse, err error) {
	response = CreateUpdateCustomLiveStreamTranscodeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCustomLiveStreamTranscodeWithChan invokes the live.UpdateCustomLiveStreamTranscode API asynchronously
func (client *Client) UpdateCustomLiveStreamTranscodeWithChan(request *UpdateCustomLiveStreamTranscodeRequest) (<-chan *UpdateCustomLiveStreamTranscodeResponse, <-chan error) {
	responseChan := make(chan *UpdateCustomLiveStreamTranscodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCustomLiveStreamTranscode(request)
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

// UpdateCustomLiveStreamTranscodeWithCallback invokes the live.UpdateCustomLiveStreamTranscode API asynchronously
func (client *Client) UpdateCustomLiveStreamTranscodeWithCallback(request *UpdateCustomLiveStreamTranscodeRequest, callback func(response *UpdateCustomLiveStreamTranscodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCustomLiveStreamTranscodeResponse
		var err error
		defer close(result)
		response, err = client.UpdateCustomLiveStreamTranscode(request)
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

// UpdateCustomLiveStreamTranscodeRequest is the request struct for api UpdateCustomLiveStreamTranscode
type UpdateCustomLiveStreamTranscodeRequest struct {
	*requests.RpcRequest
	Template          string           `position:"Query" name:"Template"`
	ResWithSource     string           `position:"Query" name:"ResWithSource"`
	Lazy              string           `position:"Query" name:"Lazy"`
	Gop               string           `position:"Query" name:"Gop"`
	AudioCodec        string           `position:"Query" name:"AudioCodec"`
	TemplateType      string           `position:"Query" name:"TemplateType"`
	AudioProfile      string           `position:"Query" name:"AudioProfile"`
	Height            requests.Integer `position:"Query" name:"Height"`
	App               string           `position:"Query" name:"App"`
	EncryptParameters string           `position:"Query" name:"EncryptParameters"`
	AudioChannelNum   requests.Integer `position:"Query" name:"AudioChannelNum"`
	Profile           requests.Integer `position:"Query" name:"Profile"`
	FPS               requests.Integer `position:"Query" name:"FPS"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
	ExtWithSource     string           `position:"Query" name:"ExtWithSource"`
	BitrateWithSource string           `position:"Query" name:"BitrateWithSource"`
	AudioRate         requests.Integer `position:"Query" name:"AudioRate"`
	FpsWithSource     string           `position:"Query" name:"FpsWithSource"`
	AudioBitrate      requests.Integer `position:"Query" name:"AudioBitrate"`
	Domain            string           `position:"Query" name:"Domain"`
	Width             requests.Integer `position:"Query" name:"Width"`
	VideoBitrate      requests.Integer `position:"Query" name:"VideoBitrate"`
}

// UpdateCustomLiveStreamTranscodeResponse is the response struct for api UpdateCustomLiveStreamTranscode
type UpdateCustomLiveStreamTranscodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCustomLiveStreamTranscodeRequest creates a request to invoke UpdateCustomLiveStreamTranscode API
func CreateUpdateCustomLiveStreamTranscodeRequest() (request *UpdateCustomLiveStreamTranscodeRequest) {
	request = &UpdateCustomLiveStreamTranscodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateCustomLiveStreamTranscode", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateCustomLiveStreamTranscodeResponse creates a response to parse from UpdateCustomLiveStreamTranscode response
func CreateUpdateCustomLiveStreamTranscodeResponse() (response *UpdateCustomLiveStreamTranscodeResponse) {
	response = &UpdateCustomLiveStreamTranscodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
