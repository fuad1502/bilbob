package sessions

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

type session struct {
	username  string
	logged_in bool
}

var sessions = make(map[string]session)

func CreateSession(username string) string {
	sessionId := generateSessionId()
	newSession := session{username, true}
	sessions[sessionId] = newSession
	return sessionId
}

func IsLoggedIn(sessionId string) bool {
	sess, ok := sessions[sessionId]
	if ok {
		return sess.logged_in
	} else {
		return false
	}
}

func GetUsername(sessionId string) (string, bool) {
	sess, ok := sessions[sessionId]
	return sess.username, ok
}

func generateSessionId() string {
	sessionIdBytes := make([]byte, 16)
	for i := 0; i < 16; i += 1 {
		r, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			panic("")
		}
		sessionIdBytes[i] = byte(r.Int64())
	}
	return hex.EncodeToString(sessionIdBytes)
}
