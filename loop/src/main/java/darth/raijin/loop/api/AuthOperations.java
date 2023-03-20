package darth.raijin.loop.api;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import darth.raijin.loop.dtos.exceptions.domainError.EmailNotUnique;
import darth.raijin.loop.dtos.exceptions.domainError.UsernameNotUnique;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/auth")
public interface AuthOperations {

    @PostMapping("/register")
    @Schema(description = "Used for registering a user")
    @ApiResponse(responseCode = "201", description = "Registered successfully")
    @ApiResponse(responseCode = UsernameNotUnique.domainErrorCode, description = UsernameNotUnique.message)
    @ApiResponse(responseCode = EmailNotUnique.domainErrorCode, description = EmailNotUnique.message)
    public ResponseEntity registerUser();

    @PostMapping("/login")
    public ResponseEntity loginUser();
}
