package funcs

import (
	"net/http"

	"github.com/yusuke0701/goutils/gcp"
	"github.com/yusuke0701/goutils/googleapi/userinfo"
)

var defaultHTTPClient = http.DefaultClient

func setHeaderForCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.WriteHeader(http.StatusNoContent)
}

func getGoogleID(token string) (string, error) {
	var h userinfo.Helper
	if gcp.OnGCP() {
		h = userinfo.NewHelper(defaultHTTPClient)
	} else {
		h = userinfo.NewMockHelper()
		userinfo.CallUserInfoMeAPIMockData = &userinfo.CallUserInfoMeAPIRes{
			ID:            "sampleID",
			Email:         "sample@example.com",
			VerifiedEmail: true,
			Picture:       "samplePicture",
		}
	}
	res, err := h.CallUserInfoMeAPI(token)
	if err != nil {
		return "", err
	}
	return res.ID, nil
}
