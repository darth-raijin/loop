package darth.raijin.loop.dtos.exceptions.domainError;

import darth.raijin.loop.dtos.exceptions.DomainError;

public class EmailNotUnique extends DomainError {
    public final static String statusCode = "409";
    public final static String domainErrorCode = "2";
    public final static String message = "Email %s is not unique";

    public EmailNotUnique(String email) {
        DomainError.message = String.format("Email %s is not unique", email);  
    }
}
