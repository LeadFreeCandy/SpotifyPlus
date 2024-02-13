package internal

import (
	"github.com/SpotifyPlus/internal/scope"
	"net/url"
	"reflect"
	"testing"
)

func TestAppState_GenerateAuthenticationURL(t *testing.T) {
	regular, _ := url.Parse("https://accounts.spotify.com/authorize?client_id=sampleID&redirect_uri=localhost%2Fcallback&response_type=code")
	regularScoped, _ := url.Parse("https://accounts.spotify.com/authorize?client_id=sampleID&redirect_uri=localhost%2Fcallback&response_type=code&scope=user-read-private&scope=user-create-partner")
	tests := []struct {
		name    string
		config  Config
		scopes  []scope.Scope
		want    *url.URL
		wantErr bool
	}{
		{
			name:   "Regular Case No scope",
			config: Config{clientID: "sampleID", redirectURI: "localhost/callback"},
			scopes: []scope.Scope{},
			want:   regular,
		},
		{
			name:   "Regular Case Scopes",
			config: Config{clientID: "sampleID", redirectURI: "localhost/callback"},
			scopes: []scope.Scope{scope.UserReadPrivate, scope.UserCreatePartner},
			want:   regularScoped,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &AppState{
				config: tt.config,
			}
			got, err := app.GenerateAuthenticationURL(tt.scopes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuthenticationURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			query := got.Query()
			query.Del("state") // State is random each time, it's only useful for when we handle the response
			got.RawQuery = query.Encode()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateAuthenticationURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
