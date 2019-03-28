#!/bin/sh

sample_text="golabl"

foo(){
	local sample_text="local"
	echo $sample_text

	echo "funtion is "
}

echo "script starting"
echo $sample_text
foo
echo "script ended"
exit 0
