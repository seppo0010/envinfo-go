FROM envinfo

FROM ubuntu:20.04
RUN apt-get update && apt-get install -y \
    zsh=5.8-3ubuntu1 \
    && rm -rf /var/lib/apt/lists/*
COPY --from=0 /build/envinfo ./
ENV SHELL=/bin/zsh
CMD ["./envinfo", "--system"]
