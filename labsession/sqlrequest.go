package labsession

import (
	"github.com/6lab/6lib/labcast"
	"github.com/6lab/6lib/labcrypt"
	"github.com/6lab/6lib/labstring"
)

func getSessionSQLCreationScript() string {
	return `
CREATE TABLE IF NOT EXISTS tSession (
	idSession TEXT PRIMARY KEY,
	params TEXT
);

CREATE INDEX IF NOT EXISTS xidSession ON tSession (idSession);
--------------------------
`
}

func querySessionSQLReadAll() string {
	return `SELECT idSession, params FROM tSession`
}

func querySessionSQLReadOne(idSession string) string {
	query := `SELECT idSession, params FROM tSession WHERE idSession = '%1'`

	return xtstring.StringBuild(query, idSession)
}

func querySessionSQLDelete(idSession string) string {
	query := `DELETE FROM tSession WHERE idSession = '%1'`

	return xtstring.StringBuild(query, idSession)
}

func querySessionSQLDeleteAll() string {
	return `DELETE FROM tSession`
}

func querySessionSQLDeleteList(sessionsList []string) string {
	sessions := xtcast.ConvertArrayStringToListOfParams(sessionsList)
	query := `DELETE FROM tSession WHERE idSession IN (%1)`

	return xtstring.StringBuild(query, sessions)
}

func querySessionSQLSave(session *SESSION) string {
	query := `
INSERT OR REPLACE INTO tSession (
	idSession,
	params
) VALUES (
	'%1',
	'%2'
)
`
	return xtstring.StringBuild(query,
		session.IdSession,
		xtcrypt.EncryptJSON(session))
}
