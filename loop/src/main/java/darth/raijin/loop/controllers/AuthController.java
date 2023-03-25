package darth.raijin.loop.controllers;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;
import darth.raijin.loop.api.AuthOperations;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserRequest;
import darth.raijin.loop.dtos.users.registerUsers.RegisterUserResponse;
import darth.raijin.loop.services.AuthService;

@RestController
public class AuthController implements AuthOperations {
    
    private AuthService auth;

    @Autowired
    public AuthController(AuthService authBean) {
        this.auth = authBean;
    }

    @Override
    public ResponseEntity loginUser() {
        
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'loginUser'");
    }

    @Override
    public ResponseEntity<RegisterUserResponse> registerUser(RegisterUserRequest user) {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'registerUser'");
    }
    
}
