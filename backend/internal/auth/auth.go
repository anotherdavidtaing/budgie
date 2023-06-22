package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
)

func returnActiveSession(w http.ResponseWriter, req *http.Request) {
	session := req.Context().Value(clerk.ActiveSession)
	jsonResp, _ := json.Marshal(session)

	fmt.Fprintf(w, string(jsonResp))
}

func New() (clerk.Client, *http.ServeMux, error) {
	apiKey := os.Getenv("CLERK_API_KEY")
	client, err := clerk.NewClient(apiKey)

	if err != nil {
		return nil, nil, err
	}

	mux := http.NewServeMux()
	injectActiveSession := clerk.WithSession(client)
	mux.Handle("/session", injectActiveSession(http.HandlerFunc(returnActiveSession)))

	return client, mux, nil
}
