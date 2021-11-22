FROM envinfo

FROM postgres:10.19-alpine
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
