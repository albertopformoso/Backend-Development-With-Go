# Testing with Go

## Test types
### Unit Test

Validates that a functionality executed what expected from it.

### Table Driven Test

Writing table-driven tests is a common pattern for unit tests in Go.

With a table-driven test approach, we would define all test cases in a struct slice, which represents the test table. We then reduce the test implementation to one test function that loops over the test cases in the struct slice and makes appropriate assertions

### Integration Test

Validates the functionality which depends on an external service. For example: DBs, mailing services, communications (slack, telegram)...

Note: It's only integration if we use that external service.

```go
func ControllerSave(db databaseInterface, name string) {
    db.Save(name)
}
```

+ If I connect to the DB, then I'm making an integration test.
+ If I fake a DB, then I'm making a unit test.

### Mock Test

Mock testing is an approach to unit testing that lets you make assertions about how the code under test is interacting with other system modules. In mock testing, the dependencies are replaced with objects that simulate the behaviour of the real ones. Mainly in mock testing, we simulate the behaviour of a third-party service, we are using in our project.

---
# Go Useful Test Commands

## Test only one function 

With Regular Expressions you can test an specified test.
```go
go test -run <function name>
```

## Test performance 

```go
go test -bench=.
```

## Test subtest

```go
go test -run <function name>/<subtest name>
````

**Additional testing package**
```go
go get -u github.com/stretchr/testify
```