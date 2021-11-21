FROM envinfo

FROM httpd:2.4
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--servers"]
