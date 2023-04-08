package darth.raijin.loop.dtos.events.createEvent;

import java.time.Instant;

public record CreateEventResponse(
        String id,
        String title,
        String description,
        String location,
        Instant date

) {
}
