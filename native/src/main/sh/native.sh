#!/bin/sh
JAR=native-jar-with-dependencies.jar
BIN="`readlink $0`"
DIR="`dirname "$BIN"`"
exec java -jar "$DIR/$JAR"
