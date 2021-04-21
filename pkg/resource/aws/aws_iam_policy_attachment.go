// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	rescty "github.com/cloudskiff/driftctl/pkg/resource/cty"
	"github.com/zclconf/go-cty/cty"
)

const AwsIamPolicyAttachmentResourceType = "aws_iam_policy_attachment"

type AwsIamPolicyAttachment struct {
	Groups    *[]string  `cty:"groups"`
	Id        string     `cty:"id" diff:"Id, identifier" computed:"true"`
	Name      *string    `cty:"name" diff:"-"`
	PolicyArn *string    `cty:"policy_arn"`
	Roles     *[]string  `cty:"roles"`
	Users     *[]string  `cty:"users"`
	CtyVal    *cty.Value `diff:"-"`
}

func (r *AwsIamPolicyAttachment) TerraformId() string {
	return r.Id
}

func (r *AwsIamPolicyAttachment) TerraformType() string {
	return AwsIamPolicyAttachmentResourceType
}

func (r *AwsIamPolicyAttachment) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsIamPolicyAttachmentMetaData(resourceSchemaRepository *resource.SchemaRepository) {
	resourceSchemaRepository.SetNormalizeFunc(AwsIamPolicyAttachmentResourceType, func(val *rescty.CtyAttributes) {
		val.SafeDelete([]string{"name"})
	})
}
