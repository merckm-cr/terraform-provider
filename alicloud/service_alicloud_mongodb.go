package alicloud

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/denverdino/aliyungo/common"
)

type MongoDBInstance struct {
	ChargeType            string `json:"ChargeType"`
	CreationTime          string `json:"CreationTime"`
	DBInstanceClass       string `json:"DBInstanceClass"`
	DBInstanceDescription string `json:"DBInstanceDescription"`
	DBInstanceID          string `json:"DBInstanceId"`
	DBInstanceStatus      string `json:"DBInstanceStatus"`
	DBInstanceStorage     int    `json:"DBInstanceStorage"`
	DBInstanceType        string `json:"DBInstanceType"`
	Engine                string `json:"Engine"`
	EngineVersion         string `json:"EngineVersion"`
	ExpireTime            string `json:"ExpireTime"`
	LockMode              string `json:"LockMode"`
	NetworkType           string `json:"NetworkType"`
	RegionID              string `json:"RegionId"`
	ReplicationFactor     string `json:"ReplicationFactor"`
	ZoneID                string `json:"ZoneId"`
}

type ItemsInDescribeMongoDBInstances struct {
	DBInstances []MongoDBInstance `json:"DBInstance"`
}

type DescribeMongoDBInstancesResponse struct {
	PageNumber int                             `json:"PageNumber"`
	PageSize   int                             `json:"PageSize"`
	RequestID  string                          `json:"RequestId"`
	TotalCount int                             `json:"TotalCount"`
	Items      ItemsInDescribeMongoDBInstances `json:"DBInstances"`
}

type DescribeDBInstanceAttributeResponse struct {
	Items ItemsInDescribeMongoDBInstances `json:"DBInstances"`
}

type CreateMongoDBInstanceResponse struct {
	DBInstanceId string `json:"DBInstanceId"`
	OrderId      string `json:"OrderId"`
}

type DescribeMongoDBSecurityIpsResponse struct {
	SecurityIps string                                    `json:"SecurityIps"`
	Items       ItemsInDescribeMongoDBSecurityIpsResponse `json"SecurityIpGroups"`
}

type ItemsInDescribeMongoDBSecurityIpsResponse struct {
	SecurityIpGroups []SecurityMongoDBIpGroup `json"SecurityIpGroup"`
}

type SecurityMongoDBIpGroup struct {
	SecurityIpGroupName string `json:"SecurityIpGroupName"`
	SecurityIpList      string `json:"SecurityIpList"`
	SecurityIpAttribute string `json:"SecurityIpAttribute"`
}

type DescribeMongoDBBackupPolicyResponse struct {
	BackupRetentionPeriod string `json:"BackupRetentionPeriod"`
	PreferredBackupTime   string `json:"PreferredBackupTime"`
	PreferredBackupPeriod string `json:"PreferredBackupPeriod"`
}

func (client *AliyunClient) DescribeMongoDBInstances(request *requests.CommonRequest) (response *DescribeMongoDBInstancesResponse, err error) {
	request.Version = ApiVersion20151201
	request.ApiName = "DescribeDBInstances"
	resp, err := client.ecsconn.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}
	response = new(DescribeMongoDBInstancesResponse)
	err = json.Unmarshal(resp.BaseResponse.GetHttpContentBytes(), &response)

	return response, err
}

func (client *AliyunClient) CreateMongoDBInstance(request *requests.CommonRequest) (response *CreateMongoDBInstanceResponse, err error) {
	request.Version = ApiVersion20151201
	request.ApiName = "CreateDBInstance"
	resp, err := client.ecsconn.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}
	response = new(CreateMongoDBInstanceResponse)
	err = json.Unmarshal(resp.BaseResponse.GetHttpContentBytes(), &response)

	return response, err
}

// WaitForInstance waits for instance to given status
func (client *AliyunClient) WaitForMongoDBInstance(instanceId string, regionId string, status Status, timeout int) error {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	for {
		instance, err := client.DescribeMongoDBInstanceById(instanceId, regionId)
		if err != nil && !NotFoundError(err) && !IsExceptedError(err, InvalidDBInstanceIdNotFound) {
			return err
		}
		if instance != nil && instance.DBInstanceStatus == string(status) {
			break
		}

		if timeout <= 0 {
			return common.GetClientErrorFromString("Timeout")
		}

		timeout = timeout - DefaultIntervalMedium
		time.Sleep(DefaultIntervalMedium * time.Second)
	}
	return nil
}

func (client *AliyunClient) DescribeMongoDBInstanceById(id string, regionId string) (*MongoDBInstance, error) {
	request := CommonRequestInit(regionId, MONGODBCode, MongoDBDomain)
	request.RegionId = regionId
	request.Version = ApiVersion20151201
	request.ApiName = "DescribeDBInstanceAttribute"
	request.QueryParams["DBInstanceId"] = id

	resp, err := client.ecsconn.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}

	response := new(DescribeDBInstanceAttributeResponse)
	err = json.Unmarshal(resp.BaseResponse.GetHttpContentBytes(), &response)

	attr := response.Items.DBInstances

	if len(attr) <= 0 {
		return nil, GetNotFoundErrorFromString(fmt.Sprintf("DB instance %s is not found.", id))
	}

	return &attr[0], nil
}

func (client *AliyunClient) DescribeMongoDBSecurityIps(request *requests.CommonRequest) (*DescribeMongoDBSecurityIpsResponse, error) {
	request.Version = ApiVersion20151201
	request.ApiName = "DescribeSecurityIps"
	resp, err := client.ecsconn.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}
	response := new(DescribeMongoDBSecurityIpsResponse)
	err = json.Unmarshal(resp.BaseResponse.GetHttpContentBytes(), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *AliyunClient) DeleteMongoDBInstance(request *requests.CommonRequest) error {
	request.Version = ApiVersion20151201
	request.ApiName = "DeleteDBInstance"
	_, err := client.ecsconn.ProcessCommonRequest(request)
	return err
}

func (client *AliyunClient) ModifyMongoDBSecurityIps(request *requests.CommonRequest) error {
	request.Version = ApiVersion20151201
	request.ApiName = "ModifySecurityIps"
	if _, ok := request.QueryParams["foo"]; ok {
		request.QueryParams["ModifyMode"] = "Cover"
	}
	_, err := client.ecsconn.ProcessCommonRequest(request)
	return err
}

func (client *AliyunClient) ModifyMongoDBInstanceSpec(request *requests.CommonRequest) error {
	request.Version = ApiVersion20151201
	request.ApiName = "ModifyDBInstanceSpec"
	_, err := client.ecsconn.ProcessCommonRequest(request)
	return err
}

func (client *AliyunClient) ModifyMongoDBInstanceDescription(request *requests.CommonRequest) error {
	request.Version = ApiVersion20151201
	request.ApiName = "ModifyDBInstanceDescription"
	_, err := client.ecsconn.ProcessCommonRequest(request)
	return err
}

func (client *AliyunClient) ModifyMongoDBBackupPolicy(request *requests.CommonRequest) error {
	request.Version = ApiVersion20151201
	request.ApiName = "ModifyBackupPolicy"
	_, err := client.ecsconn.ProcessCommonRequest(request)
	return err
}

func (client *AliyunClient) DescribeMongoDBBackupPolicy(request *requests.CommonRequest) (*DescribeMongoDBBackupPolicyResponse, error) {
	request.Version = ApiVersion20151201
	request.ApiName = "DescribeBackupPolicy"
	resp, err := client.ecsconn.ProcessCommonRequest(request)
	response := new(DescribeMongoDBBackupPolicyResponse)
	err = json.Unmarshal(resp.BaseResponse.GetHttpContentBytes(), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}