#!/bin/sh
mysqldef -u $MYSQL_USER -p $MYSQL_PASSWORD -h $MYSQL_HOST -P $MYSQL_PORT go_database < ./01_schema.sql

