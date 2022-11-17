package oceanbasepro

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

// IncrSyncStepTransferConfig is a nested struct in oceanbasepro response
type IncrSyncStepTransferConfig struct {
	StoreTransactionEnabled bool     `json:"StoreTransactionEnabled" xml:"StoreTransactionEnabled"`
	StartTimestamp          int64    `json:"StartTimestamp" xml:"StartTimestamp"`
	StoreLogKeptHour        int64    `json:"StoreLogKeptHour" xml:"StoreLogKeptHour"`
	TransferStepType        string   `json:"TransferStepType" xml:"TransferStepType"`
	RecordTypeList          []string `json:"RecordTypeList" xml:"RecordTypeList"`
}
