FROM envinfo

FROM mongo:5.0.4-focal
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--databases"]
