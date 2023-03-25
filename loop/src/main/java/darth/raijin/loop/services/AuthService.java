package darth.raijin.loop.services;

import darth.raijin.loop.dtos.users.loginUsers.LoginUserRequest;
import darth.raijin.loop.dtos.users.loginUsers.LoginUserResponse;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;
import darth.raijin.loop.services.interfaces.AuthInterface;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataAccessException;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import darth.raijin.loop.dtos.exceptions.domainError.errors.Users.EmailNotUnique;
import darth.raijin.loop.dtos.exceptions.domainError.errors.Users.UsernameNotUnique;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.entities.UserEntity;
import darth.raijin.loop.repositories.AuthRepository;

@Service
public class AuthService implements AuthInterface {

    private AuthRepository authRepository;

    @Autowired
    public AuthService(AuthRepository authBean) {
        this.authRepository = authBean;
    }

    @Override
    public RegisterUserResponse createUser(RegisterUserRequest dto) throws DomainErrorWrapperException {
        dto.validatePassword();

        UserEntity user = new UserEntity(dto.name(), dto.username(), dto.email(), dto.password());

        try {
            user = authRepository.save(user);

        } catch (DataIntegrityViolationException ex) {
            DomainErrorWrapperException wrapper = new DomainErrorWrapperException(
                    "Provided data is not unique",
                    HttpStatus.CONFLICT
            );

            if (ex.getMessage().contains("username"))
                wrapper.appendError(new UsernameNotUnique(dto.username()));

            if (ex.getMessage().contains("email"))
                wrapper.appendError(new EmailNotUnique(dto.email()));

            throw wrapper;

        } catch (DataAccessException ex) {
            throw new DomainErrorWrapperException("Unable to persist to database", HttpStatus.INTERNAL_SERVER_ERROR);
        }

        return new RegisterUserResponse();
    }

    @Override
    public LoginUserResponse loginUser(LoginUserRequest user) throws DomainErrorWrapperException {
        return null;
    }
}
