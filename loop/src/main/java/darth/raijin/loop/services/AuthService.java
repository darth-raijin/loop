package darth.raijin.loop.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.stereotype.Service;

import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequestDto;
import darth.raijin.loop.entities.UserEntity;
import darth.raijin.loop.repositories.AuthRepository;

@Service
public class AuthService {

    @Autowired
    private AuthRepository authRepository;

    public UserEntity registerUser(RegisterUserRequestDto dto) {
        try {
            UserEntity user = new UserEntity(null, null, null, null)
            return authRepository.save(user);

        } catch (DataIntegrityViolationException ex) {
            String errorMessage;
            if (ex.getMessage().contains("username")) {
                errorMessage = "Username already exists";
            } else if (ex.getMessage().contains("email")) {
                errorMessage = "Email already exists";
            } else {
                errorMessage = "User registration failed";
            }
            throw new IllegalArgumentException(errorMessage);
        }
    }

}
