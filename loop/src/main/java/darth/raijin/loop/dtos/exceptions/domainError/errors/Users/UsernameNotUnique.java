package darth.raijin.loop.dtos.exceptions.domainError.errors.Users;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;

public class UsernameNotUnique extends DomainError {
    public final static String domainErrorCode = "1";
    public final static String message = "Username %s is not unique";

    public UsernameNotUnique(String username) {
        DomainError.message = String.format(message, username);  
    }
}
