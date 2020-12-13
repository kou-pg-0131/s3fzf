package cmd

// Copy ...
func (c *Command) Copy(bucket, output string) error {
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

	f, err := c.s3Controller.GetObject(bucket, *o.Key)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := c.fileWriter.Write(output, f); err != nil {
		return err
	}

	return nil
}
