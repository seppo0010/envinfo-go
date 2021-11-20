FROM envinfo

FROM ruby:2.6.8-alpine3.13
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages"]
