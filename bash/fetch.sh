OS_NAME=`/bin/cat /etc/os-release | grep PRETTY_NAME | sed 's/\"//g' | sed 's/.*=//g'`

echo $OS_NAME
