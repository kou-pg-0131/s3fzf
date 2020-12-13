package cmd

import (
	"errors"
	"fmt"
)

// Remove .
func (c *Command) Remove(bucket string) error {
	if bucket == "" {
		b, err := c.s3Controller.FindBucket()
		if err != nil {
			return err
		}
		bucket = *b.Name
	}

	o, err := c.s3Controller.FindObject(bucket)
	if err != nil {
		return err
	}

	yes, err := c.confirmer.Confirm(fmt.Sprintf("remove `s3://%s/%s`?(y/N): ", bucket, *o.Key))
	if err != nil {
		return err
	}
	if !yes {
		return errors.New("cancelled")
	}

	if err := c.s3Controller.DeleteObject(bucket, *o.Key); err != nil {
		return err
	}

	return nil
}
