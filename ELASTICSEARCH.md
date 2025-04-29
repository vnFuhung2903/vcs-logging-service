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
### Index API (Create, Update)
Adds a JSON document to the specified data stream or index and makes it searchable. If the target is an index and the document already exists, the request updates the document and increments its version.

### Get API, Search API (Read)
The Get API retrieves a document directly by its unique ID, providing a simple and fast way to access known records.\
The Search API are used to search and aggregate data stored in Elasticsearch indices and data streams. A search consists of one or more queries that are combined and sent to Elasticsearch. Documents that match a search’s queries are returned in the *hits*, or *search results*, of the response.A search may also contain additional information used to better process its queries.

### Update API, Update By Query API (Update)
Enables script document updates. The update API also supports passing a partial document, which is merged into the existing document. To fully replace an existing document, the **Index API** should be used.\
**Update By Query API** is used to update documents that match the specified query. If no query is specified, performs an update on every document in the data stream or index without modifying the source, which is useful for picking up mapping changes.

### Delete API, Delete By Query API (Delete)
Remove a document from an index. The index name and document ID must be specified when using **Delete API**.\
When using **Delete By Query API**, Elasticsearch gets a snapshot of the data stream or index when it begins processing the request and deletes matching documents using internal versioning. If a document changes between the time that the snapshot is taken and the delete operation is processed, it results in a version conflict and the delete operation fails. While processing a delete by query request, Elasticsearch performs multiple search requests sequentially to find all of the matching documents to delete. A bulk delete request is performed for each batch of matching documents. If a search or bulk request is rejected, the requests are retried up to 10 times, with exponential back off. If the maximum retry limit is reached, processing halts and all failed requests are returned in the response. Any delete requests that completed successfully still stick, they are not rolled back.

## Aggregate
Elasticsearch organizes aggregations into three categories:
- **Metric aggregations** that calculate metrics, such as a sum or average, from field values.
- **Bucket aggregations** that group documents into buckets, also called bins, based on field values, ranges, or other criteria.
- **Pipeline aggregations** that take input from other aggregations instead of documents or fields.

Aggregations is part of a search by specifying the **Search API**'s aggs parameter. Aggregation results are in the response’s *aggregations* object

## Index / Search algorithm
### Index
An Elasticsearch index is a logical namespace that holds a collection of documents, where each document is a collection of fields, which are key-value pairs that contain your data.\
In Elasticsearch, denormalization is a common practice. Instead of splitting data across multiple tables, all the relevant information is stored in a single JSON document. This allows for faster and more efficient retrieval of data in Elasticsearch during search operations. As a general rule of thumb, storage can be cheaper than compute costs for joining data.\
Each index is identified by a unique name and is divided into one or more shards, which are smaller subsets of the index that allow for parallel processing and distributed storage across a cluster of Elasticsearch nodes.  Shards have a primary and a replica shard. Replicas provide redundant copies of your data to protect against hardware failure and increase capacity to serve read requests like searching or retrieving a document. Adding more nodes into the cluster gives more capacity for indexing and searching, something that’s not so easily achieved with a relational database.\
Elasticsearch uses a data structure called an **inverted index** that supports very fast full-text searches. An inverted index lists every unique word that appears in any document and identifies all of the documents each word occurs in. By default, Elasticsearch indexes all data in every field and each indexed field has a dedicated, optimized data structure. For example, text fields are stored in inverted indices, and numeric and geo fields are stored in BKD trees. The ability to use the per-field data structures to assemble and return search results is what makes Elasticsearch so fast.

### Search algorithm
- **Lexical/Full-text search** algorithms: Using inverted indices for full-text searches
- **Semantic search** algorithms: Leveraging machine learning to capture meaning and content of unstructured data, including text and images, transforming it into embeddings, which are numerical representations stored in vectors. Semantic search uses **approximate nearest neighbor (ANN) algorithms**, then matches vectors of existing documents (a semantic search concerns text) to the query vectors
