package validators

import (
    "net"
    "net/http"
    "errors"
    "fmt"

    "github.com/jmcarbo/cas-server/types"
)

// ValidateRequest executes validation against plain request
func ValidateRequest(r *http.Request, config *types.Config) *types.CasError {
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return &types.CasError{Error: errors.New("Could not parse remote IP:Port."), CasErrorCode: types.CAS_ERROR_CODE_INTERNAL_ERROR}
    }

    casError := isRemoteAddrAllowed(ip)
    if casError != nil {
        return casError
    }

    fmt.Print(config.Ldap.Addresses[0])

    return nil
}

func isRemoteAddrAllowed(ip string) *types.CasError {
    return nil
    //return &types.CasError{Error: errors.New("The IP is currently not allowed."), CasErrorCode: types.CAS_ERROR_CODE_INTERNAL_ERROR}
}
