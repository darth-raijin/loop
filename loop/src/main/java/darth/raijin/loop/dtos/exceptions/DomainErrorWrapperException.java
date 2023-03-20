package darth.raijin.loop.dtos.exceptions;

import java.time.Instant;
import java.util.List;

import org.springframework.http.HttpStatusCode;


public class DomainErrorWrapperException extends Exception {

    public DomainErrorWrapperException(String message, Instant timestamp, List<Error> errors, HttpStatusCode statusCode) {
        super(message);
        this.timestamp = timestamp;
        this.errors = errors;
        this.statusCode = statusCode;
    }

    private HttpStatusCode statusCode;
    private Instant timestamp;
    private List<Error> errors;

    public void appendError(Error error) {
        errors.add(error);
    }

    public void setTimestamp(Instant timestamp) {
        this.timestamp = timestamp;
    }

    public HttpStatusCode getStatusCode() {
        return statusCode;
    }

    public Instant getTimestamp() {
        return timestamp;
    }

    public List<Error> getErrors() {
        return errors;
    }

    
}
