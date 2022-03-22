#!/usr/bin/env bash

sql_path=.
output_path=..
files=(`find $sql_path -name '*.sql'`)

for file in "${files[@]}"
do
    filename=$(basename $file)
    # https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html
    filename=${filename%.*}
    build_tools --target=orm_v3 --ddlpath=$filename.sql --package=model > $output_path/$filename.ddl.go
done