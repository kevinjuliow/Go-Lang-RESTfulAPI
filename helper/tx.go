package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover() //If there's a panic . it wont directly stop the program .
	if err != nil {
		errorRollBack := tx.Rollback()
		PanicIfError(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
