// fix package contains the logics of the fix operations, fix operations
// are in charge of adapting our source code to comply with newer versions
// of the CLI.
package fix

// Things to Fix:

// 1. models/models.go has changed its structure not to use an init function to
// set the database, it now provides a method to return the database connection
