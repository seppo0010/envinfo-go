FROM envinfo

FROM mysql:8.0.27
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
