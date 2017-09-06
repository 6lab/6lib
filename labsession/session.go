package labsession

import (
	"encoding/json"

	"github.com/6lab/6lib/labcrypt"
	"github.com/6lab/6lib/labrand"
	"github.com/6lab/6lib/labtime"
)

const (
	// MaxAge contains the default Validity Time for the Session (24 hours)
	MaxAge = 3600 * 24
)

type SESSION struct {
	IDSession  string
	IDUser     string
	Expire     int64
	IpAddr     string
	RemoteAddr string
}

func getSessionInvalid(internalSessionOnly bool) *INVALIDSESSION {
	return &INVALIDSESSION{
		Valid:        false,
		InternalOnly: internalSessionOnly,
	}
}

// NewSession creates a new SESSION
func newSession(idUser, ipAddr, remoteAddr string, maxAge int) *SESSION {
	session := new(SESSION)

	if maxAge == 0 {
		maxAge = MaxAge
	}

	// Init
	session.IDSession = xtrand.GetUUID()
	session.IDUser = idUser
	session.Expire = xttime.GetTimeAfterOffset(maxAge)
	session.IpAddr = ipAddr
	session.RemoteAddr = remoteAddr

	return session
}

func decryptJSONSession(s string, session *SESSION) error {
	err := json.Unmarshal(xtcrypt.DecodeBase64(s), &session)

	return err
}
