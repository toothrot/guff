package services

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"google.golang.org/api/oauth2/v2"

	"github.com/toothrot/guff/go/core"
)

type OAuth struct {
	Config *core.Config
}

func (o *OAuth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	s, err := o.Config.CookieStore.Get(r, "glogin")
	if err != nil {
		glog.Errorf("CookieStore.Get(_, %q) = %q", "glogin", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	u, err := uuid.NewRandom()
	if err != nil {
		glog.Errorf("uuid.NewRandom() = _, %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	s.Values["oauth-state"] = u.String()
	if err := sessions.Save(r, w); err != nil {
		glog.Errorf("s.Save() = %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	o.Config.OAuthConfig.AuthCodeURL(u.String())

	http.Redirect(w, r, o.Config.OAuthConfig.AuthCodeURL(u.String()), http.StatusFound)
}

func (o *OAuth) OAuth2Callback(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		glog.Errorf("parseform: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	l, err := o.Config.CookieStore.Get(r, "glogin")
	if err != nil {
		glog.Errorf("CookieStore.Get(_%q): %q", "glogin", err)
		// TODO: delete cookie? We're trying to delete it anyway.
	}
	l.Options.MaxAge = -1

	if r.Form.Get("state") != l.Values["oauth-state"].(string) {
		glog.Errorf("failed to validate oauth state")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	token, err := o.Config.OAuthConfig.Exchange(context.Background(), r.Form.Get("code"))
	if err != nil {
		glog.Errorf("exchange: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		glog.Errorf("userinfo: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	d := json.NewDecoder(resp.Body)
	user := &oauth2.Userinfoplus{}
	if err := d.Decode(user); err != nil {
		glog.Errorf("decode: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	session, err := o.Config.CookieStore.Get(r, "guff")
	if err != nil {
		glog.Errorf("cookiestore get error %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	session.Values["email"] = user.Email
	if err := sessions.Save(r, w); err != nil {
		glog.Errorf("session save: %q", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
