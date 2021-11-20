FROM envinfo

FROM znly/protoc:0.2.0
COPY --from=0 /build/envinfo ./
ENTRYPOINT ["./envinfo", "--languages"]
