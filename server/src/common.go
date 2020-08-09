package funcs

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"google.golang.org/api/idtoken"
)

func setHeaderForCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.WriteHeader(http.StatusNoContent)
}

// callFunction は、別のGCF関数を呼び出す関数
// functionURL = "https://REGION-PROJECT.cloudfunctions.net/RECEIVING_FUNCTION"
func callFunction(ctx context.Context, method, functionURL string, headers map[string]string) ([]byte, error) {
	client, err := idtoken.NewClient(ctx, functionURL)
	if err != nil {
		return nil, fmt.Errorf("idtoken.NewClient: %v", err)
	}

	req, err := http.NewRequest(method, functionURL, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Get: %v", err)
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// callGetGoogleIDFunction は、GetGoogleID関数を呼び出す
func callGetGoogleIDFunction(ctx context.Context, tokenFromUser string) (string, error) {
	u := fmt.Sprintf("https://us-central1-%s.cloudfunctions.net/GetGoogleID", os.Getenv("GOOGLE_CLOUD_PROJECT"))
	h := map[string]string{"AuthorizationFromUser": tokenFromUser}

	b, err := callFunction(ctx, http.MethodGet, u, h)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
