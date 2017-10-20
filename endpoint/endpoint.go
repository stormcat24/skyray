package endpoint

type SkyrayEndpoint struct {
}

func NewSkyrayEndpoint() SkyrayServiceServer {
	return &SkyrayEndpoint{}
}

func (e SkyrayEndpoint) Connect(SkyrayService_ConnectServer) error {
	return nil
}
