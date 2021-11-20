FROM envinfo

FROM perl:5.34.0-slim
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
