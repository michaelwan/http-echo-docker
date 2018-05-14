#/bin/bash
rm -f nohup.out
nohup docker run -p 18007:8080 echo &
