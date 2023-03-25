package darth.raijin.loop.dtos.users.loginUsers;

import jakarta.validation.constraints.NotBlank;

import java.util.UUID;

public record LoginUserResponse(
        UUID id
) {
    public UUID id() {
        return id;
    }
}
