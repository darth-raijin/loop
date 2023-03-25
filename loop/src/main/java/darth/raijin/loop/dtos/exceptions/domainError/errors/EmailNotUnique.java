package darth.raijin.loop.dtos.exceptions.domainError.errors;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;

public class EmailNotUnique extends DomainError {
    public final static String domainErrorCode = "3";
    public final static String message = "Email %s is not unique";

    public EmailNotUnique(String email) {
        DomainError.message = String.format("Email %s is not unique", email);  
    }
}
