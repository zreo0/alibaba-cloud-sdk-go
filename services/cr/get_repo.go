package cr

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

// GetRepo invokes the cr.GetRepo API synchronously
// api document: https://help.aliyun.com/api/cr/getrepo.html
func (client *Client) GetRepo(request *GetRepoRequest) (response *GetRepoResponse, err error) {
	response = CreateGetRepoResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepoWithChan invokes the cr.GetRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoWithChan(request *GetRepoRequest) (<-chan *GetRepoResponse, <-chan error) {
	responseChan := make(chan *GetRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepo(request)
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

// GetRepoWithCallback invokes the cr.GetRepo API asynchronously
// api document: https://help.aliyun.com/api/cr/getrepo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRepoWithCallback(request *GetRepoRequest, callback func(response *GetRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepoResponse
		var err error
		defer close(result)
		response, err = client.GetRepo(request)
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

// GetRepoRequest is the request struct for api GetRepo
type GetRepoRequest struct {
	*requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

// GetRepoResponse is the response struct for api GetRepo
type GetRepoResponse struct {
	*responses.BaseResponse
}

// CreateGetRepoRequest creates a request to invoke GetRepo API
func CreateGetRepoRequest() (request *GetRepoRequest) {
	request = &GetRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "GetRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetRepoResponse creates a response to parse from GetRepo response
func CreateGetRepoResponse() (response *GetRepoResponse) {
	response = &GetRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}