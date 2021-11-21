FROM envinfo

FROM nginx:1.21
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--servers"]
