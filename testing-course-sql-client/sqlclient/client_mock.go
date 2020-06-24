package sqlclient

type clientMock struct {
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {
	return nil, nil
}
