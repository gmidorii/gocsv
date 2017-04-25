#!bin/bash

# $1 = search condition
# $2 = csv file

for str in `cat $1`
do
	# C0 coverage
	./gocsv -f $2 -c 3:4 -g $str
done
