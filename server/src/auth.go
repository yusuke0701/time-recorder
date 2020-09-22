package funcs

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/yusuke0701/goutils/firebase"
)

func init() {
	ctx := context.Background()

	if err := firebase.SetupWithoutAPIKey(ctx); err != nil {
		log.Fatalf("failed to setup firebase client: %v", err)
	}
}

// VerifyIDToken は、クライアント側から渡されたトークンの検証を行います。
// FirebaseAuth のトークンを想定しており、検証に成功すると、ログインしているユーザーの UID を返却します。
func VerifyIDToken(w http.ResponseWriter, r *http.Request) {
	// pre process
	ctx := r.Context()

	idToken := strings.TrimPrefix(r.Header.Get("AuthorizationFromUser"), "Bearer ")

	token, err := firebase.VerifyIDToken(ctx, idToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error verifying ID token: %v\n", err)
		return
	}

	// main process

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token.UID)
}
