package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

func commRedirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	var start, end int
	for i, segment := range segments {
		if segment == "oauth" {
			start = i + 1
		} else if segment == "redirect" && start > 0 {
			end = i
			break
		}
	}
	brand := strings.Join(segments[start:end], "")

	if commConfig == nil {
		logger.Errorf("%s oauth access attempt without configuration", brand)
		fmt.Fprintf(w, "error: this website has not configured %s OAuth", brand)
		return
	}

	commenterToken := r.FormValue("commenterToken")

	_, err := commenterGetByCommenterToken(commenterToken)
	if err != nil && err != errorNoSuchToken {
		fmt.Fprintf(w, "error: %s\n", err.Error())
		return
	}

	var url string
	if brand == "google" {
		// oidc
		nonce, err := randString(16)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}
		setCallbackCookie(w, r, "nonce", nonce)

		url = commConfig.AuthCodeURL(commenterToken, oauth2.SetAuthURLParam("nonce", nonce))
	} else {
		url = commConfig.AuthCodeURL(commenterToken)
	}

	http.Redirect(w, r, url, http.StatusFound)
}
