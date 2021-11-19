FROM envinfo

FROM golang:1.17.3-alpine3.14
COPY --from=0 /build/envinfo ./
CMD ["./envinfo"]
