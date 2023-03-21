package darth.raijin.loop.dtos.exceptions.domainError;

public class UsernameNotUnique extends DomainError {
    public final static String statusCode = "409";
    public final static String domainErrorCode = "4090";
    public final static String message = "Username %s is not unique";

    public UsernameNotUnique(String username) {
        DomainError.message = String.format(message, username);  
    }
}
