#!/bin/sh

echo "is it morining?"
read timeofday
case "$timeofday" in
	yes) echo "good"
		echo "test"
		echo "okok";;
	no)  echo "after";;
        y)   echo "short";;
        n)   echo "sort";;
        *)   echo "sorry";;
esac

exit 0
