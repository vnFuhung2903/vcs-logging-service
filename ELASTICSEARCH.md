# elasticsearch
## Table of content
- [Comparison](#comparison-postgresql-vs-mongodb)
- [CRUD](#crud)
- [Aggregate](#aggregate)
- [Index / Search algorithm](#table-of-content)

## Comparison: Elasticsearch vs SQL databases (MySQL, PostgreSQL)
| Feature | Elasticsearch | SQL Databases |
|---------|----------------|----------------|
| **Data Model** | Document-oriented (JSON-based) | Relational (tables, rows, columns) |
| **Query Language** | DSL (JSON-based Domain Specific Language) | SQL (Structured Query Language) |
| **Indexing** | Inverted index, optimized for text search | B-Tree / Hash indexing... |
| **Schema Flexibility** | Schema-less / dynamic mapping | Strongly typed, rigid schema |
| **Scalability**  Horizontally scalable via sharding | Typically vertical scalable |
| **Search Performance** | Fast for full-text and fuzzy search | Not optimized for full-text search |

## CRUD
CRUD stands for **Create**, **Read**, **Update**, and **Delete**, which are the four basic operations for managing document in Elasticsearch.
### Create
Used to create indices, create an instance of document and index it via the client.

### Read
Used to retrieve documents in Elasticsearch

### Update
Used to update documents or add documents to existing indices in Elasticsearch

### Delete
Used to delete documents from Elasticsearch
## Aggregate
Elasticsearch organizes aggregations into three categories:
- **Metric aggregations** that calculate metrics, such as a sum or average, from field values.
- **Bucket aggregations** that group documents into buckets, also called bins, based on field values, ranges, or other criteria.
- **Pipeline aggregations** that take input from other aggregations instead of documents or fields.

## Index / Search algorithm
### Index
