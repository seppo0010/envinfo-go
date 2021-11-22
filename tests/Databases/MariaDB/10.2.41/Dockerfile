FROM envinfo

FROM mariadb:10.2.41-bionic
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
