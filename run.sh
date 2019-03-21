#/bin/bash
rm -f nohup.out
nohup docker run --rm --name echo -p 3322:3322 echo &
