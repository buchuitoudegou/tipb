#!/bin/bash

if [ -z "$PROTOC" ]; then
    echo "use system protoc."
    PROTOC="protoc"
fi

set -euo pipefail

rm -rf proto-cpp && mkdir -p proto-cpp
rm -rf cpp/tipb && mkdir cpp/tipb

cp proto/*.proto proto-cpp/

function sed_inplace()
{
    if sed --help 2>/dev/null | grep GNU > /dev/null; then
        sed -i "$@"
    else
        sed -i '' "$@"
    fi
}

sed_inplace '/gogo.proto/d' proto-cpp/*
sed_inplace '/option\ (gogoproto/d' proto-cpp/*
sed_inplace -e 's/\[.*gogoproto.*\]//g' proto-cpp/*

cd proto-cpp
echo "generate cpp code..."
${PROTOC} --cpp_out=../cpp/tipb/ *.proto
cd ..

rm -rf proto-cpp
