package darth.raijin.loop.loop.controllerTests;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.ResultActions;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@SpringBootTest
public class Auth {

  private final String baseUrl = "/auth";

  @Autowired private MockMvc mockMvc;

  @Test
  void createOrder() throws Exception {
    mockMvc
        .perform(
            post("").contentType(MediaType.APPLICATION_JSON).content("{\"amount\": \"EUR100.0\"}"))
        .andExpect(status().is2xxSuccessful());
  }
}
