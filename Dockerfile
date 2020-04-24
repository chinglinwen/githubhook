FROM centos:7
WORKDIR /app
RUN yum install -y git wget && \
    wget https://github.com/gohugoio/hugo/releases/download/v0.69.2/hugo_0.69.2_Linux-64bit.tar.gz && \
    tar -xzf *.gz && \
    mv hugo /bin/ && \
    rm -f LICENSE  README.md *.gz && \
    cd /app && \
    git clone https://github.com/chinglinwen/blogs.git

COPY githubhook trigger.sh start.sh /app/

CMD /app/start.sh 
EXPOSE 8080 3000