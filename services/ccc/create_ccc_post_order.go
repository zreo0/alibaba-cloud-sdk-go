package ccc

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

// CreateCCCPostOrder invokes the ccc.CreateCCCPostOrder API synchronously
// api document: https://help.aliyun.com/api/ccc/createcccpostorder.html
func (client *Client) CreateCCCPostOrder(request *CreateCCCPostOrderRequest) (response *CreateCCCPostOrderResponse, err error) {
	response = CreateCreateCCCPostOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCCCPostOrderWithChan invokes the ccc.CreateCCCPostOrder API asynchronously
// api document: https://help.aliyun.com/api/ccc/createcccpostorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCCCPostOrderWithChan(request *CreateCCCPostOrderRequest) (<-chan *CreateCCCPostOrderResponse, <-chan error) {
	responseChan := make(chan *CreateCCCPostOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCCCPostOrder(request)
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

// CreateCCCPostOrderWithCallback invokes the ccc.CreateCCCPostOrder API asynchronously
// api document: https://help.aliyun.com/api/ccc/createcccpostorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCCCPostOrderWithCallback(request *CreateCCCPostOrderRequest, callback func(response *CreateCCCPostOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCCCPostOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateCCCPostOrder(request)
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

// CreateCCCPostOrderRequest is the request struct for api CreateCCCPostOrder
type CreateCCCPostOrderRequest struct {
	*requests.RpcRequest
	OwnerId string `position:"Query" name:"OwnerId"`
}

// CreateCCCPostOrderResponse is the response struct for api CreateCCCPostOrder
type CreateCCCPostOrderResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	OrderId        string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateCCCPostOrderRequest creates a request to invoke CreateCCCPostOrder API
func CreateCreateCCCPostOrderRequest() (request *CreateCCCPostOrderRequest) {
	request = &CreateCCCPostOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateCCCPostOrder", "", "")
	return
}

// CreateCreateCCCPostOrderResponse creates a response to parse from CreateCCCPostOrder response
func CreateCreateCCCPostOrderResponse() (response *CreateCCCPostOrderResponse) {
	response = &CreateCCCPostOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}