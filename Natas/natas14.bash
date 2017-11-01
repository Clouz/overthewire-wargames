usr="natas14"
psw="Lg96M10TdfaPyVBkJdjymbllQ5L6qdl1"

curl --user "$usr:$psw" -X POST --form "username=1000" --form "password=xx" -H "debug"  http://$usr.natas.labs.overthewire.org/ 
