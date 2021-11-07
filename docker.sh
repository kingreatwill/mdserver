docker build -t mdserve:golang .
docker stop mdserve_golang
docker rm mdserve_golang
docker run -itd -p 8002:8080 -v /root/github/open:/root/mdfolder --name mdserve_golang mdserve:golang
# 资源文件路径
# docker run -itd -p 8002:8080 -v /root/github/open:/root/mdfolder -v /xx/static:/root/static --name mdserve_golang mdserve:golang