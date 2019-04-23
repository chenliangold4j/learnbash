#!bin/sh

#create this file is use to make a swapfile 
dd if=/dev/zero of=/root/swapfile bs=1M count=10240

mkswap /root/swapfile

chmod 0600 /root/swapfile
swapoff #old swapfile
swapon /root/swapfile
