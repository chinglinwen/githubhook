set -e

go build

image=asia.gcr.io/alpha2-273207/hugo:$(TZ=Asia/Shanghai date +%Y%m%d_%H%M%S)
docker build -t $image .
# docker push $image