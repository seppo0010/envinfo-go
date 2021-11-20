FROM envinfo

FROM php:7.4.26-cli-alpine3.14
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
