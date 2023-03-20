package darth.raijin.loop.api;

import java.net.http.HttpResponse;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/auth")
public interface AuthOperations {

    @PostMapping("/register")
    @Schema(description = "Used for registering a user")
    @ApiResponse(responseCode = "201", description = "Registered successfully")
    @ApiResponse(responseCode = "409", description = "Username is not unique: 4090")
    @ApiResponse(responseCode = "409", description = "Email is not unique: 4090")
    public ResponseEntity registerUser();

    @PostMapping("/login")
    public ResponseEntity loginUser();
}
