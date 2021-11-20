FROM envinfo

FROM ubuntu:20.04
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--system"]
