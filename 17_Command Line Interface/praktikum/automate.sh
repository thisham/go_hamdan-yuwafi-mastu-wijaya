#! /bin/bash

pathName="$1 $(date)" 
mkdir "$pathName"
mkdir "$pathName/about_me"
mkdir "$pathName/about_me/personal"
mkdir "$pathName/about_me/professional"
mkdir "$pathName/my_friends"
mkdir "$pathName/my_system_info"

# about system
echo "My username: $(uname)" >> "$pathName/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$pathName/my_system_info/about_this_laptop.txt"

# socmed
echo "https://www.facebook.com/$2" >> "$pathName/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" >> "$pathName/about_me/professional/linkedin.txt"

# ping
ping 8.8.8.8 -c 3 >> "$pathName/my_system_info/internet_connection.txt"

# friends
curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt >> "$pathName/my_friends/list_of_my_friends.txt"
