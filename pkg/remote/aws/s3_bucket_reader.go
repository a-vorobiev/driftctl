package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/aws"
	"github.com/cloudskiff/driftctl/pkg/terraform"
)

type S3BucketDetailFetcher struct {
	reader       terraform.ResourceReader
	deserializer *resource.Deserializer
}

func NewS3BucketReader(provider *AWSTerraformProvider, deserializer *resource.Deserializer) *S3BucketDetailFetcher {
	return &S3BucketDetailFetcher{
		reader:       provider,
		deserializer: deserializer,
	}
}

func (r *S3BucketDetailFetcher) ReadDetails(res resource.Resource) (resource.Resource, error) {
	ctyVal, err := r.reader.ReadResource(terraform.ReadResourceArgs{
		Ty: aws.AwsS3BucketResourceType,
		ID: res.TerraformId(),
		Attributes: map[string]string{
			"alias": *res.Attributes().GetString("region"),
		},
	})
	if err != nil {
		return nil, err
	}
	deserializedRes, err := r.deserializer.DeserializeOne(aws.AwsS3BucketResourceType, *ctyVal)
	if err != nil {
		return nil, err
	}

	return deserializedRes, nil
}
