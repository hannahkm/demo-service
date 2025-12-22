#!/bin/bash

export DD_APM_ENABLE_V1_TRACE_ENDPOINT=true
export DD_TRACE_DEBUG=true

DD_TRACE_V1_PAYLOAD_FORMAT_ENABLED=true go run ./sub-service1/main.go &
PID1=$!

DD_TRACE_V1_PAYLOAD_FORMAT_ENABLED=false go run ./sub-service2/main.go &
PID2=$!

wait $PID1
wait $PID2