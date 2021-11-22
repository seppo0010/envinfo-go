FROM envinfo

FROM mariadb:10.7.1-focal
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
