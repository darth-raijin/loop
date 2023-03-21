package darth.raijin.loop.api;

import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequestDto;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponseDto;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/auth")
public interface AuthOperations {

    @PostMapping("/register")
    @Schema(description = "Used for registering a user")
    @ApiResponse(responseCode = "201", description = "Registered successfully")
    @ApiResponse(responseCode = "409", description = "Either Username or Email is not unique")
    public ResponseEntity<RegisterUserResponseDto> 
    registerUser(@RequestBody @Validated RegisterUserRequestDto user);

    @PostMapping("/login")
    public ResponseEntity loginUser();
}
