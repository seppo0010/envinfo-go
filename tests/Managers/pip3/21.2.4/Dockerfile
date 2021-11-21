FROM envinfo

FROM python:3.10.0-alpine3.14
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--managers"]
