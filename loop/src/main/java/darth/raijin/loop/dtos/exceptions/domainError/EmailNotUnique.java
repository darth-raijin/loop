package darth.raijin.loop.dtos.exceptions.domainError;

public class EmailNotUnique extends DomainError {
    public final static String statusCode = "409";
    public final static String domainErrorCode = "4092";
    public final static String message = "Email %s is not unique";

    public EmailNotUnique(String email) {
        DomainError.message = String.format("Email %s is not unique", email);  
    }
}
