# postgreesql
## Comparison: PostgreSQL vs MySQL, PostgreSQL vs SQL Server
### PostgreSQL vs MySQL
Both PostgreSQL and MySQL are open-source relational database management systems (RDBMS), but they have different use cases and strengths.

| Feature | PostgreSQL | MySQL |
|---------|--------------|--------------|
| **ACID Compliance** | Fully ACID-compliant in all configurations | Fully ACID-compliant when using InnoDB (default engine) |
| **Performance** | Better for complex queries, analytics, and write-heavy applications | Faster for read-heavy workloads and simple queries |
| **Indexing** | Supports advanced indexing: GIN, GiST, BRIN, and full-text search | Basic indexing with B-Tree and some full-text search |
| **Replication & Clustering** | Supports streaming replication and logical replication natively | Supports replication but requires third-party tools for clustering |
| **JSON Support** | Advanced JSONB support for semi-structured data | JSON functions available but not as optimized as PostgreSQL |
| **Stored Procedures & Functions** | Supports procedural languages (PL/pgSQL, PL/Python, etc.) | Supports stored procedures but with fewer built-in language options |
| **Concurrency Control** | Uses MVCC (Multi-Version Concurrency Control) to handle concurrent transactions efficiently | Also uses MVCC but may have issues under high concurrency |
| **Extensibility** | Highly extensible with custom data types, operators, and procedural languages | Limited extensibility compared to PostgreSQL |
| **Community & Support** | Large and active community, with strong enterprise support | Larger adoption, strong community support |

### PostgreSQL vs SQL Server
PostgreSQL and Microsoft SQL Server (SQL Server) are both powerful RDBMS, but SQL Server is a proprietary system mainly used in enterprise environments.

| Feature | PostgreSQL | SQL Server |
|---------|--------------|--------------|
| **License** | Open-source (PostgreSQL License) | Proprietary (Microsoft) |
| **Platform Compatibility** | Runs on Linux, Windows, macOS | Primarily Windows, with some support for Linux |
| **ACID Compliance** | Fully ACID-compliant | Fully ACID-compliant |
| **Performance** | Optimized for complex queries and analytics | Strong performance, optimized for Windows environments |
| **JSON Support** | Full JSONB support | JSON support but not as advanced as PostgreSQL |
| **Stored Procedures & Functions** | Supports PL/pgSQL, PL/Python, and other procedural languages | Uses T-SQL for stored procedures and functions |
| **Scalability** | High scalability with horizontal scaling | Supports vertical scaling with enterprise features |
| **Security** | Advanced role-based access control, SSL encryption | Enterprise-level security, integration with Active Directory |
| **Cost** | Free and open-source | Requires licensing fees, can be expensive for large deployments |

## CRUD
## Foreign key
## Join
## Sub query
## Index
## Partition
## Transaction