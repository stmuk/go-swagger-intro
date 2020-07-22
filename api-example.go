package api

import (
	"context"
	"net/http"
	"time"

	// START OMIT
	"wibble.co.uk/foobar/newexttests/shim/client"
	apiclient "wibble.co.uk/foobar/newexttests/shim/client"
	"wibble.co.uk/foobar/newexttests/shim/client/clients"
	"wibble.co.uk/foobar/newexttests/shim/models"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func (c Config) AddClient(body *models.NewClientSettings) (resp *clients.CreateClientCreated, err error) {
	hclient := &http.Client{Transport: &http.Transport{Proxy: nil}}
	transport := httptransport.NewWithClient(c.Endpoint, client.DefaultBasePath, []string{"http"}, hclient)
	transport.SetDebug(true)
	client := apiclient.New(transport, strfmt.Default)

	ctx, cancel := context.WithTimeout(context.Background(), ctxto*time.Second)
	defer cancel()

	param := &clients.CreateClientParams{Body: body, Context: ctx}
	resp, err = client.Clients.CreateClient(param, httptransport.BasicAuth(c.User, c.Pass))
	return resp, err
}

// END OMIT
