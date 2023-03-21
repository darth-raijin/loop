package darth.raijin.loop.dtos.users.registerUsers;

import com.fasterxml.jackson.annotation.JsonProperty;

import jakarta.validation.constraints.NotBlank;

public record RegisterUserRequestDto(
        @NotBlank String name,
        @NotBlank String username,
        @NotBlank String email,
        @NotBlank String password) {

    @JsonProperty("name")
    public String name() {
        return name;
    }

    @JsonProperty("username")
    public String username() {
        return username;
    }

    @JsonProperty("email")
    public String email() {
        return email;
    }

    @JsonProperty("password")
    public String password() {
        return password;
    }
}

