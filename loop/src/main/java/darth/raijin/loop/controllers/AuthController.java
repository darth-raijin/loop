package darth.raijin.loop.controllers;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;
import darth.raijin.loop.api.AuthOperations;

@RestController
public class AuthController implements AuthOperations {

    @Override
    public ResponseEntity registerUser() {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'registerUser'");
    }

    @Override
    public ResponseEntity loginUser() {
        
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'loginUser'");
    }
    
}
