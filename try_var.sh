#!/bin/sh

salutation="test"
echo $salutation
echo "The program $0 is now running"
echo "th second parameter was $2"
echo "the first parameter was $1"
echo "the parameter list was $*"
echo "the user's home directory is $HOME"
echo "Please enter anew greeting"
read salutation

echo $salutation
echo "The script is now comlete"
exit 0

