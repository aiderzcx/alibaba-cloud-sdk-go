package oos

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

// UpdatePatchBaseline invokes the oos.UpdatePatchBaseline API synchronously
func (client *Client) UpdatePatchBaseline(request *UpdatePatchBaselineRequest) (response *UpdatePatchBaselineResponse, err error) {
	response = CreateUpdatePatchBaselineResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePatchBaselineWithChan invokes the oos.UpdatePatchBaseline API asynchronously
func (client *Client) UpdatePatchBaselineWithChan(request *UpdatePatchBaselineRequest) (<-chan *UpdatePatchBaselineResponse, <-chan error) {
	responseChan := make(chan *UpdatePatchBaselineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePatchBaseline(request)
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

// UpdatePatchBaselineWithCallback invokes the oos.UpdatePatchBaseline API asynchronously
func (client *Client) UpdatePatchBaselineWithCallback(request *UpdatePatchBaselineRequest, callback func(response *UpdatePatchBaselineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePatchBaselineResponse
		var err error
		defer close(result)
		response, err = client.UpdatePatchBaseline(request)
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

// UpdatePatchBaselineRequest is the request struct for api UpdatePatchBaseline
type UpdatePatchBaselineRequest struct {
	*requests.RpcRequest
	Sources                          *[]string                  `position:"Query" name:"Sources"  type:"Json"`
	ClientToken                      string                     `position:"Query" name:"ClientToken"`
	ApprovalRules                    string                     `position:"Query" name:"ApprovalRules"`
	Description                      string                     `position:"Query" name:"Description"`
	ResourceGroupId                  string                     `position:"Query" name:"ResourceGroupId"`
	RejectedPatchesAction            string                     `position:"Query" name:"RejectedPatchesAction"`
	ApprovedPatchesEnableNonSecurity requests.Boolean           `position:"Query" name:"ApprovedPatchesEnableNonSecurity"`
	Tags                             *[]UpdatePatchBaselineTags `position:"Query" name:"Tags"  type:"Json"`
	RejectedPatches                  *[]string                  `position:"Query" name:"RejectedPatches"  type:"Json"`
	Name                             string                     `position:"Query" name:"Name"`
	ApprovedPatches                  *[]string                  `position:"Query" name:"ApprovedPatches"  type:"Json"`
}

// UpdatePatchBaselineTags is a repeated param struct in UpdatePatchBaselineRequest
type UpdatePatchBaselineTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// UpdatePatchBaselineResponse is the response struct for api UpdatePatchBaseline
type UpdatePatchBaselineResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	PatchBaseline PatchBaseline `json:"PatchBaseline" xml:"PatchBaseline"`
}

// CreateUpdatePatchBaselineRequest creates a request to invoke UpdatePatchBaseline API
func CreateUpdatePatchBaselineRequest() (request *UpdatePatchBaselineRequest) {
	request = &UpdatePatchBaselineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "UpdatePatchBaseline", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePatchBaselineResponse creates a response to parse from UpdatePatchBaseline response
func CreateUpdatePatchBaselineResponse() (response *UpdatePatchBaselineResponse) {
	response = &UpdatePatchBaselineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
