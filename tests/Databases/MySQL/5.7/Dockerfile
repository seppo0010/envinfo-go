FROM envinfo

FROM mysql:5.7
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
