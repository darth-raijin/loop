package darth.raijin.loop.dtos.exceptions.domainError.errors;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;

public class PasswordNotEqual extends DomainError {
    public final static String domainErrorCode = "2";
    public final static String message = "Supplied passwords are not equal";

    public PasswordNotEqual() {

    }
}
