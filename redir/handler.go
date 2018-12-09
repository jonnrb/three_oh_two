package redir // import "go.jonnrb.io/three_oh_two/redir"

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// Returns an http.Handler that returns a 302 where the Location is the request
// path rewritten to be rooted at `base`.
func Handler(base string) (http.Handler, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newURL, err := url.Parse(u.String())
		if err != nil {
			panic(fmt.Sprintf("Error copying redirect URL root: %v", err))
		}
		newURL.Path = path.Join(u.Path, r.URL.Path)
		http.Redirect(w, r, newURL.String(), http.StatusFound)
	}), nil
}
