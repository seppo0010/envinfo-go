FROM envinfo

FROM ruby:3.0.2-alpine3.13
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--managers"]
