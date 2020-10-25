package r_kvstore

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

// CreateUserClusterHost invokes the r_kvstore.CreateUserClusterHost API synchronously
func (client *Client) CreateUserClusterHost(request *CreateUserClusterHostRequest) (response *CreateUserClusterHostResponse, err error) {
	response = CreateCreateUserClusterHostResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserClusterHostWithChan invokes the r_kvstore.CreateUserClusterHost API asynchronously
func (client *Client) CreateUserClusterHostWithChan(request *CreateUserClusterHostRequest) (<-chan *CreateUserClusterHostResponse, <-chan error) {
	responseChan := make(chan *CreateUserClusterHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUserClusterHost(request)
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

// CreateUserClusterHostWithCallback invokes the r_kvstore.CreateUserClusterHost API asynchronously
func (client *Client) CreateUserClusterHostWithCallback(request *CreateUserClusterHostRequest, callback func(response *CreateUserClusterHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserClusterHostResponse
		var err error
		defer close(result)
		response, err = client.CreateUserClusterHost(request)
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

// CreateUserClusterHostRequest is the request struct for api CreateUserClusterHost
type CreateUserClusterHostRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	OrderPeriod          requests.Integer `position:"Query" name:"OrderPeriod"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AgentId              string           `position:"Query" name:"AgentId"`
	HostClass            string           `position:"Query" name:"HostClass"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OrderNum             requests.Integer `position:"Query" name:"OrderNum"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AutoRenew            requests.Boolean `position:"Query" name:"AutoRenew"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	ChargeType           string           `position:"Query" name:"ChargeType"`
}

// CreateUserClusterHostResponse is the response struct for api CreateUserClusterHost
type CreateUserClusterHostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
	HostId    string `json:"HostId" xml:"HostId"`
}

// CreateCreateUserClusterHostRequest creates a request to invoke CreateUserClusterHost API
func CreateCreateUserClusterHostRequest() (request *CreateUserClusterHostRequest) {
	request = &CreateUserClusterHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateUserClusterHost", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateUserClusterHostResponse creates a response to parse from CreateUserClusterHost response
func CreateCreateUserClusterHostResponse() (response *CreateUserClusterHostResponse) {
	response = &CreateUserClusterHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
