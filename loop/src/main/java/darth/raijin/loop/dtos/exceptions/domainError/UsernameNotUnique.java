package darth.raijin.loop.dtos.exceptions.domainError;

import darth.raijin.loop.dtos.exceptions.DomainError;

public class UsernameNotUnique extends DomainError {
    public final static String statusCode = "409";
    public final static String domainErrorCode = "2";
    public final static String message = "Username %s is not unique";

    public UsernameNotUnique(String username) {
        DomainError.message = String.format(message, username);  
    }
}
