#/bin/bash

nohup docker run -d --rm --name echo -p 3322:3322 echo
