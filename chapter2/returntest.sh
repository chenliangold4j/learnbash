#!/bin/sh

yes_or_no(){
	echo "is you name $* ?"
	while true
	do
		echo -n "enter yes or no: "
		read x
		case "$x" in
			[yY]*)return 0;;
			[nN]*)return 1;;
			*) echo "error";;
		esac
	done
}

echo "parameters are $*"

if yes_or_no "$1"
then
	echo "hi $1,nice name"
else
	echo "never mind"
fi
exit 0
