FROM envinfo

FROM mongo:4.4.10-focal
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
