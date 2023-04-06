package darth.raijin.loop.services;

import darth.raijin.loop.dtos.users.loginUsers.LoginUserRequest;
import darth.raijin.loop.dtos.users.loginUsers.LoginUserResponse;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;
import darth.raijin.loop.services.interfaces.AuthInterface;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataAccessException;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.HttpStatus;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import darth.raijin.loop.dtos.exceptions.domainError.errors.Users.EmailNotUnique;
import darth.raijin.loop.dtos.exceptions.domainError.errors.Users.UsernameNotUnique;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.entities.UserEntity;
import darth.raijin.loop.repositories.AuthRepository;

import java.util.Optional;

@Service
public class AuthService implements AuthInterface {

  private AuthRepository authRepository;
  private BCryptPasswordEncoder bCryptPasswordEncoder;

  @Autowired
  public AuthService(AuthRepository authBean, BCryptPasswordEncoder bCrypt) {
    this.authRepository = authBean;
    this.bCryptPasswordEncoder = bCrypt;
  }

  @Override
  public RegisterUserResponse createUser(RegisterUserRequest dto)
      throws DomainErrorWrapperException {
    dto.validatePassword();

    UserEntity user =
        new UserEntity(
            dto.name(), dto.username(), dto.email(), bCryptPasswordEncoder.encode(dto.password()));

    try {
      user = authRepository.save(user);

    } catch (DataIntegrityViolationException ex) {
      DomainErrorWrapperException wrapper =
          new DomainErrorWrapperException("Provided data is not unique", HttpStatus.CONFLICT);

      if (ex.getMessage().contains("username"))
        wrapper.appendError(new UsernameNotUnique(dto.username()));

      if (ex.getMessage().contains("email")) wrapper.appendError(new EmailNotUnique(dto.email()));

      throw wrapper;

    } catch (DataAccessException ex) {
      throw new DomainErrorWrapperException(
          "Unable to persist to database", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    return new RegisterUserResponse(user.getUsername(), user.getEmail());
  }

  @Override
  public LoginUserResponse loginUser(LoginUserRequest user) throws DomainErrorWrapperException {
    UserEntity login = new UserEntity(null, null, null, null);
    Optional<UserEntity> foundUser;

    if (login.getUsername() != null) {
      foundUser =
          authRepository.findByEmailAndPassword(
              login.getUsername(), bCryptPasswordEncoder.encode(login.getPassword()));

    } else {
      foundUser =
          authRepository.findByUsernameAndPassword(
              login.getUsername(), bCryptPasswordEncoder.encode(login.getPassword()));
    }

    if (foundUser.isEmpty()) {
      throw new DomainErrorWrapperException("Invalid credentials", HttpStatus.UNAUTHORIZED);
    }

    return new LoginUserResponse(
        foundUser.get().getId(), foundUser.get().getUsername(), foundUser.get().getEmail());
  }

  @Override
  public Object resetPassowrd() throws DomainErrorWrapperException {
    return null;
  }
}
