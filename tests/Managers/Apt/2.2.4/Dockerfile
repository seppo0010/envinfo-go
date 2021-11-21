FROM envinfo

FROM debian:bullseye-20211115
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--managers"]
