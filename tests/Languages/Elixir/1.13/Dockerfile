FROM envinfo

FROM elixir:1.13
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
