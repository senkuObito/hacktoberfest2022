FROM maven:3.6.3-jdk-11 as builder
WORKDIR /usr/home/app
COPY pom.xml .
RUN mvn -e -B dependency:resolve
COPY src ./src
RUN mvn -e -B package

FROM adoptopenjdk/openjdk11:alpine-jre
COPY --from=builder /usr/home/app/target/app-*.jar app.jar
EXPOSE 8080
CMD ["java","-Xms256 -Xmx512 -XX:MaxGCPauseMillis=400 -XX:InitiatingHeapOccupancyPercent=90 -XX:ActiveProcessorCount=2 -Djava.security.egd=file:/dev/./urandom", "-jar", "app.jar"]