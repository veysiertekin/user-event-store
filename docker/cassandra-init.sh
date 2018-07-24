cat >/import.cql <<EOF
drop keyspace test;
CREATE keyspace test with replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
CREATE TABLE test.api_keys (api_key text PRIMARY KEY);
CREATE TABLE test.user_events (api_key text, user_id int, event_time timestamp, PRIMARY KEY ((api_key, user_id), event_time));
INSERT into test.api_keys (api_key) values ('TEST_API_KEY');
EOF

until cqlsh -f /import.cql; do
  echo "cqlsh: Cassandra is unavailable to initialize - will retry later"
  sleep 2
done &

exec /docker-entrypoint.sh "$@"