package darth.raijin.loop.controllers;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;
import darth.raijin.loop.api.AuthOperations;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;
import darth.raijin.loop.services.AuthService;

@RestController
public class AuthController implements AuthOperations {

  private AuthService auth;

  public AuthController(AuthService authBean) {
    this.auth = authBean;
  }

  @Override
  public ResponseEntity loginUser() {

    // TODO Auto-generated method stub
    throw new UnsupportedOperationException("Unimplemented method 'loginUser'");
  }

  @Override
  public ResponseEntity<RegisterUserResponse> registerUser(
      @Valid @RequestBody RegisterUserRequest user) throws DomainErrorWrapperException {
    return new ResponseEntity<>(auth.createUser(user), HttpStatus.CREATED);
  }
}
