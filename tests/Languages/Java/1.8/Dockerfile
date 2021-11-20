FROM envinfo

FROM openjdk:8-jdk-slim-bullseye
COPY --from=0 /build/envinfo ./
CMD ["./envinfo", "--languages", "-vvvv"]
