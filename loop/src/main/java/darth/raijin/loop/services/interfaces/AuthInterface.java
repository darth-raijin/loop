package darth.raijin.loop.services.interfaces;

import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import darth.raijin.loop.dtos.users.loginUsers.LoginUserRequest;
import darth.raijin.loop.dtos.users.loginUsers.LoginUserResponse;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;

public interface AuthInterface {
    RegisterUserResponse createUser(RegisterUserRequest user) throws DomainErrorWrapperException;
    LoginUserResponse loginUser(LoginUserRequest user) throws DomainErrorWrapperException;
}
