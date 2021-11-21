FROM envinfo

FROM python:2.7.9
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--managers"]
