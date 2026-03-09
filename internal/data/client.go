package data

type Client struct{}

func NewClient(cfg Config) (c *Client, cleanup func(), err error) {
	c = &Client{}

	return
}
