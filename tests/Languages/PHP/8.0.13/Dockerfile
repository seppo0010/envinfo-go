FROM envinfo

FROM php:8.0.13-cli-alpine3.14
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
