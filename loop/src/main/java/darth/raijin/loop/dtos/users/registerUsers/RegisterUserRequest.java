package darth.raijin.loop.dtos.users.registerUsers;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import darth.raijin.loop.dtos.exceptions.domainError.errors.Users.PasswordNotEqual;
import darth.raijin.loop.dtos.exceptions.domainError.errors.Users.WeakPassword;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

import java.util.Objects;

public record RegisterUserRequest(
        @NotBlank String name,
         String username,
        @Email(message = "Email is not valid") String email,
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

    public void validatePassword() throws DomainErrorWrapperException {
        DomainErrorWrapperException wrapper = new DomainErrorWrapperException();

        if (!Objects.equals(this.password, this.repeatedPassword)) {
            wrapper.appendError(new PasswordNotEqual());
        }

        int minimumLength = 8;
        // Assumption will be that 'password' will be the main password that the user wants
        if (password.length() < minimumLength) {
            wrapper.appendError(new WeakPassword("Length is " +
                    password.length() + ", must be at least " + minimumLength
                    ));
        }

        // regex explanation:
        // ^               start of string
        // (?=.*[a-z])     at least one lowercase letter
        // (?=.*[A-Z])     at least one uppercase letter
        // (?=.*\d)        at least one digit
        // [a-zA-Z\d]{8,}  any combination of letters and digits with a length of at least 8
        // $               end of string
        String passwordRegex =
                String.format("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)[a-zA-Z\\d]{%s,}$", minimumLength);

        if (!password.matches(passwordRegex)) {
            wrapper.appendError(new WeakPassword("Password too weak :/"));
        }

        if (!wrapper.getErrors().isEmpty()) {
            throw wrapper;
        }

    }
}

