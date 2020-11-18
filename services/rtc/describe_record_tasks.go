package rtc

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

// DescribeRecordTasks invokes the rtc.DescribeRecordTasks API synchronously
func (client *Client) DescribeRecordTasks(request *DescribeRecordTasksRequest) (response *DescribeRecordTasksResponse, err error) {
	response = CreateDescribeRecordTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecordTasksWithChan invokes the rtc.DescribeRecordTasks API asynchronously
func (client *Client) DescribeRecordTasksWithChan(request *DescribeRecordTasksRequest) (<-chan *DescribeRecordTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeRecordTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecordTasks(request)
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

// DescribeRecordTasksWithCallback invokes the rtc.DescribeRecordTasks API asynchronously
func (client *Client) DescribeRecordTasksWithCallback(request *DescribeRecordTasksRequest, callback func(response *DescribeRecordTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecordTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecordTasks(request)
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

// DescribeRecordTasksRequest is the request struct for api DescribeRecordTasks
type DescribeRecordTasksRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	TaskIds   *[]string        `position:"Query" name:"TaskIds"  type:"Repeated"`
	PageNum   requests.Integer `position:"Query" name:"PageNum"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
	Status    string           `position:"Query" name:"Status"`
}

// DescribeRecordTasksResponse is the response struct for api DescribeRecordTasks
type DescribeRecordTasksResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	TotalNum    int64        `json:"TotalNum" xml:"TotalNum"`
	TotalPage   int64        `json:"TotalPage" xml:"TotalPage"`
	RecordTasks []RecordTask `json:"RecordTasks" xml:"RecordTasks"`
}

// CreateDescribeRecordTasksRequest creates a request to invoke DescribeRecordTasks API
func CreateDescribeRecordTasksRequest() (request *DescribeRecordTasksRequest) {
	request = &DescribeRecordTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRecordTasks", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRecordTasksResponse creates a response to parse from DescribeRecordTasks response
func CreateDescribeRecordTasksResponse() (response *DescribeRecordTasksResponse) {
	response = &DescribeRecordTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
