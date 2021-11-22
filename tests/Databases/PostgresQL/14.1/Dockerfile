FROM envinfo

FROM postgres:14.1-alpine
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
