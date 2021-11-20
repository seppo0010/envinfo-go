FROM envinfo

FROM jlesage/firefox:v1.17.1
COPY --from=0 /build/envinfo ./
ENTRYPOINT ["./envinfo", "--browsers"]
