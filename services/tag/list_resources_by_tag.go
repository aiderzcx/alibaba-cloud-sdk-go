package tag

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

// ListResourcesByTag invokes the tag.ListResourcesByTag API synchronously
func (client *Client) ListResourcesByTag(request *ListResourcesByTagRequest) (response *ListResourcesByTagResponse, err error) {
	response = CreateListResourcesByTagResponse()
	err = client.DoAction(request, response)
	return
}

// ListResourcesByTagWithChan invokes the tag.ListResourcesByTag API asynchronously
func (client *Client) ListResourcesByTagWithChan(request *ListResourcesByTagRequest) (<-chan *ListResourcesByTagResponse, <-chan error) {
	responseChan := make(chan *ListResourcesByTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourcesByTag(request)
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

// ListResourcesByTagWithCallback invokes the tag.ListResourcesByTag API asynchronously
func (client *Client) ListResourcesByTagWithCallback(request *ListResourcesByTagRequest, callback func(response *ListResourcesByTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourcesByTagResponse
		var err error
		defer close(result)
		response, err = client.ListResourcesByTag(request)
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

// ListResourcesByTagRequest is the request struct for api ListResourcesByTag
type ListResourcesByTagRequest struct {
	*requests.RpcRequest
	TagFilterKey         string           `position:"Query" name:"TagFilter.Key"`
	NextToken            string           `position:"Query" name:"NextToken"`
	IncludeAllTags       requests.Boolean `position:"Query" name:"IncludeAllTags"`
	TagFilterValue       string           `position:"Query" name:"TagFilter.Value"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	MaxResult            requests.Integer `position:"Query" name:"MaxResult"`
	FuzzyType            string           `position:"Query" name:"FuzzyType"`
}

// ListResourcesByTagResponse is the response struct for api ListResourcesByTag
type ListResourcesByTagResponse struct {
	*responses.BaseResponse
	NextToken string        `json:"NextToken" xml:"NextToken"`
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Resources []TagResource `json:"Resources" xml:"Resources"`
}

// CreateListResourcesByTagRequest creates a request to invoke ListResourcesByTag API
func CreateListResourcesByTagRequest() (request *ListResourcesByTagRequest) {
	request = &ListResourcesByTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "ListResourcesByTag", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListResourcesByTagResponse creates a response to parse from ListResourcesByTag response
func CreateListResourcesByTagResponse() (response *ListResourcesByTagResponse) {
	response = &ListResourcesByTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
