package middlewares

import (
	"github.com/sirupsen/logrus"

	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/aws"
)

// When scanning a brand new AWS account, some users may see irrelevant results about default AWS role policies.
// We ignore these resources by default when strict mode is disabled.
type AwsIamRolePolicyDefaults struct{}

var ignoredIamRolePolicyIds = []string{
	"OrganizationAccountAccessRole:AdministratorAccess",
}

func NewAwsIamRolePolicyDefaults() AwsIamRolePolicyDefaults {
	return AwsIamRolePolicyDefaults{}
}

func (m AwsIamRolePolicyDefaults) Execute(remoteResources, resourcesFromState *[]resource.Resource) error {
	for _, remoteResource := range *remoteResources {
		// Ignore all resources other than role policy
		if remoteResource.TerraformType() != aws.AwsIamRolePolicyResourceType {
			continue
		}

		existInState := false
		for _, stateResource := range *resourcesFromState {
			if resource.IsSameResource(remoteResource, stateResource) {
				existInState = true
				break
			}
		}

		if existInState {
			continue
		}

		for _, id := range ignoredIamRolePolicyIds {
			if remoteResource.TerraformId() == id {
				*resourcesFromState = append(*resourcesFromState, remoteResource)

				logrus.WithFields(logrus.Fields{
					"id":   remoteResource.TerraformId(),
					"type": remoteResource.TerraformType(),
				}).Debug("Ignoring default iam role policy as it is not managed by IaC")
			}
		}
	}

	return nil
}