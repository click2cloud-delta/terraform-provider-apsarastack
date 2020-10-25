package cr_ee

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

// CreateChartRepository invokes the cr.CreateChartRepository API synchronously
// api document: https://help.aliyun.com/api/cr/createchartrepository.html
func (client *Client) CreateChartRepository(request *CreateChartRepositoryRequest) (response *CreateChartRepositoryResponse, err error) {
	response = CreateCreateChartRepositoryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateChartRepositoryWithChan invokes the cr.CreateChartRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/createchartrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChartRepositoryWithChan(request *CreateChartRepositoryRequest) (<-chan *CreateChartRepositoryResponse, <-chan error) {
	responseChan := make(chan *CreateChartRepositoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateChartRepository(request)
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

// CreateChartRepositoryWithCallback invokes the cr.CreateChartRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/createchartrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChartRepositoryWithCallback(request *CreateChartRepositoryRequest, callback func(response *CreateChartRepositoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateChartRepositoryResponse
		var err error
		defer close(result)
		response, err = client.CreateChartRepository(request)
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

// CreateChartRepositoryRequest is the request struct for api CreateChartRepository
type CreateChartRepositoryRequest struct {
	*requests.RpcRequest
	RepoType          string `position:"Query" name:"RepoType"`
	Summary           string `position:"Query" name:"Summary"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	RepoName          string `position:"Query" name:"RepoName"`
	RepoNamespaceName string `position:"Query" name:"RepoNamespaceName"`
}

// CreateChartRepositoryResponse is the response struct for api CreateChartRepository
type CreateChartRepositoryResponse struct {
	*responses.BaseResponse
	CreateChartRepositoryIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string `json:"Code" xml:"Code"`
	RequestId                      string `json:"RequestId" xml:"RequestId"`
	RepoId                         string `json:"RepoId" xml:"RepoId"`
}

// CreateCreateChartRepositoryRequest creates a request to invoke CreateChartRepository API
func CreateCreateChartRepositoryRequest() (request *CreateChartRepositoryRequest) {
	request = &CreateChartRepositoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "CreateChartRepository", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateChartRepositoryResponse creates a response to parse from CreateChartRepository response
func CreateCreateChartRepositoryResponse() (response *CreateChartRepositoryResponse) {
	response = &CreateChartRepositoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
