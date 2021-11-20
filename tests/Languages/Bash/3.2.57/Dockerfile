FROM envinfo

FROM debian:bullseye-20211115
RUN apt-get update && apt-get install -y \
     byacc \
     wget \
     gcc \
     make \
     && rm -rf /var/lib/apt/lists/
RUN mkdir /bash &&\
     cd /bash && \
     wget http://ftp.gnu.org/gnu/bash/bash-3.2.57.tar.gz && \
     tar xvzf bash-3.2.57.tar.gz && \
     cd bash-3.2.57 && \
     ./configure && \
     make && \
     make install &&\
     cd / &&\
     rm -rf /bash
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
