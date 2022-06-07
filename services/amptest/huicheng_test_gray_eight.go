package amptest

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

// HuichengTestGrayEight invokes the amptest.HuichengTestGrayEight API synchronously
func (client *Client) HuichengTestGrayEight(request *HuichengTestGrayEightRequest) (response *HuichengTestGrayEightResponse, err error) {
	response = CreateHuichengTestGrayEightResponse()
	err = client.DoAction(request, response)
	return
}

// HuichengTestGrayEightWithChan invokes the amptest.HuichengTestGrayEight API asynchronously
func (client *Client) HuichengTestGrayEightWithChan(request *HuichengTestGrayEightRequest) (<-chan *HuichengTestGrayEightResponse, <-chan error) {
	responseChan := make(chan *HuichengTestGrayEightResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HuichengTestGrayEight(request)
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

// HuichengTestGrayEightWithCallback invokes the amptest.HuichengTestGrayEight API asynchronously
func (client *Client) HuichengTestGrayEightWithCallback(request *HuichengTestGrayEightRequest, callback func(response *HuichengTestGrayEightResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HuichengTestGrayEightResponse
		var err error
		defer close(result)
		response, err = client.HuichengTestGrayEight(request)
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

// HuichengTestGrayEightRequest is the request struct for api HuichengTestGrayEight
type HuichengTestGrayEightRequest struct {
	*requests.RpcRequest
	Home HuichengTestGrayEightHome `position:"Query" name:"Home"  type:"Struct"`
}

// HuichengTestGrayEightHome is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHome struct {
	Address      HuichengTestGrayEightHomeAddress             `name:"Address" type:"Struct"`
	T            HuichengTestGrayEightHomeT                   `name:"T" type:"Struct"`
	PhoneNumbers *[]string                                    `name:"PhoneNumbers" type:"Repeated"`
	DMap         map[string]HuichengTestGrayEightHomeDMapItem `name:"DMap" type:"Map"`
	NameToAge    map[string]string                            `name:"NameToAge" type:"Map"`
	Locations    *[]HuichengTestGrayEightHomeLocationsItem    `name:"Locations" type:"Repeated"`
}

// HuichengTestGrayEightHomeAddress is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeAddress struct {
	T        HuichengTestGrayEightHomeAddressT        `name:"T" type:"Struct"`
	Location HuichengTestGrayEightHomeAddressLocation `name:"Location" type:"Struct"`
	Detail   string                                   `name:"Detail"`
}

// HuichengTestGrayEightHomeT is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeT struct {
	Class string `name:"Class"`
}

// HuichengTestGrayEightHomeDMapItem is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeDMapItem struct {
	Location HuichengTestGrayEightHomeDMapItemLocation `name:"Location" type:"Struct"`
	Detail   string                                    `name:"Detail"`
}

// HuichengTestGrayEightHomeLocationsItem is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeLocationsItem struct {
	Late string `name:"Late"`
	Lon  string `name:"Lon"`
}

// HuichengTestGrayEightHomeAddressT is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeAddressT struct {
	Class string `name:"Class"`
}

// HuichengTestGrayEightHomeAddressLocation is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeAddressLocation struct {
	Late string `name:"Late"`
	Lon  string `name:"Lon"`
}

// HuichengTestGrayEightHomeDMapItemLocation is a repeated param struct in HuichengTestGrayEightRequest
type HuichengTestGrayEightHomeDMapItemLocation struct {
	Late string `name:"Late"`
	Lon  string `name:"Lon"`
}

// HuichengTestGrayEightResponse is the response struct for api HuichengTestGrayEight
type HuichengTestGrayEightResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Size      int    `json:"Size" xml:"Size"`
	Value     string `json:"Value" xml:"Value"`
}

// CreateHuichengTestGrayEightRequest creates a request to invoke HuichengTestGrayEight API
func CreateHuichengTestGrayEightRequest() (request *HuichengTestGrayEightRequest) {
	request = &HuichengTestGrayEightRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("AmpTest", "2020-12-30", "HuichengTestGrayEight", "AmpTest", "openAPI")
	request.Method = requests.POST
	return
}

// CreateHuichengTestGrayEightResponse creates a response to parse from HuichengTestGrayEight response
func CreateHuichengTestGrayEightResponse() (response *HuichengTestGrayEightResponse) {
	response = &HuichengTestGrayEightResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
