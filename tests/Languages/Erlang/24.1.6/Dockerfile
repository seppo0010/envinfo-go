FROM envinfo

FROM erlang:24.1.6.0-alpine
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
