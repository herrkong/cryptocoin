#!/usr/bin/env bash
PROJECT_FOLDER=$(cd "$(dirname "$0")";cd ..;pwd)
cd $PROJECT_FOLDER/bin
go build -o $PROJECT_FOLDER/bin/cobra_linux ../src

