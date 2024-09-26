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

// UpdateRtsLiveStreamTranscode invokes the live.UpdateRtsLiveStreamTranscode API synchronously
func (client *Client) UpdateRtsLiveStreamTranscode(request *UpdateRtsLiveStreamTranscodeRequest) (response *UpdateRtsLiveStreamTranscodeResponse, err error) {
	response = CreateUpdateRtsLiveStreamTranscodeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRtsLiveStreamTranscodeWithChan invokes the live.UpdateRtsLiveStreamTranscode API asynchronously
func (client *Client) UpdateRtsLiveStreamTranscodeWithChan(request *UpdateRtsLiveStreamTranscodeRequest) (<-chan *UpdateRtsLiveStreamTranscodeResponse, <-chan error) {
	responseChan := make(chan *UpdateRtsLiveStreamTranscodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRtsLiveStreamTranscode(request)
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

// UpdateRtsLiveStreamTranscodeWithCallback invokes the live.UpdateRtsLiveStreamTranscode API asynchronously
func (client *Client) UpdateRtsLiveStreamTranscodeWithCallback(request *UpdateRtsLiveStreamTranscodeRequest, callback func(response *UpdateRtsLiveStreamTranscodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRtsLiveStreamTranscodeResponse
		var err error
		defer close(result)
		response, err = client.UpdateRtsLiveStreamTranscode(request)
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

// UpdateRtsLiveStreamTranscodeRequest is the request struct for api UpdateRtsLiveStreamTranscode
type UpdateRtsLiveStreamTranscodeRequest struct {
	*requests.RpcRequest
	Template        string           `position:"Query" name:"Template"`
	DeleteBframes   requests.Boolean `position:"Query" name:"DeleteBframes"`
	Lazy            string           `position:"Query" name:"Lazy"`
	Gop             string           `position:"Query" name:"Gop"`
	Opus            requests.Boolean `position:"Query" name:"Opus"`
	AudioCodec      string           `position:"Query" name:"AudioCodec"`
	TemplateType    string           `position:"Query" name:"TemplateType"`
	AudioProfile    string           `position:"Query" name:"AudioProfile"`
	Height          requests.Integer `position:"Query" name:"Height"`
	App             string           `position:"Query" name:"App"`
	AudioChannelNum requests.Integer `position:"Query" name:"AudioChannelNum"`
	Profile         requests.Integer `position:"Query" name:"Profile"`
	FPS             requests.Integer `position:"Query" name:"FPS"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	AudioRate       requests.Integer `position:"Query" name:"AudioRate"`
	AudioBitrate    requests.Integer `position:"Query" name:"AudioBitrate"`
	Domain          string           `position:"Query" name:"Domain"`
	Width           requests.Integer `position:"Query" name:"Width"`
	VideoBitrate    requests.Integer `position:"Query" name:"VideoBitrate"`
}

// UpdateRtsLiveStreamTranscodeResponse is the response struct for api UpdateRtsLiveStreamTranscode
type UpdateRtsLiveStreamTranscodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateRtsLiveStreamTranscodeRequest creates a request to invoke UpdateRtsLiveStreamTranscode API
func CreateUpdateRtsLiveStreamTranscodeRequest() (request *UpdateRtsLiveStreamTranscodeRequest) {
	request = &UpdateRtsLiveStreamTranscodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateRtsLiveStreamTranscode", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRtsLiveStreamTranscodeResponse creates a response to parse from UpdateRtsLiveStreamTranscode response
func CreateUpdateRtsLiveStreamTranscodeResponse() (response *UpdateRtsLiveStreamTranscodeResponse) {
	response = &UpdateRtsLiveStreamTranscodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
