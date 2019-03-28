#!/bin/sh

read timeofday

case "$timeofday" in
	yes | y | Yes) echo "good";;
	n* | N* )      echo "aff";;
	*)            echo "sorry";;
esac

exit 0
