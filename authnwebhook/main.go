package main

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	authentication "k8s.io/api/authentication/v1beta1"
	"log"
	"net/http"
)

func authfunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tr authentication.TokenReview
	err := decoder.Decode(&tr)
	if err != nil {
		log.Println("[Error]", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind":       "TokenReview",
			"status": authentication.TokenReviewStatus{
				Authenticated: false,
			},
		})
		return
	}
	log.Printf("receving request", tr)
	// Check User
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tr.Spec.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		log.Println("[Error]", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind":       "TokenReview",
			"status": authentication.TokenReviewStatus{
				Authenticated: false,
			},
		})
		return
	}
	log.Printf("[Success] login as %s", *user.Login)
	w.WriteHeader(http.StatusOK)
	trs := authentication.TokenReviewStatus{
		Authenticated: true,
		User: authentication.UserInfo{
			Username: *user.Login,
			UID:      *user.Login,
		},
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"apiVersion": "authentication.k8s.io/v1beta1",
		"kind":       "TokenReview",
		"status":     trs,
	})
}

func main() {
	http.HandleFunc("/authenticate", authfunc)
	if err := http.ListenAndServe("0.0.0.0:3000", nil); err != nil {
		log.Println(err)
	}
}
