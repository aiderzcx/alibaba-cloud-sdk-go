package cdrs

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

// RecognizeImage invokes the cdrs.RecognizeImage API synchronously
func (client *Client) RecognizeImage(request *RecognizeImageRequest) (response *RecognizeImageResponse, err error) {
	response = CreateRecognizeImageResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeImageWithChan invokes the cdrs.RecognizeImage API asynchronously
func (client *Client) RecognizeImageWithChan(request *RecognizeImageRequest) (<-chan *RecognizeImageResponse, <-chan error) {
	responseChan := make(chan *RecognizeImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeImage(request)
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

// RecognizeImageWithCallback invokes the cdrs.RecognizeImage API asynchronously
func (client *Client) RecognizeImageWithCallback(request *RecognizeImageRequest, callback func(response *RecognizeImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeImageResponse
		var err error
		defer close(result)
		response, err = client.RecognizeImage(request)
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

// RecognizeImageRequest is the request struct for api RecognizeImage
type RecognizeImageRequest struct {
	*requests.RpcRequest
	CorpId       string `position:"Body" name:"CorpId"`
	Vendor       string `position:"Body" name:"Vendor"`
	ImageUrl     string `position:"Body" name:"ImageUrl"`
	ImageContent string `position:"Body" name:"ImageContent"`
}

// RecognizeImageResponse is the response struct for api RecognizeImage
type RecognizeImageResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeImageRequest creates a request to invoke RecognizeImage API
func CreateRecognizeImageRequest() (request *RecognizeImageRequest) {
	request = &RecognizeImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "RecognizeImage", "", "")
	request.Method = requests.POST
	return
}

// CreateRecognizeImageResponse creates a response to parse from RecognizeImage response
func CreateRecognizeImageResponse() (response *RecognizeImageResponse) {
	response = &RecognizeImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
