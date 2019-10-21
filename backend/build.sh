#!/usr/bin/env bash
pushd src > /dev/null
go build dslic
#mv dslic ../
popd > /dev/null
echo build done.
