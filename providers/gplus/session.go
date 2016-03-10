package gplus

import (
	"encoding/json"
	"errors"
	"time"
	"github.com/markbates/goth"
	"golang.org/x/oauth2"
)

// Session stores data during the auth process with Facebook.
type Session struct {
	AuthURL      string
	AccessToken  string
	RefreshToken string
	ExpiresIn    time.Time
}

// GetAuthURL will return the URL set by calling the `BeginAuth` function on the Google+ provider.
func (s Session) GetAuthURL() (string, error) {
	if s.AuthURL == "" {
		return "", errors.New("an AuthURL has not be set")
	}
	return s.AuthURL, nil
}

// Authorize the session with Google+ and return the access token to be stored for future use.
func (s *Session) Authorize(provider goth.Provider, params goth.Params) (string, error) {
	p := provider.(*Provider)
	token, err := p.config.Exchange(oauth2.NoContext, params.Get("code"))
	if err != nil {
		return "", err
	}
	s.AccessToken = token.AccessToken
	s.RefreshToken = token.RefreshToken
	s.ExpiresIn = token.Expiry
	return token.AccessToken, err
}

// Marshal the session into a string
func (s Session) Marshal() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s Session) String() string {
	return s.Marshal()
}
