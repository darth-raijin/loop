package darth.raijin.loop.dtos.exceptions;

import java.time.Instant;
import java.util.List;


public class DomainErrorWrapperException extends Exception {

    public DomainErrorWrapperException(String message, Instant timestamp, List<Error> errors) {
        super(message);
        this.timestamp = timestamp;
        this.errors = errors;
    }

    private Instant timestamp;
    private List<Error> errors;

    public void appendError(Error error) {
        errors.add(error);
    }

    public void setTimestamp(Instant timestamp) {
        this.timestamp = timestamp;
    }
}
