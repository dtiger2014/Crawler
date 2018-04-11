#!/bin/sh
export GOPATH=$PWD
export GOBIN=$PWD/bin
rm -rf $PWD/bin/*

go install crawler

echo "----build finish----"
