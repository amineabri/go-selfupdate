package selfupdate

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/blang/semver"
	"github.com/google/go-github/v30/github"
	"github.com/tcnksm/go-gitconfig"
	"golang.org/x/oauth2"
)

type UpdaterIn interface {
	DetectLatest(slug string) (release *Release, found bool, err error)
	DetectVersion(slug string, version string) (release *Release, found bool, err error)
	downloadDirectlyFromURL(assetURL string) (io.ReadCloser, error)
	UpdateTo(rel *Release, cmdPath string) error
	UpdateCommand(cmdPath string, current semver.Version, slug string) (*Release, error)
	UpdateSelf(current semver.Version, slug string) (*Release, error)
}

// Updater is responsible for managing the context of self-update.
// It contains GitHub client and its context.
type Updater struct {
	api       *github.Client
	apiCtx    context.Context //nolint:containedctx
	validator Validator
	filters   []*regexp.Regexp
	pre       bool
	draft     bool
}

// Config represents the configuration of self-update.
type Config struct {
	// APIToken represents GitHub API token. If it's not empty, it will be used for authentication of GitHub API
	APIToken string
	// EnterpriseBaseURL is a base URL of GitHub API. If you want to use this library with GitHub Enterprise,
	// please set "https://{your-organization-address}/api/v3/" to this field.
	EnterpriseBaseURL string
	// EnterpriseUploadURL is a URL to upload stuffs to GitHub Enterprise instance. This is often the same as an API base URL.
	// So if this field is not set and EnterpriseBaseURL is set, EnterpriseBaseURL is also set to this field.
	EnterpriseUploadURL string
	// Validator represents types which enable additional validation of downloaded release.
	Validator Validator
	// Filters are regexp used to filter on specific assets for releases with multiple assets.
	// An asset is selected if it matches any of those, in addition to the regular tag, os, arch, extensions.
	// Please make sure that your filter(s) uniquely match an asset.
	Filters []string
	// PreRelease indicates if pre-releases are allowed.
	PreRelease bool
	// Draft indicates if drafts are allowed.
	Draft bool
}

func newHTTPClient(ctx context.Context, token string) *http.Client {
	if token == "" {
		return http.DefaultClient
	}

	src := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	return oauth2.NewClient(ctx, src)
}

// NewUpdater creates a new updater instance. It initializes GitHub API client.
// If you set your API token to $GITHUB_TOKEN, the client will use it.
func NewUpdater(config Config) (*Updater, error) {
	token := config.APIToken
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}

	if token == "" {
		token, _ = gitconfig.GithubToken()
	}

	ctx := context.Background()

	hc := newHTTPClient(ctx, token)

	filtersRe := make([]*regexp.Regexp, 0, len(config.Filters))

	for _, filter := range config.Filters {
		re, err := regexp.Compile(filter)
		if err != nil {
			return nil, fmt.Errorf("could not compile regular expression %q for filtering releases: %w", filter, err)
		}

		filtersRe = append(filtersRe, re)
	}

	if config.EnterpriseBaseURL == "" {
		client := github.NewClient(hc)

		return &Updater{api: client, apiCtx: ctx, validator: config.Validator, filters: filtersRe, pre: config.PreRelease, draft: config.Draft}, nil
	}

	u := config.EnterpriseUploadURL
	if u == "" {
		u = config.EnterpriseBaseURL
	}

	client, err := github.NewEnterpriseClient(config.EnterpriseBaseURL, u, hc)
	if err != nil {
		return nil, err
	}

	return &Updater{api: client, apiCtx: ctx, validator: config.Validator, filters: filtersRe, pre: config.PreRelease, draft: config.Draft}, nil
}

// DefaultUpdater creates a new updater instance with default configuration.
// It initializes GitHub API client with default API base URL.
// If you set your API token to $GITHUB_TOKEN, the client will use it.
func DefaultUpdater() *Updater {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		token, _ = gitconfig.GithubToken()
	}

	ctx := context.Background()

	client := newHTTPClient(ctx, token)

	return &Updater{api: github.NewClient(client), apiCtx: ctx}
}
