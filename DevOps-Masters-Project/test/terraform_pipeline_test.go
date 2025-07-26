package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformS3Bucket(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../terraform", // path to your Terraform code
    }

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    bucketName := terraform.Output(t, terraformOptions, "bucket_name")
    assert.Contains(t, bucketName, "abismruta-test")
}
