package typos

import (
	catcher "github.com/ArseniySavin/catcher/pkg"
)

var (
	newErr                 = catcher.BaseError.New
	ErrClientNotFound      = newErr("", "Client cannot found in cache")
	ErrClaimsNotFound      = newErr("", "Claims cannot found in cache")
	ErrConversion          = newErr("", "Client data into cache is not consistency")
	ErrInvalidScope        = newErr("", "Client have not permit for requested scope")
	ErrInvalidGrant        = newErr("", "Client have not permit for endpoint")
	ErrAccessTokenNotFound = newErr("", "access_token have not into cache")

	ErrEmptyGrantType             = newErr("", "Do not set grant_type")
	ErrInvalidGrantType           = newErr("", "grant_type is invalid")
	ErrEmptyScope                 = newErr("", "Do not set scope")
	ErrEmptyToken                 = newErr("", "Do not set access_token")
	ErrEmptyAuthorization         = newErr("", "Do not set HEADER Authorization")
	ErrInvalidFormatAuthorization = newErr("", "Authorization header do not equal auth_type auth_data")
	ErrInvalidTypeAuthorization   = newErr("", "You use unsupported type Authorization. Use Basic authorization")
	ErrConversionAuthorization    = newErr("", "Authorization header do not equal base64")
	ErrInvalidFormatAuthData      = newErr("", "Authorization header do not equal base64 in password")
	ErrInvalidToken               = newErr("", "Token expired")
	ErrTokenIsDemolished          = newErr("", "Token is disaster or format is invalid")

	// jwt errors
	ErrEmptySecret = newErr("", "secret is empty")
	ErrEmptyJwt    = newErr("", "jwt is empty")
	ErrEmptySalt   = newErr("", "salt is empty")

	// cache
	ErrCacheNotOpen          = newErr("", "Can not open connection for cache")
	ErrDefaultCacheUndefined = newErr("", "Default cache do not undefined")
	ErrUnknownCache          = newErr("", "Unknown cache")
	ErrStringConversion      = newErr("", "Can not convert cache data to string")
)
