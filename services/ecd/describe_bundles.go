package ecd

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

// DescribeBundles invokes the ecd.DescribeBundles API synchronously
func (client *Client) DescribeBundles(request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
	response = CreateDescribeBundlesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBundlesWithChan invokes the ecd.DescribeBundles API asynchronously
func (client *Client) DescribeBundlesWithChan(request *DescribeBundlesRequest) (<-chan *DescribeBundlesResponse, <-chan error) {
	responseChan := make(chan *DescribeBundlesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBundles(request)
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

// DescribeBundlesWithCallback invokes the ecd.DescribeBundles API asynchronously
func (client *Client) DescribeBundlesWithCallback(request *DescribeBundlesRequest, callback func(response *DescribeBundlesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBundlesResponse
		var err error
		defer close(result)
		response, err = client.DescribeBundles(request)
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

// DescribeBundlesRequest is the request struct for api DescribeBundles
type DescribeBundlesRequest struct {
	*requests.RpcRequest
	GpuCount                requests.Float   `position:"Query" name:"GpuCount"`
	ImageId                 *[]string        `position:"Query" name:"ImageId"  type:"Repeated"`
	BundleId                *[]string        `position:"Query" name:"BundleId"  type:"Repeated"`
	DesktopTypeFamily       string           `position:"Query" name:"DesktopTypeFamily"`
	SelectedBundle          requests.Boolean `position:"Query" name:"SelectedBundle"`
	NextToken               string           `position:"Query" name:"NextToken"`
	FromDesktopGroup        requests.Boolean `position:"Query" name:"FromDesktopGroup"`
	Scope                   string           `position:"Query" name:"Scope"`
	BundleType              string           `position:"Query" name:"BundleType"`
	FotaChannel             string           `position:"Query" name:"FotaChannel"`
	VolumeEncryptionEnabled requests.Boolean `position:"Query" name:"VolumeEncryptionEnabled"`
	MemorySize              requests.Integer `position:"Query" name:"MemorySize"`
	SessionType             string           `position:"Query" name:"SessionType"`
	OsType                  string           `position:"Query" name:"OsType"`
	MaxResults              requests.Integer `position:"Query" name:"MaxResults"`
	CheckStock              requests.Boolean `position:"Query" name:"CheckStock"`
	ProtocolType            string           `position:"Query" name:"ProtocolType"`
	CpuCount                requests.Integer `position:"Query" name:"CpuCount"`
	SupportMultiSession     requests.Boolean `position:"Query" name:"SupportMultiSession"`
	GpuDriverType           string           `position:"Query" name:"GpuDriverType"`
}

// DescribeBundlesResponse is the response struct for api DescribeBundles
type DescribeBundlesResponse struct {
	*responses.BaseResponse
	NextToken string   `json:"NextToken" xml:"NextToken"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Bundles   []Bundle `json:"Bundles" xml:"Bundles"`
}

// CreateDescribeBundlesRequest creates a request to invoke DescribeBundles API
func CreateDescribeBundlesRequest() (request *DescribeBundlesRequest) {
	request = &DescribeBundlesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeBundles", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBundlesResponse creates a response to parse from DescribeBundles response
func CreateDescribeBundlesResponse() (response *DescribeBundlesResponse) {
	response = &DescribeBundlesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
