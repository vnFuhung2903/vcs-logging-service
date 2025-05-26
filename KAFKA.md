# kafka
## Table of content
- [Definition, purpose](#definition-purpose)
- [Comparison](#comparison-with-other-message-queue)
- [Producer / Consumer](#producer--consumer)

## Definition, purpose
Apache Kafka is a distributed streaming platform used designed to handle large volumes of real-time data. It’s an open-source system used for stream processing, real-time data pipelines and data integration. It was built on the concept of publish/subscribe model and provides high throughput, reliability and fault tolerance.
![alt text](image.png)
Kafka is built on 5 purposes:
- Real-time Data Pipelines
- Messaging Systems
- Stream Processing
- Event-driven Architecture
- Log Aggregation

## Comparison with other message queue
| Feature                     | Apache Kafka                                     | Redis Streams                                  | RabbitMQ                                        |
|----------------------------|--------------------------------------------------|------------------------------------------------|-------------------------------------------------|
| **Category** | Distributed event streaming platform | In-memory data structure with stream support  | Traditional message broker (AMQP-based) |
| **Persistence** | Durable log storage on disk | In-memory | Persistent queues supported (optional) |
| **Retention Model** | Time-based or size-based retention, configurable | Manual trimming required | Messages removed after acknowledgment |
| **Consumer Model** | Pull-based | Pull-based | Push/pull (configurable prefetching) |
| **Replay Support** | Yes (via offsets) | Yes (via stream IDs) | No (unless requeued manually) |
| **Scalability** | Partitioned, horizontal scaling | Scales with Redis Cluster | Clustering supported |
| **Latency** | Low | Very low (in-memory) | Low to medium |
| **Throughput** | Very high (millions of messages/sec) | High (memory-limited) | Medium |
| **Protocol** | Custom TCP | RESP (Redis protocol) | AMQP, MQTT, STOMP |
| **Message Ordering** | Per partition | Guaranteed by ID sequence | Queue-level FIFO |

## Producer / Consumer
### Producer
Producers are client applications that write event messages to topics in the Kafka cluster. Messages are stored in the form of key-values in the partitions of the topics. Messages produced with the same key are written to the same partition. Sending keys during production is not mandatory. In case the key (`key=null`) is not specified by the producer, messages are distributed evenly across partitions in a topic. This means messages are sent in a round-robin fashion (partition *p0* then *p1* then *p2*, etc... then back to *p0* and so on...). Kafka message keys are commonly used when there is a need for message ordering for all messages sharing the same field.

### Consumer
Kafka Consumers is used to reading data from a topic and remember a topic again is identified by its name.\
A Consumer Group will manage a set of single consumers, allowing Kafka and the consumers to distribute messages based on Kafka partitions. When a consumer group contains just one consumer, it will get all messages from all partitions in the Kafka topic. Each consumer group is identified by a group id. It must be unique for each group, that is, two consumers that have different group ids will not be in the same group.\
Kafka does not have an explicit way of tracking which message has been read by a consumer of a group. Instead, it allows consumers to track the offset (the position in the queue) of what messages it has read for a given partition in a topic. To do this, a consumer in a consumer group will publish a special message topic for that topic/partition with the committed offset for each partition it has gotten up to.\
When a new consumer is added to a consumer group, it will start consuming messages from partitions previously assigned to another consumer. It will generally pick up from the most recently committed offset. If a consumer leaves or crashes, the partitions it used to consume will be picked up by one of the consumers that is still in the consumer group. This change in partitions can also occur when a topic gets modified or more partitions are added to the topic. The process of changing which consumer is assigned to what partition is called rebalancing. Rebalances are a normal part of Apache Kafka operations and will occur during configuration changes, scaling, or if a broker or consumer crashes.\
Consumers in a group share ownership of a topic based on the partitions within that topic. Consumers are considered to be members of a consumer group and membership is maintained by a heartbeat mechanism. Consumers are considered to be part of the consumer group if they continue to send heartbeats to a Kafka broker designated as the group coordinator. Brokers are different for different consumer groups. If the Group Coordinator does not see a heartbeat from a consumer within a certain amount of time, it will consider the consumer to be dead and will start a rebalance. During the time from when a coordinator has seen the last heartbeat and when it marks a consumer as dead, messages for that partition are likely to build up as they are not being processed. If a consumer is shutdown cleanly it will notify a coordinator that it is leaving and minimize this window of unprocessed messages.