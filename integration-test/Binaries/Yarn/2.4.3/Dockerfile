FROM envinfo

FROM node:17.1.0-alpine3.12
RUN npm uninstall --global yarn && npm install --global --force yarn@2.4.3
COPY --from=0 /build/envinfo ./
CMD ["./envinfo"]
