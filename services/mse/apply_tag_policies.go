package mse

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

// ApplyTagPolicies invokes the mse.ApplyTagPolicies API synchronously
func (client *Client) ApplyTagPolicies(request *ApplyTagPoliciesRequest) (response *ApplyTagPoliciesResponse, err error) {
	response = CreateApplyTagPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyTagPoliciesWithChan invokes the mse.ApplyTagPolicies API asynchronously
func (client *Client) ApplyTagPoliciesWithChan(request *ApplyTagPoliciesRequest) (<-chan *ApplyTagPoliciesResponse, <-chan error) {
	responseChan := make(chan *ApplyTagPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyTagPolicies(request)
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

// ApplyTagPoliciesWithCallback invokes the mse.ApplyTagPolicies API asynchronously
func (client *Client) ApplyTagPoliciesWithCallback(request *ApplyTagPoliciesRequest, callback func(response *ApplyTagPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyTagPoliciesResponse
		var err error
		defer close(result)
		response, err = client.ApplyTagPolicies(request)
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

// ApplyTagPoliciesRequest is the request struct for api ApplyTagPolicies
type ApplyTagPoliciesRequest struct {
	*requests.RpcRequest
	MseSessionId   string                           `position:"Query" name:"MseSessionId"`
	Rules          map[string]ApplyTagPoliciesRules `position:"Query" name:"Rules"  type:"Map"`
	Source         string                           `position:"Query" name:"Source"`
	AppName        string                           `position:"Query" name:"AppName"`
	NamespaceId    string                           `position:"Query" name:"NamespaceId"`
	Enable         requests.Boolean                 `position:"Query" name:"Enable"`
	AppId          string                           `position:"Query" name:"AppId"`
	Namespace      string                           `position:"Query" name:"Namespace"`
	AcceptLanguage string                           `position:"Query" name:"AcceptLanguage"`
	Region         string                           `position:"Query" name:"Region"`
}

// ApplyTagPoliciesRules is a repeated param struct in ApplyTagPoliciesRequest
type ApplyTagPoliciesRules struct {
	Rate        string                     `name:"Rate"`
	Enable      string                     `name:"Enable"`
	InstanceNum string                     `name:"InstanceNum"`
	Name        string                     `name:"Name"`
	CarryData   string                     `name:"CarryData"`
	Rules       ApplyTagPoliciesRulesRules `name:"Rules" type:"Struct"`
	Tag         string                     `name:"Tag"`
	Id          string                     `name:"Id"`
	Remove      string                     `name:"remove"`
	Status      string                     `name:"Status"`
}

// ApplyTagPoliciesRulesRules is a repeated param struct in ApplyTagPoliciesRequest
type ApplyTagPoliciesRulesRules struct {
	Springcloud *[]ApplyTagPoliciesRulesRulesSpringcloudItem `name:"springcloud" type:"Repeated"`
	Dubbo       *[]ApplyTagPoliciesRulesRulesDubboItem       `name:"dubbo" type:"Repeated"`
}

// ApplyTagPoliciesRulesRulesSpringcloudItem is a repeated param struct in ApplyTagPoliciesRequest
type ApplyTagPoliciesRulesRulesSpringcloudItem struct {
	RestItems     *[]ApplyTagPoliciesRulesRulesSpringcloudItemRestItemsItem `name:"restItems" type:"Repeated"`
	Path          string                                                    `name:"path"`
	Condition     string                                                    `name:"condition"`
	Enable        string                                                    `name:"enable"`
	Paths         *[]string                                                 `name:"paths" type:"Repeated"`
	AppId         string                                                    `name:"appId"`
	Priority      string                                                    `name:"priority"`
	TriggerPolicy string                                                    `name:"triggerPolicy"`
	Tags          *[]string                                                 `name:"tags" type:"Repeated"`
}

// ApplyTagPoliciesRulesRulesDubboItem is a repeated param struct in ApplyTagPoliciesRequest
type ApplyTagPoliciesRulesRulesDubboItem struct {
	ParamTypes    *[]string                                               `name:"paramTypes" type:"Repeated"`
	Condition     string                                                  `name:"condition"`
	AppId         string                                                  `name:"appId"`
	CarryData     string                                                  `name:"carryData"`
	ServiceName   string                                                  `name:"serviceName"`
	TriggerPolicy string                                                  `name:"triggerPolicy"`
	Version       string                                                  `name:"version"`
	ArgumentItems *[]ApplyTagPoliciesRulesRulesDubboItemArgumentItemsItem `name:"argumentItems" type:"Repeated"`
	Tags          *[]string                                               `name:"tags" type:"Repeated"`
	Group         string                                                  `name:"group"`
	MethodName    string                                                  `name:"methodName"`
}

// ApplyTagPoliciesRulesRulesSpringcloudItemRestItemsItem is a repeated param struct in ApplyTagPoliciesRequest
type ApplyTagPoliciesRulesRulesSpringcloudItemRestItemsItem struct {
	Datum     string    `name:"datum"`
	Divisor   string    `name:"divisor"`
	Rate      string    `name:"rate"`
	NameList  *[]string `name:"nameList" type:"Repeated"`
	Name      string    `name:"name"`
	Cond      string    `name:"cond"`
	Type      string    `name:"type"`
	Remainder string    `name:"remainder"`
	Value     string    `name:"value"`
	Operator  string    `name:"operator"`
}

// ApplyTagPoliciesRulesRulesDubboItemArgumentItemsItem is a repeated param struct in ApplyTagPoliciesRequest
type ApplyTagPoliciesRulesRulesDubboItemArgumentItemsItem struct {
	Datum     string    `name:"datum"`
	Divisor   string    `name:"divisor"`
	Rate      string    `name:"rate"`
	NameList  *[]string `name:"nameList" type:"Repeated"`
	Index     string    `name:"index"`
	Expr      string    `name:"expr"`
	Cond      string    `name:"cond"`
	Remainder string    `name:"remainder"`
	Value     string    `name:"value"`
	Operator  string    `name:"operator"`
}

// ApplyTagPoliciesResponse is the response struct for api ApplyTagPolicies
type ApplyTagPoliciesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string        `json:"Message" xml:"Message"`
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Code           int           `json:"Code" xml:"Code"`
	Success        bool          `json:"Success" xml:"Success"`
	Data           []RouteRuleVO `json:"Data" xml:"Data"`
}

// CreateApplyTagPoliciesRequest creates a request to invoke ApplyTagPolicies API
func CreateApplyTagPoliciesRequest() (request *ApplyTagPoliciesRequest) {
	request = &ApplyTagPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ApplyTagPolicies", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateApplyTagPoliciesResponse creates a response to parse from ApplyTagPolicies response
func CreateApplyTagPoliciesResponse() (response *ApplyTagPoliciesResponse) {
	response = &ApplyTagPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
