package alidns

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

// DescribePdnsRequestStatistics invokes the alidns.DescribePdnsRequestStatistics API synchronously
func (client *Client) DescribePdnsRequestStatistics(request *DescribePdnsRequestStatisticsRequest) (response *DescribePdnsRequestStatisticsResponse, err error) {
	response = CreateDescribePdnsRequestStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePdnsRequestStatisticsWithChan invokes the alidns.DescribePdnsRequestStatistics API asynchronously
func (client *Client) DescribePdnsRequestStatisticsWithChan(request *DescribePdnsRequestStatisticsRequest) (<-chan *DescribePdnsRequestStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribePdnsRequestStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePdnsRequestStatistics(request)
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

// DescribePdnsRequestStatisticsWithCallback invokes the alidns.DescribePdnsRequestStatistics API asynchronously
func (client *Client) DescribePdnsRequestStatisticsWithCallback(request *DescribePdnsRequestStatisticsRequest, callback func(response *DescribePdnsRequestStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePdnsRequestStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribePdnsRequestStatistics(request)
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

// DescribePdnsRequestStatisticsRequest is the request struct for api DescribePdnsRequestStatistics
type DescribePdnsRequestStatisticsRequest struct {
	*requests.RpcRequest
	DomainName     string           `position:"Query" name:"DomainName"`
	NeedThreatInfo requests.Boolean `position:"Query" name:"NeedThreatInfo"`
	Type           string           `position:"Query" name:"Type"`
	StartDate      string           `position:"Query" name:"StartDate"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	EndDate        string           `position:"Query" name:"EndDate"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	SubDomain      string           `position:"Query" name:"SubDomain"`
	Lang           string           `position:"Query" name:"Lang"`
}

// DescribePdnsRequestStatisticsResponse is the response struct for api DescribePdnsRequestStatistics
type DescribePdnsRequestStatisticsResponse struct {
	*responses.BaseResponse
	TotalCount int64           `json:"TotalCount" xml:"TotalCount"`
	PageSize   int64           `json:"PageSize" xml:"PageSize"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	PageNumber int64           `json:"PageNumber" xml:"PageNumber"`
	Data       []StatisticItem `json:"Data" xml:"Data"`
}

// CreateDescribePdnsRequestStatisticsRequest creates a request to invoke DescribePdnsRequestStatistics API
func CreateDescribePdnsRequestStatisticsRequest() (request *DescribePdnsRequestStatisticsRequest) {
	request = &DescribePdnsRequestStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribePdnsRequestStatistics", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePdnsRequestStatisticsResponse creates a response to parse from DescribePdnsRequestStatistics response
func CreateDescribePdnsRequestStatisticsResponse() (response *DescribePdnsRequestStatisticsResponse) {
	response = &DescribePdnsRequestStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
