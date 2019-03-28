#!/bin/sh

read timeofday
if test "$timeofday" = "yes"
then
	echo "moring"
elif test "$timeofday" = "no"
then
	echo "good "
else
	echo "sorry"
	exit 1
fi
exit 0

