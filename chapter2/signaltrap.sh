#!/bin/sh

trap 'rm -f /tmp/my_tmp_file' INT
echo creating file 
date > /tmp/my_tmp_file

echo "press interpt "
while test -f /tmp/my_tmp_file ;
do
	echo file exists
	sleep 1
done
echo no longer exists

trap INT
echo createing
date > /tmp/my_tmp_file

echo "press inter"
while test -f /tmp/my_tmp_file;
do
	echo file exists
	sleep 1
done

echo we nerver get here
exit 0
