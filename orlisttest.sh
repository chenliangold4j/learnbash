#!/bin/sh

rm -f file_one

if test -f file_one || echo "hello" || echo "there"
then
	echo "in if"
else
	echo "iin else"
fi

exit 0
