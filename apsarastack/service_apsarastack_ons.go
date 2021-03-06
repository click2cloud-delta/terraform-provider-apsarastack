package apsarastack

import (
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ons"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
)

type OnsService struct {
	client *connectivity.ApsaraStackClient
}

func (s *OnsService) DescribeOnsInstance(id string) (*ons.OnsInstanceBaseInfoResponse, error) {
	response := &ons.OnsInstanceBaseInfoResponse{}
	request := ons.CreateOnsInstanceBaseInfoRequest()
	request.RegionId = s.client.RegionId
	request.Headers = map[string]string{"RegionId": s.client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": s.client.SecretKey, "Product": "ons", "Department": s.client.Department, "ResourceGroup": s.client.ResourceGroup}
	request.InstanceId = id

	raw, err := s.client.WithOnsClient(func(onsClient *ons.Client) (interface{}, error) {
		return onsClient.OnsInstanceBaseInfo(request)
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidDomainName.NoExist"}) {
			return response, WrapErrorf(err, NotFoundMsg, ApsaraStackSdkGoERROR)
		}
		return response, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), ApsaraStackSdkGoERROR)
	}

	response, _ = raw.(*ons.OnsInstanceBaseInfoResponse)
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	return response, nil

}

func (s *OnsService) DescribeOnsTopic(id string) (*ons.PublishInfoDo, error) {
	onsTopic := &ons.PublishInfoDo{}
	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return onsTopic, WrapError(err)
	}
	instanceId := parts[0]
	topic := parts[1]

	request := ons.CreateOnsTopicListRequest()
	request.RegionId = s.client.RegionId
	request.Headers = map[string]string{"RegionId": s.client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": s.client.SecretKey, "Product": "ons", "Department": s.client.Department, "ResourceGroup": s.client.ResourceGroup}
	request.InstanceId = instanceId

	raw, err := s.client.WithOnsClient(func(onsClient *ons.Client) (interface{}, error) {
		return onsClient.OnsTopicList(request)
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"AUTH_RESOURCE_OWNER_ERROR", "INSTANCE_NOT_FOUND"}) {
			return onsTopic, WrapErrorf(err, NotFoundMsg, ApsaraStackSdkGoERROR)
		}
		return onsTopic, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), ApsaraStackSdkGoERROR)
	}

	topicListResp, _ := raw.(*ons.OnsTopicListResponse)
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)

	for _, v := range topicListResp.Data.PublishInfoDo {
		if v.Topic == topic {
			return &v, nil
		}
	}
	return onsTopic, WrapErrorf(err, NotFoundMsg, ApsaraStackSdkGoERROR)
}

func (s *OnsService) DescribeOnsGroup(id string) (*ons.SubscribeInfoDo, error) {
	onsGroup := &ons.SubscribeInfoDo{}
	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return onsGroup, WrapError(err)
	}
	instanceId := parts[0]
	groupId := parts[1]

	request := ons.CreateOnsGroupListRequest()
	request.RegionId = s.client.RegionId
	request.Headers = map[string]string{"RegionId": s.client.RegionId}
	request.QueryParams = map[string]string{"AccessKeySecret": s.client.SecretKey, "Product": "ons", "Department": s.client.Department, "ResourceGroup": s.client.ResourceGroup}
	request.InstanceId = instanceId

	raw, err := s.client.WithOnsClient(func(onsClient *ons.Client) (interface{}, error) {
		return onsClient.OnsGroupList(request)
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"AUTH_RESOURCE_OWNER_ERROR", "INSTANCE_NOT_FOUND"}) {
			return onsGroup, WrapErrorf(err, NotFoundMsg, ApsaraStackSdkGoERROR)
		}
		return onsGroup, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), ApsaraStackSdkGoERROR)
	}

	groupListResp, _ := raw.(*ons.OnsGroupListResponse)
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)

	for _, v := range groupListResp.Data.SubscribeInfoDo {
		if v.GroupId == groupId {
			return &v, nil
		}
	}

	return onsGroup, WrapErrorf(err, NotFoundMsg, ApsaraStackSdkGoERROR)
}

func (s *OnsService) WaitForOnsInstance(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		response, err := s.DescribeOnsInstance(id)

		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}

		if response.InstanceBaseInfo.InstanceId == id && status != Deleted {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, response.InstanceBaseInfo.InstanceId, id, ProviderERROR)
		}
		time.Sleep(DefaultIntervalShort * time.Second)
	}

}

func (s *OnsService) WaitForOnsTopic(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return WrapError(err)
	}
	instanceId := parts[0]
	topic := parts[1]
	for {
		response, err := s.DescribeOnsTopic(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}

		if response.InstanceId+":"+response.Topic == id && status != Deleted {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, instanceId+":"+topic, id, ProviderERROR)
		}
		time.Sleep(DefaultIntervalShort * time.Second)
	}
}

func (s *OnsService) WaitForOnsGroup(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	for {
		response, err := s.DescribeOnsGroup(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}

		if response.InstanceId+":"+response.GroupId == id && status != Deleted {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, response.InstanceId+":"+response.GroupId, id, ProviderERROR)
		}
		time.Sleep(DefaultIntervalShort * time.Second)
	}
}
