package darth.raijin.loop.dtos.events.createEvent;

import java.time.Instant;

public record CreateEventRequest(
        String title,
        String description,
        String location,
        Instant date

) {
}
