package darth.raijin.loop.dtos.users.registerUsers;

import com.fasterxml.jackson.annotation.JsonProperty;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import org.springframework.web.client.HttpClientErrorException;

import java.util.Objects;

public record RegisterUserRequest(
        @NotBlank String name,
         String username,
         String email,
        @NotNull(message = "Password cannot be null") String password,
        @NotNull(message = "Repeated password cannot be null") String repeatedPassword) {

    public String name() {
        return name;
    }

    public String username() {
        return username;
    }

    public String email() {
        return email;
    }

    public String password() {
        return password;
    }

    public String repeatedPassword() {
        return password;
    }

    public Boolean validatePassword() throws DomainErrorWrapperException {
        DomainErrorWrapperException wrapper = new DomainErrorWrapperException();

        if (!Objects.equals(this.password, this.repeatedPassword)) {
            wrapper.appendError(HttpClientErrorException.UnprocessableEntity(""));
        }

        return true;
    }
}

