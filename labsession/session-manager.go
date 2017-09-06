package labsession

import (
	"log"
	"time"

	"github.com/6lab/6lib/labcrypt"
	"github.com/6lab/6lib/labtime"
)

var (
	database *DATABASE
)


	// Connect to Couchbase
	sl, err := newSQLite("/", )
	if err != nil {
		log.Fatalln("Couchbase connection error:", err.Error())
	}

		// Set Database
	database = &DATABASE{accesser: sl}

// SESSIONSMANAGER allow to manage the Sessions
type SESSIONMANAGER struct {
	sync.RWMutex
	strategy int8
	sessionsList map[string]*SESSION
}

// NewSessionManager creates a new SESSIONMANAGER
func NewSessionManger() *SESSIONMANAGER {
	sessions := new(SESSIONMANAGER)

	sessions.sessionsList = make(map[string]*SESSION)

	return sessions
}

func (sessions *SESSIONMANAGER) add(idUser, ipAddr, remoteAddr string, maxAge int) string {
	user := app.users.getUser(idUser)

	if user == nil {
		return ""
	}

	session := newSession(idUser, ipAddr, remoteAddr, maxAge)
	if session == nil {
		return ""
	}

	// Save Session
	DBExec(querySessionSQLSave(session))

	// Lock sessions before adding
	sessions.Lock()
	defer sessions.Unlock()

	sessions.sessionsList[session.IdSession] = session

	return session.IDSession
}

func (sessions *SESSIONMANAGER) readAllSessions(clean bool) {
	var sessions2Clean []string

	// Execute query
	rows, err := app.db.Query(querySessionSQLReadAll())

	// Read Sessions
	if err != nil {
		log.Println(err)
	} else {
		defer rows.Close()

		for rows.Next() {
			var idSession string
			var params string

			rows.Scan(&idSession, &params)

			session := new(SESSION)
			decryptJSONSession(params, session)

			// Add User in Memory
			if session != nil && session.Expire > xttime.GetCurrentTime() {
				if !clean {
					// Lock is not needed because the load is made only one time at start before concurrency is possible
					sessions.sessionsList[idSession] = session
				}
			} else {
				sessions2Clean = append(sessions2Clean, idSession)
			}
		}
	}
}

func (sessions *SESSIONMANAGER) deleteAllSessions() {
	// Delete in DB
	DBExec(querySessionSQLDeleteAll())
	// No need to delete in memory because this method can only be used before to load in memory
}

func (sessions *SESSIONMANAGER) deleteSession(idSession string) {
	if idSession != "" {
		// Delete in DB
		DBExec(querySessionSQLDelete(idSession))

		// Lock sessions before deleting
		sessions.Lock()
		defer sessions.Unlock()

		// Delete in Memory
		delete(sessions.sessionsList, idSession)
	}
}

// func (sessions *SESSIONMANAGER) cleanSessions() {
// 	t := time.NewTicker(time.Hour * 1)
// 	apply := false

// 	for {
// 		if apply {
// 			sessions.readAllSessions(true)
// 		}

// 		// Avoid first clean
// 		apply = true

// 		// Lock until next Period
// 		<-t.C
// 	}
// }
