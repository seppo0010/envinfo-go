FROM envinfo

FROM python:2.7.18
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
