package darth.raijin.loop.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.stereotype.Service;

import darth.raijin.loop.dtos.exceptions.domainError.DomainError;
import darth.raijin.loop.dtos.exceptions.domainError.DomainErrorWrapperException;
import darth.raijin.loop.dtos.exceptions.domainError.EmailNotUnique;
import darth.raijin.loop.dtos.exceptions.domainError.UsernameNotUnique;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequestDto;
import darth.raijin.loop.entities.UserEntity;
import darth.raijin.loop.repositories.AuthRepository;

@Service
public class AuthService {

    @Autowired
    private AuthRepository authRepository;

    public UserEntity registerUser(RegisterUserRequestDto dto) {
        UserEntity user = new UserEntity(dto.name(), dto.username(), dto.email(), dto.password());

        try {
            return authRepository.save(user);

        } catch (DataIntegrityViolationException ex) {
            DomainErrorWrapperException wrapper = new DomainErrorWrapperException();

            if (ex.getMessage().contains("username"))
                wrapper.appendError(new UsernameNotUnique(dto.username()));

            if (ex.getMessage().contains("email"))
                wrapper.appendError(new EmailNotUnique(dto.email()));

            throw wrapper;
        }
    }
}
