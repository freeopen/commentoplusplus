package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

func googleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{
		ClientID: googleConfig.ClientID,
	}
	verifier := provider.Verifier(oidcConfig)

	commenterToken := r.FormValue("state")
	code := r.FormValue("code")

	_, err = commenterGetByCommenterToken(commenterToken)
	if err != nil && err != errorNoSuchToken {
		fmt.Fprintf(w, "Error: %s\n", err.Error())
		return
	}

	// token, err := googleConfig.Exchange(oauth2.NoContext, code)
	token, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := r.Cookie("nonce")
	if err != nil {
		http.Error(w, "nonce not found", http.StatusBadRequest)
		return
	}
	if idToken.Nonce != nonce.Value {
		http.Error(w, "nonce did not match", http.StatusBadRequest)
		return
	}

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// defer resp.Body.Close()
	//
	// contents, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, "Error: %s", errorCannotReadResponse.Error())
	// 	return
	// }
	//
	// user := make(map[string]interface{})
	// if err := json.Unmarshal(contents, &user); err != nil {
	// 	fmt.Fprintf(w, "Error: %s", errorInternal.Error())
	// 	return
	// }
	//
	// if user["email"] == nil {
	// 	fmt.Fprintf(w, "Error: no email address returned by Github")
	// 	return
	// }
	//
	// email := user["email"].(string)
	//

	c, err := commenterGetByEmail("google", email)
	if err != nil && err != errorNoSuchCommenter {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	//
	// name := user["name"].(string)
	//
	// link := "undefined"
	// if user["link"] != nil {
	// 	link = user["link"].(string)
	// }
	//
	// photo := "undefined"
	// if user["picture"] != nil {
	// 	photo = user["picture"].(string)
	// }

	var commenterHex string

	if err == errorNoSuchCommenter {
		commenterHex, err = commenterNew(userInfo.Email, userInfo.Profile, link, photo, "google", "")
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
			return
		}
	} else {
		if err = commenterUpdate(c.CommenterHex, email, name, link, photo, "google"); err != nil {
			logger.Warningf("cannot update commenter: %s", err)
			// not a serious enough to exit with an error
		}

		commenterHex = c.CommenterHex
	}

	if err := commenterSessionUpdate(commenterToken, commenterHex); err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "<html><script>window.parent.close()</script></html>")
}
