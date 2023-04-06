package darth.raijin.loop;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;

@SpringBootApplication
@ConfigurationProperties
@EnableWebSecurity
public class LoopApplication {

  public static void main(String[] args) {
    SpringApplication.run(LoopApplication.class, args);
  }
}
