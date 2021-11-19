FROM envinfo

FROM golang:1.16.10-alpine3.14
COPY --from=0 /build/envinfo ./
CMD ["./envinfo"]
