package funcs

import (
	"fmt"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var (
	firebaseApp        *firebase.App
	firebaseAuthClient *auth.Client
)

// VerifyIDToken は、クライアント側から渡されたトークンの検証を行います。
// FirebaseAuth のトークンを想定しており、検証に成功すると、ログインしているユーザーの UID を返却します。
func VerifyIDToken(w http.ResponseWriter, r *http.Request) {
	// pre process
	ctx := r.Context()

	if firebaseApp == nil {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error initializing app: %v\n", err)
			return
		}
		firebaseApp = app
	}

	if firebaseAuthClient == nil {
		client, err := firebaseApp.Auth(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error getting Auth client: %v\n", err)
			return
		}
		firebaseAuthClient = client
	}

	idToken := strings.TrimPrefix(r.Header.Get("AuthorizationFromUser"), "Bearer ")

	token, err := firebaseAuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error verifying ID token: %v\n", err)
		return
	}

	// main process

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token.UID)
}
