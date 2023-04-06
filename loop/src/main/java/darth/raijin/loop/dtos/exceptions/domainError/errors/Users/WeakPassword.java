package darth.raijin.loop.dtos.exceptions.domainError.errors.Users;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;

public class WeakPassword extends DomainError {
  public static final String domainErrorCode = "4";
  public static final String message = "Password is weak: %s";

  public WeakPassword(String error) {
    DomainError.message = String.format(message, error);
  }
}
