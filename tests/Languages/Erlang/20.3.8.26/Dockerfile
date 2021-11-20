FROM envinfo

FROM erlang:20.3.8.26-alpine
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
