package common

import (
	"fmt"

	"context"

	operatorv1 "github.com/openshift/api/operator/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/client-go/listers/core/v1"
)

func GetOAuthServerService(ctx context.Context, serviceLister v1.ServiceLister, conditionPrefix string) (*corev1.Service, []operatorv1.OperatorCondition) {
	service, err := serviceLister.Services("openshift-authentication").Get(ctx, "oauth-openshift")
	if err != nil {
		return nil, []operatorv1.OperatorCondition{
			{
				Type:    conditionPrefix + "Degraded",
				Status:  operatorv1.ConditionTrue,
				Reason:  "GetFailed",
				Message: fmt.Sprintf("Unable to get oauth server service: %v", err),
			},
		}
	}
	return service, nil
}
