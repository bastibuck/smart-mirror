#!/bin/bash

set -e

echo -e "\nProvide GitHub token with repo access"
read -s GITHUB_TOKEN

# get newest artifact ID
REPO_ARTIFACTS=$(curl -L -H "Accept: application/vnd.github+json" -H "X-GitHub-Api-Version: 2022-11-28" https://api.github.com/repos/bastibuck/smart-mirror/actions/artifacts?per_page=1)

ARTIFACT_ID=$(echo $REPO_ARTIFACTS | python -c "import sys, json; print(json.load(sys.stdin)['artifacts'][0]['id'])")

# download and unzip artifact
curl -L \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $GITHUB_TOKEN" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/bastibuck/smart-mirror/actions/artifacts/$ARTIFACT_ID/zip --output smart-mirror.zip

unzip -o smart-mirror.zip

# load image into docker
docker load -i smart-mirror.image.tar

# remove unused old images
docker image prune -f

# create and run new container, auto-remove container after stopping
docker run -d -v ./mounted/db:/app/db -p 3000:3000 -e DATABASE_URL="file:db/db.sqlite" --name smart-mirror smart-mirror-image