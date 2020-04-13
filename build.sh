docker build -t cloudbuild --target app .
docker rmi $(docker images -f "dangling=true" -q)
docker run -d cloudbuild