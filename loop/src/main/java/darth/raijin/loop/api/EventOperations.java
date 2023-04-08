package darth.raijin.loop.api;

import darth.raijin.loop.dtos.events.createEvent.CreateEventRequest;
import darth.raijin.loop.dtos.events.createEvent.CreateEventResponse;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

@RequestMapping("/events")
public interface EventOperations {

  @PostMapping("/")
  @Schema(description = "Used for registering an user")
  @ApiResponse(responseCode = "201", description = "Created event")
  @ApiResponse(responseCode = "422", description = "Failed creating event")
  public ResponseEntity<CreateEventResponse> createEvent(
      @RequestBody @Validated CreateEventRequest event);

}
