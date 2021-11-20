FROM envinfo

FROM r-base:4.1.2
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
