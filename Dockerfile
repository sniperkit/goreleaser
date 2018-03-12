FROM golang:1.10
RUN apt-get update && \
	apt-get install -y rpm git gnupg docker && \
	rm -rf /var/lib/apt/lists/*
COPY goreleaser /usr/local/bin/goreleaser

