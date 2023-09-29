package tokens

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Scope allows a created token to be limited to a specific domain or project.
type Scope struct {
	ProjectID   string
	ProjectName string
	DomainID    string
	DomainName  string
}

// AuthOptionsBuilder provides the ability for extensions to add additional
// parameters to AuthOptions. Extensions must satisfy all required methods.
type AuthOptionsBuilder interface {
	// ToTokenV3CreateMap assembles the Create request body, returning an error
	// if parameters are missing or inconsistent.
	ToTokenV3CreateMap(map[string]interface{}) (map[string]interface{}, error)
	ToTokenV3ScopeMap() (map[string]interface{}, error)
	CanReauth() bool
	AuthTokenID() string
	AuthHeaderDomainID() string
}

// AuthOptions represents options for authenticating a user.
type AuthOptions struct {
	// IdentityEndpoint specifies the HTTP endpoint that is required to work with
	// the Identity API of the appropriate version. While it's ultimately needed
	// by all of the identity services, it will often be populated by a
	// provider-level function.
	IdentityEndpoint string `json:"-"`

	// Username is required if using Identity V2 API. Consult with your provider's
	// control panel to discover your account's username. In Identity V3, either
	// UserID or a combination of Username and DomainID or DomainName are needed.
	Username string `json:"username,omitempty"`
	UserID   string `json:"id,omitempty"`

	Password string `json:"password,omitempty"`

	// At most one of DomainID and DomainName must be provided if using Username
	// with Identity V3. Otherwise, either are optional.
	DomainID   string `json:"-"`
	DomainName string `json:"name,omitempty"`

	// AllowReauth should be set to true if you grant permission for Gophercloud
	// to cache your credentials in memory, and to allow Gophercloud to attempt
	// to re-authenticate automatically if/when your token expires.  If you set
	// it to false, it will not cache these settings, but re-authentication will
	// not be possible.  This setting defaults to false.
	AllowReauth bool `json:"-"`

	// TokenID allows users to authenticate (possibly as another user) with an
	// authentication token ID.
	TokenID string `json:"-"`

	// Passcode is a Virtual MFA device verification code, which can be obtained on the MFA app.
	Passcode string `json:"-"`

	Scope Scope `json:"-"`
}

// ToTokenV3CreateMap builds a request body from AuthOptions.
func (opts *AuthOptions) ToTokenV3CreateMap(scope map[string]interface{}) (map[string]interface{}, error) {
	golangsdkAuthOpts := golangsdk.AuthOptions{
		Username:    opts.Username,
		UserID:      opts.UserID,
		Password:    opts.Password,
		DomainID:    opts.DomainID,
		DomainName:  opts.DomainName,
		AllowReauth: opts.AllowReauth,
		TokenID:     opts.TokenID,
		Passcode:    opts.Passcode,
	}

	return golangsdkAuthOpts.ToTokenV3CreateMap(scope)
}

// ToTokenV3ScopeMap builds a scope request body from AuthOptions.
func (opts *AuthOptions) ToTokenV3ScopeMap() (map[string]interface{}, error) {
	if opts.Scope.ProjectName != "" {
		// ProjectName provided: either DomainID or DomainName must also be supplied.
		// ProjectID may not be supplied.
		if opts.Scope.DomainID == "" && opts.Scope.DomainName == "" {
			return nil, golangsdk.ErrScopeDomainIDOrDomainName{}
		}
		if opts.Scope.ProjectID != "" {
			return nil, golangsdk.ErrScopeProjectIDOrProjectName{}
		}

		if opts.Scope.DomainID != "" {
			// ProjectName + DomainID
			return map[string]interface{}{
				"project": map[string]interface{}{
					"name":   &opts.Scope.ProjectName,
					"domain": map[string]interface{}{"id": &opts.Scope.DomainID},
				},
			}, nil
		}

		if opts.Scope.DomainName != "" {
			// ProjectName + DomainName
			return map[string]interface{}{
				"project": map[string]interface{}{
					"name":   &opts.Scope.ProjectName,
					"domain": map[string]interface{}{"name": &opts.Scope.DomainName},
				},
			}, nil
		}
	} else if opts.Scope.ProjectID != "" {
		// ProjectID provided. ProjectName, DomainID, and DomainName may not be provided.
		if opts.Scope.DomainID != "" {
			return nil, golangsdk.ErrScopeProjectIDAlone{}
		}
		if opts.Scope.DomainName != "" {
			return nil, golangsdk.ErrScopeProjectIDAlone{}
		}

		// ProjectID
		return map[string]interface{}{
			"project": map[string]interface{}{
				"id": &opts.Scope.ProjectID,
			},
		}, nil
	} else if opts.Scope.DomainID != "" {
		// DomainID provided. ProjectID, ProjectName, and DomainName may not be provided.
		if opts.Scope.DomainName != "" {
			return nil, golangsdk.ErrScopeDomainIDOrDomainName{}
		}

		// DomainID
		return map[string]interface{}{
			"domain": map[string]interface{}{
				"id": &opts.Scope.DomainID,
			},
		}, nil
	} else if opts.Scope.DomainName != "" {
		// DomainName
		return map[string]interface{}{
			"domain": map[string]interface{}{
				"name": &opts.Scope.DomainName,
			},
		}, nil
	}

	return nil, nil
}

func (opts *AuthOptions) CanReauth() bool {
	return opts.AllowReauth
}

func (opts *AuthOptions) AuthTokenID() string {
	return ""
}

func (opts *AuthOptions) AuthHeaderDomainID() string {
	return ""
}

func subjectTokenHeaders(_ *golangsdk.ServiceClient, subjectToken string) map[string]string {
	return map[string]string{
		"X-Subject-Token": subjectToken,
	}
}

// Create authenticates and either generates a new token, or changes the Scope
// of an existing token.
func Create(c *golangsdk.ServiceClient, opts AuthOptionsBuilder) (r CreateResult) {
	scope, err := opts.ToTokenV3ScopeMap()
	if err != nil {
		r.Err = err
		return
	}

	b, err := opts.ToTokenV3CreateMap(scope)
	if err != nil {
		r.Err = err
		return
	}

	resp, err := c.Post(tokenURL(c), b, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"X-Auth-Token": opts.AuthTokenID(),
			"X-Domain-Id":  opts.AuthHeaderDomainID(),
		},
	})
	r.Err = err
	if resp != nil {
		r.Header = resp.Header
	}
	return
}

// Get validates and retrieves information about another token.
func Get(c *golangsdk.ServiceClient, token string) (r GetResult) {
	resp, err := c.Get(tokenURL(c), &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: subjectTokenHeaders(c, token),
		OkCodes:     []int{200, 203},
	})
	if resp != nil {
		r.Err = err
		r.Header = resp.Header
	}
	return
}

// Validate determines if a specified token is valid or not.
func Validate(c *golangsdk.ServiceClient, token string) (bool, error) {
	resp, err := c.Request("HEAD", tokenURL(c), &golangsdk.RequestOpts{
		MoreHeaders: subjectTokenHeaders(c, token),
		OkCodes:     []int{200, 204, 404},
	})
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200 || resp.StatusCode == 204, nil
}

// Revoke immediately makes specified token invalid.
func Revoke(c *golangsdk.ServiceClient, token string) (r RevokeResult) {
	_, r.Err = c.Delete(tokenURL(c), &golangsdk.RequestOpts{
		MoreHeaders: subjectTokenHeaders(c, token),
	})
	return
}
