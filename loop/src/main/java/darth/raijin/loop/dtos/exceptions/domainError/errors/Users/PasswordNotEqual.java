package darth.raijin.loop.dtos.exceptions.domainError.errors.Users;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;

public class PasswordNotEqual extends DomainError {
  public static final String domainErrorCode = "2";
  public static final String message = "Supplied passwords are not equal";

  public PasswordNotEqual() {}
}
