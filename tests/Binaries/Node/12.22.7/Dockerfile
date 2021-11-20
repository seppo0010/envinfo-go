FROM envinfo


FROM node:12.22.7-alpine3.14
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--binaries"]
