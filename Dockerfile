FROM golang:1.10
RUN apt-get update && \
	apt-get install -y rpm git gnupg docker snapd && \
	rm -rf /var/lib/apt/lists/* && \
	snap install snapcraft --classic
ENV PATH=/snap/bin:$PATH
COPY goreleaser /usr/local/bin/goreleaser

