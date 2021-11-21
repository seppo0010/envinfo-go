FROM envinfo

FROM centos:centos6
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--managers"]
