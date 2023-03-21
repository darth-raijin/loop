package darth.raijin.loop.repositories;

import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;
import darth.raijin.loop.entities.UserEntity;

public interface AuthRepository extends JpaRepository<UserEntity, UUID> {
    
}
