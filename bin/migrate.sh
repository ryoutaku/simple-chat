#!/bin/bash

cd `dirname $0`
export $(cat ../.env | grep -v ^# | xargs);

OPTIONS="-config=../db/mysql/dbconfig.yml -env=$GO_ENV"
sql-migrate ${@} $OPTIONS
