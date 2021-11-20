FROM envinfo

FROM znly/protoc:0.4.0
COPY --from=0 /build/envinfo ./
ENTRYPOINT ["./envinfo", "--languages"]
