package darth.raijin.loop.dtos.users.loginUsers;

import jakarta.validation.constraints.NotBlank;

import java.util.UUID;

public record LoginUserResponse(UUID id, String username, String email) {

  @Override
  public String username() {
    return username;
  }

  @Override
  public String email() {
    return email;
  }

  @Override
  public UUID id() {
    return id;
  }
}
