FROM envinfo

FROM node:17.1.0-alpine3.12
COPY --from=0 /build/envinfo ./
CMD ["./envinfo"]
