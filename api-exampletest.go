

func TestaddUser(t *testing.T) {
	body := &models.NewClientSettings{ClientID: &shim.ClientID}
	body.Service = "isp_consumer"
	resp, err := shim.AddClient(body)
	if err != nil {
		log.Print(err)
	}
	p := resp.Payload
	if p.Service != "isp_consumer" || !*(p.FilterEnabled) || *(p.NetworkSecurityFilterEnabled) {
		t.Fail()
	}

}
