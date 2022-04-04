#!/bin/bash

if [ $# != 1 ]; then
    echo "missing proto type(pb, pbserver)"
    exit 1;
fi

# base dir
BASEDIR=$(dirname "$0")

PROTO_DIR=$BASEDIR/../../proto/$1
PROTOBUF_DIR=$BASEDIR/../../protobuf/$1

if [ ! -d "$PROTO_DIR" ] || [ ! -d "$PROTOBUF_DIR" ]; then
    echo "proto or protobuf dir not exist"
    exit 1;
fi

rm $PROTOBUF_DIR/*.pb.*
protoc -I=$PROTO_DIR --micro_out=$PROTOBUF_DIR --gofast_out=$PROTOBUF_DIR  $PROTO_DIR/*.proto