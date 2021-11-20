FROM envinfo

FROM rust:1.48.0-alpine3.12
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
