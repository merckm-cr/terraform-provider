package cloudapi

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

// SetDomain invokes the cloudapi.SetDomain API synchronously
// api document: https://help.aliyun.com/api/cloudapi/setdomain.html
func (client *Client) SetDomain(request *SetDomainRequest) (response *SetDomainResponse, err error) {
	response = CreateSetDomainResponse()
	err = client.DoAction(request, response)
	return
}

// SetDomainWithChan invokes the cloudapi.SetDomain API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/setdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainWithChan(request *SetDomainRequest) (<-chan *SetDomainResponse, <-chan error) {
	responseChan := make(chan *SetDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomain(request)
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

// SetDomainWithCallback invokes the cloudapi.SetDomain API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/setdomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainWithCallback(request *SetDomainRequest, callback func(response *SetDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainResponse
		var err error
		defer close(result)
		response, err = client.SetDomain(request)
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

// SetDomainRequest is the request struct for api SetDomain
type SetDomainRequest struct {
	*requests.RpcRequest
	GroupId               string `position:"Query" name:"GroupId"`
	DomainName            string `position:"Query" name:"DomainName"`
	CertificateName       string `position:"Query" name:"CertificateName"`
	CertificateBody       string `position:"Query" name:"CertificateBody"`
	CertificatePrivateKey string `position:"Query" name:"CertificatePrivateKey"`
}

// SetDomainResponse is the response struct for api SetDomain
type SetDomainResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	GroupId               string `json:"GroupId" xml:"GroupId"`
	DomainName            string `json:"DomainName" xml:"DomainName"`
	SubDomain             string `json:"SubDomain" xml:"SubDomain"`
	DomainBindingStatus   string `json:"DomainBindingStatus" xml:"DomainBindingStatus"`
	DomainLegalStatus     string `json:"DomainLegalStatus" xml:"DomainLegalStatus"`
	DomainWebSocketStatus string `json:"DomainWebSocketStatus" xml:"DomainWebSocketStatus"`
	DomainRemark          string `json:"DomainRemark" xml:"DomainRemark"`
}

// CreateSetDomainRequest creates a request to invoke SetDomain API
func CreateSetDomainRequest() (request *SetDomainRequest) {
	request = &SetDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SetDomain", "apigateway", "openAPI")
	return
}

// CreateSetDomainResponse creates a response to parse from SetDomain response
func CreateSetDomainResponse() (response *SetDomainResponse) {
	response = &SetDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}