package funcs

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"google.golang.org/api/idtoken"
)

// callFunction は、別のGCF関数を呼び出す関数
// functionURL = "https://REGION-PROJECT.cloudfunctions.net/RECEIVING_FUNCTION"
func callFunction(ctx context.Context, method, functionURL string, headers map[string]string) (string, error) {
	client, err := idtoken.NewClient(ctx, functionURL)
	if err != nil {
		return "", fmt.Errorf("idtoken.NewClient: %v", err)
	}

	req, err := http.NewRequest(method, functionURL, nil)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("client.Get: %v", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", errors.New(string(b))
	}

	return string(b), nil
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
