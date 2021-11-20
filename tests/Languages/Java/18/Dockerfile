FROM envinfo

FROM openjdk:18-jdk-slim-bullseye
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages", "-vvvv"]
