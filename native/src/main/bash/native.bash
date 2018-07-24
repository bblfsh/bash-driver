#!/usr/bin/env bash
JAR=native-jar-with-dependencies.jar
BIN="`readlink -f $0`"
DIR="`dirname "$BIN"`"
exec java -jar "$DIR/$JAR"
