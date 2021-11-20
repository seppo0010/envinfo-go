FROM envinfo

FROM elixir:1.10
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
