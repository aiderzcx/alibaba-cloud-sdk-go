package imageprocess

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

// DetectRibFracture invokes the imageprocess.DetectRibFracture API synchronously
func (client *Client) DetectRibFracture(request *DetectRibFractureRequest) (response *DetectRibFractureResponse, err error) {
	response = CreateDetectRibFractureResponse()
	err = client.DoAction(request, response)
	return
}

// DetectRibFractureWithChan invokes the imageprocess.DetectRibFracture API asynchronously
func (client *Client) DetectRibFractureWithChan(request *DetectRibFractureRequest) (<-chan *DetectRibFractureResponse, <-chan error) {
	responseChan := make(chan *DetectRibFractureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectRibFracture(request)
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

// DetectRibFractureWithCallback invokes the imageprocess.DetectRibFracture API asynchronously
func (client *Client) DetectRibFractureWithCallback(request *DetectRibFractureRequest, callback func(response *DetectRibFractureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectRibFractureResponse
		var err error
		defer close(result)
		response, err = client.DetectRibFracture(request)
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

// DetectRibFractureRequest is the request struct for api DetectRibFracture
type DetectRibFractureRequest struct {
	*requests.RpcRequest
	OrgName    string                      `position:"Body" name:"OrgName"`
	SourceType string                      `position:"Body" name:"SourceType"`
	DataFormat string                      `position:"Body" name:"DataFormat"`
	URLList    *[]DetectRibFractureURLList `position:"Body" name:"URLList"  type:"Repeated"`
	OrgId      string                      `position:"Body" name:"OrgId"`
	Async      requests.Boolean            `position:"Body" name:"Async"`
}

// DetectRibFractureURLList is a repeated param struct in DetectRibFractureRequest
type DetectRibFractureURLList struct {
	URL string `name:"URL"`
}

// DetectRibFractureResponse is the response struct for api DetectRibFracture
type DetectRibFractureResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectRibFractureRequest creates a request to invoke DetectRibFracture API
func CreateDetectRibFractureRequest() (request *DetectRibFractureRequest) {
	request = &DetectRibFractureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "DetectRibFracture", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectRibFractureResponse creates a response to parse from DetectRibFracture response
func CreateDetectRibFractureResponse() (response *DetectRibFractureResponse) {
	response = &DetectRibFractureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
