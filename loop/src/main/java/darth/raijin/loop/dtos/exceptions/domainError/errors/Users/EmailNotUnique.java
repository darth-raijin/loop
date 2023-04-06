package darth.raijin.loop.dtos.exceptions.domainError.errors.Users;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;

public class EmailNotUnique extends DomainError {
  public static final String domainErrorCode = "3";
  public static final String message = "Email %s is not unique";

  public EmailNotUnique(String email) {
    DomainError.message = String.format(message, email);
  }
}
