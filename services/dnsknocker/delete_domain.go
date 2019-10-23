package dnsknocker

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

// DeleteDomain invokes the dnsknocker.DeleteDomain API synchronously
// api document: https://help.aliyun.com/api/dnsknocker/deletedomain.html
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
	response = CreateDeleteDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDomainWithChan invokes the dnsknocker.DeleteDomain API asynchronously
// api document: https://help.aliyun.com/api/dnsknocker/deletedomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDomainWithChan(request *DeleteDomainRequest) (<-chan *DeleteDomainResponse, <-chan error) {
	responseChan := make(chan *DeleteDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDomain(request)
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

// DeleteDomainWithCallback invokes the dnsknocker.DeleteDomain API asynchronously
// api document: https://help.aliyun.com/api/dnsknocker/deletedomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDomainWithCallback(request *DeleteDomainRequest, callback func(response *DeleteDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDomainResponse
		var err error
		defer close(result)
		response, err = client.DeleteDomain(request)
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

// DeleteDomainRequest is the request struct for api DeleteDomain
type DeleteDomainRequest struct {
	*requests.RpcRequest
	AccessID      string `position:"Body" name:"AccessID"`
	AccessSecret  string `position:"Body" name:"AccessSecret"`
	Line          string `position:"Body" name:"Line"`
	DomainName    string `position:"Body" name:"DomainName"`
	ZoneName      string `position:"Body" name:"ZoneName"`
	TransactionId string `position:"Body" name:"TransactionId"`
	Group         string `position:"Body" name:"Group"`
}

// DeleteDomainResponse is the response struct for api DeleteDomain
type DeleteDomainResponse struct {
	*responses.BaseResponse
	ResultCode    string `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string `json:"ResultMessage" xml:"ResultMessage"`
	Success       bool   `json:"Success" xml:"Success"`
	TransactionId string `json:"TransactionId" xml:"TransactionId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDomainRequest creates a request to invoke DeleteDomain API
func CreateDeleteDomainRequest() (request *DeleteDomainRequest) {
	request = &DeleteDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DnsKnocker", "2019-09-10", "DeleteDomain", "dns_knocker", "openAPI")
	return
}

// CreateDeleteDomainResponse creates a response to parse from DeleteDomain response
func CreateDeleteDomainResponse() (response *DeleteDomainResponse) {
	response = &DeleteDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}