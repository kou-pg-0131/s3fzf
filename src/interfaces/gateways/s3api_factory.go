package gateways

import (
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// IS3APIFactory .
type IS3APIFactory interface {
	Create(region string) s3iface.S3API
}
