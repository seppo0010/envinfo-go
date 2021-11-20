FROM envinfo

FROM r-base:3.3.2
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
