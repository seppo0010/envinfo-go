FROM envinfo

FROM debian:11.1
COPY --from=0 /build/envinfo ./
CMD ["./envinfo"]
