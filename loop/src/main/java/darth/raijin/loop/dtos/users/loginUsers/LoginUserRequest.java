package darth.raijin.loop.dtos.users.loginUsers;

import jakarta.validation.constraints.NotBlank;

public record LoginUserRequest(String username, String email, @NotBlank String password) {

  public String username() {
    return username;
  }

  public String email() {
    return email;
  }

  public String password() {
    return password;
  }
}
