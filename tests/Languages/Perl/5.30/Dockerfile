FROM envinfo

FROM perl:5.30.3-slim
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
