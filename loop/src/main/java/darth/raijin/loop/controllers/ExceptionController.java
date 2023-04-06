package darth.raijin.loop.controllers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;

@ControllerAdvice
public class ExceptionController {

  @ExceptionHandler(DomainErrorWrapperException.class)
  public ResponseEntity<DomainErrorWrapperException> handleDomainErrors(
      DomainErrorWrapperException wrapper) {
    ResponseEntity<DomainErrorWrapperException> response =
        new ResponseEntity<DomainErrorWrapperException>(wrapper, null, wrapper.getStatusCode());

    return response;
  }
}
