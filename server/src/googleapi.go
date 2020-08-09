package funcs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yusuke0701/goutils/googleapi/userinfo"
)

var userinfoHelper userinfo.Helper

// GetGoogleID は、GoogleIDを取得するための関数です。クライアント側から渡されたトークンが必要です。
func GetGoogleID(w http.ResponseWriter, r *http.Request) {
	// pre process

	if userinfoHelper == nil {
		userinfoHelper = userinfo.NewHelper(http.DefaultClient, true, 100)
	}

	token := strings.TrimPrefix(r.Header.Get("AuthorizationFromUser"), "Bearer ")
	res, err := userinfoHelper.CallUserInfoMeAPI(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	// main process

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res.ID)
}
