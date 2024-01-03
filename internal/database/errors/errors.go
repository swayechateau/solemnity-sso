package errors

import "errors"

var (
	ErrIdEmpty                   = errors.New("id is empty")
	ErrIdInvalid                 = errors.New("id is invalid")
	ErrUriEmpty                  = errors.New("uri is empty")
	ErrUriInvalid                = errors.New("uri is invalid")
	ErrExtensionEmpty            = errors.New("extension is empty")
	ErrExtensionInvalid          = errors.New("extension is invalid")
	ErrEmailEmpty                = errors.New("email is empty")
	ErrEmailInvalid              = errors.New("email is invalid")
	ErrEmailHashEmpty            = errors.New("email hash is empty")
	ErrEmailHashInvalid          = errors.New("email hash is invalid")
	ErrProviderEmpty             = errors.New("provider is empty")
	ErrProviderInvalid           = errors.New("provider is invalid")
	ErrTokenEmpty                = errors.New("token is empty")
	ErrTokenInvalid              = errors.New("token is invalid")
	ErrProviderNameEmpty         = errors.New("provider name is empty")
	ErrProviderNameInvalid       = errors.New("provider name is invalid")
	ErrProviderUserIdEmpty       = errors.New("provider user id is empty")
	ErrProviderUserIdInvalid     = errors.New("provider user id is invalid")
	ErrProviderUserIdHashEmpty   = errors.New("provider user id hash is empty")
	ErrProviderUserIdHashInvalid = errors.New("provider user id hash is invalid")
	ErrProviderPrincipalEmpty    = errors.New("provider principal is empty")
	ErrProviderPrincipalInvalid  = errors.New("provider principal is invalid")
)
