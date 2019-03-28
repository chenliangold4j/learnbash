#!/bin/bash

read unvar

until test  "$unvar" = "ttt"
do
	echo $unvar
	read unvar
done
exit 0
