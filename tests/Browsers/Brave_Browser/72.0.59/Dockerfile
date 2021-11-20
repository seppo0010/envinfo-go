FROM envinfo

FROM sjawhar/brave-browser:latest
COPY --from=0 /build/envinfo ./
ENTRYPOINT ["./envinfo", "--browsers"]
