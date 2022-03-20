package router

import (
	"github.com/plaid/plaid-go/plaid"
	"os"
)

var config *plaid.Configuration
var client *plaid.APIClient

func setup() {
	config = plaid.NewConfiguration()
	config.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("PLAID_CLIENT_ID"))
	config.AddDefaultHeader("PLAID-SECRET", os.Getenv("PLAID_SECRET")
	client = plaid.NewAPIClient(config)
}
