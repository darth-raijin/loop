package darth.raijin.loop.repositories;

import java.util.Optional;
import java.util.UUID;

import org.springframework.data.jpa.repository.JpaRepository;
import darth.raijin.loop.entities.UserEntity;

public interface AuthRepository extends JpaRepository<UserEntity, UUID> {
  Optional<UserEntity> findByEmailAndPassword(String email, String password);

  Optional<UserEntity> findByUsernameAndPassword(String username, String password);
}
