drop keyspace test;
CREATE keyspace test with replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
CREATE TABLE test.api_key (api_key text PRIMARY KEY);
CREATE TABLE test.user_event (api_key text, user_id int, event_time timestamp, PRIMARY KEY ((api_key, user_id), event_time));
INSERT into test.api_key (api_key) values ('TEST_API_KEY');