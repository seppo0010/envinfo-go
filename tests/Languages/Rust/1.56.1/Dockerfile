FROM envinfo

FROM rust:1.56.1-alpine3.13
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
