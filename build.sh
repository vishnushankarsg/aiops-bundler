#!/bin/bash
set -o errexit
BRANCH=$1
BUILD=$2

if [ $BRANCH = "stable" ]
then
  DOCKER_REGISTRY="docker-registry.prod.qw"
else
   DOCKER_REGISTRY="docker-registry.lab.qw"
fi


#write version info
cat <<EOF > version.json
{
   "component": "aiops-bundler",
   "buildBranch": "$BRANCH",
   "buildNumber": "$BUILD",
   "buildDate": "$(date +"%Y-%m-%dT%H:%M:%S")",
   "buildChangeSet": "$(git rev-parse HEAD)"
}
EOF

docker build . -t aiops-bundler -f Dockerfile

docker tag aiops-bundler $DOCKER_REGISTRY/aiops-bundler:$BRANCH-$BUILD
docker tag aiops-bundler $DOCKER_REGISTRY/aiops-bundler:$BRANCH

docker push $DOCKER_REGISTRY/aiops-bundler:$BRANCH-$BUILD
docker push $DOCKER_REGISTRY/aiops-bundler:$BRANCH
