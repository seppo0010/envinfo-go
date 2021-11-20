FROM envinfo

FROM python:3.6.15-alpine3.13
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
