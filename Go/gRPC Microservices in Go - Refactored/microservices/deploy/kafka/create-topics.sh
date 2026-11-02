#!/bin/sh
# Creates the topics the services use. Safe to run again (--if-not-exists).
set -eu
BOOTSTRAP="${BOOTSTRAP:-kafka:19092}"
TOPICS=/opt/kafka/bin/kafka-topics.sh

# Event streams: 3 partitions so 3 instances of a consumer can share the work.
# Records are keyed by order_id, so one order's events stay in order.
for topic in orders shipments; do
  $TOPICS --bootstrap-server "$BOOTSTRAP" --create --if-not-exists \
    --topic "$topic" --partitions 3 --replication-factor 1 \
    --config retention.ms=604800000 # 7 days
done

# Dead-letter topics: low volume, kept 14 days for inspection and replay.
for topic in order.dlq shipping.dlq; do
  $TOPICS --bootstrap-server "$BOOTSTRAP" --create --if-not-exists \
    --topic "$topic" --partitions 1 --replication-factor 1 \
    --config retention.ms=1209600000
done

$TOPICS --bootstrap-server "$BOOTSTRAP" --list
