#!/bin/sh

x=0
while test "$x"  -ne 10;
do
	echo $x
	x=$(($x+1))
done
exit 0
