FROM envinfo

FROM nginx:1.20
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--servers"]
