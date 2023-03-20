plugins {
	java
	id("org.springframework.boot") version "3.1.0-SNAPSHOT"
	id("io.spring.dependency-management") version "1.1.0"
}

group = "darth.raijin.loop"
version = "0.0.1-SNAPSHOT"

java {
    toolchain {
        languageVersion.set(JavaLanguageVersion.of(17))
    }
}

repositories {
	mavenCentral()
	maven { url = uri("https://repo.spring.io/milestone") }
	maven { url = uri("https://repo.spring.io/snapshot") }
}

dependencies {
	implementation("org.springframework.boot:spring-boot-starter-web")
	implementation("org.springframework.boot:spring-boot-starter-data-jpa")
	implementation("org.springframework.boot:spring-boot-starter-jdbc")
	implementation("org.springframework.boot:spring-boot-starter-security")
	implementation("org.springframework.boot:spring-boot-starter-validation")
	implementation("org.liquibase:liquibase-core")
	developmentOnly("org.springframework.boot:spring-boot-devtools")

	runtimeOnly("com.h2database:h2")
	runtimeOnly("org.postgresql:postgresql")

	implementation ("org.yaml:snakeyaml") 

	testImplementation("org.springframework.boot:spring-boot-starter-test")
	testImplementation("org.springframework.security:spring-security-test")
}

tasks.withType<Test> {
	useJUnitPlatform()
}
