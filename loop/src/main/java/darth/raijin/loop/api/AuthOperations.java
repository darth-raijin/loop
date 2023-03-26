package darth.raijin.loop.api;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import darth.raijin.loop.dtos.users.loginUsers.LoginUserResponse;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/auth")
public interface AuthOperations {

    @PostMapping("/register")
    @Schema(description = "Used for registering a user")
    @ApiResponse(responseCode = "201", description = "Registered successfully")
    @ApiResponse(responseCode = "409", description = "Either Username or Email is not unique")
    public ResponseEntity<RegisterUserResponse>
    registerUser(@RequestBody @Validated RegisterUserRequest user) throws DomainErrorWrapperException;

    @PostMapping("/login")
    public ResponseEntity<LoginUserResponse> loginUser();
}
