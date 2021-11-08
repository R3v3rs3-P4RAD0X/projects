OS_NAME=`/bin/cat /etc/os-release | grep PRETTY_NAME | sed 's/\"//g' | sed 's/.*=//g'`
USER_NAME=$(whoami)
PACKAGES=`pacman -Qq | wc -l`

echo -e "Operating System:  " $OS_NAME
echo -e "Username:          " $USER_NAME
echo -e "Packages:          " $PACKAGES