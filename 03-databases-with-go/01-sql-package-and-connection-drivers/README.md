# 3.1 SQL Package and Connection Drivers

## Databases drivers
+ Go's database/sql package works internally with another package called [driver](https://pkg.go.dev/database/sql/driver).
+ The _driver_ package **contains only interfaces**, which must be implemented by a specific driver (these specific implementations are community-made packages).
+ There is a [list of drivers on the Go GitHub wiki](https://github.com/golang/go/wiki/SQLDrivers), it will work with _lib/pq_ for PostgreSQL and _go-sql-driver/MySQL_ for MySQL.


## DB Methods
+ `Query()`, `QueryContext()`, `QueryRow()`: They allow performing SQL queries (`SELECT`) to receive information from the database.
+ `Exec()`, `ExecContext()`: They allow executing an SQL statement (`INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `ALTER`) to alter information in the database. They return a [Result](https://pkg.go.dev/database/sql#Result).

**Note**: Actually, it can be used in any SQL statement for any of the methods presented, **but it is NOT a good practice, in some cases**.

### `Prepare()` Method
It allows communicating to the DB to create an instruction where a series of parameters will be received, in this way, the DB is already prepared and expects to receive only the arguments.

Returns a `Stmt` (statement) structure that has the same methods as the `DB` structure, it also includes a connection to the database, which should be closed later.

**Advantages**:
+ The statement can be repeatedly executed by passing the arguments.
+ SQL injection attacks are avoided, since the database already knows which SQL statement to execute and only waits for the arguments to be passed.