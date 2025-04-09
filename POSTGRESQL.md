# postgresql
## Table of content
- [Comparison](#comparison-postgresql-vs-mysql-postgresql-vs-sql-server)
- [CRUD](#crud)
- [Join](#join)
- [Sub query](#sub-query)
- [Index](#index)
- [Partition](#partition)
- [Transaction](#transaction)

## Comparison: PostgreSQL vs MySQL, PostgreSQL vs SQL Server
### PostgreSQL vs MySQL
Both PostgreSQL and MySQL are open-source relational database management systems (RDBMS), but they have different use cases and strengths.

| Feature | PostgreSQL | MySQL |
|---------|--------------|--------------|
| **Database Technology** | Object-Relational Database Management | Relational Database Management |
| **ACID Compliance** | Fully ACID-compliant in all configurations | Fully ACID-compliant when using InnoDB (default engine) |
| **Performance** | Better for complex queries, analytics, and write-heavy applications | Faster for read-heavy workloads and simple queries |
| **Indexing** | Supports advanced indexing: GIN, GiST, BRIN, and full-text search | Basic indexing with B-Tree and R-Tree |
| **Data types** | Supports varous data types: Character, numeric, Date/Time, spartial, JSON, monetary, binary, boolean, geometric, list, network address, array, UUID, XML, domain and composity. | Simpler data type: Character, numeric, Date and Time, spartial, JSON |
| **JSON Support** | Advanced JSONB support for semi-structured data | JSON functions available but not as optimized as PostgreSQL |
| **Stored Procedures & Functions** | Supports procedural languages (PL/pgSQL, PL/Python, etc.) | Supports stored procedures but with fewer built-in language options |

### PostgreSQL vs SQL Server
PostgreSQL and Microsoft SQL Server (SQL Server) are both powerful RDBMS, but SQL Server is a proprietary system mainly used in enterprise environments.

| Feature | PostgreSQL | SQL Server |
|---------|--------------|--------------|
| **Platform Compatibility** | Runs on Linux, Windows, macOS | Primarily Windows, with some support for Linux |
| **ACID Compliance** | Fully ACID-compliant | Fully ACID-compliant |
| **Performance** | Optimized for complex queries and analytics | Strong performance, optimized for Windows environments |
| **Indexing** | Supports advanced indexing | Supports advanced indexing |
| **Data types** | Supports varous data types | Supports varous data types |
| **JSON Support** | Full JSONB support | JSON support but not as advanced as PostgreSQL |
| **Stored Procedures & Functions** | Supports PL/pgSQL, PL/Python, and other procedural languages | Uses T-SQL for stored procedures and functions |
| **Cost** | Free and open-source | Requires licensing fees, can be expensive for large deployments |

## CRUD
CRUD stands for **Create**, **Read**, **Update**, and **Delete**, which are the four basic operations for managing data in a database, including PostgreSQL.
### Create
Used to add new records to a table
```
newUser := &model.User{
	Email:    email,
	Password: password,
}
res := ur.Db.Create(newUser)
```

### Read
Used to retrieve records from a table
```
res = ur.Db.Find(user, model.User{Email: email})
```

### Update
Used to modify existing records
```
res := ur.Db.Model(user).Update("password", password)
```

### Delete
Used to remove records from a table
```
res := ur.Db.Delete(user)

```

## Foreign key
A **foreign key** is a column (or a group of columns) in one table that references the **primary key** of another table, establishing a link between the two tables, ensuring that any data entered into the foreign key column must already exist in the parent table.

In PostgreSQL, foreign keys come with several constraints when creating or altering a table:
- **ON DELETE CASCADE**: Automatically deletes any rows in the child table when the corresponding row in the parent table is deleted.
- **ON DELETE SET NULL**: Sets the foreign key value in the child table to NULL when the corresponding row in the parent table is deleted.
- **ON UPDATE CASCADE**: Updates the foreign key in the child table when the corresponding primary key in the parent table is updated

```
type User struct {
	Id        uint `gorm:"primaryKey"`
	Password  string
	Email     string         `gorm:"unique;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Logs      []Log          `gorm:"foreignKey:UserId;"`
}
```

## Join
**JOIN** is used to combine columns from one (self-join) or more tables based on the values of the common columns between related tables.
### Inner join
`INNER JOIN` returns only matching rows from both tables. If there is no match, the row is excluded.

### Left join
`LEFT JOIN` returns all rows from the left table and only matching rows from the right table. If no match is found, NULL is returned for columns from the right table.

### Right join
`RIGHT JOIN` returns all rows from the right table and only matching rows from the left table. If no match is found, NULL is returned for columns from the left table.

### Full outer join
`FULL OUTER JOIN` combine data from two tables and returns all rows from both tables, including matching and non-matching rows from both sides. If no match is found, NULL is returned for missing values.

### Cross join
`CROSS JOIN` returns the cartesian product of two tables. Its clause does not have a join predicate.

### Natural join
`NATURAL JOIN` is a join that creates an implicit join based on the same column names in the joined tables.

### Self-join
A **self-join** is a regular join that joins a table to itself.

## Sub query
A subquery is a query nested inside another query. Subqueries are used to perform complex data retrieval operations and can return individual values or a set of rows that the main query uses for its conditions.

## Index
### B-Tree Index
B-Tree is the default index type in PostgreSQL and is well-suited for most scenarios. In particular, the PostgreSQL query planner will consider using a B-tree index whenever an indexed column is involved in a comparison.

### GiST (Generalized Search Tree) Index
GiST indices are flexible and support a wide range of data types and search operations. They are particularly useful for spatial and full-text search queries.

### GIN (Generalized Inverted Tree) Index
GIN indices are designed for handling complex data types such as arrays and full-text searches.

### BRIN
BRIN indices are suitable for large tables with ordered data. They divide the table into blocks and store summarized information for each block, making them efficient for range queries on sorted data.

```
type Log struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"not null;index"`
	Operation string `gorm:"not null;index"`
	Collumn   string
	OldData   string
	NewData   string
	CreatedAt time.Time
}
```

## Partition
Partitioning refers to splitting logically one large table into smaller physical pieces. The table that is divided is referred to as a **partitioned table**. The partitioned table itself is a “virtual” table having no storage of its own. Instead, the storage belongs to **partitions**, which are tables associated with the partitioned table. Each partition stores a subset of the data as defined by its partition bounds. All rows inserted into a partitioned table will be routed to the appropriate one of the partitions based on the values of the partition key column(s). Updating the partition key of a row will cause it to be moved into a different partition if it no longer satisfies the partition bounds of its original partition.
### Range Partitioning
The table is partitioned into “ranges” defined by a key column or set of columns, with no overlap between the ranges of values assigned to different partitions. Each range's bounds are understood as being inclusive at the lower end and exclusive at the upper end.

### List Partitioning
The table is partitioned by explicitly listing which key value(s) appear in each partition.

### Hash Partitioning
The table is partitioned by specifying a modulus and a remainder for each partition. Each partition will hold the rows for which the hash value of the partition key divided by the specified modulus will produce the specified remainder.

## Transaction
A database transaction is a single unit of work that consists of one or more operations.
A classical example of a transaction is a bank transfer from one account to another. A complete transaction must ensure a balance between the sender and receiver accounts.
A PostgreSQL transaction is atomic, consistent, isolated, and durable. These properties are often referred to collectively as ACID:
- **Atomicity** guarantees that the transaction is completed in an all-or-nothing manner.
- **Consistency** ensures that changes to data written to the database are valid and adhere to predefined rules.
- **Isolation** determines how the integrity of a transaction is visible to other transactions.
- **Durability** ensures that transactions that have been committed are permanently stored in the database.
```
db.Transaction(func(tx *gorm.DB) error {
	userService := config.ConnectServices(tx)
	for i := range 500 {
		email := fmt.Sprint(i, "@gmail.com")
		_, err := userService.Register(email, string(rune(i)))
		if err != nil {
			return err
		}
	}
	return nil
})
```
