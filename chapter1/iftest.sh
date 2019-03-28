#!/bin/sh

if test -f /bin/bash
then
	echo "file /bin/bash exists"
fi

if test -d /bin/bash
then
	echo "/bin/bash is a directory"
else
	echo "/bin/bash is Not a direcotory"
fi
exit 0
