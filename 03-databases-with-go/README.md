# 3 Databases with Go

## Section Objectives
+ Use the database/sql package (specifically the implementation in PostgreSQL and MySQL).
+ Using the DAO pattern.
+ Deploy through an ORM.

## Shortcuts
+ [3.1 SQL Package and Connection Drivers](https://github.com/albertopformoso/Backend-Development-With-Go/tree/master/03-databases-with-go/01-sql-package-and-connection-drivers)
+ [3.2 CRUD with PostgreSQL](https://github.com/albertopformoso/Backend-Development-With-Go/tree/master/03-databases-with-go/02-crud-postgresql)
+ [3.3 CRUD with MySQL](https://github.com/albertopformoso/Backend-Development-With-Go/tree/master/03-databases-with-go/03-crud-mysql)
+ [3.4 ORM](https://github.com/albertopformoso/Backend-Development-With-Go/tree/master/03-databases-with-go/04-orm)

## DAO Pattern (implemented on section 3.3)

(Data Access Object) This pattern allows data to be collected from a site and stored somewhere, it decouples the business logic from the storage system.

### Implementation of the DAO Pattern

+ It must have an interface with a series of methods that will communicate with the business logic.
+ Then you must have a series of objects in charge of implementing said interface (the specific implementation of the storage system).
+ It is complemented by the Creational Factory Pattern, which allows obtaining instances of specific implementations.

## ORM 
It allows to map the structure of the database with structures in Go, it is through these structures that the ORM will interact with the DB (for example, defining the tables such as indicating their primary keys, their foreign keys, the types of fields, etc).

### Advantages
+ It will have encapsulated all the complex logic of the SQL instructions that will be carried out in the DB and will only expose the necessary methods to work (read, insert, update, delete records), in this way **you do not have to write SQL instructions**.
+ **It brings implemented the most used drivers** (Postgres, MySQL, SQL Server, SQLite), in this way, through the methods that are exposed, it does not matter the DBMS with which you work, you could change the DB and neither the logic nor the code will need to be changed.
+ It allows to simplify the work and accelerate the development of the application. **The code will be much shorter and faster to perform**.

### Disadvantages
+ You have to inspect the structures in Go to be able to create the SQL statements, when you have very large projects where there are **too many structures, too many relationships, too many tables**, this affects the overall performance of the application.
+ When working with large applications, you will often need to do advanced queries. Through the methods that the ORM offers, a series of such methods are nested to create the instruction, **when doing so it can be complicated** and an ORM option must be enabled to see the SQL instructions that are being built below. **It's easier to write the SQL directly than to use ORM**.