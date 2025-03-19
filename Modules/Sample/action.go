package Sample

func (c *StructSamples) SetValue(value string) {

	c.Value = append(c.Value, value)
}

func (c *StructSamples) GetValue() []string {
	return c.Value
}
