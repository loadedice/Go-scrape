#! /bin/sh

while read line
do
    ./scrape $1 $line
    sleep 0.5
done < $2
