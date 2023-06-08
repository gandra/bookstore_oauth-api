
## Cassandra

### Docker
```
docker run -d --name cassandra_bookstore_oauth -p 9042:9042 cassandra
docker container start cassandra_bookstore_oauth
docker exec -it cassandra_bookstore_oauth bash
```

```
docker rm cassandra_bookstore_oauth
```

### CQSQL
```
docker exec -it cassandra_bookstore_oauth bash -c 'cqlsh'

CREATE KEYSPACE oauth WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};
describe keyspaces;

use oauth;
describe tables;

CREATE TABLE access_tokens(access_token varchar PRIMARY KEY, user_id bigint,  client_id bigint, expires bigint);
select * from access_tokens where access_token = 'abc';
```