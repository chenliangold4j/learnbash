#!/bin/sh

rm -rf fred*
echo 1> fred1
echo 2>fred2
mkdir fred3
echo 4>fred4

for file in fred*
do
	if test -d "$file";
       	then
		echo "skpping direcotry $file"
		continue
	fi
	echo file is $file
done
rm -rf fred*
exit 0
