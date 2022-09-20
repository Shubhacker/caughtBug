package db

import (
	"caughtBug/caughtBug/fields"
	"database/sql"
)

func FetchBugInformation(database *sql.DB) (fields.BugInformation, error) {
	var bug fields.BugInformation
	sqlQuery := `select * from dbo.bugInfo`
	row, err := database.Query(sqlQuery)
	if err != nil {
		return bug, err
	}

	for row.Next() {
		err := row.Scan(
			&bug.UniqueId,
			&bug.BugDescription,
			&bug.BugTopic,
			&bug.ApplicationName,
			&bug.StillPresent,
			&bug.PostedBy,
		)
		if err != nil {
			return bug, err
		}
	}
	return bug, nil
}

func FetchBugInformationWithFilter(database *sql.DB, bugFilter fields.BugFilter) (fields.BugInformation, error) {
	var bug fields.BugInformation
	var isFilterPresent bool
	sqlQuery := `select * from dbo.bugInfo`
	row, err := database.Query(sqlQuery)
	if err != nil {
		return bug, err
	}

	if bugFilter.ApplicationName != nil || *bugFilter.ApplicationName != "" {
		if isFilterPresent {
			sqlQuery += ` and applicationName = ` + *bugFilter.ApplicationName
		} else {
			sqlQuery += ` where applicationName = ` + *bugFilter.ApplicationName
		}
	}

	if bugFilter.UniqueId != nil || *bugFilter.UniqueId != "" {
		if isFilterPresent {
			sqlQuery += ` and UniqueId = ` + *bugFilter.UniqueId
		} else {
			sqlQuery += ` where UniqueId = ` + *bugFilter.UniqueId
		}
	}

	if bugFilter.StillPresent != nil || *bugFilter.StillPresent != "" {
		if isFilterPresent {
			sqlQuery += ` and stillPresent = ` + *bugFilter.StillPresent
		} else {
			sqlQuery += ` where stillPresent = ` + *bugFilter.StillPresent
		}
	}

	if bugFilter.BugDescription != nil || *bugFilter.BugDescription != "" {
		if isFilterPresent {
			sqlQuery += ` and bugDescription = ` + *bugFilter.BugDescription
		} else {
			sqlQuery += ` where bugDescription = ` + *bugFilter.BugDescription
		}
	}

	for row.Next() {
		err := row.Scan(
			&bug.UniqueId,
			&bug.BugDescription,
			&bug.BugTopic,
			&bug.ApplicationName,
			&bug.StillPresent,
			&bug.PostedBy,
		)
		if err != nil {
			return bug, err
		}
	}
	return bug, nil
}
