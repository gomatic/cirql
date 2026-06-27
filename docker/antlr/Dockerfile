# ANTLR4 code generator, isolated so grammar authors need only Docker (no JDK).
# Generated Go is committed; normal build/test/CI never touches this image.
FROM eclipse-temurin:21-jre
ARG ANTLR_VERSION=4.13.2
ADD https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar /antlr.jar
WORKDIR /work
ENTRYPOINT ["java", "-jar", "/antlr.jar"]
