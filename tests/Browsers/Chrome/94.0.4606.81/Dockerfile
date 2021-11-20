FROM envinfo

FROM browserless/chrome:1.49-chrome-stable
COPY --from=0 /build/envinfo ./
ENTRYPOINT ["./envinfo", "--browsers"]
