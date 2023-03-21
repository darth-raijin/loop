package darth.raijin.loop.controllers;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;
import darth.raijin.loop.api.AuthOperations;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequestDto;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponseDto;
import darth.raijin.loop.services.AuthService;

@RestController
public class AuthController implements AuthOperations {
    
    @Autowired
    private AuthService auth;

    @Override
    public ResponseEntity loginUser() {
        
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'loginUser'");
    }

    @Override
    public ResponseEntity<RegisterUserResponseDto> registerUser(RegisterUserRequestDto user) {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'registerUser'");
    }
    
}
